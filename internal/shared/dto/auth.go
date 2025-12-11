package sharedDTO

type (
	SharedAuthResponseDTO[T any] struct {
		Token string `json:"token"`
		Data  *T     `json:"data"`
	}

	SharedAuthRequestDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewSharedAuthResponseDTO[T any](token string, data *T) *SharedAuthResponseDTO[T] {
	return &SharedAuthResponseDTO[T]{
		Token: token,
		Data:  data,
	}
}
