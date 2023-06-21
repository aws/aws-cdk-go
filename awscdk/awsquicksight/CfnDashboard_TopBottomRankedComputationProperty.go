package awsquicksight


// The top ranked and bottom ranked computation configuration.
//
// Example:
//
//
type CfnDashboard_TopBottomRankedComputationProperty struct {
	// The category field that is used in a computation.
	Category interface{} `field:"required" json:"category" yaml:"category"`
	// The ID for a computation.
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The computation type. Choose one of the following options:.
	//
	// - TOP: A top ranked computation.
	// - BOTTOM: A bottom ranked computation.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of a computation.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The result size of a top and bottom ranked computation.
	ResultSize *float64 `field:"optional" json:"resultSize" yaml:"resultSize"`
	// The value field that is used in a computation.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

