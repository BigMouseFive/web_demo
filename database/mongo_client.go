package database

import (
	"context"
	"github.com/web_demo/v2/config"
	"github.com/web_demo/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var mctx context.Context
var cancel context.CancelFunc

var (
	MongoCli      *mongo.Client
	MongoCollList map[string]*mongo.Collection
)

func CreateMongo() {
	// 创建mongo客户端
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	log.Sugar.Infow("CreateMongo ", "mongo_address", config.Config.GetString("mongo.address"))
	var err error
	credential := options.Credential{
		AuthSource: config.Config.GetString("mongo.db_name"),
		Username:   config.Config.GetString("mongo.username"),
		Password:   config.Config.GetString("mongo.password"),
	}
	MongoCli, err = mongo.Connect(mctx, options.Client().ApplyURI(config.Config.GetString("mongo.address")).SetAuth(credential))
	if err != nil {
		log.Sugar.Fatalw("Can't connect to mongo server", "error", err.Error())
	}

	// Ping mongo server
	if err = MongoCli.Ping(mctx, readpref.Primary()); err != nil {
		log.Sugar.Fatalw("Mongo Ping error", "error", err.Error())
	}
	MongoCollList = make(map[string]*mongo.Collection)

	// 创建约束唯一值约束
	CreateIndex("area", "id", true, false)
	CreateIndex("auth", "id", true, false)
	CreateIndex("corporation", "id", true, false)
	CreateIndex("hd_map", "id", true, false)
	CreateIndex("device", "id", true, false)
	CreateIndex("instance", "id", true, false)
	CreateIndex("map", "id", true, false)
	CreateIndex("project", "id", true, false)
	CreateIndex("role", "id", true, false)
	CreateIndex("slam_map", "id", true, false)
	CreateIndex("user", "id", true, false)

	CreateTextIndex("area", []string{"id", "name"})
	CreateTextIndex("auth", []string{"id", "name"})
	CreateTextIndex("corporation", []string{"id", "name"})
	CreateTextIndex("hd_map", []string{"id", "name"})
	CreateTextIndex("device", []string{"id", "name"})
	CreateTextIndex("instance", []string{"id", "name"})
	CreateTextIndex("map", []string{"id", "name"})
	CreateTextIndex("project", []string{"id", "name"})
	CreateTextIndex("role", []string{"id", "name"})
	CreateTextIndex("slam_map", []string{"id", "name"})
	CreateTextIndex("user", []string{"id", "name"})
}

// CreateIndex - creates an index for a specific field in a collection
// ref: https://christiangiacomi.com/posts/mongodb-index-using-go/
func CreateIndex(collectionName string, field string, unique bool, sparse bool) bool {
	createFlag := true
	// 1. Lets define the keys for the index we want to create
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique).SetSparse(sparse),
	}

	// 2. Create the context for this operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3. Create a single index
	_, err := Coll(collectionName).Indexes().CreateOne(ctx, mod)
	if err != nil {
		// 4. Something went wrong, we log it and return false
		log.Sugar.Warn(err.Error())
		createFlag = false
	} else {
		// 5. All went well, we return true
		log.Sugar.Infow(collectionName+" create index "+field+" success",
			"unique", unique, "sparse", sparse)
	}

	return createFlag
}

// CreateTextIndex - creates an text index for a specific field in a collection
// ref: https://medium.com/easyread/today-i-learn-text-search-on-mongodb-6b87cd8497c9
func CreateTextIndex(collectionName string, fieldList []string) bool {
	createFlag := true

	keys := bson.D{}
	for _, v := range fieldList {
		keys = append(keys, bson.E{Key: v, Value: "text"})
	}
	_, err := Coll(collectionName).Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: keys,
	})
	if err != nil {
		log.Sugar.Warn(err.Error())
		createFlag = false
	} else {
		log.Sugar.Infof(collectionName+" create text index success. %v", fieldList)
	}

	return createFlag
}

func Coll(collName string) *mongo.Collection {
	coll, ok := MongoCollList[collName]
	if !ok {
		MongoCollList[collName] = MongoCli.Database(config.Config.GetString("mongo.db_name")).Collection(collName)
		coll = MongoCollList[collName]
	}
	return coll
}
