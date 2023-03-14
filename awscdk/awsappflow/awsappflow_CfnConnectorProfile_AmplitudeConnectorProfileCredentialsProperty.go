package awsappflow


// The connector-specific credentials required when using Amplitude.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amplitudeConnectorProfileCredentialsProperty := &amplitudeConnectorProfileCredentialsProperty{
//   	apiKey: jsii.String("apiKey"),
//   	secretKey: jsii.String("secretKey"),
//   }
//
type CfnConnectorProfile_AmplitudeConnectorProfileCredentialsProperty struct {
	// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The Secret Access Key portion of the credentials.
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
}

