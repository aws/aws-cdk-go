package awsiotevents


// Information that defines a state of a detector.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html
//
type CfnDetectorModel_StateProperty struct {
	// The name of the state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-statename
	//
	StateName *string `field:"required" json:"stateName" yaml:"stateName"`
	// When entering this state, perform these `actions` if the `condition` is TRUE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-onenter
	//
	OnEnter interface{} `field:"optional" json:"onEnter" yaml:"onEnter"`
	// When exiting this state, perform these `actions` if the specified `condition` is `TRUE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-onexit
	//
	OnExit interface{} `field:"optional" json:"onExit" yaml:"onExit"`
	// When an input is received and the `condition` is TRUE, perform the specified `actions` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-oninput
	//
	OnInput interface{} `field:"optional" json:"onInput" yaml:"onInput"`
}

