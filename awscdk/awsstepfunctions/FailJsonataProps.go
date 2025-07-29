package awsstepfunctions


// Properties for defining a Fail state that using JSONata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failJsonataProps := &FailJsonataProps{
//   	Cause: jsii.String("cause"),
//   	Comment: jsii.String("comment"),
//   	Error: jsii.String("error"),
//   	QueryLanguage: awscdk.Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	StateName: jsii.String("stateName"),
//   }
//
type FailJsonataProps struct {
	// A comment describing this state.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The name of the query language used by the state.
	//
	// If the state does not contain a `queryLanguage` field,
	// then it will use the query language specified in the top-level `queryLanguage` field.
	// Default: - JSONPath.
	//
	QueryLanguage QueryLanguage `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
	// A description for the cause of the failure.
	// Default: - No description.
	//
	Cause *string `field:"optional" json:"cause" yaml:"cause"`
	// Error code used to represent this failure.
	// Default: - No error code.
	//
	Error *string `field:"optional" json:"error" yaml:"error"`
}

