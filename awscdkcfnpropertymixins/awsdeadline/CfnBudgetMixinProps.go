package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnBudgetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnBudgetMixinProps := &CfnBudgetMixinProps{
//   	Actions: []interface{}{
//   		&BudgetActionToAddProperty{
//   			Description: jsii.String("description"),
//   			ThresholdPercentage: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ApproximateDollarLimit: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	FarmId: jsii.String("farmId"),
//   	Schedule: &BudgetScheduleProperty{
//   		Fixed: &FixedBudgetScheduleProperty{
//   			EndTime: jsii.String("endTime"),
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UsageTrackingResource: &UsageTrackingResourceProperty{
//   		QueueId: jsii.String("queueId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html
//
type CfnBudgetMixinProps struct {
	// The budget actions to specify what happens when the budget runs out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The dollar limit based on consumed usage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-approximatedollarlimit
	//
	ApproximateDollarLimit *float64 `field:"optional" json:"approximateDollarLimit" yaml:"approximateDollarLimit"`
	// The description of the budget.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the budget.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The farm ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-farmid
	//
	FarmId *string `field:"optional" json:"farmId" yaml:"farmId"`
	// The start and end time of the budget.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-schedule
	//
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The usage details of the allotted budget.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-usagetrackingresource
	//
	UsageTrackingResource interface{} `field:"optional" json:"usageTrackingResource" yaml:"usageTrackingResource"`
}

