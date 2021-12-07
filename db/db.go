package db

import (
	"log"

	. "github.com/cleslley/api-rest-go/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AntispamDB struct {
	Server   string
	Database string
}

var db *mgo.Database


const Collection = "antispam"

func (a *AntispamDB) Connect() {
	session, err := mgo.Dial(a.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(a.Database)

}

func (a *AntispamDb) FindById(id string) (Antispam, error) {
	var antispam Antispam
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&antispam)
	return antispam, err
}

func (a *AntispamDb) Insert(antispam Antispam) error {
	err := db.C(COLLECTION).Insert(&antispam)
	return err
}

func (a *AntispamDb) Delete(antispam Antispam) error {
	err := db.C(COLLECTION).Remove(&antispam)
	return err
}

func (a *AntispamDb) Update(antispam Antispam) error {
	err := db.C(COLLECTION).UpdateId(antispam.ID, &antispam)
	return err
}
