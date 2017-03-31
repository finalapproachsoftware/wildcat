package biz

import (
	"github.com/finalapproachsoftware/wildcat/models"
	"gopkg.in/mgo.v2"
)

type TenantManager struct {
	Tenants ICollection
}

func NewTenantManager(s *mgo.Session) *TenantManager {
	tmgr := new(TenantManager)
	_, tmgr.Tenants = NewMongoCollection(s, "wildcat", "Tenants")
	return tmgr
}

func (tmgr *TenantManager) Create(t *models.Tenant) error {
	err := tmgr.Tenants.Add(t)
	return err
}
