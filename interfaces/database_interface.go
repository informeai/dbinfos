package interfaces

// IDatabase is interface for databases
type IDatabase interface {
	Save(topic string, infos any) error
	Connect() error
}
