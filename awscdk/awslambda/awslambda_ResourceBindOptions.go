package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceBindOptions := &resourceBindOptions{
//   	resourceProperty: jsii.String("resourceProperty"),
//   }
//
type ResourceBindOptions struct {
	// The name of the CloudFormation property to annotate with asset metadata.
	// See: https://github.com/aws/aws-cdk/issues/1432
	//
	ResourceProperty *string `field:"optional" json:"resourceProperty" yaml:"resourceProperty"`
}

