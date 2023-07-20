package awsevents


// Contains the OAuth authorization parameters to use for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientParametersProperty := &ClientParametersProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-clientparameters.html
//
type CfnConnection_ClientParametersProperty struct {
	// The client ID to use for OAuth authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-clientparameters.html#cfn-events-connection-clientparameters-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret assciated with the client ID to use for OAuth authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-clientparameters.html#cfn-events-connection-clientparameters-clientsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

