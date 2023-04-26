package awsappflow


// The basic auth credentials required for basic authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basicAuthCredentialsProperty := &BasicAuthCredentialsProperty{
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//   }
//
type CfnConnectorProfile_BasicAuthCredentialsProperty struct {
	// The password to use to connect to a resource.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username to use to connect to a resource.
	Username *string `field:"required" json:"username" yaml:"username"`
}

