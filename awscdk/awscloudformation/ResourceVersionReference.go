package awscloudformation


// A reference to a ResourceVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceVersionReference := &ResourceVersionReference{
//   	ResourceVersionArn: jsii.String("resourceVersionArn"),
//   }
//
type ResourceVersionReference struct {
	// The Arn of the ResourceVersion resource.
	ResourceVersionArn *string `field:"required" json:"resourceVersionArn" yaml:"resourceVersionArn"`
}

