package awsquicksight


// The configuration of an insight visual.
//
// Example:
//
//
type CfnTemplate_InsightConfigurationProperty struct {
	// The computations configurations of the insight visual.
	Computations interface{} `field:"optional" json:"computations" yaml:"computations"`
	// The custom narrative of the insight visual.
	CustomNarrative interface{} `field:"optional" json:"customNarrative" yaml:"customNarrative"`
}

