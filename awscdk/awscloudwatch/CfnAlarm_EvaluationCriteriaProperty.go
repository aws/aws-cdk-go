package awscloudwatch


// The evaluation criteria for an alarm.
//
// This is a union type that currently supports ``PromQLCriteria``.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationCriteriaProperty := &EvaluationCriteriaProperty{
//   	PromQlCriteria: &AlarmPromQLCriteriaProperty{
//   		PendingPeriod: jsii.Number(123),
//   		Query: jsii.String("query"),
//   		RecoveryPeriod: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-evaluationcriteria.html
//
type CfnAlarm_EvaluationCriteriaProperty struct {
	// Contains the configuration that determines how a PromQL alarm evaluates its contributors, including the query to run and the durations that define when contributors transition between states.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-evaluationcriteria.html#cfn-cloudwatch-alarm-evaluationcriteria-promqlcriteria
	//
	PromQlCriteria interface{} `field:"optional" json:"promQlCriteria" yaml:"promQlCriteria"`
}

