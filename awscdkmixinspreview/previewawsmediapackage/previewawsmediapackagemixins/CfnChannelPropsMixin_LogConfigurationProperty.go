package previewawsmediapackagemixins


// The access log configuration parameters for your channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	LogGroupName: jsii.String("logGroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-logconfiguration.html
//
type CfnChannelPropsMixin_LogConfigurationProperty struct {
	// Sets a custom Amazon CloudWatch log group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-logconfiguration.html#cfn-mediapackage-channel-logconfiguration-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

