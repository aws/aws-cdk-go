package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The base class for log drivers.
//
// Example:
//   var secret secret
//
//
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
//   	MemoryLimitMiB: jsii.Number(256),
//   	Logging: ecs.LogDrivers_Splunk(&SplunkLogDriverProps{
//   		SecretToken: secret,
//   		Url: jsii.String("my-splunk-url"),
//   	}),
//   })
//
type LogDrivers interface {
}

// The jsii proxy struct for LogDrivers
type jsiiProxy_LogDrivers struct {
	_ byte // padding
}

func NewLogDrivers() LogDrivers {
	_init_.Initialize()

	j := jsiiProxy_LogDrivers{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewLogDrivers_Override(l LogDrivers) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		nil, // no parameters
		l,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func LogDrivers_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateLogDrivers_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to firelens log router.
//
// For detail configurations, please refer to Amazon ECS FireLens Examples:
// https://github.com/aws-samples/amazon-ecs-firelens-examples
func LogDrivers_Firelens(props *FireLensLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateLogDrivers_FirelensParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"firelens",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to fluentd Logs.
func LogDrivers_Fluentd(props *FluentdLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateLogDrivers_FluentdParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"fluentd",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to gelf Logs.
func LogDrivers_Gelf(props *GelfLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateLogDrivers_GelfParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"gelf",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to journald Logs.
func LogDrivers_Journald(props *JournaldLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateLogDrivers_JournaldParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"journald",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to json-file Logs.
func LogDrivers_JsonFile(props *JsonFileLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateLogDrivers_JsonFileParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"jsonFile",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to splunk Logs.
func LogDrivers_Splunk(props *SplunkLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateLogDrivers_SplunkParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"splunk",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a log driver configuration that sends log information to syslog Logs.
func LogDrivers_Syslog(props *SyslogLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateLogDrivers_SyslogParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDrivers",
		"syslog",
		[]interface{}{props},
		&returns,
	)

	return returns
}

