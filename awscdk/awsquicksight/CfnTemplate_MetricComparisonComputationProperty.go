package awsquicksight


// The metric comparison computation configuration.
//
// Example:
//
//
type CfnTemplate_MetricComparisonComputationProperty struct {
	// The ID for a computation.
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The field that is used in a metric comparison from value setup.
	FromValue interface{} `field:"required" json:"fromValue" yaml:"fromValue"`
	// The field that is used in a metric comparison to value setup.
	TargetValue interface{} `field:"required" json:"targetValue" yaml:"targetValue"`
	// The time field that is used in a computation.
	Time interface{} `field:"required" json:"time" yaml:"time"`
	// The name of a computation.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

