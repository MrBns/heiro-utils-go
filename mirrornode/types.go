package mirrornode

type MirrorNodeAPIErrors struct {
	Status struct {
		Messages []struct {
			Message string `json:"message"`
		} `json:"messages"`
	} `json:"_status"`
}

// Links
type Links struct {
	Next *string `json:"next"`
}

type MirrorNodeResponse[T any] struct {
	Data   *T
	Errors *MirrorNodeAPIErrors
}

type MirrorNodeListResponse[T any] struct {
	Data   *T
	Errors *MirrorNodeAPIErrors
}
