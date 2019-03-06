package db

import (
	"errors"
	"github.com/dy-dayan/common-srv-atomicid/util/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	defaultMgo *mgo.Session
)

func Mgo() *mgo.Session {
	return defaultMgo
}

func Init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     uconfig.DefaultMgoConf.Addr,
		Direct:    false,
		Timeout:   time.Second * 3,
		PoolLimit: uconfig.DefaultMgoConf.PoolLimit,
		Username:  uconfig.DefaultMgoConf.Username,
		Password:  uconfig.DefaultMgoConf.Password,
	}

	ses, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logrus.Fatalf("dail mgo server error:%v", err)
	}

	ses.SetMode(mgo.Monotonic, true)
	defaultMgo = ses
}

type ID struct {
	ID        string  `bson:"_id"`
	Value     int64 `bson:"value"`
}

var (
	DBCommonID = "dayan_common_id"
	CAtomicID  = "atomic_id"
)

func GetID(label string) (id int64, err error) {

	ses := defaultMgo.Copy()
	if ses == nil {
		return 0, errors.New("mgo session is nil")
	}
	defer ses.Close()

	query := bson.M{
		"_id": label,
	}

	ret := &ID{}
	change := mgo.Change{
		Update: bson.M{
			"$inc":bson.M{
				"value":1,
			},
		},
		Upsert:    true,
		Remove:    false,
		ReturnNew: true,
	}

	_, err = ses.DB(DBCommonID).C(CAtomicID).Find(query).Apply(change, ret)
	if err != nil {
		return 0, err
	}

	return ret.Value, nil
}