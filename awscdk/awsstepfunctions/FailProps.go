package awsstepfunctions


// Properties for defining a Fail state.
//
// Example:
//   fail := sfn.NewFail(this, jsii.String("Fail"), &FailProps{
//   	ErrorPath: sfn.JsonPath_StringAt(jsii.String("$.someError")),
//   	CausePath: sfn.JsonPath_*StringAt(jsii.String("$.someCause")),
//   })
//
type FailProps struct {
	// A description for the cause of the failure.
	// Default: - No description.
	//
	Cause *string `field:"optional" json:"cause" yaml:"cause"`
	// JsonPath expression to select part of the state to be the cause to this state.
	// Default: - No cause path.
	//
	CausePath *string `field:"optional" json:"causePath" yaml:"causePath"`
	// An optional description for this state.
	// Default: - No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Error code used to represent this failure.
	// Default: - No error code.
	//
	Error *string `field:"optional" json:"error" yaml:"error"`
	// JsonPath expression to select part of the state to be the error to this state.
	// Default: - No error path.
	//
	ErrorPath *string `field:"optional" json:"errorPath" yaml:"errorPath"`
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
}

