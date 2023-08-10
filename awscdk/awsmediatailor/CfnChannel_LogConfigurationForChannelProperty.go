package awsmediatailor


// <p>The log configuration for the channel.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationForChannelProperty := &LogConfigurationForChannelProperty{
//   	LogTypes: []*string{
//   		jsii.String("logTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-logconfigurationforchannel.html
//
type CfnChannel_LogConfigurationForChannelProperty struct {
	// <p>The log types.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-logconfigurationforchannel.html#cfn-mediatailor-channel-logconfigurationforchannel-logtypes
	//
	LogTypes *[]*string `field:"optional" json:"logTypes" yaml:"logTypes"`
}

