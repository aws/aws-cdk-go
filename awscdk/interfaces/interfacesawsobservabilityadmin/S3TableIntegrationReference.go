package interfacesawsobservabilityadmin


// A reference to a S3TableIntegration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3TableIntegrationReference := &S3TableIntegrationReference{
//   	S3TableIntegrationArn: jsii.String("s3TableIntegrationArn"),
//   }
//
type S3TableIntegrationReference struct {
	// The Arn of the S3TableIntegration resource.
	S3TableIntegrationArn *string `field:"required" json:"s3TableIntegrationArn" yaml:"s3TableIntegrationArn"`
}

