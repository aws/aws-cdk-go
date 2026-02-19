package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A construct that adds CloudWatch Agent as a container to an ECS task definition.
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
type CloudWatchAgentIntegration interface {
	constructs.Construct
	// The CloudWatch Agent container definition.
	// Experimental.
	AgentContainer() awsecs.ContainerDefinition
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

// The jsii proxy struct for CloudWatchAgentIntegration
type jsiiProxy_CloudWatchAgentIntegration struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CloudWatchAgentIntegration) AgentContainer() awsecs.ContainerDefinition {
	var returns awsecs.ContainerDefinition
	_jsii_.Get(
		j,
		"agentContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudWatchAgentIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates a new CloudWatch Agent integration.
// Experimental.
func NewCloudWatchAgentIntegration(scope constructs.Construct, id *string, props *CloudWatchAgentIntegrationProps) CloudWatchAgentIntegration {
	_init_.Initialize()

	if err := validateNewCloudWatchAgentIntegrationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudWatchAgentIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new CloudWatch Agent integration.
// Experimental.
func NewCloudWatchAgentIntegration_Override(c CloudWatchAgentIntegration, scope constructs.Construct, id *string, props *CloudWatchAgentIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentIntegration",
		[]interface{}{scope, id, props},
		c,
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
func CloudWatchAgentIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudWatchAgentIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudWatchAgentIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudWatchAgentIntegration) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

