package awsecs


// The interface for Runtime Platform.
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
type RuntimePlatform struct {
	// The CpuArchitecture for Fargate Runtime Platform.
	CpuArchitecture CpuArchitecture `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// The operating system for Fargate Runtime Platform.
	OperatingSystemFamily OperatingSystemFamily `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}

