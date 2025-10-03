package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A ManagedComputeEnvironment that uses ECS orchestration on EC2 instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var instanceType instanceType
//   var launchTemplate launchTemplate
//   var machineImage iMachineImage
//   var placementGroup placementGroup
//   var role role
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   managedEc2EksComputeEnvironment := awscdk.Aws_batch.NewManagedEc2EksComputeEnvironment(this, jsii.String("MyManagedEc2EksComputeEnvironment"), &ManagedEc2EksComputeEnvironmentProps{
//   	EksCluster: cluster,
//   	KubernetesNamespace: jsii.String("kubernetesNamespace"),
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	AllocationStrategy: awscdk.*Aws_batch.AllocationStrategy_BEST_FIT,
//   	ComputeEnvironmentName: jsii.String("computeEnvironmentName"),
//   	Enabled: jsii.Boolean(false),
//   	Images: []eksMachineImage{
//   		&eksMachineImage{
//   			Image: machineImage,
//   			ImageType: awscdk.*Aws_batch.EksMachineImageType_EKS_AL2,
//   		},
//   	},
//   	InstanceClasses: []instanceClass{
//   		awscdk.Aws_ec2.*instanceClass_STANDARD3,
//   	},
//   	InstanceRole: role,
//   	InstanceTypes: []*instanceType{
//   		instanceType,
//   	},
//   	LaunchTemplate: launchTemplate,
//   	MaxvCpus: jsii.Number(123),
//   	MinvCpus: jsii.Number(123),
//   	PlacementGroup: placementGroup,
//   	ReplaceComputeEnvironment: jsii.Boolean(false),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	ServiceRole: role,
//   	Spot: jsii.Boolean(false),
//   	SpotBidPercentage: jsii.Number(123),
//   	TerminateOnUpdate: jsii.Boolean(false),
//   	UpdateTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	UpdateToLatestImageVersion: jsii.Boolean(false),
//   	UseOptimalInstanceClasses: jsii.Boolean(false),
//   	VpcSubnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.*Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   })
//
type ManagedEc2EksComputeEnvironment interface {
	awscdk.Resource
	IComputeEnvironment
	IManagedComputeEnvironment
	// The allocation strategy to use if not enough instances of the best fitting instance type can be allocated.
	AllocationStrategy() AllocationStrategy
	// The ARN of this compute environment.
	ComputeEnvironmentArn() *string
	// The name of the ComputeEnvironment.
	ComputeEnvironmentName() *string
	// The network connections associated with this resource.
	Connections() awsec2.Connections
	// The cluster that backs this Compute Environment. Required for Compute Environments running Kubernetes jobs.
	//
	// Please ensure that you have followed the steps at
	//
	// https://docs.aws.amazon.com/batch/latest/userguide/getting-started-eks.html
	//
	// before attempting to deploy a `ManagedEc2EksComputeEnvironment` that uses this cluster.
	// If you do not follow the steps in the link, the deployment fail with a message that the
	// compute environment did not stabilize.
	EksCluster() awseks.ICluster
	// Whether or not this ComputeEnvironment can accept jobs from a Queue.
	//
	// Enabled ComputeEnvironments can accept jobs from a Queue and
	// can scale instances up or down.
	// Disabled ComputeEnvironments cannot accept jobs from a Queue or
	// scale instances up or down.
	//
	// If you change a ComputeEnvironment from enabled to disabled while it is executing jobs,
	// Jobs in the `STARTED` or `RUNNING` states will not
	// be interrupted. As jobs complete, the ComputeEnvironment will scale instances down to `minvCpus`.
	//
	// To ensure you aren't billed for unused capacity, set `minvCpus` to `0`.
	Enabled() *bool
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Configure which AMIs this Compute Environment can launch.
	Images() *[]*EksMachineImage
	// The instance types that this Compute Environment can launch.
	//
	// Which one is chosen depends on the `AllocationStrategy` used.
	InstanceClasses() *[]awsec2.InstanceClass
	// The execution Role that instances launched by this Compute Environment will use.
	InstanceRole() awsiam.IRole
	// The instance types that this Compute Environment can launch.
	//
	// Which one is chosen depends on the `AllocationStrategy` used.
	InstanceTypes() *[]awsec2.InstanceType
	// The namespace of the Cluster.
	//
	// Cannot be 'default', start with 'kube-', or be longer than 64 characters.
	KubernetesNamespace() *string
	// The Launch Template that this Compute Environment will use to provision EC2 Instances.
	//
	// *Note*: if `securityGroups` is specified on both your
	// launch template and this Compute Environment, **the
	// `securityGroup`s on the Compute Environment override the
	// ones on the launch template.
	LaunchTemplate() awsec2.ILaunchTemplate
	// The maximum vCpus this `ManagedComputeEnvironment` can scale up to.
	//
	// *Note*: if this Compute Environment uses EC2 resources (not Fargate) with either `AllocationStrategy.BEST_FIT_PROGRESSIVE` or
	// `AllocationStrategy.SPOT_CAPACITY_OPTIMIZED`, or `AllocationStrategy.BEST_FIT` with Spot instances,
	// The scheduler may exceed this number by at most one of the instances specified in `instanceTypes`
	// or `instanceClasses`.
	MaxvCpus() *float64
	// The minimum vCPUs that an environment should maintain, even if the compute environment is DISABLED.
	MinvCpus() *float64
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The EC2 placement group to associate with your compute resources.
	//
	// If you intend to submit multi-node parallel jobs to this Compute Environment,
	// you should consider creating a cluster placement group and associate it with your compute resources.
	// This keeps your multi-node parallel job on a logical grouping of instances
	// within a single Availability Zone with high network flow potential.
	PlacementGroup() awsec2.IPlacementGroup
	// Specifies whether this Compute Environment is replaced if an update is made that requires replacing its instances.
	//
	// To enable more properties to be updated,
	// set this property to `false`. When changing the value of this property to false,
	// do not change any other properties at the same time.
	// If other properties are changed at the same time,
	// and the change needs to be rolled back but it can't,
	// it's possible for the stack to go into the UPDATE_ROLLBACK_FAILED state.
	// You can't update a stack that is in the UPDATE_ROLLBACK_FAILED state.
	// However, if you can continue to roll it back,
	// you can return the stack to its original settings and then try to update it again.
	//
	// The properties which require a replacement of the Compute Environment are:.
	ReplaceComputeEnvironment() *bool
	// The security groups this Compute Environment will launch instances in.
	SecurityGroups() *[]awsec2.ISecurityGroup
	// The role Batch uses to perform actions on your behalf in your account, such as provision instances to run your jobs.
	ServiceRole() awsiam.IRole
	// Whether or not to use spot instances.
	//
	// Spot instances are less expensive EC2 instances that can be
	// reclaimed by EC2 at any time; your job will be given two minutes
	// of notice before reclamation.
	Spot() *bool
	// The maximum percentage that a Spot Instance price can be when compared with the On-Demand price for that instance type before instances are launched.
	//
	// For example, if your maximum percentage is 20%, the Spot price must be
	// less than 20% of the current On-Demand price for that Instance.
	// You always pay the lowest market price and never more than your maximum percentage.
	// For most use cases, Batch recommends leaving this field empty.
	//
	// Implies `spot == true` if set.
	SpotBidPercentage() *float64
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// TagManager to set, remove and format tags.
	Tags() awscdk.TagManager
	// Whether or not any running jobs will be immediately terminated when an infrastructure update occurs.
	//
	// If this is enabled, any terminated jobs may be retried, depending on the job's
	// retry policy.
	TerminateOnUpdate() *bool
	// Only meaningful if `terminateOnUpdate` is `false`.
	//
	// If so,
	// when an infrastructure update is triggered, any running jobs
	// will be allowed to run until `updateTimeout` has expired.
	UpdateTimeout() awscdk.Duration
	// Whether or not the AMI is updated to the latest one supported by Batch when an infrastructure update occurs.
	//
	// If you specify a specific AMI, this property will be ignored.
	//
	// Note: the CDK will never set this value by default, `false` will set by CFN.
	// This is to avoid a deployment failure that occurs when this value is set.
	UpdateToLatestImageVersion() *bool
	// Add an instance class to this compute environment.
	AddInstanceClass(instanceClass awsec2.InstanceClass)
	// Add an instance type to this compute environment.
	AddInstanceType(instanceType awsec2.InstanceType)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ManagedEc2EksComputeEnvironment
type jsiiProxy_ManagedEc2EksComputeEnvironment struct {
	internal.Type__awscdkResource
	jsiiProxy_IComputeEnvironment
	jsiiProxy_IManagedComputeEnvironment
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) AllocationStrategy() AllocationStrategy {
	var returns AllocationStrategy
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) ComputeEnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) ComputeEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) EksCluster() awseks.ICluster {
	var returns awseks.ICluster
	_jsii_.Get(
		j,
		"eksCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) Images() *[]*EksMachineImage {
	var returns *[]*EksMachineImage
	_jsii_.Get(
		j,
		"images",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) InstanceClasses() *[]awsec2.InstanceClass {
	var returns *[]awsec2.InstanceClass
	_jsii_.Get(
		j,
		"instanceClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) InstanceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"instanceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) InstanceTypes() *[]awsec2.InstanceType {
	var returns *[]awsec2.InstanceType
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) KubernetesNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) LaunchTemplate() awsec2.ILaunchTemplate {
	var returns awsec2.ILaunchTemplate
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) MaxvCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxvCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) MinvCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minvCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) PlacementGroup() awsec2.IPlacementGroup {
	var returns awsec2.IPlacementGroup
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) ReplaceComputeEnvironment() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"replaceComputeEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) Spot() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"spot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) SpotBidPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotBidPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) TerminateOnUpdate() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"terminateOnUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) UpdateTimeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"updateTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedEc2EksComputeEnvironment) UpdateToLatestImageVersion() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"updateToLatestImageVersion",
		&returns,
	)
	return returns
}


