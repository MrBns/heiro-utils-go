package mirrornode

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func api_url(path string) string {
	return GetConfig().BASE_URL + path
}

func checkApiStatus(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		var api_err MirrorNodeAPIErrors
		if resp.Body != nil {
			defer resp.Body.Close()
			if err := json.NewDecoder(resp.Body).Decode(&api_err); err != nil {
				return fmt.Errorf("failed to decode response body; status=%v", resp.StatusCode)
			}
			errorsString := ""
			for _, e := range api_err.Status.Messages {
				errorsString += e.Message + "; "
			}
			if errorsString != "" {
				return errors.New(errorsString)
			}
			return fmt.Errorf("unexpected API error format; status=%v", resp.StatusCode)
		}
		return fmt.Errorf("unexpected API errors; status=%#v", resp.StatusCode)
	}
	return nil
}
