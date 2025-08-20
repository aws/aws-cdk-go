package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Available versions for Java instrumentation.
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
type JavaInstrumentationVersion interface {
	InstrumentationVersion
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

// The jsii proxy struct for JavaInstrumentationVersion
type jsiiProxy_JavaInstrumentationVersion struct {
	jsiiProxy_InstrumentationVersion
}

func (j *jsiiProxy_JavaInstrumentationVersion) ImageRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaInstrumentationVersion) MemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaInstrumentationVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewJavaInstrumentationVersion(imageRepo *string, version *string, memoryLimit *float64) JavaInstrumentationVersion {
	_init_.Initialize()

	if err := validateNewJavaInstrumentationVersionParameters(imageRepo, version, memoryLimit); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaInstrumentationVersion{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentationVersion",
		[]interface{}{imageRepo, version, memoryLimit},
		&j,
	)

	return &j
}

// Experimental.
func NewJavaInstrumentationVersion_Override(j JavaInstrumentationVersion, imageRepo *string, version *string, memoryLimit *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentationVersion",
		[]interface{}{imageRepo, version, memoryLimit},
		j,
	)
}

func JavaInstrumentationVersion_DEFAULT_MEMORY_LIMIT_MIB() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentationVersion",
		"DEFAULT_MEMORY_LIMIT_MIB",
		&returns,
	)
	return returns
}

func JavaInstrumentationVersion_IMAGE_REPO() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentationVersion",
		"IMAGE_REPO",
		&returns,
	)
	return returns
}

func JavaInstrumentationVersion_V1_32_6() JavaInstrumentationVersion {
	_init_.Initialize()
	var returns JavaInstrumentationVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentationVersion",
		"V1_32_6",
		&returns,
	)
	return returns
}

func JavaInstrumentationVersion_V1_33_0() JavaInstrumentationVersion {
	_init_.Initialize()
	var returns JavaInstrumentationVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentationVersion",
		"V1_33_0",
		&returns,
	)
	return returns
}

func JavaInstrumentationVersion_V2_10_0() JavaInstrumentationVersion {
	_init_.Initialize()
	var returns JavaInstrumentationVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentationVersion",
		"V2_10_0",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaInstrumentationVersion) ImageURI() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"imageURI",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaInstrumentationVersion) MemoryLimitMiB() *float64 {
	var returns *float64

	_jsii_.Invoke(
		j,
		"memoryLimitMiB",
		nil, // no parameters
		&returns,
	)

	return returns
}

