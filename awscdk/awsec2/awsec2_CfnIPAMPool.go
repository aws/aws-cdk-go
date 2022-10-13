package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::EC2::IPAMPool`.
//
// In IPAM, a pool is a collection of contiguous IP addresses CIDRs. Pools enable you to organize your IP addresses according to your routing and security needs. For example, if you have separate routing and security needs for development and production applications, you can create a pool for each.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPAMPool := awscdk.Aws_ec2.NewCfnIPAMPool(this, jsii.String("MyCfnIPAMPool"), &cfnIPAMPoolProps{
//   	addressFamily: jsii.String("addressFamily"),
//   	ipamScopeId: jsii.String("ipamScopeId"),
//
//   	// the properties below are optional
//   	allocationDefaultNetmaskLength: jsii.Number(123),
//   	allocationMaxNetmaskLength: jsii.Number(123),
//   	allocationMinNetmaskLength: jsii.Number(123),
//   	allocationResourceTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	autoImport: jsii.Boolean(false),
//   	awsService: jsii.String("awsService"),
//   	description: jsii.String("description"),
//   	locale: jsii.String("locale"),
//   	provisionedCidrs: []interface{}{
//   		&provisionedCidrProperty{
//   			cidr: jsii.String("cidr"),
//   		},
//   	},
//   	publiclyAdvertisable: jsii.Boolean(false),
//   	sourceIpamPoolId: jsii.String("sourceIpamPoolId"),
//   	tags: []*cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnIPAMPool interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The address family of the pool.
	AddressFamily() *string
	SetAddressFamily(val *string)
	// The default netmask length for allocations added to this pool.
	//
	// If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16.
	AllocationDefaultNetmaskLength() *float64
	SetAllocationDefaultNetmaskLength(val *float64)
	// The maximum netmask length possible for CIDR allocations in this IPAM pool to be compliant.
	//
	// The maximum netmask length must be greater than the minimum netmask length. Possible netmask lengths for IPv4 addresses are 0 - 32. Possible netmask lengths for IPv6 addresses are 0 - 128.
	AllocationMaxNetmaskLength() *float64
	SetAllocationMaxNetmaskLength(val *float64)
	// The minimum netmask length required for CIDR allocations in this IPAM pool to be compliant.
	//
	// The minimum netmask length must be less than the maximum netmask length. Possible netmask lengths for IPv4 addresses are 0 - 32. Possible netmask lengths for IPv6 addresses are 0 - 128.
	AllocationMinNetmaskLength() *float64
	SetAllocationMinNetmaskLength(val *float64)
	// Tags that are required for resources that use CIDRs from this IPAM pool.
	//
	// Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
	AllocationResourceTags() interface{}
	SetAllocationResourceTags(val interface{})
	// The ARN of the IPAM pool.
	AttrArn() *string
	// The ARN of the IPAM.
	AttrIpamArn() *string
	// The ID of the IPAM pool.
	AttrIpamPoolId() *string
	// The ARN of the scope of the IPAM pool.
	AttrIpamScopeArn() *string
	// The scope of the IPAM.
	AttrIpamScopeType() *string
	// The depth of pools in your IPAM pool.
	//
	// The pool depth quota is 10.
	AttrPoolDepth() *float64
	// The state of the IPAM pool.
	AttrState() *string
	// A message related to the failed creation of an IPAM pool.
	AttrStateMessage() *string
	// If selected, IPAM will continuously look for resources within the CIDR range of this pool and automatically import them as allocations into your IPAM.
	//
	// The CIDRs that will be allocated for these resources must not already be allocated to other resources in order for the import to succeed. IPAM will import a CIDR regardless of its compliance with the pool's allocation rules, so a resource might be imported and subsequently marked as noncompliant. If IPAM discovers multiple CIDRs that overlap, IPAM will import the largest CIDR only. If IPAM discovers multiple CIDRs with matching CIDRs, IPAM will randomly import one of them only.
	//
	// A locale must be set on the pool for this feature to work.
	AutoImport() interface{}
	SetAutoImport(val interface{})
	// `AWS::EC2::IPAMPool.AwsService`.
	AwsService() *string
	SetAwsService(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the IPAM pool.
	Description() *string
	SetDescription(val *string)
	// The ID of the scope in which you would like to create the IPAM pool.
	IpamScopeId() *string
	SetIpamScopeId(val *string)
	// The locale of the IPAM pool.
	//
	// In IPAM, the locale is the AWS Region where you want to make an IPAM pool available for allocations. Only resources in the same Region as the locale of the pool can get IP address allocations from the pool. You can only allocate a CIDR for a VPC, for example, from an IPAM pool that shares a locale with the VPC’s Region. Note that once you choose a Locale for a pool, you cannot modify it. If you choose an AWS Region for locale that has not been configured as an operating Region for the IPAM, you'll get an error.
	Locale() *string
	SetLocale(val *string)
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
	// The tree node.
	Node() constructs.Node
	// Information about the CIDRs provisioned to an IPAM pool.
	ProvisionedCidrs() interface{}
	SetProvisionedCidrs(val interface{})
	// Determines if a pool is publicly advertisable.
	//
	// This option is not available for pools with AddressFamily set to `ipv4` .
	PubliclyAdvertisable() interface{}
	SetPubliclyAdvertisable(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the source IPAM pool.
	//
	// You can use this option to create an IPAM pool within an existing source pool.
	SourceIpamPoolId() *string
	SetSourceIpamPoolId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The key/value combination of a tag assigned to the resource.
	//
	// Use the tag key in the filter name and the tag value as the filter value. For example, to find all resources that have a tag with the key `Owner` and the value `TeamA` , specify `tag:Owner` for the filter name and `TeamA` for the filter value.
	Tags() awscdk.TagManager
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
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnIPAMPool
type jsiiProxy_CfnIPAMPool struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIPAMPool) AddressFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AllocationDefaultNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationDefaultNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AllocationMaxNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMaxNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AllocationMinNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMinNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AllocationResourceTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allocationResourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AttrIpamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIpamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AttrIpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AttrIpamScopeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIpamScopeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AttrIpamScopeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIpamScopeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AttrPoolDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrPoolDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AttrStateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AutoImport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoImport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) AwsService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) IpamScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) ProvisionedCidrs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) PubliclyAdvertisable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAdvertisable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) SourceIpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPool) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::EC2::IPAMPool`.
func NewCfnIPAMPool(scope constructs.Construct, id *string, props *CfnIPAMPoolProps) CfnIPAMPool {
	_init_.Initialize()

	if err := validateNewCfnIPAMPoolParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIPAMPool{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnIPAMPool",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EC2::IPAMPool`.
func NewCfnIPAMPool_Override(c CfnIPAMPool, scope constructs.Construct, id *string, props *CfnIPAMPoolProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnIPAMPool",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetAddressFamily(val *string) {
	if err := j.validateSetAddressFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressFamily",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetAllocationDefaultNetmaskLength(val *float64) {
	_jsii_.Set(
		j,
		"allocationDefaultNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetAllocationMaxNetmaskLength(val *float64) {
	_jsii_.Set(
		j,
		"allocationMaxNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetAllocationMinNetmaskLength(val *float64) {
	_jsii_.Set(
		j,
		"allocationMinNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetAllocationResourceTags(val interface{}) {
	if err := j.validateSetAllocationResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationResourceTags",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetAutoImport(val interface{}) {
	if err := j.validateSetAutoImportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoImport",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetAwsService(val *string) {
	_jsii_.Set(
		j,
		"awsService",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetIpamScopeId(val *string) {
	if err := j.validateSetIpamScopeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipamScopeId",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetLocale(val *string) {
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetProvisionedCidrs(val interface{}) {
	if err := j.validateSetProvisionedCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedCidrs",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetPubliclyAdvertisable(val interface{}) {
	if err := j.validateSetPubliclyAdvertisableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAdvertisable",
		val,
	)
}

func (j *jsiiProxy_CfnIPAMPool)SetSourceIpamPoolId(val *string) {
	_jsii_.Set(
		j,
		"sourceIpamPoolId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnIPAMPool_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIPAMPool_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnIPAMPool",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnIPAMPool_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnIPAMPool_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnIPAMPool",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnIPAMPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIPAMPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnIPAMPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIPAMPool_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.CfnIPAMPool",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIPAMPool) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnIPAMPool) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIPAMPool) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnIPAMPool) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnIPAMPool) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnIPAMPool) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnIPAMPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnIPAMPool) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIPAMPool) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnIPAMPool) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnIPAMPool) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIPAMPool) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnIPAMPool) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIPAMPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIPAMPool) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

