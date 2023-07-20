package awspinpoint


// Properties for defining a `CfnBaiduChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBaiduChannelProps := &CfnBaiduChannelProps{
//   	ApiKey: jsii.String("apiKey"),
//   	ApplicationId: jsii.String("applicationId"),
//   	SecretKey: jsii.String("secretKey"),
//
//   	// the properties below are optional
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-baiduchannel.html
//
type CfnBaiduChannelProps struct {
	// The API key that you received from the Baidu Cloud Push service to communicate with the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-baiduchannel.html#cfn-pinpoint-baiduchannel-apikey
	//
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The unique identifier for the Amazon Pinpoint application that you're configuring the Baidu channel for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-baiduchannel.html#cfn-pinpoint-baiduchannel-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The secret key that you received from the Baidu Cloud Push service to communicate with the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-baiduchannel.html#cfn-pinpoint-baiduchannel-secretkey
	//
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
	// Specifies whether to enable the Baidu channel for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-baiduchannel.html#cfn-pinpoint-baiduchannel-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

