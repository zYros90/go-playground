package house

type House struct {
	numberFloors  int
	withConcrete  bool
	withFireplace bool
}

type Option func(*House)

const (
	defaultNumberFloors  = 1
	defaultWithConcrete  = false
	defaultWithFireplace = false
)

func New(opts ...Option) *House {
	h := &House{
		withConcrete:  defaultWithConcrete,
		withFireplace: defaultWithFireplace,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func WithConcrete() Option {
	return func(h *House) {
		h.withConcrete = true
	}
}

func WithFireplace() Option {
	return func(h *House) {
		h.withFireplace = true
	}
}

func WithNumberFloors(n int) Option {
	return func(h *House) {
		h.numberFloors = n
	}
}

func (h *House) GetNumberFloors() int {
	return h.numberFloors
}
