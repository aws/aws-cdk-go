package awscdk


// A reference to a ResourceDefaultVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDefaultVersionReference := &ResourceDefaultVersionReference{
//   	ResourceDefaultVersionArn: jsii.String("resourceDefaultVersionArn"),
//   }
//
type ResourceDefaultVersionReference struct {
	// The Arn of the ResourceDefaultVersion resource.
	ResourceDefaultVersionArn *string `field:"required" json:"resourceDefaultVersionArn" yaml:"resourceDefaultVersionArn"`
}

