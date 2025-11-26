package previewawsmediatailormixins


// The log configuration for the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logConfigurationForChannelProperty := &LogConfigurationForChannelProperty{
//   	LogTypes: []*string{
//   		jsii.String("logTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-logconfigurationforchannel.html
//
type CfnChannelPropsMixin_LogConfigurationForChannelProperty struct {
	// The log types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-logconfigurationforchannel.html#cfn-mediatailor-channel-logconfigurationforchannel-logtypes
	//
	LogTypes *[]*string `field:"optional" json:"logTypes" yaml:"logTypes"`
}

