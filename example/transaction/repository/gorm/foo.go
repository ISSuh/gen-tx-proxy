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

package infra

import (
	"context"
	"errors"

	"github.com/ISSuh/gen-go-proxy/example/transaction/entity"
	"gorm.io/gorm"
)

type FooGORMRepository struct {
	db *gorm.DB
}

func NewFooGORMRepository(db *gorm.DB) *FooGORMRepository {
	return &FooGORMRepository{
		db: db,
	}
}

func (r *FooGORMRepository) Create(c context.Context, value int) (int, error) {
	conn, err := FromContext(c)
	if err != nil {
		return 0, err
	}

	if value < 0 {
		return 0, errors.New("value must be greater than 0")
	}

	f := &entity.Foo{
		Value: int(value),
	}

	tx := conn.Create(f)
	if err := tx.Error; err != nil {
		return 0, err
	}

	return f.ID, nil
}

func (r *FooGORMRepository) Find(id int) (*entity.Foo, error) {
	f := &entity.Foo{}
	tx := r.db.Where("id = ?", id)
	if err := tx.First(f).Error; err != nil {
		return nil, err
	}
	return f, nil
}
