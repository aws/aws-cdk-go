package previewawsxraymixins


// The structure containing configurations related to insights.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   insightsConfigurationProperty := &InsightsConfigurationProperty{
//   	InsightsEnabled: jsii.Boolean(false),
//   	NotificationsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-group-insightsconfiguration.html
//
type CfnGroupPropsMixin_InsightsConfigurationProperty struct {
	// Set the InsightsEnabled value to true to enable insights or false to disable insights.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-group-insightsconfiguration.html#cfn-xray-group-insightsconfiguration-insightsenabled
	//
	InsightsEnabled interface{} `field:"optional" json:"insightsEnabled" yaml:"insightsEnabled"`
	// Set the NotificationsEnabled value to true to enable insights notifications.
	//
	// Notifications can only be enabled on a group with InsightsEnabled set to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-group-insightsconfiguration.html#cfn-xray-group-insightsconfiguration-notificationsenabled
	//
	NotificationsEnabled interface{} `field:"optional" json:"notificationsEnabled" yaml:"notificationsEnabled"`
}

