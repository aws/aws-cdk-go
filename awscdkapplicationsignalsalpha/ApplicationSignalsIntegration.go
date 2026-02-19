package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Class for integrating Application Signals into an ECS task definition.
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
type ApplicationSignalsIntegration interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for ApplicationSignalsIntegration
type jsiiProxy_ApplicationSignalsIntegration struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApplicationSignalsIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewApplicationSignalsIntegration(scope constructs.Construct, id *string, props *ApplicationSignalsIntegrationProps) ApplicationSignalsIntegration {
	_init_.Initialize()

	if err := validateNewApplicationSignalsIntegrationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationSignalsIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.ApplicationSignalsIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApplicationSignalsIntegration_Override(a ApplicationSignalsIntegration, scope constructs.Construct, id *string, props *ApplicationSignalsIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.ApplicationSignalsIntegration",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func ApplicationSignalsIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationSignalsIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-applicationsignals-alpha.ApplicationSignalsIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSignalsIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSignalsIntegration) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

