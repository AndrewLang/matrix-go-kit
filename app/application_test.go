package matrix

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewApplication(t *testing.T) {
	app := NewApplication()

	assert.NotNil(t, app)
	assert.NotNil(t, app.logger)
	assert.NotNil(t, app.Context)

}

func TestApplicationInitialize(t *testing.T) {
	app := NewApplication()

	assert.NotNil(t, app)

	app.Initialize()
}

func TestApplicationStart(t *testing.T) {
	app := NewApplication()

	assert.NotNil(t, app)
	app.Start()
	app.Stop()
}

func TestApplicationStop(t *testing.T) {
	app := NewApplication()

	assert.NotNil(t, app)
	app.Stop()
}

func TestApplicationUseAction(t *testing.T) {
	app := NewApplication()

	assert.Equal(t, 0, len(app.Actions))

	app.Use(NewStartSaturnAction())

	assert.Equal(t, 1, len(app.Actions))

	app.Use(NewStartJupiterAction())

	assert.Equal(t, 2, len(app.Actions))
}

func TestApplicationSortActions(t *testing.T) {
	app := NewApplication().
		Use(NewStartSaturnAction()).
		Use(NewStartJupiterAction()).
		Initialize()

	assert.Equal(t, 2, len(app.Actions))

	assert.Equal(t, 200, app.Actions[0].Priority())
	assert.Equal(t, 500, app.Actions[1].Priority())
}

func TestApplicationWait(t *testing.T) {
	app := NewApplication().Initialize()
	app.allowWait = false

	go func() {
		time.Sleep(3 * time.Second)
		app.Stop()
	}()
	app.Start().Wait()
}
