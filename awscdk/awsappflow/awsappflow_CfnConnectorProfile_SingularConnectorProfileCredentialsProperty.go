package awsappflow


// The connector-specific profile credentials required when using Singular.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singularConnectorProfileCredentialsProperty := &singularConnectorProfileCredentialsProperty{
//   	apiKey: jsii.String("apiKey"),
//   }
//
type CfnConnectorProfile_SingularConnectorProfileCredentialsProperty struct {
	// A unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
}

