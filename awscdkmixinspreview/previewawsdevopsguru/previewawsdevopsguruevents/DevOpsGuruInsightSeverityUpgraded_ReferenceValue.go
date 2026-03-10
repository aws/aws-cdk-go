package previewawsdevopsguruevents


// Type definition for ReferenceValue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceValue := &ReferenceValue{
//   	ComparisonOperator: []*string{
//   		jsii.String("comparisonOperator"),
//   	},
//   	DatapointsToAlarm: []*string{
//   		jsii.String("datapointsToAlarm"),
//   	},
//   	EvaluationPeriod: []*string{
//   		jsii.String("evaluationPeriod"),
//   	},
//   	Threshold: []*string{
//   		jsii.String("threshold"),
//   	},
//   }
//
// Experimental.
type DevOpsGuruInsightSeverityUpgraded_ReferenceValue struct {
	// comparisonOperator property.
	//
	// Specify an array of string values to match this event if the actual value of comparisonOperator is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ComparisonOperator *[]*string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// datapointsToAlarm property.
	//
	// Specify an array of string values to match this event if the actual value of datapointsToAlarm is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatapointsToAlarm *[]*string `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// evaluationPeriod property.
	//
	// Specify an array of string values to match this event if the actual value of evaluationPeriod is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EvaluationPeriod *[]*string `field:"optional" json:"evaluationPeriod" yaml:"evaluationPeriod"`
	// threshold property.
	//
	// Specify an array of string values to match this event if the actual value of threshold is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Threshold *[]*string `field:"optional" json:"threshold" yaml:"threshold"`
}

