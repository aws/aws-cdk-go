package awsxray

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupProps := &CfnGroupProps{
//   	GroupName: jsii.String("groupName"),
//
//   	// the properties below are optional
//   	FilterExpression: jsii.String("filterExpression"),
//   	InsightsConfiguration: &InsightsConfigurationProperty{
//   		InsightsEnabled: jsii.Boolean(false),
//   		NotificationsEnabled: jsii.Boolean(false),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-group.html
//
type CfnGroupProps struct {
	// The unique case-sensitive name of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-group.html#cfn-xray-group-groupname
	//
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The filter expression defining the parameters to include traces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-group.html#cfn-xray-group-filterexpression
	//
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
	// The structure containing configurations related to insights.
	//
	// - The InsightsEnabled boolean can be set to true to enable insights for the group or false to disable insights for the group.
	// - The NotificationsEnabled boolean can be set to true to enable insights notifications through Amazon EventBridge for the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-group.html#cfn-xray-group-insightsconfiguration
	//
	InsightsConfiguration interface{} `field:"optional" json:"insightsConfiguration" yaml:"insightsConfiguration"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-group.html#cfn-xray-group-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

