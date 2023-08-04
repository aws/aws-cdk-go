package awsstepfunctions


// Options for finding reachable states.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   findStateOptions := &FindStateOptions{
//   	IncludeErrorHandlers: jsii.Boolean(false),
//   }
//
type FindStateOptions struct {
	// Whether or not to follow error-handling transitions.
	// Default: false.
	//
	IncludeErrorHandlers *bool `field:"optional" json:"includeErrorHandlers" yaml:"includeErrorHandlers"`
}

