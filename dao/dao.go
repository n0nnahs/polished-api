package dao

import (
	"gopkg.in/mgo.v2/bson"
)


func (m *PolishDAO) FindAll() ([]polish.Polish, error) {
	var polish []polish.Polish
	err := db.C(COLLECTION).Find(bson.M{}).All(&polish)
	return polish, err
}

func (m *PolishDAO) FindById(id string) (polish.Polish, error) {
	var polish polish.Polish
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&polish)
	return polish, err
}

func (m *PolishDAO) Insert(polish polish.Polish) error {
	err := db.C(COLLECTION).Insert(&polish)
	return err
}

func (m *PolishDAO) Delete(polish polish.Polish) error {
	err := db.C(COLLECTION).Remove(&polish)
	return err
}

func (m *PolishDAO) Update(polish polish.Polish) error {
	err := db.C(COLLECTION).UpdateId(polish.ID, &polish)
	return err
}