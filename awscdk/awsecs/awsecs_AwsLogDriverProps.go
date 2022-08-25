package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Specifies the awslogs log driver configuration options.
//
// Example:
//   // Create a Task Definition for the Windows container to start
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
//   	runtimePlatform: &runtimePlatform{
//   		operatingSystemFamily: ecs.operatingSystemFamily_WINDOWS_SERVER_2019_CORE(),
//   		cpuArchitecture: ecs.cpuArchitecture_X86_64(),
//   	},
//   	cpu: jsii.Number(1024),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   taskDefinition.addContainer(jsii.String("windowsservercore"), &containerDefinitionOptions{
//   	logging: ecs.logDriver.awsLogs(&awsLogDriverProps{
//   		streamPrefix: jsii.String("win-iis-on-fargate"),
//   	}),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(80),
//   		},
//   	},
//   	image: ecs.containerImage.fromRegistry(jsii.String("mcr.microsoft.com/windows/servercore/iis:windowsservercore-ltsc2019")),
//   })
//
type AwsLogDriverProps struct {
	// Prefix for the log streams.
	//
	// The awslogs-stream-prefix option allows you to associate a log stream
	// with the specified prefix, the container name, and the ID of the Amazon
	// ECS task to which the container belongs. If you specify a prefix with
	// this option, then the log stream takes the following format:
	//
	// prefix-name/container-name/ecs-task-id.
	StreamPrefix *string `field:"required" json:"streamPrefix" yaml:"streamPrefix"`
	// This option defines a multiline start pattern in Python strftime format.
	//
	// A log message consists of a line that matches the pattern and any
	// following lines that don’t match the pattern. Thus the matched line is
	// the delimiter between log messages.
	DatetimeFormat *string `field:"optional" json:"datetimeFormat" yaml:"datetimeFormat"`
	// The log group to log to.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The number of days log events are kept in CloudWatch Logs when the log group is automatically created by this construct.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The delivery mode of log messages from the container to awslogs.
	Mode AwsLogDriverMode `field:"optional" json:"mode" yaml:"mode"`
	// This option defines a multiline start pattern using a regular expression.
	//
	// A log message consists of a line that matches the pattern and any
	// following lines that don’t match the pattern. Thus the matched line is
	// the delimiter between log messages.
	//
	// This option is ignored if datetimeFormat is also configured.
	MultilinePattern *string `field:"optional" json:"multilinePattern" yaml:"multilinePattern"`
}

