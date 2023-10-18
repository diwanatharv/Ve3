package service

import (
	Mongo "awesomeProject/pkg/dataaccess/mongo"
	"awesomeProject/pkg/dataaccess/redis"
	"awesomeProject/pkg/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

func Createtask(task models.Task) error {
	manager := Mongo.MongoManager()

	idcount, err := manager.Totalcount(context.Background())
	if err != nil {
		log.Error(err.Error())
		return err
	}
	idcount++
	task.Id = int(idcount)
	_, err = manager.Insert(context.Background(), task)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
func Getalltask() ([]models.Task, error) {
	var res []models.Task
	filter := bson.M{}
	manager := Mongo.MongoManager()
	cur, err := manager.Findusers(context.Background(), filter)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	for cur.Next(context.Background()) {
		var temp models.Task
		err := cur.Decode(&temp)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		res = append(res, temp)
	}
	return res, nil
}
func Updatetask(user models.Task, Id int) error {
	manager := Mongo.MongoManager()
	redimanager := redis.Redismanager()

	str := strconv.Itoa(Id)
	filter := bson.M{"Id": Id}
	update := bson.M{"$set": bson.M{}}
	opts := options.Update().SetUpsert(true)
	_, err := manager.Updateone(context.Background(), filter, update, opts)
	asli, err := json.Marshal(&user)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err = redimanager.Setredis(context.Background(), str, asli, time.Second*10*500).Err()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func Gettask(id int) (models.Task, error, error) {
	var user models.Task
	manager := Mongo.MongoManager()
	redismanager := redis.Redismanager()
	str := strconv.Itoa(id)
	ans, err := redismanager.Getredis(context.Background(), str).Result()
	if err == nil {
		fmt.Println("yei value redis sei aaya hai")
		err = json.Unmarshal([]byte(ans), &user)
		if err != nil {
			log.Error(err.Error())
			return user, nil, err
		}
		return user, nil, nil
	} else {
		fmt.Println("yei value humlong ko mila hai mongosei")
		filter := bson.M{"Id": id}

		err := manager.Findone(context.Background(), filter).Decode(&user)
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Error(err)
			return user, errors.New("no  such task exists"), nil
		}
		asli, err := json.Marshal(&user)
		err = redismanager.Setredis(context.Background(), str, asli, time.Second*10*500).Err()
		if err != nil {
			log.Error(err.Error())
		}
		return user, nil, nil

	}
}
func Deletetask(id int) (error, error) {
	manager := Mongo.MongoManager()
	redismanager := redis.Redismanager()
	str := strconv.Itoa(id)
	err := redismanager.Getredis(context.Background(), str).Err()
	if err != nil {
		log.Error(err.Error())
	}
	if err == nil {
		err = redismanager.Deletekey(context.Background(), str).Err()
		if err != nil {
			log.Error(err.Error())
		}
	}
	filter := bson.M{"Id": id}
	val := manager.Findanddelete(context.Background(), filter)

	if val == nil {
		fmt.Println(err.Error())
		return err, err
	}
	return nil, nil
}
