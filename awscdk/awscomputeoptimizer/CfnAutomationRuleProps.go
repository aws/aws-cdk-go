package awscomputeoptimizer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAutomationRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAutomationRuleProps := &CfnAutomationRuleProps{
//   	Name: jsii.String("name"),
//   	RecommendedActionTypes: []*string{
//   		jsii.String("recommendedActionTypes"),
//   	},
//   	RuleType: jsii.String("ruleType"),
//   	Schedule: &ScheduleProperty{
//   		ExecutionWindowInMinutes: jsii.Number(123),
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   		ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   	},
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	Criteria: &CriteriaProperty{
//   		EbsVolumeSizeInGib: []interface{}{
//   			&IntegerCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		EbsVolumeType: []interface{}{
//   			&StringCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		EstimatedMonthlySavings: []interface{}{
//   			&DoubleCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		LookBackPeriodInDays: []interface{}{
//   			&IntegerCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		Region: []interface{}{
//   			&StringCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		ResourceArn: []interface{}{
//   			&StringCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		ResourceTag: []interface{}{
//   			&ResourceTagsCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Key: jsii.String("key"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		RestartNeeded: []interface{}{
//   			&StringCriteriaConditionProperty{
//   				Comparison: jsii.String("comparison"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	OrganizationConfiguration: &OrganizationConfigurationProperty{
//   		AccountIds: []*string{
//   			jsii.String("accountIds"),
//   		},
//   		RuleApplyOrder: jsii.String("ruleApplyOrder"),
//   	},
//   	Priority: jsii.String("priority"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html
//
type CfnAutomationRuleProps struct {
	// The name of the automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The types of recommended actions this rule will implement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-recommendedactiontypes
	//
	RecommendedActionTypes *[]*string `field:"required" json:"recommendedActionTypes" yaml:"recommendedActionTypes"`
	// The type of automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-ruletype
	//
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-schedule
	//
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// The status of the automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-criteria
	//
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// The description of the automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-organizationconfiguration
	//
	OrganizationConfiguration interface{} `field:"optional" json:"organizationConfiguration" yaml:"organizationConfiguration"`
	// Rule priority within its group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-priority
	//
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Tags associated with the automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

