package awsiotevents


// Information that defines how a detector operates.
//
// Example:
//
//
type CfnDetectorModel_DetectorModelDefinitionProperty struct {
	// The state that is entered at the creation of each detector (instance).
	InitialStateName *string `field:"required" json:"initialStateName" yaml:"initialStateName"`
	// Information about the states of the detector.
	States interface{} `field:"required" json:"states" yaml:"states"`
}

