package awsmedialive


// A reference to a Input resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputReference := &InputReference{
//   	InputArn: jsii.String("inputArn"),
//   	InputId: jsii.String("inputId"),
//   }
//
type InputReference struct {
	// The ARN of the Input resource.
	InputArn *string `field:"required" json:"inputArn" yaml:"inputArn"`
	// The Id of the Input resource.
	InputId *string `field:"required" json:"inputId" yaml:"inputId"`
}

