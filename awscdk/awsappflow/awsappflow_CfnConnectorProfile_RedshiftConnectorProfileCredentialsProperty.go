package awsappflow


// The connector-specific profile credentials required when using Amazon Redshift.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftConnectorProfileCredentialsProperty := &redshiftConnectorProfileCredentialsProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//   }
//
type CfnConnectorProfile_RedshiftConnectorProfileCredentialsProperty struct {
	// The password that corresponds to the user name.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The name of the user.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

