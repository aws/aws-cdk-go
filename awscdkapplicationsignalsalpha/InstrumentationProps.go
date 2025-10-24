package awscdkapplicationsignalsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Interface for instrumentation properties.
//
// Example:
//   import "github.com/aws/constructs-go/constructs"
//   import "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   type myStack struct {
//   	Stack
//   }
//
//   func newMyStack(scope Construct, id *string, props StackProps) *myStack {
//   	if props == nil {
//   		props = &StackProps{
//   		}
//   	}
//   	this := &myStack{}
//   	cdk.NewStack_Override(this, scope, id, props)
//
//   	vpc := ec2.NewVpc(this, jsii.String("TestVpc"), &VpcProps{
//   	})
//   	cluster := ecs.NewCluster(this, jsii.String("TestCluster"), &ClusterProps{
//   		Vpc: Vpc,
//   	})
//
//   	// Define Task Definition for CloudWatch agent (Daemon)
//   	cwAgentTaskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("CloudWatchAgentTaskDefinition"), &Ec2TaskDefinitionProps{
//   		NetworkMode: ecs.NetworkMode_HOST,
//   	})
//
//   	appsignals.NewCloudWatchAgentIntegration(this, jsii.String("CloudWatchAgentIntegration"), &CloudWatchAgentIntegrationProps{
//   		TaskDefinition: cwAgentTaskDefinition,
//   		ContainerName: jsii.String("ecs-cwagent"),
//   		EnableLogging: jsii.Boolean(false),
//   		Cpu: jsii.Number(128),
//   		MemoryLimitMiB: jsii.Number(64),
//   		PortMappings: []PortMapping{
//   			&PortMapping{
//   				ContainerPort: jsii.Number(4316),
//   				HostPort: jsii.Number(4316),
//   			},
//   			&PortMapping{
//   				ContainerPort: jsii.Number(2000),
//   				HostPort: jsii.Number(2000),
//   			},
//   		},
//   	})
//
//   	// Create the CloudWatch Agent daemon service
//   	// Create the CloudWatch Agent daemon service
//   	ecs.NewEc2Service(this, jsii.String("CloudWatchAgentDaemon"), &Ec2ServiceProps{
//   		Cluster: Cluster,
//   		TaskDefinition: cwAgentTaskDefinition,
//   		Daemon: jsii.Boolean(true),
//   	})
//
//   	// Define Task Definition for user application
//   	sampleAppTaskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("SampleAppTaskDefinition"), &Ec2TaskDefinitionProps{
//   		NetworkMode: ecs.NetworkMode_HOST,
//   	})
//
//   	sampleAppTaskDefinition.AddContainer(jsii.String("app"), &ContainerDefinitionOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("test/sample-app")),
//   		Cpu: jsii.Number(0),
//   		MemoryLimitMiB: jsii.Number(512),
//   	})
//
//   	// No CloudWatch Agent side car is needed as application container communicates to CloudWatch Agent daemon through host network
//   	// No CloudWatch Agent side car is needed as application container communicates to CloudWatch Agent daemon through host network
//   	appsignals.NewApplicationSignalsIntegration(this, jsii.String("ApplicationSignalsIntegration"), &ApplicationSignalsIntegrationProps{
//   		TaskDefinition: sampleAppTaskDefinition,
//   		Instrumentation: &InstrumentationProps{
//   			SdkVersion: appsignals.PythonInstrumentationVersion_V0_8_0(),
//   		},
//   		ServiceName: jsii.String("sample-app"),
//   	})
//
//   	ecs.NewEc2Service(this, jsii.String("MySampleApp"), &Ec2ServiceProps{
//   		Cluster: Cluster,
//   		TaskDefinition: sampleAppTaskDefinition,
//   		DesiredCount: jsii.Number(1),
//   	})
//   	return this
//   }
//
// Experimental.
type InstrumentationProps struct {
	// The version of the instrumentation.
	// Experimental.
	SdkVersion InstrumentationVersion `field:"required" json:"sdkVersion" yaml:"sdkVersion"`
	// The runtime platform of the instrumentation.
	// Default: - the runtime platform specified through the input TaskDefinition.
	//
	// Experimental.
	RuntimePlatform *awsecs.RuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
}

