package quick

// Application represents an GoQuick application.
type Application interface {
        Load(source interface{}) error
        Exec() int
}

type View interface {
        SetSource(source interface{}) error
        Show()
}
