package main

import "github.com/rogelioConsejo/Consus/domain"

type StorableSaver interface {
	Save(storable domain.Storable) (error, uint)
}

type StorableRetriever interface {
	Retrieve(id uint) (error, domain.Storable)
}