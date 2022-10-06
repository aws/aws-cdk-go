package awscloudformation


// Configuration options for custom resource providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResourceProviderConfig := &customResourceProviderConfig{
//   	serviceToken: jsii.String("serviceToken"),
//   }
//
// Deprecated: used in {@link ICustomResourceProvider} which is now deprecated.
type CustomResourceProviderConfig struct {
	// The ARN of the SNS topic or the AWS Lambda function which implements this provider.
	// Deprecated: used in {@link ICustomResourceProvider} which is now deprecated.
	ServiceToken *string `field:"required" json:"serviceToken" yaml:"serviceToken"`
}

