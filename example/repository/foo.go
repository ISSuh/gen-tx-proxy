﻿// MIT License

// Copyright (c) 2025 ISSuh

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package repository

import (
	"github.com/ISSuh/simple-gen-proxy/example/entity"
	"gorm.io/gorm"
)

type Foo interface {
	Create(value int64) error
	Find(id int) (*entity.Foo, error)
}

type fooRepository struct {
	db *gorm.DB
}

func NewFooRepository(db *gorm.DB) *fooRepository {
	return &fooRepository{
		db: db,
	}
}

func (r *fooRepository) Create(value int64) error {
	f := &entity.Foo{
		Value: int(value),
	}

	if err := r.db.Create(f).Error; err != nil {
		return err
	}
	return nil
}

func (r *fooRepository) Find(id int) (*entity.Foo, error) {
	f := &entity.Foo{}
	if err := r.db.Where("id = ?", id).First(f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

func (r *fooRepository) RunRX(f func() error) {
	db, err := r.db.DB()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	if err := f(); err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}
