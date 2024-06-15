package interfaces

import "github.com/informeai/dbinfos/entities"

type IDatabase interface {
	Save(entities.DBInfo) error
}
