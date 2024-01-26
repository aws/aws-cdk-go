package awskinesisfirehose


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakevpcconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakevpcconfiguration-privatelinkvpceid
	//
	PrivateLinkVpceId *string `field:"required" json:"privateLinkVpceId" yaml:"privateLinkVpceId"`
}

