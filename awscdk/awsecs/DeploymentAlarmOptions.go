package awsecs


// Options for deployment alarms.
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
type DeploymentAlarmOptions struct {
	// Default rollback on alarm.
	Behavior AlarmBehavior `field:"optional" json:"behavior" yaml:"behavior"`
}

