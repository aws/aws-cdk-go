package awsrtbfabric


// Describes the settings for a link log.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-linklogsettings.html
//
type CfnLink_LinkLogSettingsProperty struct {
	// Describes the configuration of a link application log.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-linklogsettings.html#cfn-rtbfabric-link-linklogsettings-applicationlogs
	//
	ApplicationLogs interface{} `field:"required" json:"applicationLogs" yaml:"applicationLogs"`
}

