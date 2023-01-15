package main

import (
	"context"
	"time"
)

const (
	waitDur    = 1 * time.Second
	cancelDur  = 2000 * time.Millisecond
	timeoutDur = 500 * time.Millisecond
)

type Config struct {
	SelectTimeout time.Duration
}

type DB struct {
	cfg Config
}

type User struct {
	Name string
}

func (d *DB) SelectUser(ctx context.Context, email string) (User, error) {
	ctx2, cancel := context.WithTimeout(ctx, timeoutDur)
	//* 1 — допишите здесь создание контекста с тайм-аутом */;
	defer cancel()

	timer := time.NewTimer(waitDur)
	select {
	case <-timer.C:
		return User{Name: "Gosha"}, nil
	case <-ctx2.Done():
		return User{}, ctx2.Err()
	}
}

type Handler struct {
	db *DB
}

type Request struct {
	Email string
}
