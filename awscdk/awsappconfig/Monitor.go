package awsappconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Defines monitors that will be associated with an AWS AppConfig environment.
//
// Example:
//   var application application
//   var alarm alarm
//   var compositeAlarm compositeAlarm
//
//
//   appconfig.NewEnvironment(this, jsii.String("MyEnvironment"), &EnvironmentProps{
//   	Application: Application,
//   	Monitors: []monitor{
//   		appconfig.*monitor_FromCloudWatchAlarm(alarm),
//   		appconfig.*monitor_*FromCloudWatchAlarm(compositeAlarm),
//   	},
//   })
//
type Monitor interface {
	// The alarm ARN for AWS AppConfig to monitor.
	AlarmArn() *string
	// The IAM role ARN for AWS AppConfig to view the alarm state.
	AlarmRoleArn() *string
	// Indicates whether a CloudWatch alarm is a composite alarm.
	IsCompositeAlarm() *bool
	// The type of monitor.
	MonitorType() MonitorType
}

// The jsii proxy struct for Monitor
type jsiiProxy_Monitor struct {
	_ byte // padding
}

func (j *jsiiProxy_Monitor) AlarmArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) AlarmRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) IsCompositeAlarm() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isCompositeAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) MonitorType() MonitorType {
	var returns MonitorType
	_jsii_.Get(
		j,
		"monitorType",
		&returns,
	)
	return returns
}


func NewMonitor_Override(m Monitor) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.Monitor",
		nil, // no parameters
		m,
	)
}

// Creates a Monitor from a CfnEnvironment.MonitorsProperty construct.
func Monitor_FromCfnMonitorsProperty(monitorsProperty *CfnEnvironment_MonitorsProperty) Monitor {
	_init_.Initialize()

	if err := validateMonitor_FromCfnMonitorsPropertyParameters(monitorsProperty); err != nil {
		panic(err)
	}
	var returns Monitor

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.Monitor",
		"fromCfnMonitorsProperty",
		[]interface{}{monitorsProperty},
		&returns,
	)

	return returns
}

// Creates a Monitor from a CloudWatch alarm.
//
// If the alarm role is not specified, a role will
// be generated.
func Monitor_FromCloudWatchAlarm(alarm awscloudwatch.IAlarm, alarmRole awsiam.IRoleRef) Monitor {
	_init_.Initialize()

	if err := validateMonitor_FromCloudWatchAlarmParameters(alarm); err != nil {
		panic(err)
	}
	var returns Monitor

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.Monitor",
		"fromCloudWatchAlarm",
		[]interface{}{alarm, alarmRole},
		&returns,
	)

	return returns
}

