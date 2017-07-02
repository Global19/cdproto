package animation

// AUTOGENERATED. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
)

// EventAnimationCreated event for each animation that has been created.
type EventAnimationCreated struct {
	ID string `json:"id"` // Id of the animation that was created.
}

// EventAnimationStarted event for animation that has been started.
type EventAnimationStarted struct {
	Animation *Animation `json:"animation"` // Animation that was started.
}

// EventAnimationCanceled event for when an animation has been cancelled.
type EventAnimationCanceled struct {
	ID string `json:"id"` // Id of the animation that was cancelled.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventAnimationAnimationCreated,
	cdp.EventAnimationAnimationStarted,
	cdp.EventAnimationAnimationCanceled,
}
