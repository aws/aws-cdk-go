package previewawsrbinmixins


// Information about the retention period for which the retention rule is to retain resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retentionPeriodProperty := &RetentionPeriodProperty{
//   	RetentionPeriodUnit: jsii.String("retentionPeriodUnit"),
//   	RetentionPeriodValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-retentionperiod.html
//
type CfnRulePropsMixin_RetentionPeriodProperty struct {
	// The unit of time in which the retention period is measured.
	//
	// Currently, only `DAYS` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-retentionperiod.html#cfn-rbin-rule-retentionperiod-retentionperiodunit
	//
	RetentionPeriodUnit *string `field:"optional" json:"retentionPeriodUnit" yaml:"retentionPeriodUnit"`
	// The period value for which the retention rule is to retain resources, measured in days.
	//
	// The supported retention periods are:
	//
	// - EBS volumes: 1 - 7 days
	// - EBS snapshots and EBS-backed AMIs: 1 - 365 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-retentionperiod.html#cfn-rbin-rule-retentionperiod-retentionperiodvalue
	//
	RetentionPeriodValue *float64 `field:"optional" json:"retentionPeriodValue" yaml:"retentionPeriodValue"`
}

