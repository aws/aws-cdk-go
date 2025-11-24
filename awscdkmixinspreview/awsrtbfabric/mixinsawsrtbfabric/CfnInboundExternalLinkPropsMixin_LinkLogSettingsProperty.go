package mixinsawsrtbfabric


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-linklogsettings.html
//
type CfnInboundExternalLinkPropsMixin_LinkLogSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-linklogsettings.html#cfn-rtbfabric-inboundexternallink-linklogsettings-applicationlogs
	//
	ApplicationLogs interface{} `field:"optional" json:"applicationLogs" yaml:"applicationLogs"`
}

