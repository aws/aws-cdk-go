package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Base class for instrumentation versions.
//
// Provides functionality to generate image URIs for different instrumentation types.
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
//   		PortMappings: []portMapping{
//   			&portMapping{
//   				ContainerPort: jsii.Number(4316),
//   				HostPort: jsii.Number(4316),
//   			},
//   			&portMapping{
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
type InstrumentationVersion interface {
	// Experimental.
	ImageRepo() *string
	// Experimental.
	MemoryLimit() *float64
	// Experimental.
	Version() *string
	// Get the image URI for the instrumentation version.
	//
	// Returns: The image URI.
	// Experimental.
	ImageURI() *string
	// Get the memory limit in MiB for the instrumentation version.
	//
	// Returns: The memory limit.
	// Experimental.
	MemoryLimitMiB() *float64
}

// The jsii proxy struct for InstrumentationVersion
type jsiiProxy_InstrumentationVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_InstrumentationVersion) ImageRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstrumentationVersion) MemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstrumentationVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewInstrumentationVersion_Override(i InstrumentationVersion, imageRepo *string, version *string, memoryLimit *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.InstrumentationVersion",
		[]interface{}{imageRepo, version, memoryLimit},
		i,
	)
}

func (i *jsiiProxy_InstrumentationVersion) ImageURI() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"imageURI",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstrumentationVersion) MemoryLimitMiB() *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"memoryLimitMiB",
		nil, // no parameters
		&returns,
	)

	return returns
}

