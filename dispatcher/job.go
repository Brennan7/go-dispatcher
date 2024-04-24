package dispatcher

// Job defines an interface that allows different types of jobs to be created with their own processing logic.
type Job interface {
	Process() error
}
