package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The CpuArchitecture for Fargate Runtime Platform.
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
// Experimental.
type CpuArchitecture interface {
}

// The jsii proxy struct for CpuArchitecture
type jsiiProxy_CpuArchitecture struct {
	_ byte // padding
}

// Other cpu architecture.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-runtimeplatform.html#cfn-ecs-taskdefinition-runtimeplatform-cpuarchitecture for all available cpu architecture.
//
// Experimental.
func CpuArchitecture_Of(cpuArchitecture *string) CpuArchitecture {
	_init_.Initialize()

	if err := validateCpuArchitecture_OfParameters(cpuArchitecture); err != nil {
		panic(err)
	}
	var returns CpuArchitecture

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.CpuArchitecture",
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
		"monocdk.aws_ecs.CpuArchitecture",
		"ARM64",
		&returns,
	)
	return returns
}

func CpuArchitecture_X86_64() CpuArchitecture {
	_init_.Initialize()
	var returns CpuArchitecture
	_jsii_.StaticGet(
		"monocdk.aws_ecs.CpuArchitecture",
		"X86_64",
		&returns,
	)
	return returns
}

