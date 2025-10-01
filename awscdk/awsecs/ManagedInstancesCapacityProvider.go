package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Managed Instances Capacity Provider.
//
// This allows an ECS cluster to use
// Managed Instances for task placement with managed infrastructure.
//
// Example:
//   var vpc vpc
//   var infrastructureRole role
//   var instanceProfile instanceProfile
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   // Create a Managed Instances Capacity Provider
//   miCapacityProvider := ecs.NewManagedInstancesCapacityProvider(this, jsii.String("MICapacityProvider"), &ManagedInstancesCapacityProviderProps{
//   	InfrastructureRole: InfrastructureRole,
//   	Ec2InstanceProfile: instanceProfile,
//   	Subnets: vpc.PrivateSubnets,
//   	SecurityGroups: []iSecurityGroup{
//   		ec2.NewSecurityGroup(this, jsii.String("MISecurityGroup"), &SecurityGroupProps{
//   			Vpc: *Vpc,
//   		}),
//   	},
//   	InstanceRequirements: &InstanceRequirementsConfig{
//   		VCpuCountMin: jsii.Number(1),
//   		MemoryMin: awscdk.Size_Gibibytes(jsii.Number(2)),
//   		CpuManufacturers: []cpuManufacturer{
//   			ec2.*cpuManufacturer_INTEL,
//   		},
//   		AcceleratorManufacturers: []acceleratorManufacturer{
//   			ec2.*acceleratorManufacturer_NVIDIA,
//   		},
//   	},
//   	PropagateTags: ecs.PropagateManagedInstancesTags_CAPACITY_PROVIDER,
//   })
//
//   // Add the capacity provider to the cluster
//   cluster.AddManagedInstancesCapacityProvider(miCapacityProvider)
//
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryReservationMiB: jsii.Number(256),
//   })
//
//   ecs.NewEc2Service(this, jsii.String("EC2Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	MinHealthyPercent: jsii.Number(100),
//   	CapacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			CapacityProvider: miCapacityProvider.CapacityProviderName,
//   			Weight: jsii.Number(1),
//   		},
//   	},
//   })
//
type ManagedInstancesCapacityProvider interface {
	constructs.Construct
	// Capacity provider name.
	CapacityProviderName() *string
	// The tree node.
	Node() constructs.Node
	// Associates the capacity provider with the specified cluster.
	//
	// This method is called by the cluster when adding the capacity provider.
	Bind(cluster ICluster)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ManagedInstancesCapacityProvider
type jsiiProxy_ManagedInstancesCapacityProvider struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ManagedInstancesCapacityProvider) CapacityProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedInstancesCapacityProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewManagedInstancesCapacityProvider(scope constructs.Construct, id *string, props *ManagedInstancesCapacityProviderProps) ManagedInstancesCapacityProvider {
	_init_.Initialize()

	if err := validateNewManagedInstancesCapacityProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedInstancesCapacityProvider{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ManagedInstancesCapacityProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewManagedInstancesCapacityProvider_Override(m ManagedInstancesCapacityProvider, scope constructs.Construct, id *string, props *ManagedInstancesCapacityProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ManagedInstancesCapacityProvider",
		[]interface{}{scope, id, props},
		m,
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
func ManagedInstancesCapacityProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedInstancesCapacityProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ManagedInstancesCapacityProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ManagedInstancesCapacityProvider_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.ManagedInstancesCapacityProvider",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ManagedInstancesCapacityProvider) Bind(cluster ICluster) {
	if err := m.validateBindParameters(cluster); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"bind",
		[]interface{}{cluster},
	)
}

func (m *jsiiProxy_ManagedInstancesCapacityProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

