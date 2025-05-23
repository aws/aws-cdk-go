package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a new VPC subnet resource.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(1024),
//   	DesiredCount: jsii.Number(1),
//   	Cpu: jsii.Number(512),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	MinHealthyPercent: jsii.Number(100),
//   	TaskSubnets: &SubnetSelection{
//   		Subnets: []iSubnet{
//   			ec2.Subnet_FromSubnetId(this, jsii.String("subnet"), jsii.String("VpcISOLATEDSubnet1Subnet80F07FA0")),
//   		},
//   	},
//   })
//
type Subnet interface {
	awscdk.Resource
	ISubnet
	// The Availability Zone the subnet is located in.
	AvailabilityZone() *string
	// Parts of this VPC subnet.
	DependencyElements() *[]constructs.IDependable
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Dependable that can be depended upon to force internet connectivity established on the VPC.
	InternetConnectivityEstablished() constructs.IDependable
	// The IPv4 CIDR block for this subnet.
	Ipv4CidrBlock() *string
	// Network ACL associated with this Subnet.
	//
	// Upon creation, this is the default ACL which allows all traffic, except
	// explicit DENY entries that you add.
	//
	// You can replace it with a custom ACL which denies all traffic except
	// the explicit ALLOW entries that you add by creating a `NetworkAcl`
	// object and calling `associateNetworkAcl()`.
	NetworkAcl() INetworkAcl
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
	// The routeTableId attached to this subnet.
	RouteTable() IRouteTable
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	SubnetAvailabilityZone() *string
	// The subnetId for this particular subnet.
	SubnetId() *string
	SubnetIpv6CidrBlocks() *[]*string
	SubnetNetworkAclAssociationId() *string
	// The Amazon Resource Name (ARN) of the Outpost for this subnet (if one exists).
	SubnetOutpostArn() *string
	SubnetVpcId() *string
	// Create a default route that points to a passed IGW, with a dependency on the IGW's attachment to the VPC.
	AddDefaultInternetRoute(gatewayId *string, gatewayAttachment constructs.IDependable)
	// Adds an entry to this subnets route table that points to the passed NATGatewayId.
	AddDefaultNatRoute(natGatewayId *string)
	// Create a default IPv6 route that points to a passed EIGW.
	AddIpv6DefaultEgressOnlyInternetRoute(gatewayId *string)
	// Create a default IPv6 route that points to a passed IGW.
	AddIpv6DefaultInternetRoute(gatewayId *string)
	// Adds an entry to this subnets route table that points to the passed NATGatewayId.
	//
	// Uses the known 64:ff9b::/96 prefix.
	AddIpv6Nat64Route(natGatewayId *string)
	// Adds an entry to this subnets route table.
	AddRoute(id *string, options *AddRouteOptions)
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
	// Associate a Network ACL with this subnet.
	AssociateNetworkAcl(id *string, networkAcl INetworkAcl)
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

// The jsii proxy struct for Subnet
type jsiiProxy_Subnet struct {
	internal.Type__awscdkResource
	jsiiProxy_ISubnet
}

func (j *jsiiProxy_Subnet) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) DependencyElements() *[]constructs.IDependable {
	var returns *[]constructs.IDependable
	_jsii_.Get(
		j,
		"dependencyElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) InternetConnectivityEstablished() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"internetConnectivityEstablished",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) Ipv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) NetworkAcl() INetworkAcl {
	var returns INetworkAcl
	_jsii_.Get(
		j,
		"networkAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) RouteTable() IRouteTable {
	var returns IRouteTable
	_jsii_.Get(
		j,
		"routeTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) SubnetAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) SubnetIpv6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIpv6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) SubnetNetworkAclAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetNetworkAclAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) SubnetOutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetOutpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Subnet) SubnetVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetVpcId",
		&returns,
	)
	return returns
}


func NewSubnet(scope constructs.Construct, id *string, props *SubnetProps) Subnet {
	_init_.Initialize()

	if err := validateNewSubnetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Subnet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.Subnet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSubnet_Override(s Subnet, scope constructs.Construct, id *string, props *SubnetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.Subnet",
		[]interface{}{scope, id, props},
		s,
	)
}

func Subnet_FromSubnetAttributes(scope constructs.Construct, id *string, attrs *SubnetAttributes) ISubnet {
	_init_.Initialize()

	if err := validateSubnet_FromSubnetAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns ISubnet

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Subnet",
		"fromSubnetAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import existing subnet from id.
func Subnet_FromSubnetId(scope constructs.Construct, id *string, subnetId *string) ISubnet {
	_init_.Initialize()

	if err := validateSubnet_FromSubnetIdParameters(scope, id, subnetId); err != nil {
		panic(err)
	}
	var returns ISubnet

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Subnet",
		"fromSubnetId",
		[]interface{}{scope, id, subnetId},
		&returns,
	)

	return returns
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
func Subnet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSubnet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Subnet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Subnet_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSubnet_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Subnet",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Subnet_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSubnet_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Subnet",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Subnet_IsVpcSubnet(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSubnet_IsVpcSubnetParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Subnet",
		"isVpcSubnet",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Subnet_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Subnet",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Subnet) AddDefaultInternetRoute(gatewayId *string, gatewayAttachment constructs.IDependable) {
	if err := s.validateAddDefaultInternetRouteParameters(gatewayId, gatewayAttachment); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addDefaultInternetRoute",
		[]interface{}{gatewayId, gatewayAttachment},
	)
}

func (s *jsiiProxy_Subnet) AddDefaultNatRoute(natGatewayId *string) {
	if err := s.validateAddDefaultNatRouteParameters(natGatewayId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addDefaultNatRoute",
		[]interface{}{natGatewayId},
	)
}

func (s *jsiiProxy_Subnet) AddIpv6DefaultEgressOnlyInternetRoute(gatewayId *string) {
	if err := s.validateAddIpv6DefaultEgressOnlyInternetRouteParameters(gatewayId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addIpv6DefaultEgressOnlyInternetRoute",
		[]interface{}{gatewayId},
	)
}

func (s *jsiiProxy_Subnet) AddIpv6DefaultInternetRoute(gatewayId *string) {
	if err := s.validateAddIpv6DefaultInternetRouteParameters(gatewayId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addIpv6DefaultInternetRoute",
		[]interface{}{gatewayId},
	)
}

func (s *jsiiProxy_Subnet) AddIpv6Nat64Route(natGatewayId *string) {
	if err := s.validateAddIpv6Nat64RouteParameters(natGatewayId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addIpv6Nat64Route",
		[]interface{}{natGatewayId},
	)
}

func (s *jsiiProxy_Subnet) AddRoute(id *string, options *AddRouteOptions) {
	if err := s.validateAddRouteParameters(id, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addRoute",
		[]interface{}{id, options},
	)
}

func (s *jsiiProxy_Subnet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_Subnet) AssociateNetworkAcl(id *string, networkAcl INetworkAcl) {
	if err := s.validateAssociateNetworkAclParameters(id, networkAcl); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"associateNetworkAcl",
		[]interface{}{id, networkAcl},
	)
}

func (s *jsiiProxy_Subnet) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Subnet) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Subnet) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Subnet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

