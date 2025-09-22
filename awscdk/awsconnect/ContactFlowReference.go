package awsconnect


// A reference to a ContactFlow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactFlowReference := &ContactFlowReference{
//   	ContactFlowArn: jsii.String("contactFlowArn"),
//   }
//
type ContactFlowReference struct {
	// The ContactFlowArn of the ContactFlow resource.
	ContactFlowArn *string `field:"required" json:"contactFlowArn" yaml:"contactFlowArn"`
}

