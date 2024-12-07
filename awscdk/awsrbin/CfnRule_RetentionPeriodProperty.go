package awsrbin


// Information about the retention period for which the retention rule is to retain resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retentionPeriodProperty := &RetentionPeriodProperty{
//   	RetentionPeriodUnit: jsii.String("retentionPeriodUnit"),
//   	RetentionPeriodValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-retentionperiod.html
//
type CfnRule_RetentionPeriodProperty struct {
	// The unit of time in which the retention period is measured.
	//
	// Currently, only `DAYS` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-retentionperiod.html#cfn-rbin-rule-retentionperiod-retentionperiodunit
	//
	RetentionPeriodUnit *string `field:"required" json:"retentionPeriodUnit" yaml:"retentionPeriodUnit"`
	// The period value for which the retention rule is to retain resources.
	//
	// The period is measured using the unit specified for *RetentionPeriodUnit* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-retentionperiod.html#cfn-rbin-rule-retentionperiod-retentionperiodvalue
	//
	RetentionPeriodValue *float64 `field:"required" json:"retentionPeriodValue" yaml:"retentionPeriodValue"`
}

