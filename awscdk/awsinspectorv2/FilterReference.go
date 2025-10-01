package awsinspectorv2


// A reference to a Filter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterReference := &FilterReference{
//   	FilterArn: jsii.String("filterArn"),
//   }
//
type FilterReference struct {
	// The Arn of the Filter resource.
	FilterArn *string `field:"required" json:"filterArn" yaml:"filterArn"`
}

