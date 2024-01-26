package awskinesisfirehose


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeRoleConfigurationProperty := &SnowflakeRoleConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	SnowflakeRole: jsii.String("snowflakeRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakeroleconfiguration.html
//
type CfnDeliveryStream_SnowflakeRoleConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakeroleconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakeroleconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakeroleconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakeroleconfiguration-snowflakerole
	//
	SnowflakeRole *string `field:"optional" json:"snowflakeRole" yaml:"snowflakeRole"`
}

