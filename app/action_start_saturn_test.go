package app

type StartSaturnAction struct {
}

func NewStartSaturnAction() IApplicationStartupAction {
	action := &StartSaturnAction{}
	return action
}

func (action *StartSaturnAction) Initialize(context *ApplicationContext) IApplicationStartupAction {
	return action
}

func (action *StartSaturnAction) Start() {

}

func (action *StartSaturnAction) Stop() {

}

func (action *StartSaturnAction) Priority() int {
	return 500
}
