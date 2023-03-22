package awsevents


// Contains the Basic authorization parameters for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basicAuthParametersProperty := &basicAuthParametersProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//   }
//
type CfnConnection_BasicAuthParametersProperty struct {
	// The password associated with the user name to use for Basic authorization.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name to use for Basic authorization.
	Username *string `field:"required" json:"username" yaml:"username"`
}

