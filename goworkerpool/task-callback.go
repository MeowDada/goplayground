package workerpool

type TaskCallbackReturn interface{}
type TaskCallbackInput  interface{}

type TaskCallback interface {
	OnStart(TaskCallbackInput) TaskCallbackReturn
	OnExit(TaskCallbackInput) TaskCallbackReturn
	OnSucceed(TaskCallbackInput) TaskCallbackReturn
	OnFail(TaskCallbackInput) TaskCallbackReturn
	OnTimeOut(TaskCallbackInput) TaskCallbackReturn
}

type CustomTaskCallback struct {
	OnStartFn   func(TaskCallbackInput) TaskCallbackReturn
	OnExitFn    func(TaskCallbackInput) TaskCallbackReturn
	OnSucceedFn func(TaskCallbackInput) TaskCallbackReturn
	OnFailFn    func(TaskCallbackInput) TaskCallbackReturn
	OnTimeOutFn func(TaskCallbackInput) TaskCallbackReturn
}

func (c *CustomTaskCallback) OnStart(input TaskCallbackInput) TaskCallbackReturn {
	if c.OnStartFn != nil {
		return c.OnStartFn(input)
	}
	return nil
}

func (c *CustomTaskCallback) OnExit(input TaskCallbackInput) TaskCallbackReturn {
	if c.OnExitFn != nil {
		return c.OnExitFn(input)
	}
	return nil
}

func (c *CustomTaskCallback) OnSucceed(input TaskCallbackInput) TaskCallbackReturn {
	if c.OnSucceedFn != nil {
		return c.OnSucceedFn(input)
	}
	return nil
}

func (c *CustomTaskCallback) OnFail(input TaskCallbackInput) TaskCallbackReturn {
	if c.OnFailFn != nil {
		return c.OnFailFn(input)
	}
	return nil
}

func (c *CustomTaskCallback) OnTimeOut(input TaskCallbackInput) TaskCallbackReturn {
	if c.OnTimeOutFn != nil {
		return c.OnTimeOutFn(input)
	}
	return nil
}