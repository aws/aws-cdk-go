package awsiotevents


// Information that defines a state of a detector.
//
// Example:
//
//
type CfnDetectorModel_StateProperty struct {
	// The name of the state.
	StateName *string `field:"required" json:"stateName" yaml:"stateName"`
	// When entering this state, perform these `actions` if the `condition` is TRUE.
	OnEnter interface{} `field:"optional" json:"onEnter" yaml:"onEnter"`
	// When exiting this state, perform these `actions` if the specified `condition` is `TRUE` .
	OnExit interface{} `field:"optional" json:"onExit" yaml:"onExit"`
	// When an input is received and the `condition` is TRUE, perform the specified `actions` .
	OnInput interface{} `field:"optional" json:"onInput" yaml:"onInput"`
}

