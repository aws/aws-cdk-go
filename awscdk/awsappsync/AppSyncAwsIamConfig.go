package awsappsync


// The authorization config in case the HTTP endpoint requires authorization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appSyncAwsIamConfig := &AppSyncAwsIamConfig{
//   	SigningRegion: jsii.String("signingRegion"),
//   	SigningServiceName: jsii.String("signingServiceName"),
//   }
//
type AppSyncAwsIamConfig struct {
	// The signing region for AWS IAM authorization.
	SigningRegion *string `field:"required" json:"signingRegion" yaml:"signingRegion"`
	// The signing service name for AWS IAM authorization.
	SigningServiceName *string `field:"required" json:"signingServiceName" yaml:"signingServiceName"`
}

