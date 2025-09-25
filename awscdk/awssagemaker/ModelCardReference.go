package awssagemaker


// A reference to a ModelCard resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelCardReference := &ModelCardReference{
//   	ModelCardArn: jsii.String("modelCardArn"),
//   	ModelCardName: jsii.String("modelCardName"),
//   }
//
type ModelCardReference struct {
	// The ARN of the ModelCard resource.
	ModelCardArn *string `field:"required" json:"modelCardArn" yaml:"modelCardArn"`
	// The ModelCardName of the ModelCard resource.
	ModelCardName *string `field:"required" json:"modelCardName" yaml:"modelCardName"`
}

