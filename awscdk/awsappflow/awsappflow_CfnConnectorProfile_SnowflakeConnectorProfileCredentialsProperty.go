package awsappflow


// The connector-specific profile credentials required when using Snowflake.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeConnectorProfileCredentialsProperty := &snowflakeConnectorProfileCredentialsProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//   }
//
type CfnConnectorProfile_SnowflakeConnectorProfileCredentialsProperty struct {
	// The password that corresponds to the user name.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The name of the user.
	Username *string `field:"required" json:"username" yaml:"username"`
}

