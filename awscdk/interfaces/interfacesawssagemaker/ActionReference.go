package interfacesawssagemaker


// A reference to a Action resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionReference := &ActionReference{
//   	ActionArn: jsii.String("actionArn"),
//   }
//
type ActionReference struct {
	// The Arn of the Action resource.
	ActionArn *string `field:"required" json:"actionArn" yaml:"actionArn"`
}

