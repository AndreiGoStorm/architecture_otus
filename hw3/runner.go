package hw3

import "reflect"

type Runner struct {
	q  *Queue
	eh *ExceptionHandler
}

func NewRunner() *Runner {
	q := NewQueue()
	return &Runner{q, NewExceptionHandler(q)}
}

func (r *Runner) run() {
	for {
		elem := r.q.Dequeue()
		if elem == nil {
			break
		}

		cmd := elem.(Command)
		r.execute(cmd)
	}
}

func (r *Runner) execute(cmd Command) {
	err := cmd.execute()
	if err != nil {
		cmd = r.eh.handle(cmd, err)
		if cmd.execute() != nil {
			r.q.Enqueue(cmd)
		}
	}
}

func getTypeCmd(cmd Command) string {
	return reflect.TypeOf(cmd).Elem().Name()
}
