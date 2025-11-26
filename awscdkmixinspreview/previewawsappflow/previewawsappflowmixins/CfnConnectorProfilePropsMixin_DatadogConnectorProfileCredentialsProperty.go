package previewawsappflowmixins


// The connector-specific credentials required by Datadog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   datadogConnectorProfileCredentialsProperty := &DatadogConnectorProfileCredentialsProperty{
//   	ApiKey: jsii.String("apiKey"),
//   	ApplicationKey: jsii.String("applicationKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-datadogconnectorprofilecredentials.html
//
type CfnConnectorProfilePropsMixin_DatadogConnectorProfileCredentialsProperty struct {
	// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-datadogconnectorprofilecredentials.html#cfn-appflow-connectorprofile-datadogconnectorprofilecredentials-apikey
	//
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Application keys, in conjunction with your API key, give you full access to Datadog’s programmatic API.
	//
	// Application keys are associated with the user account that created them. The application key is used to log all requests made to the API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-datadogconnectorprofilecredentials.html#cfn-appflow-connectorprofile-datadogconnectorprofilecredentials-applicationkey
	//
	ApplicationKey *string `field:"optional" json:"applicationKey" yaml:"applicationKey"`
}

