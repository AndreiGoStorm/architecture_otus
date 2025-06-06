package hw3

var repeatTries = 2

type ExceptionHandler struct {
	queue *Queue
	store map[string]map[string]func(Command, error) Command
}

func NewExceptionHandler(q *Queue) *ExceptionHandler {
	h := &ExceptionHandler{
		queue: q,
		store: make(map[string]map[string]func(Command, error) Command, 10),
	}
	h.init()
	return h
}

func (h *ExceptionHandler) handle(cmd Command, err error) Command {
	ct := getTypeCmd(cmd)
	et := err.Error()
	value := h.getStoreValue(ct, et)

	return value(cmd, err)
}

func (h *ExceptionHandler) getStoreValue(ct, et string) func(Command, error) Command {
	return h.store[ct][et]
}

func (h *ExceptionHandler) init() {
	h.register("MoveCommand", "ErrMoveRepeatHandler", h.repeatHandler())
	h.register("MoveCommand", "ErrMoveLogHandler", h.logHandler())
	h.register("MoveCommand", "ErrMoveRepeatLogHandler", h.repeatLogHandler(&RepeatLogCommand{tries: repeatTries}))
	h.register("RotateCommand", "ErrRotateLogHandler", h.logHandler())
	h.register("RotateCommand", "ErrRotateRepeatLogHandler", h.repeatLogHandler(&RepeatLogCommand{tries: repeatTries - 1}))
}

func (h *ExceptionHandler) repeatHandler() func(Command, error) Command {
	return func(cmd Command, err error) Command {
		return &RepeatCommand{cmd, err, h.queue}
	}
}

func (h *ExceptionHandler) logHandler() func(Command, error) Command {
	return func(cmd Command, err error) Command {
		return &LogCommand{cmd, err}
	}
}

func (h *ExceptionHandler) repeatLogHandler(rlc *RepeatLogCommand) func(Command, error) Command {
	return func(cmd Command, err error) Command {
		rlc.addParams(cmd, err, h.queue)
		return rlc
	}
}

func (h *ExceptionHandler) register(ct, et string, handler func(Command, error) Command) {
	_, ok := h.store[ct]
	if !ok {
		h.store[ct] = make(map[string]func(Command, error) Command, 5)
	}
	h.store[ct][et] = handler
}
