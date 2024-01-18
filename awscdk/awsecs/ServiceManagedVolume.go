package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a service-managed volume and always configured at launch.
//
// Example:
//   var cluster cluster
//
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"))
//
//   container := taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	PortMappings: []portMapping{
//   		&portMapping{
//   			ContainerPort: jsii.Number(80),
//   			Protocol: ecs.Protocol_TCP,
//   		},
//   	},
//   })
//
//   volume := ecs.NewServiceManagedVolume(this, jsii.String("EBSVolume"), &ServiceManagedVolumeProps{
//   	Name: jsii.String("ebs1"),
//   	ManagedEBSVolume: &ServiceManagedEBSVolumeConfiguration{
//   		Size: awscdk.Size_Gibibytes(jsii.Number(15)),
//   		VolumeType: ec2.EbsDeviceVolumeType_GP3,
//   		FileSystemType: ecs.FileSystemType_XFS,
//   		TagSpecifications: []eBSTagSpecification{
//   			&eBSTagSpecification{
//   				Tags: map[string]*string{
//   					"purpose": jsii.String("production"),
//   				},
//   				PropagateTags: ecs.EbsPropagatedTagSource_SERVICE,
//   			},
//   		},
//   	},
//   })
//
//   volume.MountIn(container, &ContainerMountPoint{
//   	ContainerPath: jsii.String("/var/lib"),
//   	ReadOnly: jsii.Boolean(false),
//   })
//
//   taskDefinition.AddVolume(volume)
//
//   service := ecs.NewFargateService(this, jsii.String("FargateService"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   })
//
//   service.AddVolume(volume)
//
type ServiceManagedVolume interface {
	constructs.Construct
	// Volume configuration.
	Config() *ServiceManagedEBSVolumeConfiguration
	// configuredAtLaunch indicates volume at launch time, referenced by taskdefinition volume.
	ConfiguredAtLaunch() *bool
	// Name of the volume, referenced by taskdefintion and mount point.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// An IAM role that allows ECS to make calls to EBS APIs.
	//
	// If not provided, a new role with appropriate permissions will be created by default.
	Role() awsiam.IRole
	// Mounts the service managed volume to a specified container at a defined mount point.
	MountIn(container ContainerDefinition, mountPoint *ContainerMountPoint)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ServiceManagedVolume
type jsiiProxy_ServiceManagedVolume struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ServiceManagedVolume) Config() *ServiceManagedEBSVolumeConfiguration {
	var returns *ServiceManagedEBSVolumeConfiguration
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceManagedVolume) ConfiguredAtLaunch() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"configuredAtLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceManagedVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceManagedVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceManagedVolume) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}


func NewServiceManagedVolume(scope constructs.Construct, id *string, props *ServiceManagedVolumeProps) ServiceManagedVolume {
	_init_.Initialize()

	if err := validateNewServiceManagedVolumeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceManagedVolume{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ServiceManagedVolume",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewServiceManagedVolume_Override(s ServiceManagedVolume, scope constructs.Construct, id *string, props *ServiceManagedVolumeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ServiceManagedVolume",
		[]interface{}{scope, id, props},
		s,
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
func ServiceManagedVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceManagedVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ServiceManagedVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceManagedVolume) MountIn(container ContainerDefinition, mountPoint *ContainerMountPoint) {
	if err := s.validateMountInParameters(container, mountPoint); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"mountIn",
		[]interface{}{container, mountPoint},
	)
}

func (s *jsiiProxy_ServiceManagedVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

