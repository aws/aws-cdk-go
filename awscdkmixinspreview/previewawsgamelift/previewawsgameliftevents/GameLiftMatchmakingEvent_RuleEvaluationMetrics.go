package previewawsgameliftevents


// Type definition for RuleEvaluationMetrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleEvaluationMetrics := &RuleEvaluationMetrics{
//   	FailedCount: []*string{
//   		jsii.String("failedCount"),
//   	},
//   	PassedCount: []*string{
//   		jsii.String("passedCount"),
//   	},
//   	RuleName: []*string{
//   		jsii.String("ruleName"),
//   	},
//   }
//
// Experimental.
type GameLiftMatchmakingEvent_RuleEvaluationMetrics struct {
	// failedCount property.
	//
	// Specify an array of string values to match this event if the actual value of failedCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailedCount *[]*string `field:"optional" json:"failedCount" yaml:"failedCount"`
	// passedCount property.
	//
	// Specify an array of string values to match this event if the actual value of passedCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PassedCount *[]*string `field:"optional" json:"passedCount" yaml:"passedCount"`
	// ruleName property.
	//
	// Specify an array of string values to match this event if the actual value of ruleName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RuleName *[]*string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

