package repository

import (
	"container/list"
	"encoding/json"
	"kpi/internal/entity"
	"sync"
)

type ListQueue struct {
	list  *list.List
	mutex *sync.Mutex
}

var queue *ListQueue

// GetQueue создает и возвращает простую очерель
func GetQueue() *ListQueue {
	if queue == nil {
		q := ListQueue{
			list:  list.New(),
			mutex: &sync.Mutex{},
		}
		queue = &q
	}

	return queue
}

// Add добавляет элемент очереди в конец
func (q *ListQueue) Add(v *[]byte) {
	q.mutex.Lock()
	q.list.PushBack(*v)
	q.mutex.Unlock()
}

// SaveFact получает первый элемент очереди, сохраняет его в бд, если при сохранении не возникло ошибок - удаляет элемент из очереди
func (q *ListQueue) SaveFact() error {
	q.mutex.Lock()
	if q.list.Len() > 0 {
		var fact entity.Fact
		e := q.list.Front() // Первый элемент
		err := json.Unmarshal(e.Value.([]byte), &fact)
		if err != nil {
			return err
		}

		db := GetDb()
		_, err = db.saveFact(&fact)
		if err != nil {
			return err
		}

		q.list.Remove(e) // Удаление из очереди, если успешно сохранили
	}
	q.mutex.Unlock()

	return nil
}
