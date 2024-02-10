package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a subnet for the specified VPC.
//
// For an IPv4 only subnet, specify an IPv4 CIDR block. If the VPC has an IPv6 CIDR block, you can create an IPv6 only subnet or a dual stack subnet instead. For an IPv6 only subnet, specify an IPv6 CIDR block. For a dual stack subnet, specify both an IPv4 CIDR block and an IPv6 CIDR block.
//
// For more information, see [Subnets for your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/configure-subnets.html) in the *Amazon VPC User Guide* .
//
// Example:
//   var vpc vpc
//
//
//   func associateSubnetWithV6Cidr(vpc *vpc, count *f64, subnet iSubnet) {
//   	cfnSubnet := *subnet.Node.defaultChild.(cfnSubnet)
//   	cfnSubnet.Ipv6CidrBlock = awscdk.Fn_Select(count, awscdk.Fn_Cidr(awscdk.Fn_Select(jsii.Number(0), vpc.VpcIpv6CidrBlocks), jsii.Number(256), (jsii.Number(128 - 64)).toString()))
//   	cfnSubnet.AssignIpv6AddressOnCreation = true
//   }
//
//   // make an ipv6 cidr
//   ipv6cidr := ec2.NewCfnVPCCidrBlock(this, jsii.String("CIDR6"), &CfnVPCCidrBlockProps{
//   	VpcId: vpc.VpcId,
//   	AmazonProvidedIpv6CidrBlock: jsii.Boolean(true),
//   })
//
//   // connect the ipv6 cidr to all vpc subnets
//   subnetcount := 0
//   subnets := vpc.PublicSubnets.concat(vpc.PrivateSubnets)
//   for _, subnet := range subnets {
//   	// Wait for the ipv6 cidr to complete
//   	subnet.Node.AddDependency(ipv6cidr)
//   	associateSubnetWithV6Cidr(vpc, subnetcount, subnet)
//   	subnetcount = subnetcount + 1
//   }
//
//   cluster := eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_29(),
//   	Vpc: vpc,
//   	IpFamily: eks.IpFamily_IP_V6,
//   	VpcSubnets: []subnetSelection{
//   		&subnetSelection{
//   			Subnets: vpc.*PublicSubnets,
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html
//
type CfnSubnet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// Indicates whether a network interface created in this subnet receives an IPv6 address.
	//
	// The default value is `false` .
	AssignIpv6AddressOnCreation() interface{}
	SetAssignIpv6AddressOnCreation(val interface{})
	// The Availability Zone of this subnet.
	//
	// For example, `us-east-1a` .
	AttrAvailabilityZone() *string
	// The Availability Zone ID of this subnet.
	//
	// For example, `use1-az1` .
	AttrAvailabilityZoneId() *string
	// The IPv4 CIDR blocks that are associated with the subnet.
	AttrCidrBlock() *string
	AttrIpv6CidrBlocks() *[]*string
	// The ID of the network ACL that is associated with the subnet's VPC, such as `acl-5fb85d36` .
	AttrNetworkAclAssociationId() *string
	// The Amazon Resource Name (ARN) of the Outpost.
	AttrOutpostArn() *string
	// The ID of the subnet.
	AttrSubnetId() *string
	// The ID of the subnet's VPC, such as `vpc-11ad4878` .
	AttrVpcId() *string
	// The Availability Zone of the subnet.
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// The AZ ID of the subnet.
	AvailabilityZoneId() *string
	SetAvailabilityZoneId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The IPv4 CIDR block assigned to the subnet.
	CidrBlock() *string
	SetCidrBlock(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations.
	EnableDns64() interface{}
	SetEnableDns64(val interface{})
	// An IPv4 IPAM pool ID for the subnet.
	Ipv4IpamPoolId() *string
	SetIpv4IpamPoolId(val *string)
	// An IPv4 netmask length for the subnet.
	Ipv4NetmaskLength() *float64
	SetIpv4NetmaskLength(val *float64)
	// The IPv6 CIDR block.
	Ipv6CidrBlock() *string
	SetIpv6CidrBlock(val *string)
	// The IPv6 network ranges for the subnet, in CIDR notation.
	Ipv6CidrBlocks() *[]*string
	SetIpv6CidrBlocks(val *[]*string)
	// An IPv6 IPAM pool ID for the subnet.
	Ipv6IpamPoolId() *string
	SetIpv6IpamPoolId(val *string)
	// Indicates whether this is an IPv6 only subnet.
	Ipv6Native() interface{}
	SetIpv6Native(val interface{})
	// An IPv6 netmask length for the subnet.
	Ipv6NetmaskLength() *float64
	SetIpv6NetmaskLength(val *float64)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// Indicates whether instances launched in this subnet receive a public IPv4 address.
	//
	// The default value is `false` .
	MapPublicIpOnLaunch() interface{}
	SetMapPublicIpOnLaunch(val interface{})
	// The tree node.
	Node() constructs.Node
	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn() *string
	SetOutpostArn(val *string)
	// The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled.
	PrivateDnsNameOptionsOnLaunch() interface{}
	SetPrivateDnsNameOptionsOnLaunch(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Any tags assigned to the subnet.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The ID of the VPC the subnet is in.
	VpcId() *string
	SetVpcId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSubnet
type jsiiProxy_CfnSubnet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnSubnet) AssignIpv6AddressOnCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) AttrAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) AttrAvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAvailabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) AttrCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) AttrIpv6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrIpv6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) AttrNetworkAclAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNetworkAclAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) AttrOutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOutpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) AttrSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) AttrVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) EnableDns64() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDns64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) Ipv4IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4IpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) Ipv4NetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4NetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) Ipv6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) Ipv6IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6IpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) Ipv6Native() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Native",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) Ipv6NetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) MapPublicIpOnLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapPublicIpOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) PrivateDnsNameOptionsOnLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateDnsNameOptionsOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnet) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


