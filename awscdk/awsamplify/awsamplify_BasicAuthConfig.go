package awsamplify


// A Basic Auth configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basicAuthConfig := &basicAuthConfig{
//   	enableBasicAuth: jsii.Boolean(false),
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//   }
//
// Experimental.
type BasicAuthConfig struct {
	// Whether to enable Basic Auth.
	// Experimental.
	EnableBasicAuth *bool `field:"required" json:"enableBasicAuth" yaml:"enableBasicAuth"`
	// The password.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

