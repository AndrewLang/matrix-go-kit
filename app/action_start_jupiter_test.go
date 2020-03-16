package matrix

type StartJupiterAction struct {
}

func NewStartJupiterAction() IApplicationStartupAction {
	action := &StartJupiterAction{}
	return action
}

func (action *StartJupiterAction) Initialize(context *ApplicationContext) IApplicationStartupAction {
	return action
}

func (action *StartJupiterAction) Start() {

}

func (action *StartJupiterAction) Stop() {

}

func (action *StartJupiterAction) Priority() int {
	return 200
}
