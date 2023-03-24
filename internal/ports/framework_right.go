package ports

type DbPort interface {
    CloseConnection()
    AddToHistory(answer int32, operation string) error
}