func NewCfnSubnet(scope constructs.Construct, id *string, props *CfnSubnetProps) CfnSubnet {
	_init_.Initialize()

	if err := validateNewCfnSubnetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSubnet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnSubnet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnSubnet_Override(c CfnSubnet, scope constructs.Construct, id *string, props *CfnSubnetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnSubnet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSubnet)SetAssignIpv6AddressOnCreation(val interface{}) {
	if err := j.validateSetAssignIpv6AddressOnCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetAvailabilityZoneId(val *string) {
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetCidrBlock(val *string) {
	_jsii_.Set(
		j,
		"cidrBlock",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetEnableDns64(val interface{}) {
	if err := j.validateSetEnableDns64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDns64",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetIpv4IpamPoolId(val *string) {
	_jsii_.Set(
		j,
		"ipv4IpamPoolId",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetIpv4NetmaskLength(val *float64) {
	_jsii_.Set(
		j,
		"ipv4NetmaskLength",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetIpv6CidrBlock(val *string) {
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetIpv6CidrBlocks(val *[]*string) {
	_jsii_.Set(
		j,
		"ipv6CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetIpv6IpamPoolId(val *string) {
	_jsii_.Set(
		j,
		"ipv6IpamPoolId",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetIpv6Native(val interface{}) {
	if err := j.validateSetIpv6NativeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Native",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetIpv6NetmaskLength(val *float64) {
	_jsii_.Set(
		j,
		"ipv6NetmaskLength",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetMapPublicIpOnLaunch(val interface{}) {
	if err := j.validateSetMapPublicIpOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapPublicIpOnLaunch",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetOutpostArn(val *string) {
	_jsii_.Set(
		j,
		"outpostArn",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetPrivateDnsNameOptionsOnLaunch(val interface{}) {
	_jsii_.Set(
		j,
		"privateDnsNameOptionsOnLaunch",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnSubnet)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSubnet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSubnet_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnSubnet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnSubnet_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSubnet_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnSubnet",
		"isCfnResource",
		[]interface{}{x},
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
func CfnSubnet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSubnet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnSubnet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubnet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.CfnSubnet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubnet) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSubnet) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSubnet) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSubnet) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSubnet) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSubnet) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSubnet) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSubnet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSubnet) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubnet) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubnet) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSubnet) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubnet) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubnet) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSubnet) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSubnet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubnet) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnSubnet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubnet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubnet) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

