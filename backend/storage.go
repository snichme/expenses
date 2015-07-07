package main

import (
	"encoding/json"

	"gopkg.in/redis.v2"
)

type PaymentStorage interface {
	Get(string) (Payments, error)
	Set(string, Payments) (bool, error)
}

type RedisStorage struct {
	Client *redis.Client
}

func (rs *RedisStorage) Get(key string) (Payments, error) {
	var p Payments
	data, err := rs.Client.Get(key).Result()
	if err != nil {
		return Payments{}, err
	}
	err2 := json.Unmarshal([]byte(data), &p)
	if err2 != nil {
		return p, err2
	}
	return p, nil
}

func (rs *RedisStorage) Set(key string, val Payments) (bool, error) {
	data, err := json.Marshal(val)
	if err != nil {
		return false, err
	}

	if err := rs.Client.Set(key, string(data)).Err(); err != nil {
		return false, err
	}
	return true, nil
}

type FakeStorage struct{}

func (s *FakeStorage) Get(key string) (Payments, error) {
	p := []Payment{
		Payment{Name: "Mange", Amount: 20},
		Payment{Name: "Natta", Amount: 30.2},
		Payment{Name: "Mange", Amount: 20.5},
	}

	return Payments{
		Id:    key,
		Title: "Title" + key,
		Items: p}, nil
}
func (rs *FakeStorage) Set(key string, val Payments) (bool, error) {
	return true, nil
}
