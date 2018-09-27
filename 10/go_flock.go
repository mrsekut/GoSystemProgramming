package main

import (
	"fmt"
	"sync"
	"syscall"
	"time"
)

type FileLock struct {
	l  sync.Mutex
	fd int
}

func NewFileLock(filename string) *FileLock {
	if filename == "" {
		panic("filename needed")
	}
	fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDONLY, 0750)
	if err != nil {
		panic(err)
	}
	return &FileLock{fd: fd}
}

func (m *FileLock) Lock() {
	m.l.Lock()
	// LOCK_EX: 排他ロック
	if err := syscall.Flock(m.fd, syscall.LOCK_SH); err != nil {
		panic(err)
	}
}

func (m *FileLock) Unlock() {
	// LOCK_UN: ロック解除
	if err := syscall.Flock(m.fd, syscall.LOCK_UN); err != nil {
		panic(err)
	}
}

func main() {
	l := NewFileLock("10.2.1.go")
	fmt.Println("try locking...")
	l.Lock()
	fmt.Println("locked!")
	time.Sleep(10 * time.Second)
	l.Unlock()
	fmt.Println("unlock")
}

// LOCK_SH: 共有ロック
// LOCK_NB: ノンブロッキングモード
