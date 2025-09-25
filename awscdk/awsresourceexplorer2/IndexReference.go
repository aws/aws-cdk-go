package awsresourceexplorer2


// A reference to a Index resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   indexReference := &IndexReference{
//   	IndexArn: jsii.String("indexArn"),
//   }
//
type IndexReference struct {
	// The Arn of the Index resource.
	IndexArn *string `field:"required" json:"indexArn" yaml:"indexArn"`
}