func NewManagedEc2EksComputeEnvironment(scope constructs.Construct, id *string, props *ManagedEc2EksComputeEnvironmentProps) ManagedEc2EksComputeEnvironment {
	_init_.Initialize()

	if err := validateNewManagedEc2EksComputeEnvironmentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedEc2EksComputeEnvironment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.ManagedEc2EksComputeEnvironment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewManagedEc2EksComputeEnvironment_Override(m ManagedEc2EksComputeEnvironment, scope constructs.Construct, id *string, props *ManagedEc2EksComputeEnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.ManagedEc2EksComputeEnvironment",
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
func ManagedEc2EksComputeEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedEc2EksComputeEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.ManagedEc2EksComputeEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ManagedEc2EksComputeEnvironment_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateManagedEc2EksComputeEnvironment_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.ManagedEc2EksComputeEnvironment",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ManagedEc2EksComputeEnvironment_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateManagedEc2EksComputeEnvironment_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.ManagedEc2EksComputeEnvironment",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func ManagedEc2EksComputeEnvironment_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_batch.ManagedEc2EksComputeEnvironment",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) AddInstanceClass(instanceClass awsec2.InstanceClass) {
	if err := m.validateAddInstanceClassParameters(instanceClass); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addInstanceClass",
		[]interface{}{instanceClass},
	)
}

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) AddInstanceType(instanceType awsec2.InstanceType) {
	if err := m.validateAddInstanceTypeParameters(instanceType); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addInstanceType",
		[]interface{}{instanceType},
	)
}

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := m.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := m.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) GetResourceNameAttribute(nameAttr *string) *string {
	if err := m.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

