package awsecs


// Deployment behavior when an ECS Service Deployment Alarm is triggered.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import cw "github.com/aws/aws-cdk-go/awscdk"
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var elbAlarm cloudwatch.Alarm
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	DeploymentAlarms: &DeploymentAlarmConfig{
//   		Alarms: []interface{}{
//   			elbAlarm.alarmName,
//   		},
//   		Behavior: alarmBehavior_ROLLBACK_ON_ALARM,
//   	},
//   })
//
//   // Defining a deployment alarm after the service has been created
//   cpuAlarmName := "MyCpuMetricAlarm"
//   cw.NewAlarm(this, jsii.String("CPUAlarm"), &AlarmProps{
//   	AlarmName: cpuAlarmName,
//   	Metric: service.MetricCpuUtilization(),
//   	EvaluationPeriods: jsii.Number(2),
//   	Threshold: jsii.Number(80),
//   })
//   service.EnableDeploymentAlarms([]*string{
//   	cpuAlarmName,
//   }, alarmBehavior_FAIL_ON_ALARM)
//
type AlarmBehavior string

const (
	// ROLLBACK_ON_ALARM causes the service to roll back to the previous deployment when any deployment alarm enters the 'Alarm' state.
	//
	// The Cloudformation stack
	// will be rolled back and enter state "UPDATE_ROLLBACK_COMPLETE".
	AlarmBehavior_ROLLBACK_ON_ALARM AlarmBehavior = "ROLLBACK_ON_ALARM"
	// FAIL_ON_ALARM causes the deployment to fail immediately when any deployment alarm enters the 'Alarm' state.
	//
	// In order to restore functionality, you must
	// roll the stack forward by pushing a new version of the ECS service.
	AlarmBehavior_FAIL_ON_ALARM AlarmBehavior = "FAIL_ON_ALARM"
)

