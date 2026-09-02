package awsdeadline


// The details of a fixed budget schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   fixedBudgetScheduleProperty := &FixedBudgetScheduleProperty{
//   	EndTime: jsii.String("endTime"),
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-fixedbudgetschedule.html
//
type CfnBudgetPropsMixin_FixedBudgetScheduleProperty struct {
	// When the budget ends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-fixedbudgetschedule.html#cfn-deadline-budget-fixedbudgetschedule-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// When the budget starts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-fixedbudgetschedule.html#cfn-deadline-budget-fixedbudgetschedule-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

