package awskinesisfirehose


// Configure a Snowflake VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeVpcConfigurationProperty := &SnowflakeVpcConfigurationProperty{
//   	PrivateLinkVpceId: jsii.String("privateLinkVpceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakevpcconfiguration.html
//
type CfnDeliveryStream_SnowflakeVpcConfigurationProperty struct {
	// The VPCE ID for Firehose to privately connect with Snowflake.
	//
	// The ID format is com.amazonaws.vpce.[region].vpce-svc-<[id]>. For more information, see [Amazon PrivateLink & Snowflake](https://docs.aws.amazon.com/https://docs.snowflake.com/en/user-guide/admin-security-privatelink)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakevpcconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakevpcconfiguration-privatelinkvpceid
	//
	PrivateLinkVpceId *string `field:"required" json:"privateLinkVpceId" yaml:"privateLinkVpceId"`
}

