package db

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var TaskBucket = []byte("PendingTasks")
var CompletedBucket = []byte("CompletedTasks")
var DateCompletion = []byte("DateCompletion")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

// Initializing Database and Bucket
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(TaskBucket)
		_ = err
		_, err = tx.CreateBucketIfNotExists(DateCompletion)
		_ = err
		_, err = tx.CreateBucketIfNotExists(CompletedBucket)
		return err
	})
}

func CreateTask(task string, bucket []byte) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		// Accessing the Pending Tasks bucket (table)
		b := tx.Bucket(bucket)

		// Creating unique key for the next element
		id64, _ := b.NextSequence()
		id = int(id64)

		key := itoB(int(id64))

		// Adding new element to the pending task list
		return b.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}
	return id, err
}

func AllTasks(bucket []byte) ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoI(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(k int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(TaskBucket)
		err := b.Delete(itoB(k))
		return err
	})
}

func MarkAsCompleted(k int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(DateCompletion)
		kByte := itoB(k)
		todayDate := TodaysDate()
		return b.Put(kByte, []byte(todayDate))
	})
}

func itoB(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoI(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func TodaysDate() string {
	today := time.Now()
	return today.Format("2006-01-02")
}
