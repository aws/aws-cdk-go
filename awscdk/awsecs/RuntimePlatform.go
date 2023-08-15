package awsecs


// The interface for Runtime Platform.
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
type RuntimePlatform struct {
	// The CpuArchitecture for Fargate Runtime Platform.
	// Default: - Undefined.
	//
	CpuArchitecture CpuArchitecture `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// The operating system for Fargate Runtime Platform.
	// Default: - Undefined.
	//
	OperatingSystemFamily OperatingSystemFamily `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}

