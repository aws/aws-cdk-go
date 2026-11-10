package interfacesawssecurityhub


// A reference to a ActionTarget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionTargetReference := &ActionTargetReference{
//   	ActionTargetArn: jsii.String("actionTargetArn"),
//   }
//
type ActionTargetReference struct {
	// The ActionTargetArn of the ActionTarget resource.
	ActionTargetArn *string `field:"required" json:"actionTargetArn" yaml:"actionTargetArn"`
}

