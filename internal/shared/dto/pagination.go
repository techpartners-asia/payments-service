package sharedDTO

type (
	SharedPaginationRequestDTO struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}

	SharedPaginationResponseDTO[T any] struct {
		Total int64 `json:"total"`
		Items []*T  `json:"items"`
	}
)

func NewSharedPaginationResponseDTO[T any](total int64, items []*T) *SharedPaginationResponseDTO[T] {
	return &SharedPaginationResponseDTO[T]{
		Total: total,
		Items: items,
	}
}
