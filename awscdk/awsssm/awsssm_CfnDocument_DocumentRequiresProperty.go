package awsssm


// An SSM document required by the current document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentRequiresProperty := &documentRequiresProperty{
//   	name: jsii.String("name"),
//   	version: jsii.String("version"),
//   }
//
type CfnDocument_DocumentRequiresProperty struct {
	// The name of the required SSM document.
	//
	// The name can be an Amazon Resource Name (ARN).
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The document version required by the current document.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

