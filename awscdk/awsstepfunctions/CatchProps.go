package awsstepfunctions


// Error handler details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   catchProps := &CatchProps{
//   	Errors: []*string{
//   		jsii.String("errors"),
//   	},
//   	ResultPath: jsii.String("resultPath"),
//   }
//
type CatchProps struct {
	// Errors to recover from by going to the given state.
	//
	// A list of error strings to retry, which can be either predefined errors
	// (for example Errors.NoChoiceMatched) or a self-defined error.
	// Default: All errors.
	//
	Errors *[]*string `field:"optional" json:"errors" yaml:"errors"`
	// JSONPath expression to indicate where to inject the error data.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the error
	// data to be discarded.
	// Default: $.
	//
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
}

