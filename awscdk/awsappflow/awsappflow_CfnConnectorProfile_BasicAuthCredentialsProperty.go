package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basicAuthCredentialsProperty := &basicAuthCredentialsProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//   }
//
type CfnConnectorProfile_BasicAuthCredentialsProperty struct {
	// `CfnConnectorProfile.BasicAuthCredentialsProperty.Password`.
	Password *string `field:"required" json:"password" yaml:"password"`
	// `CfnConnectorProfile.BasicAuthCredentialsProperty.Username`.
	Username *string `field:"required" json:"username" yaml:"username"`
}

