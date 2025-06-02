package awsapplicationsignals


// The recurrence rule for the time exclusion window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recurrenceRuleProperty := &RecurrenceRuleProperty{
//   	Expression: jsii.String("expression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-recurrencerule.html
//
type CfnServiceLevelObjective_RecurrenceRuleProperty struct {
	// The following two rules are supported:.
	//
	// - rate(value unit) - The value must be a positive integer and the unit can be hour|day|month.
	// - cron - An expression which consists of six fields separated by white spaces: (minutes hours day_of_month month day_of_week year).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-recurrencerule.html#cfn-applicationsignals-servicelevelobjective-recurrencerule-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

