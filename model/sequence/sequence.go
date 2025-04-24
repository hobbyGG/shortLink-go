package sequence

// sequence sequence表相关的功能
type Methods interface {
	Get() (uint64, error)
}
