package awscloudwatch


// Contains the configuration that determines how a PromQL alarm evaluates its contributors, including the query to run and the durations that define when contributors transition between states.
//
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
	// The duration, in seconds, that a contributor must be continuously breaching before it transitions to the ``ALARM`` state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-alarmpromqlcriteria.html#cfn-cloudwatch-alarm-alarmpromqlcriteria-pendingperiod
	//
	PendingPeriod *float64 `field:"optional" json:"pendingPeriod" yaml:"pendingPeriod"`
	// The PromQL query that the alarm evaluates.
	//
	// The query must return a result of vector type. Each entry in the vector result represents an alarm contributor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-alarmpromqlcriteria.html#cfn-cloudwatch-alarm-alarmpromqlcriteria-query
	//
	Query *string `field:"optional" json:"query" yaml:"query"`
	// The duration, in seconds, that a contributor must continuously not be breaching before it transitions back to the ``OK`` state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-alarmpromqlcriteria.html#cfn-cloudwatch-alarm-alarmpromqlcriteria-recoveryperiod
	//
	RecoveryPeriod *float64 `field:"optional" json:"recoveryPeriod" yaml:"recoveryPeriod"`
}

