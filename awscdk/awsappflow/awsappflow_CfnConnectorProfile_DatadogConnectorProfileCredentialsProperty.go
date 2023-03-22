package awsappflow


// The connector-specific credentials required by Datadog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datadogConnectorProfileCredentialsProperty := &datadogConnectorProfileCredentialsProperty{
//   	apiKey: jsii.String("apiKey"),
//   	applicationKey: jsii.String("applicationKey"),
//   }
//
type CfnConnectorProfile_DatadogConnectorProfileCredentialsProperty struct {
	// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Application keys, in conjunction with your API key, give you full access to Datadog’s programmatic API.
	//
	// Application keys are associated with the user account that created them. The application key is used to log all requests made to the API.
	ApplicationKey *string `field:"required" json:"applicationKey" yaml:"applicationKey"`
}

