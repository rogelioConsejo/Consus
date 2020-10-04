package file_saver

import (
	"errors"
	"github.com/rogelioConsejo/Consus/storage/storable"
)

type fileSaver struct {

}

func FileSaver() *fileSaver {
	return &fileSaver{}
}

func (fs *fileSaver) Create(s *storable.Type) (err error, id string)  {
	return errors.New("dto.create not implemented"), id
}

func (fs *fileSaver) Read(id string) (error, *storable.Type)  {
	return errors.New("dto.read not implemented"), nil
}

func (fs *fileSaver) Delete(id string) error {
	return errors.New("dto.delete not implemented")
}



