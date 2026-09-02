package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnBudget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBudgetProps := &CfnBudgetProps{
//   	Actions: []interface{}{
//   		&BudgetActionToAddProperty{
//   			ThresholdPercentage: jsii.Number(123),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//   	ApproximateDollarLimit: jsii.Number(123),
//   	DisplayName: jsii.String("displayName"),
//   	FarmId: jsii.String("farmId"),
//   	Schedule: &BudgetScheduleProperty{
//   		Fixed: &FixedBudgetScheduleProperty{
//   			EndTime: jsii.String("endTime"),
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	UsageTrackingResource: &UsageTrackingResourceProperty{
//   		QueueId: jsii.String("queueId"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html
//
type CfnBudgetProps struct {
	// The budget actions to specify what happens when the budget runs out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The dollar limit based on consumed usage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-approximatedollarlimit
	//
	ApproximateDollarLimit *float64 `field:"required" json:"approximateDollarLimit" yaml:"approximateDollarLimit"`
	// The display name of the budget.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The farm ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-farmid
	//
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The start and end time of the budget.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-schedule
	//
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// The usage details of the allotted budget.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-usagetrackingresource
	//
	UsageTrackingResource interface{} `field:"required" json:"usageTrackingResource" yaml:"usageTrackingResource"`
	// The description of the budget.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

