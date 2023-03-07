package runner

import (
	"log"
	"os"
	"time"
)

// 最长处理时间
const timeout = 3 * time.Second

func main() {
	r := New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case ErrInterrupt:
			log.Println("Timeout!")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Interrupted!")
			os.Exit(2)
		}
	}

	log.Println("process success!")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Process - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
