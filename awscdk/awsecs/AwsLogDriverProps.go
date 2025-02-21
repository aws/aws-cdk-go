package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Specifies the awslogs log driver configuration options.
//
// Example:
//   // Create a Task Definition for the Windows container to start
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &FargateTaskDefinitionProps{
//   	RuntimePlatform: &RuntimePlatform{
//   		OperatingSystemFamily: ecs.OperatingSystemFamily_WINDOWS_SERVER_2019_CORE(),
//   		CpuArchitecture: ecs.CpuArchitecture_X86_64(),
//   	},
//   	Cpu: jsii.Number(1024),
//   	MemoryLimitMiB: jsii.Number(2048),
//   })
//
//   taskDefinition.AddContainer(jsii.String("windowsservercore"), &ContainerDefinitionOptions{
//   	Logging: ecs.LogDriver_AwsLogs(&AwsLogDriverProps{
//   		StreamPrefix: jsii.String("win-iis-on-fargate"),
//   	}),
//   	PortMappings: []portMapping{
//   		&portMapping{
//   			ContainerPort: jsii.Number(80),
//   		},
//   	},
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("mcr.microsoft.com/windows/servercore/iis:windowsservercore-ltsc2019")),
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
	// Default: - No multiline matching.
	//
	DatetimeFormat *string `field:"optional" json:"datetimeFormat" yaml:"datetimeFormat"`
	// The log group to log to.
	// Default: - A log group is automatically created.
	//
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The number of days log events are kept in CloudWatch Logs when the log group is automatically created by this construct.
	// Default: - Logs never expire.
	//
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// When AwsLogDriverMode.NON_BLOCKING is configured, this parameter controls the size of the non-blocking buffer used to temporarily store messages. This parameter is not valid with AwsLogDriverMode.BLOCKING.
	// Default: - 1 megabyte if driver mode is non-blocking, otherwise this property is not set.
	//
	MaxBufferSize awscdk.Size `field:"optional" json:"maxBufferSize" yaml:"maxBufferSize"`
	// The delivery mode of log messages from the container to awslogs.
	// Default: - AwsLogDriverMode.BLOCKING
	//
	Mode AwsLogDriverMode `field:"optional" json:"mode" yaml:"mode"`
	// This option defines a multiline start pattern using a regular expression.
	//
	// A log message consists of a line that matches the pattern and any
	// following lines that don’t match the pattern. Thus the matched line is
	// the delimiter between log messages.
	//
	// This option is ignored if datetimeFormat is also configured.
	// Default: - No multiline matching.
	//
	MultilinePattern *string `field:"optional" json:"multilinePattern" yaml:"multilinePattern"`
}

