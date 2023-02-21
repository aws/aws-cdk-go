package awsxray


// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnGroupProps := &cfnGroupProps{
//   	filterExpression: jsii.String("filterExpression"),
//   	groupName: jsii.String("groupName"),
//   	insightsConfiguration: &insightsConfigurationProperty{
//   		insightsEnabled: jsii.Boolean(false),
//   		notificationsEnabled: jsii.Boolean(false),
//   	},
//   	tags: []interface{}{
//   		tags,
//   	},
//   }
//
type CfnGroupProps struct {
	// The filter expression defining the parameters to include traces.
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
	// The unique case-sensitive name of the group.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The structure containing configurations related to insights.
	//
	// - The InsightsEnabled boolean can be set to true to enable insights for the group or false to disable insights for the group.
	// - The NotificationsEnabled boolean can be set to true to enable insights notifications through Amazon EventBridge for the group.
	InsightsConfiguration interface{} `field:"optional" json:"insightsConfiguration" yaml:"insightsConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]interface{} `field:"optional" json:"tags" yaml:"tags"`
}

