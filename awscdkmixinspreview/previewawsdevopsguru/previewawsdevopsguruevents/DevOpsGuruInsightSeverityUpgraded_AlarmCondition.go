package previewawsdevopsguruevents


// Type definition for AlarmCondition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   alarmCondition := &AlarmCondition{
//   	DetectionBand: []*string{
//   		jsii.String("detectionBand"),
//   	},
//   	ReferenceValue: &ReferenceValue{
//   		ComparisonOperator: []*string{
//   			jsii.String("comparisonOperator"),
//   		},
//   		DatapointsToAlarm: []*string{
//   			jsii.String("datapointsToAlarm"),
//   		},
//   		EvaluationPeriod: []*string{
//   			jsii.String("evaluationPeriod"),
//   		},
//   		Threshold: []*string{
//   			jsii.String("threshold"),
//   		},
//   	},
//   }
//
// Experimental.
type DevOpsGuruInsightSeverityUpgraded_AlarmCondition struct {
	// detectionBand property.
	//
	// Specify an array of string values to match this event if the actual value of detectionBand is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DetectionBand *[]*string `field:"optional" json:"detectionBand" yaml:"detectionBand"`
	// referenceValue property.
	//
	// Specify an array of string values to match this event if the actual value of referenceValue is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReferenceValue *DevOpsGuruInsightSeverityUpgraded_ReferenceValue `field:"optional" json:"referenceValue" yaml:"referenceValue"`
}

