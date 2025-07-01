package awscdkapplicationsignalsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Configuration options for the CloudWatch Agent container.
//
// Example:
//   import "github.com/aws/constructs-go/constructs"
//   import "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   type myStack struct {
//   	stack
//   }
//
//   func newMyStack(scope construct, id *string, props stackProps) *myStack {
//   	if props == nil {
//   		props = &StackProps{
//   		}
//   	}
//   	this := &myStack{}
//   	cdk.NewStack_Override(this, )
//   	vpc := ec2.NewVpc(this, jsii.String("TestVpc"), &VpcProps{
//   	})
//   	cluster := ecs.NewCluster(this, jsii.String("TestCluster"), &ClusterProps{
//   		Vpc: Vpc,
//   	})
//
//   	fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("SampleAppTaskDefinition"), &FargateTaskDefinitionProps{
//   		Cpu: jsii.Number(2048),
//   		MemoryLimitMiB: jsii.Number(4096),
//   	})
//
//   	fargateTaskDefinition.AddContainer(jsii.String("app"), &ContainerDefinitionOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("test/sample-app")),
//   	})
//
//   	appsignals.NewApplicationSignalsIntegration(this, jsii.String("ApplicationSignalsIntegration"), &ApplicationSignalsIntegrationProps{
//   		TaskDefinition: fargateTaskDefinition,
//   		Instrumentation: &InstrumentationProps{
//   			SdkVersion: appsignals.JavaInstrumentationVersion_V2_10_0(),
//   		},
//   		ServiceName: jsii.String("sample-app"),
//   		CloudWatchAgentSidecar: &CloudWatchAgentOptions{
//   			ContainerName: jsii.String("cloudwatch-agent"),
//   			EnableLogging: jsii.Boolean(true),
//   			Cpu: jsii.Number(256),
//   			MemoryLimitMiB: jsii.Number(512),
//   		},
//   	})
//
//   	ecs.NewFargateService(this, jsii.String("MySampleApp"), &FargateServiceProps{
//   		Cluster: cluster,
//   		TaskDefinition: fargateTaskDefinition,
//   		DesiredCount: jsii.Number(1),
//   	})
//   	return this
//   }
//
// Experimental.
type CloudWatchAgentOptions struct {
	// Name of the CloudWatch Agent container.
	// Experimental.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Custom agent configuration in JSON format.
	// Default: - Uses default configuration for Application Signals.
	//
	// Experimental.
	AgentConfig *string `field:"optional" json:"agentConfig" yaml:"agentConfig"`
	// The minimum number of CPU units to reserve for the container.
	// Default: - No minimum CPU units reserved.
	//
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Whether to enable logging for the CloudWatch Agent.
	// Default: - false.
	//
	// Experimental.
	EnableLogging *bool `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// Start as an essential container.
	// Default: - true.
	//
	// Experimental.
	Essential *bool `field:"optional" json:"essential" yaml:"essential"`
	// The amount (in MiB) of memory to present to the container.
	// Default: - No memory limit.
	//
	// Experimental.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	// Default: - No memory reserved.
	//
	// Experimental.
	MemoryReservationMiB *float64 `field:"optional" json:"memoryReservationMiB" yaml:"memoryReservationMiB"`
	// Operating system family for the CloudWatch Agent.
	// Default: - Linux.
	//
	// Experimental.
	OperatingSystemFamily awsecs.OperatingSystemFamily `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
	// The port mappings to add to the container definition.
	// Default: - No ports are mapped.
	//
	// Experimental.
	PortMappings *[]*awsecs.PortMapping `field:"optional" json:"portMappings" yaml:"portMappings"`
}

