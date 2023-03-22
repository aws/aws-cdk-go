package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The CpuArchitecture for Fargate Runtime Platform.
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
type CpuArchitecture interface {
}

// The jsii proxy struct for CpuArchitecture
type jsiiProxy_CpuArchitecture struct {
	_ byte // padding
}

// Other cpu architecture.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-runtimeplatform.html#cfn-ecs-taskdefinition-runtimeplatform-cpuarchitecture for all available cpu architecture.
//
func CpuArchitecture_Of(cpuArchitecture *string) CpuArchitecture {
	_init_.Initialize()

	if err := validateCpuArchitecture_OfParameters(cpuArchitecture); err != nil {
		panic(err)
	}
	var returns CpuArchitecture

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CpuArchitecture",
		"of",
		[]interface{}{cpuArchitecture},
		&returns,
	)

	return returns
}

func CpuArchitecture_ARM64() CpuArchitecture {
	_init_.Initialize()
	var returns CpuArchitecture
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CpuArchitecture",
		"ARM64",
		&returns,
	)
	return returns
}

func CpuArchitecture_X86_64() CpuArchitecture {
	_init_.Initialize()
	var returns CpuArchitecture
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CpuArchitecture",
		"X86_64",
		&returns,
	)
	return returns
}

