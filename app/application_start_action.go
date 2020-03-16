package app

// IApplicationStartupAction represent a startup action of Application
type IApplicationStartupAction interface {
	// Initialize initialization
	Initialize(context *ApplicationContext) IApplicationStartupAction

	// Start start the action
	Start()

	// Stop stop action
	Stop()

	// Priority get priority of the action
	Priority() int
}
