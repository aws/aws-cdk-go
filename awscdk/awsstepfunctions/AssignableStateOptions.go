package awsstepfunctions


// Option properties for state that can assign variables.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//
//   assignableStateOptions := &AssignableStateOptions{
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   }
//
type AssignableStateOptions struct {
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
}

