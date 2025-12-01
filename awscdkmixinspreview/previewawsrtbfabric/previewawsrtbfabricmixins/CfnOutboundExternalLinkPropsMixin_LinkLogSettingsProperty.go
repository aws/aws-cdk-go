package previewawsrtbfabricmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   linkLogSettingsProperty := &LinkLogSettingsProperty{
//   	ApplicationLogs: &ApplicationLogsProperty{
//   		LinkApplicationLogSampling: &LinkApplicationLogSamplingProperty{
//   			ErrorLog: jsii.Number(123),
//   			FilterLog: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-outboundexternallink-linklogsettings.html
//
type CfnOutboundExternalLinkPropsMixin_LinkLogSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-outboundexternallink-linklogsettings.html#cfn-rtbfabric-outboundexternallink-linklogsettings-applicationlogs
	//
	ApplicationLogs interface{} `field:"optional" json:"applicationLogs" yaml:"applicationLogs"`
}

