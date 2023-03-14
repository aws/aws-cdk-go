package awsxray


// The structure containing configurations related to insights.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   insightsConfigurationProperty := &insightsConfigurationProperty{
//   	insightsEnabled: jsii.Boolean(false),
//   	notificationsEnabled: jsii.Boolean(false),
//   }
//
type CfnGroup_InsightsConfigurationProperty struct {
	// Set the InsightsEnabled value to true to enable insights or false to disable insights.
	InsightsEnabled interface{} `field:"optional" json:"insightsEnabled" yaml:"insightsEnabled"`
	// Set the NotificationsEnabled value to true to enable insights notifications.
	//
	// Notifications can only be enabled on a group with InsightsEnabled set to true.
	NotificationsEnabled interface{} `field:"optional" json:"notificationsEnabled" yaml:"notificationsEnabled"`
}

