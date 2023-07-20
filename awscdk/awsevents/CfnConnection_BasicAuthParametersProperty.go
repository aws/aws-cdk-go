package awsevents


// Contains the Basic authorization parameters for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basicAuthParametersProperty := &BasicAuthParametersProperty{
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-basicauthparameters.html
//
type CfnConnection_BasicAuthParametersProperty struct {
	// The password associated with the user name to use for Basic authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-basicauthparameters.html#cfn-events-connection-basicauthparameters-password
	//
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name to use for Basic authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-basicauthparameters.html#cfn-events-connection-basicauthparameters-username
	//
	Username *string `field:"required" json:"username" yaml:"username"`
}

