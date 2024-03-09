package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The operating system for Fargate Runtime Platform.
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
type OperatingSystemFamily interface {
	// Returns true if the operating system family is Windows.
	IsWindows() *bool
}

// The jsii proxy struct for OperatingSystemFamily
type jsiiProxy_OperatingSystemFamily struct {
	_ byte // padding
}

// Other operating system family.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-runtimeplatform.html#cfn-ecs-taskdefinition-runtimeplatform-operatingsystemfamily for all available operating system family.
//
func OperatingSystemFamily_Of(family *string) OperatingSystemFamily {
	_init_.Initialize()

	if err := validateOperatingSystemFamily_OfParameters(family); err != nil {
		panic(err)
	}
	var returns OperatingSystemFamily

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"of",
		[]interface{}{family},
		&returns,
	)

	return returns
}

func OperatingSystemFamily_LINUX() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"LINUX",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2004_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2004_CORE",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2016_FULL() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2016_FULL",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2019_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2019_CORE",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2019_FULL() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2019_FULL",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2022_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2022_CORE",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2022_FULL() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2022_FULL",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_20H2_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_20H2_CORE",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OperatingSystemFamily) IsWindows() *bool {
	var returns *bool

	_jsii_.Invoke(
		o,
		"isWindows",
		nil, // no parameters
		&returns,
	)

	return returns
}

