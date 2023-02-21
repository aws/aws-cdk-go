package awsevents


// Contains the OAuth authorization parameters to use for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientParametersProperty := &clientParametersProperty{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   }
//
type CfnConnection_ClientParametersProperty struct {
	// The client ID to use for OAuth authorization.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret assciated with the client ID to use for OAuth authorization.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

