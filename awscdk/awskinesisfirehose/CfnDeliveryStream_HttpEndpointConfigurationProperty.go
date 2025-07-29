package awskinesisfirehose


// Describes the configuration of the HTTP endpoint to which Kinesis Firehose delivers data.
//
// Kinesis Firehose supports any custom HTTP endpoint or HTTP endpoints owned by supported third-party service providers, including Datadog, MongoDB, and New Relic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpEndpointConfigurationProperty := &HttpEndpointConfigurationProperty{
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	AccessKey: jsii.String("accessKey"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointconfiguration.html
//
type CfnDeliveryStream_HttpEndpointConfigurationProperty struct {
	// The URL of the HTTP endpoint selected as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointconfiguration-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// The access key required for Kinesis Firehose to authenticate with the HTTP endpoint selected as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointconfiguration-accesskey
	//
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// The name of the HTTP endpoint selected as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

