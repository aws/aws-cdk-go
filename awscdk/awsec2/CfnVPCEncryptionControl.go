package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes the configuration and state of VPC encryption controls.
//
// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCEncryptionControl := awscdk.Aws_ec2.NewCfnVPCEncryptionControl(this, jsii.String("MyCfnVPCEncryptionControl"), &CfnVPCEncryptionControlProps{
//   	EgressOnlyInternetGatewayExclusionInput: jsii.String("egressOnlyInternetGatewayExclusionInput"),
//   	ElasticFileSystemExclusionInput: jsii.String("elasticFileSystemExclusionInput"),
//   	InternetGatewayExclusionInput: jsii.String("internetGatewayExclusionInput"),
//   	LambdaExclusionInput: jsii.String("lambdaExclusionInput"),
//   	Mode: jsii.String("mode"),
//   	NatGatewayExclusionInput: jsii.String("natGatewayExclusionInput"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualPrivateGatewayExclusionInput: jsii.String("virtualPrivateGatewayExclusionInput"),
//   	VpcId: jsii.String("vpcId"),
//   	VpcLatticeExclusionInput: jsii.String("vpcLatticeExclusionInput"),
//   	VpcPeeringExclusionInput: jsii.String("vpcPeeringExclusionInput"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html
//
type CfnVPCEncryptionControl interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsec2.IVPCEncryptionControlRef
	awscdk.ITaggableV2
	AttrResourceExclusions() awscdk.IResolvable
	// Specifies whether to exclude egress-only internet gateway traffic from encryption enforcement.
	AttrResourceExclusionsEgressOnlyInternetGateway() awscdk.IResolvable
	// The current state of the exclusion configuration.
	AttrResourceExclusionsEgressOnlyInternetGatewayState() *string
	// A message providing additional information about the exclusion state.
	AttrResourceExclusionsEgressOnlyInternetGatewayStateMessage() *string
	// Specifies whether to exclude Elastic File System traffic from encryption enforcement.
	AttrResourceExclusionsElasticFileSystem() awscdk.IResolvable
	// The current state of the exclusion configuration.
	AttrResourceExclusionsElasticFileSystemState() *string
	// A message providing additional information about the exclusion state.
	AttrResourceExclusionsElasticFileSystemStateMessage() *string
	// Specifies whether to exclude internet gateway traffic from encryption enforcement.
	AttrResourceExclusionsInternetGateway() awscdk.IResolvable
	// The current state of the exclusion configuration.
	AttrResourceExclusionsInternetGatewayState() *string
	// A message providing additional information about the exclusion state.
	AttrResourceExclusionsInternetGatewayStateMessage() *string
	// Specifies whether to exclude Lambda function traffic from encryption enforcement.
	AttrResourceExclusionsLambda() awscdk.IResolvable
	// The current state of the exclusion configuration.
	AttrResourceExclusionsLambdaState() *string
	// A message providing additional information about the exclusion state.
	AttrResourceExclusionsLambdaStateMessage() *string
	// Specifies whether to exclude NAT gateway traffic from encryption enforcement.
	AttrResourceExclusionsNatGateway() awscdk.IResolvable
	// The current state of the exclusion configuration.
	AttrResourceExclusionsNatGatewayState() *string
	// A message providing additional information about the exclusion state.
	AttrResourceExclusionsNatGatewayStateMessage() *string
	// Specifies whether to exclude virtual private gateway traffic from encryption enforcement.
	AttrResourceExclusionsVirtualPrivateGateway() awscdk.IResolvable
	// The current state of the exclusion configuration.
	AttrResourceExclusionsVirtualPrivateGatewayState() *string
	// A message providing additional information about the exclusion state.
	AttrResourceExclusionsVirtualPrivateGatewayStateMessage() *string
	// Specifies whether to exclude VPC Lattice traffic from encryption enforcement.
	AttrResourceExclusionsVpcLattice() awscdk.IResolvable
	// The current state of the exclusion configuration.
	AttrResourceExclusionsVpcLatticeState() *string
	// A message providing additional information about the exclusion state.
	AttrResourceExclusionsVpcLatticeStateMessage() *string
	// Specifies whether to exclude VPC peering connection traffic from encryption enforcement.
	AttrResourceExclusionsVpcPeering() awscdk.IResolvable
	// The current state of the exclusion configuration.
	AttrResourceExclusionsVpcPeeringState() *string
	// A message providing additional information about the exclusion state.
	AttrResourceExclusionsVpcPeeringStateMessage() *string
	// The current state of the VPC Encryption Control configuration.
	AttrState() *string
	// A message providing additional information about the encryption control state.
	AttrStateMessage() *string
	// The ID of the VPC Encryption Control configuration.
	AttrVpcEncryptionControlId() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies whether to exclude egress-only internet gateway traffic from encryption enforcement.
	EgressOnlyInternetGatewayExclusionInput() *string
	SetEgressOnlyInternetGatewayExclusionInput(val *string)
	// Specifies whether to exclude Elastic File System traffic from encryption enforcement.
	ElasticFileSystemExclusionInput() *string
	SetElasticFileSystemExclusionInput(val *string)
	Env() *interfaces.ResourceEnvironment
	// Specifies whether to exclude internet gateway traffic from encryption enforcement.
	InternetGatewayExclusionInput() *string
	SetInternetGatewayExclusionInput(val *string)
	// Specifies whether to exclude Lambda function traffic from encryption enforcement.
	LambdaExclusionInput() *string
	SetLambdaExclusionInput(val *string)
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
	// The encryption mode for the VPC Encryption Control configuration.
	Mode() *string
	SetMode(val *string)
	// Specifies whether to exclude NAT gateway traffic from encryption enforcement.
	NatGatewayExclusionInput() *string
	SetNatGatewayExclusionInput(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags assigned to the VPC Encryption Control configuration.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
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
	// Specifies whether to exclude virtual private gateway traffic from encryption enforcement.
	VirtualPrivateGatewayExclusionInput() *string
	SetVirtualPrivateGatewayExclusionInput(val *string)
	// A reference to a VPCEncryptionControl resource.
	VpcEncryptionControlRef() *interfacesawsec2.VPCEncryptionControlReference
	// The ID of the VPC for which to create the encryption control configuration.
	VpcId() *string
	SetVpcId(val *string)
	// Specifies whether to exclude VPC Lattice traffic from encryption enforcement.
	VpcLatticeExclusionInput() *string
	SetVpcLatticeExclusionInput(val *string)
	// Specifies whether to exclude VPC peering connection traffic from encryption enforcement.
	VpcPeeringExclusionInput() *string
	SetVpcPeeringExclusionInput(val *string)
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

// The jsii proxy struct for CfnVPCEncryptionControl
type jsiiProxy_CfnVPCEncryptionControl struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsec2IVPCEncryptionControlRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusions() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrResourceExclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsEgressOnlyInternetGateway() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrResourceExclusionsEgressOnlyInternetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsEgressOnlyInternetGatewayState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsEgressOnlyInternetGatewayState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsEgressOnlyInternetGatewayStateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsEgressOnlyInternetGatewayStateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsElasticFileSystem() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrResourceExclusionsElasticFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsElasticFileSystemState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsElasticFileSystemState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsElasticFileSystemStateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsElasticFileSystemStateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsInternetGateway() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrResourceExclusionsInternetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsInternetGatewayState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsInternetGatewayState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsInternetGatewayStateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsInternetGatewayStateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsLambda() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrResourceExclusionsLambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsLambdaState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsLambdaState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsLambdaStateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsLambdaStateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsNatGateway() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrResourceExclusionsNatGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsNatGatewayState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsNatGatewayState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsNatGatewayStateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsNatGatewayStateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsVirtualPrivateGateway() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrResourceExclusionsVirtualPrivateGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsVirtualPrivateGatewayState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsVirtualPrivateGatewayState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsVirtualPrivateGatewayStateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsVirtualPrivateGatewayStateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsVpcLattice() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrResourceExclusionsVpcLattice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsVpcLatticeState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsVpcLatticeState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsVpcLatticeStateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsVpcLatticeStateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsVpcPeering() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrResourceExclusionsVpcPeering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsVpcPeeringState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsVpcPeeringState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrResourceExclusionsVpcPeeringStateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceExclusionsVpcPeeringStateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrStateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) AttrVpcEncryptionControlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVpcEncryptionControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) EgressOnlyInternetGatewayExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyInternetGatewayExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) ElasticFileSystemExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticFileSystemExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) InternetGatewayExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internetGatewayExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) LambdaExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) NatGatewayExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) VirtualPrivateGatewayExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualPrivateGatewayExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) VpcEncryptionControlRef() *interfacesawsec2.VPCEncryptionControlReference {
	var returns *interfacesawsec2.VPCEncryptionControlReference
	_jsii_.Get(
		j,
		"vpcEncryptionControlRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) VpcLatticeExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcLatticeExclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEncryptionControl) VpcPeeringExclusionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringExclusionInput",
		&returns,
	)
	return returns
}


