package main

import (
	"os"
)

type pgnum uint64
type initialPage uint

type page struct {
	num  pgnum
	data []byte
}

type dal struct {
	file     *os.File
	pageSize int
}

type freelist struct {
	maxPage       pgnum
	releasedPages []pgnum
}

func newDal(path string, pageSize int) (*dal, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &dal{file, pageSize}, nil
}

func (d *dal) allocateEmptyPage(pageSize int) *page {
	return &page{data: make([]byte, d.pageSize)}
}

func (d *dal) readPage(pagenum int) (*page, error) {
	p := d.allocateEmptyPage(d.pageSize)
	//Find correct offset in the file to read from
	offset := pagenum * d.pageSize
	_, err := d.file.ReadAt(p.data, int64(offset))
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (d *dal) writePage(p *page) error {
	offset := p.num * pgnum(d.pageSize)
	_, err := d.file.WriteAt(p.data, int64(offset))
	return error
}

func(d *dal) newFreelist() *freelist {
	return &freelist{
		maxPage: initialPage
		releasedPages: [],
	}
}
