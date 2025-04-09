package awsapplicationsignals


// The recurrence rule for the SLO time window exclusion .
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
	// A cron or rate expression that specifies the schedule for the exclusion window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-recurrencerule.html#cfn-applicationsignals-servicelevelobjective-recurrencerule-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

