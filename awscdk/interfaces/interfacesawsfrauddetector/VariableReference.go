package interfacesawsfrauddetector


// A reference to a Variable resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variableReference := &VariableReference{
//   	VariableArn: jsii.String("variableArn"),
//   }
//
type VariableReference struct {
	// The Arn of the Variable resource.
	VariableArn *string `field:"required" json:"variableArn" yaml:"variableArn"`
}

