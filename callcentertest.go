package main

import (
	"fmt"
	"time"
)

type Call struct {
	ID        int
	Duration  time.Duration
	StartTime time.Time
	EndTime   time.Time
}

type Operator struct {
	ID     int
	Name   string
	Free   bool
	Call   *Call
	CallCh chan *Call
	QuitCh chan bool
}

func (op *Operator) startHandlingCalls() {
	for {
		select {
		case call := <-op.CallCh:
			op.Free = false
			op.Call = call
			call.StartTime = time.Now()
			fmt.Printf("Operator %s начал обработку звонка %d\n", op.Name, call.ID)
			time.Sleep(call.Duration)
			call.EndTime = time.Now()
			op.Free = true
			op.Call = nil
			fmt.Printf("Оператор %s завершил обработку звонка %d\n", op.Name, call.ID)
		case <-op.QuitCh:
			return
		}
	}
}

func main() {
	callCenter := make(chan *Call)
	operators := make([]Operator, 5)
	for i := 0; i < 5; i++ {
		operators[i] = Operator{
			ID:     i + 1,
			Name:   fmt.Sprintf("Operator %d", i+1),
			Free:   true,
			CallCh: callCenter,
			QuitCh: make(chan bool),
		}
		go operators[i].startHandlingCalls()
	}

	// Логика добавления звонков в callCenter и распределения их по операторам

	// Пример работы программы: добавление звонков в callCenter и передача их операторам для обработки
	for j := 1; j <= 10; j++ {
		callCenter <- &Call{
			ID:       j,
			Duration: time.Duration((j%3)+1) * time.Second,
		}
	}

	// Пауза для продолжения работы программы
	time.Sleep(15 * time.Second)

	// Пример завершения работы программы: остановка обработки звонков и выход из приложения
	for i := 0; i < 5; i++ {
		operators[i].QuitCh <- true
	}
}
