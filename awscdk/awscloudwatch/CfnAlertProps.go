package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAlert`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlertProps := &CfnAlertProps{
//   	Name: jsii.String("name"),
//   	ProfileId: jsii.String("profileId"),
//   	Rule: &RuleProperty{
//   		TelemetryRule: &TelemetryRuleProperty{
//   			Condition: &ConditionProperty{
//   				Comparator: jsii.String("comparator"),
//   				CriticalThreshold: jsii.Number(123),
//   				ThresholdField: jsii.String("thresholdField"),
//   				ThresholdMode: jsii.String("thresholdMode"),
//   				WarningThreshold: jsii.Number(123),
//   			},
//   			Evaluation: &EvaluationProperty{
//   				IntervalSeconds: jsii.Number(123),
//
//   				// the properties below are optional
//   				PendingDurationSeconds: jsii.Number(123),
//   				RecoveryDurationSeconds: jsii.Number(123),
//   			},
//   			Query: &QueryProperty{
//   				Expression: jsii.String("expression"),
//   				Language: jsii.String("language"),
//   			},
//
//   			// the properties below are optional
//   			NoData: &NoDataProperty{
//   				TreatAs: jsii.String("treatAs"),
//   			},
//   		},
//   	},
//   	SpaceId: jsii.String("spaceId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	NotificationRules: []interface{}{
//   		&NotificationRuleProperty{
//   			Target: &TargetProperty{
//   				Arn: jsii.String("arn"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Metadata: map[string]*string{
//   					"metadataKey": jsii.String("metadata"),
//   				},
//   			},
//   			Trigger: &TriggerProperty{
//   				StateValues: []*string{
//   					jsii.String("stateValues"),
//   				},
//   			},
//   		},
//   	},
//   	NotificationsEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alert.html
//
type CfnAlertProps struct {
	// The display name of the alert.
	//
	// Names don't need to be unique within a space; the alert is identified by AlertId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alert.html#cfn-cloudwatch-alert-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the access profile the alert uses to run its query and send notifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alert.html#cfn-cloudwatch-alert-profileid
	//
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// The rule that defines how the alert is evaluated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alert.html#cfn-cloudwatch-alert-rule
	//
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// The ID of the space to create the alert in.
	//
	// Changing this value replaces the alert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alert.html#cfn-cloudwatch-alert-spaceid
	//
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
	// A description of the alert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alert.html#cfn-cloudwatch-alert-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The notification rules that determine when and where notifications are sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alert.html#cfn-cloudwatch-alert-notificationrules
	//
	NotificationRules interface{} `field:"optional" json:"notificationRules" yaml:"notificationRules"`
	// Whether the alert sends notifications.
	//
	// When false, the alert is still evaluated and records its state, but sends nothing. Defaults to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alert.html#cfn-cloudwatch-alert-notificationsenabled
	//
	NotificationsEnabled interface{} `field:"optional" json:"notificationsEnabled" yaml:"notificationsEnabled"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alert.html#cfn-cloudwatch-alert-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

