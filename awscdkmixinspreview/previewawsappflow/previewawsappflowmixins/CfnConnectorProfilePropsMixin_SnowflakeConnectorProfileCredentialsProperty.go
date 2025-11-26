package previewawsappflowmixins


// The connector-specific profile credentials required when using Snowflake.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snowflakeConnectorProfileCredentialsProperty := &SnowflakeConnectorProfileCredentialsProperty{
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofilecredentials.html
//
type CfnConnectorProfilePropsMixin_SnowflakeConnectorProfileCredentialsProperty struct {
	// The password that corresponds to the user name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofilecredentials.html#cfn-appflow-connectorprofile-snowflakeconnectorprofilecredentials-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The name of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofilecredentials.html#cfn-appflow-connectorprofile-snowflakeconnectorprofilecredentials-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

