package awsrbin


// The retention period of the rule.
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
	// The retention period unit of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-retentionperiod.html#cfn-rbin-rule-retentionperiod-retentionperiodunit
	//
	RetentionPeriodUnit *string `field:"required" json:"retentionPeriodUnit" yaml:"retentionPeriodUnit"`
	// The retention period value of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-retentionperiod.html#cfn-rbin-rule-retentionperiod-retentionperiodvalue
	//
	RetentionPeriodValue *float64 `field:"required" json:"retentionPeriodValue" yaml:"retentionPeriodValue"`
}

