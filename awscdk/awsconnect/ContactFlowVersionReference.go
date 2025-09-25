package awsconnect


// A reference to a ContactFlowVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactFlowVersionReference := &ContactFlowVersionReference{
//   	ContactFlowVersionArn: jsii.String("contactFlowVersionArn"),
//   }
//
type ContactFlowVersionReference struct {
	// The ContactFlowVersionARN of the ContactFlowVersion resource.
	ContactFlowVersionArn *string `field:"required" json:"contactFlowVersionArn" yaml:"contactFlowVersionArn"`
}

