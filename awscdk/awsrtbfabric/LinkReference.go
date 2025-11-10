package awsrtbfabric


// A reference to a Link resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linkReference := &LinkReference{
//   	LinkArn: jsii.String("linkArn"),
//   }
//
type LinkReference struct {
	// The Arn of the Link resource.
	LinkArn *string `field:"required" json:"linkArn" yaml:"linkArn"`
}

