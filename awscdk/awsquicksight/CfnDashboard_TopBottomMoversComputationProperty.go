package awsquicksight


// The top movers and bottom movers computation setup.
//
// Example:
//
//
type CfnDashboard_TopBottomMoversComputationProperty struct {
	// The category field that is used in a computation.
	Category interface{} `field:"required" json:"category" yaml:"category"`
	// The ID for a computation.
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The time field that is used in a computation.
	Time interface{} `field:"required" json:"time" yaml:"time"`
	// The computation type. Choose from the following options:.
	//
	// - TOP: Top movers computation.
	// - BOTTOM: Bottom movers computation.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The mover size setup of the top and bottom movers computation.
	MoverSize *float64 `field:"optional" json:"moverSize" yaml:"moverSize"`
	// The name of a computation.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sort order setup of the top and bottom movers computation.
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
	// The value field that is used in a computation.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

