package main

import (
	"fmt"

	"push-notice/helper"
	"push-notice/model"

	"github.com/garyburd/redigo/redis"
)

type Category struct {
	Id       int    `db:"id"`
	Type     string `db:"type"`
	NickName string `db:"nickname"`
}

func main() {
	// 获取任务列表
	// 循环消费任务
	// 根据任务key和类型适配各个消费
	//
	fmt.Println("aaa")
	var category []Category
	err := helper.Db.Select(&category, "select id,type,nickname from fa_category")
	res, _ := model.GetEventList()
	fmt.Println("res ", res)
	if err != nil {
		fmt.Println("fail", err)
	}
	fmt.Println("success", category)

	redispool := helper.Pool.Get()
	defer redispool.Close()
	_, errs := redispool.Do("Set", "abc", 100)
	if errs != nil {
		fmt.Println(err)
		return
	}
	r, errs := redis.Int(redispool.Do("Get", "abc"))
	if errs != nil {
		fmt.Println("get abc faild :", err)
		return
	}
	fmt.Println(r)
	helper.Pool.Close()
}
