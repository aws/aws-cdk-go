package awsiotevents


// Information that defines how a detector operates.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-detectormodeldefinition.html
//
type CfnDetectorModel_DetectorModelDefinitionProperty struct {
	// The state that is entered at the creation of each detector (instance).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-detectormodeldefinition.html#cfn-iotevents-detectormodel-detectormodeldefinition-initialstatename
	//
	InitialStateName *string `field:"required" json:"initialStateName" yaml:"initialStateName"`
	// Information about the states of the detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-detectormodeldefinition.html#cfn-iotevents-detectormodel-detectormodeldefinition-states
	//
	States interface{} `field:"required" json:"states" yaml:"states"`
}

