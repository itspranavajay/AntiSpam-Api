package type

import "gopkg.in/mgo.v2/bson"

type Antispam struct {
	ID          string        `bson:"id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Reason      string        `bson:"reason" json:"reason`
}
