package awsmediapackagev2


// <p>The settings for what common media server data (CMSD) headers AWS Elemental MediaPackage includes in responses to the CDN.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputHeaderConfigurationProperty := &OutputHeaderConfigurationProperty{
//   	PublishMqcs: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-outputheaderconfiguration.html
//
type CfnChannel_OutputHeaderConfigurationProperty struct {
	// <p>When true, AWS Elemental MediaPackage includes the MQCS in responses to the CDN.
	//
	// This setting is valid only when <code>InputType</code> is <code>CMAF</code>.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-outputheaderconfiguration.html#cfn-mediapackagev2-channel-outputheaderconfiguration-publishmqcs
	//
	PublishMqcs interface{} `field:"optional" json:"publishMqcs" yaml:"publishMqcs"`
}