// Create a new `AWS::EC2::VPCEncryptionControl`.
func NewCfnVPCEncryptionControl(scope constructs.Construct, id *string, props *CfnVPCEncryptionControlProps) CfnVPCEncryptionControl {
	_init_.Initialize()

	if err := validateNewCfnVPCEncryptionControlParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVPCEncryptionControl{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnVPCEncryptionControl",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EC2::VPCEncryptionControl`.
func NewCfnVPCEncryptionControl_Override(c CfnVPCEncryptionControl, scope constructs.Construct, id *string, props *CfnVPCEncryptionControlProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnVPCEncryptionControl",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVPCEncryptionControl)SetEgressOnlyInternetGatewayExclusionInput(val *string) {
	_jsii_.Set(
		j,
		"egressOnlyInternetGatewayExclusionInput",
		val,
	)
}

func (j *jsiiProxy_CfnVPCEncryptionControl)SetElasticFileSystemExclusionInput(val *string) {
	_jsii_.Set(
		j,
		"elasticFileSystemExclusionInput",
		val,
	)
}

func (j *jsiiProxy_CfnVPCEncryptionControl)SetInternetGatewayExclusionInput(val *string) {
	_jsii_.Set(
		j,
		"internetGatewayExclusionInput",
		val,
	)
}

func (j *jsiiProxy_CfnVPCEncryptionControl)SetLambdaExclusionInput(val *string) {
	_jsii_.Set(
		j,
		"lambdaExclusionInput",
		val,
	)
}

func (j *jsiiProxy_CfnVPCEncryptionControl)SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_CfnVPCEncryptionControl)SetNatGatewayExclusionInput(val *string) {
	_jsii_.Set(
		j,
		"natGatewayExclusionInput",
		val,
	)
}

func (j *jsiiProxy_CfnVPCEncryptionControl)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnVPCEncryptionControl)SetVirtualPrivateGatewayExclusionInput(val *string) {
	_jsii_.Set(
		j,
		"virtualPrivateGatewayExclusionInput",
		val,
	)
}

