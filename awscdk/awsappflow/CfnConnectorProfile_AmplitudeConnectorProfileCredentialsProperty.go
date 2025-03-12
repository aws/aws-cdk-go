package awsappflow


// The connector-specific credentials required when using Amplitude.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amplitudeConnectorProfileCredentialsProperty := &AmplitudeConnectorProfileCredentialsProperty{
//   	ApiKey: jsii.String("apiKey"),
//   	SecretKey: jsii.String("secretKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-amplitudeconnectorprofilecredentials.html
//
type CfnConnectorProfile_AmplitudeConnectorProfileCredentialsProperty struct {
	// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-amplitudeconnectorprofilecredentials.html#cfn-appflow-connectorprofile-amplitudeconnectorprofilecredentials-apikey
	//
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The Secret Access Key portion of the credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-amplitudeconnectorprofilecredentials.html#cfn-appflow-connectorprofile-amplitudeconnectorprofilecredentials-secretkey
	//
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
}

