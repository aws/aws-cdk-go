package awscloudwatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   alarmPromQLCriteriaProperty := &AlarmPromQLCriteriaProperty{
//   	PendingPeriod: jsii.Number(123),
//   	Query: jsii.String("query"),
//   	RecoveryPeriod: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-alarmpromqlcriteria.html
//
type CfnAlarmPropsMixin_AlarmPromQLCriteriaProperty struct {
	// The pending period for the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-alarmpromqlcriteria.html#cfn-cloudwatch-alarm-alarmpromqlcriteria-pendingperiod
	//
	PendingPeriod *float64 `field:"optional" json:"pendingPeriod" yaml:"pendingPeriod"`
	// The PromQL query string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-alarmpromqlcriteria.html#cfn-cloudwatch-alarm-alarmpromqlcriteria-query
	//
	Query *string `field:"optional" json:"query" yaml:"query"`
	// The recovery period for the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-alarmpromqlcriteria.html#cfn-cloudwatch-alarm-alarmpromqlcriteria-recoveryperiod
	//
	RecoveryPeriod *float64 `field:"optional" json:"recoveryPeriod" yaml:"recoveryPeriod"`
}