func (j *jsiiProxy_CfnVPCEncryptionControl)SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_CfnVPCEncryptionControl)SetVpcLatticeExclusionInput(val *string) {
	_jsii_.Set(
		j,
		"vpcLatticeExclusionInput",
		val,
	)
}

func (j *jsiiProxy_CfnVPCEncryptionControl)SetVpcPeeringExclusionInput(val *string) {
	_jsii_.Set(
		j,
		"vpcPeeringExclusionInput",
		val,
	)
}

func CfnVPCEncryptionControl_ArnForVPCEncryptionControl(resource interfacesawsec2.IVPCEncryptionControlRef) *string {
	_init_.Initialize()

	if err := validateCfnVPCEncryptionControl_ArnForVPCEncryptionControlParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnVPCEncryptionControl",
		"arnForVPCEncryptionControl",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Creates a new IVPCEncryptionControlRef from a vpcEncryptionControlId.
func CfnVPCEncryptionControl_FromVpcEncryptionControlId(scope constructs.Construct, id *string, vpcEncryptionControlId *string) interfacesawsec2.IVPCEncryptionControlRef {
	_init_.Initialize()

	if err := validateCfnVPCEncryptionControl_FromVpcEncryptionControlIdParameters(scope, id, vpcEncryptionControlId); err != nil {
		panic(err)
	}
	var returns interfacesawsec2.IVPCEncryptionControlRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnVPCEncryptionControl",
		"fromVpcEncryptionControlId",
		[]interface{}{scope, id, vpcEncryptionControlId},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnVPCEncryptionControl_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPCEncryptionControl_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnVPCEncryptionControl",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnVPCEncryptionControl_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPCEncryptionControl_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnVPCEncryptionControl",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnVPCEncryptionControl.
func CfnVPCEncryptionControl_IsCfnVPCEncryptionControl(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPCEncryptionControl_IsCfnVPCEncryptionControlParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnVPCEncryptionControl",
		"isCfnVPCEncryptionControl",
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
func CfnVPCEncryptionControl_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPCEncryptionControl_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnVPCEncryptionControl",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVPCEncryptionControl_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.CfnVPCEncryptionControl",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVPCEncryptionControl) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnVPCEncryptionControl) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnVPCEncryptionControl) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPCEncryptionControl) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPCEncryptionControl) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnVPCEncryptionControl) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnVPCEncryptionControl) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPCEncryptionControl) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVPCEncryptionControl) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

