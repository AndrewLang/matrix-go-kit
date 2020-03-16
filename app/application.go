package app

import (
	"reflect"
	"sort"
	"sync"

	log "github.com/andrewlang/matrix-go-kit/log"
)

// Application application
type Application struct {
	Context    *ApplicationContext
	Actions    []IApplicationStartupAction
	logger     log.ILogger
	exitSignal *sync.WaitGroup
	allowWait  bool
}

// NewApplication create a new instance of application
func NewApplication() *Application {
	app := &Application{
		Context:    NewApplicationContext(),
		Actions:    make([]IApplicationStartupAction, 0),
		logger:     CreateLogger("Application"),
		exitSignal: &sync.WaitGroup{},
		allowWait:  true,
	}
	app.exitSignal.Add(1)

	return app
}

// Use add a startup action
func (app *Application) Use(action IApplicationStartupAction) *Application {
	app.logger.Info(`Use action:`, reflect.TypeOf(action))
	app.Actions = append(app.Actions, action)
	return app
}

// Initialize application
func (app *Application) Initialize() *Application {
	app.sortActions()

	for _, action := range app.Actions {
		action.Initialize(app.Context)
	}
	return app
}

// Start start application
func (app *Application) Start() *Application {
	app.logger.
		Info("Start application").
		Info("Configuration:", app.Context.Configuration)

	for _, action := range app.Actions {
		action.Start()
	}

	return app
}

// Stop stop/exit application
func (app *Application) Stop() *Application {
	for _, action := range app.Actions {
		action.Stop()
	}

	if !app.allowWait {
		app.exitSignal.Done()
	}

	return app
}

// Wait hang and waiting for message loop
func (app *Application) Wait() *Application {
	app.exitSignal.Wait()
	return app
}

// sortActions by priority
func (app *Application) sortActions() {
	sort.Slice(app.Actions, func(i, j int) bool {
		return app.Actions[i].Priority() < app.Actions[j].Priority()
	})
}
