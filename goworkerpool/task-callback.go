package workerpool

type TaskCallbackInput  interface{}

type TaskCallback interface {
	OnStart(TaskCallbackInput)
	OnExit(TaskCallbackInput)
	OnSucceed(TaskCallbackInput)
	OnFail(TaskCallbackInput)
	OnTimeout(TaskCallbackInput)
}

type CustomTaskCallback struct {
	OnStartFn   func(TaskCallbackInput)
	OnExitFn    func(TaskCallbackInput)
	OnSucceedFn func(TaskCallbackInput)
	OnFailFn    func(TaskCallbackInput)
	OnTimeuutFn func(TaskCallbackInput)
}

func (c *CustomTaskCallback) OnStart(input TaskCallbackInput) {
	if c.OnStartFn != nil {
		c.OnStartFn(input)
	}
}

func (c *CustomTaskCallback) OnExit(input TaskCallbackInput) {
	if c.OnExitFn != nil {
		c.OnExitFn(input)
	}
}

func (c *CustomTaskCallback) OnSucceed(input TaskCallbackInput) {
	if c.OnSucceedFn != nil {
		c.OnSucceedFn(input)
	}
}

func (c *CustomTaskCallback) OnFail(input TaskCallbackInput) {
	if c.OnFailFn != nil {
		c.OnFailFn(input)
	}
}

func (c *CustomTaskCallback) OnTimeuut(input TaskCallbackInput) {
	if c.OnTimeuutFn != nil {
		c.OnTimeuutFn(input)
	}
}