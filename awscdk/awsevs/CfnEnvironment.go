package awsevs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon EVS environment that runs VCF software, such as SDDC Manager, NSX Manager, and vCenter Server.
//
// During environment creation, Amazon EVS performs validations on DNS settings, provisions VLAN subnets and hosts, and deploys the supplied version of VCF.
//
// It can take several hours to create an environment. After the deployment completes, you can configure VCF in the vSphere user interface according to your needs.
//
// > You cannot use the `dedicatedHostId` and `placementGroupId` parameters together in the same `CreateEnvironment` action. This results in a `ValidationException` response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironment := awscdk.Aws_evs.NewCfnEnvironment(this, jsii.String("MyCfnEnvironment"), &CfnEnvironmentProps{
//   	ConnectivityInfo: &ConnectivityInfoProperty{
//   		PrivateRouteServerPeerings: []*string{
//   			jsii.String("privateRouteServerPeerings"),
//   		},
//   	},
//   	LicenseInfo: &LicenseInfoProperty{
//   		SolutionKey: jsii.String("solutionKey"),
//   		VsanKey: jsii.String("vsanKey"),
//   	},
//   	ServiceAccessSubnetId: jsii.String("serviceAccessSubnetId"),
//   	SiteId: jsii.String("siteId"),
//   	TermsAccepted: jsii.Boolean(false),
//   	VcfHostnames: &VcfHostnamesProperty{
//   		CloudBuilder: jsii.String("cloudBuilder"),
//   		Nsx: jsii.String("nsx"),
//   		NsxEdge1: jsii.String("nsxEdge1"),
//   		NsxEdge2: jsii.String("nsxEdge2"),
//   		NsxManager1: jsii.String("nsxManager1"),
//   		NsxManager2: jsii.String("nsxManager2"),
//   		NsxManager3: jsii.String("nsxManager3"),
//   		SddcManager: jsii.String("sddcManager"),
//   		VCenter: jsii.String("vCenter"),
//   	},
//   	VcfVersion: jsii.String("vcfVersion"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	EnvironmentName: jsii.String("environmentName"),
//   	Hosts: []interface{}{
//   		&HostInfoForCreateProperty{
//   			HostName: jsii.String("hostName"),
//   			InstanceType: jsii.String("instanceType"),
//   			KeyName: jsii.String("keyName"),
//
//   			// the properties below are optional
//   			DedicatedHostId: jsii.String("dedicatedHostId"),
//   			PlacementGroupId: jsii.String("placementGroupId"),
//   		},
//   	},
//   	InitialVlans: &InitialVlansProperty{
//   		EdgeVTep: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		ExpansionVlan1: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		ExpansionVlan2: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		Hcx: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		NsxUpLink: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VmkManagement: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VmManagement: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VMotion: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VSan: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   		VTep: &InitialVlanInfoProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//
//   		// the properties below are optional
//   		HcxNetworkAclId: jsii.String("hcxNetworkAclId"),
//   		IsHcxPublic: jsii.Boolean(false),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ServiceAccessSecurityGroups: &ServiceAccessSecurityGroupsProperty{
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evs-environment.html
//
type CfnEnvironment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsevs.IEnvironmentRef
	awscdk.ITaggableV2
	// A check on the environment to identify instance health and VMware VCF licensing issues. For example:.
	//
	// `{ "checks": [ { "type": "KEY_REUSE", "result": "PASSED" }, { "type": "KEY_COVERAGE", "result": "PASSED" }, { "type": "REACHABILITY", "result": "PASSED" }, { "type": "HOST_COUNT", "result": "PASSED" } ] }`.
	AttrChecks() awscdk.IResolvable
	// The date and time that the environment was created.
	//
	// For example: `1749081600.000` .
	AttrCreatedAt() *string
	// The VCF credentials that are stored as Amazon EVS managed secrets in AWS Secrets Manager.
	//
	// Amazon EVS stores credentials that are needed to install vCenter Server, NSX, and SDDC Manager. For example:
	//
	// `{ [ { "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_vCenterAdmin-MnTMEi" }, { "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_vCenterRoot-87VyCF" }, { "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_NSXRoot-SR3k43" }, { "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_NSXAdmin-L5LUiD" }, { "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_NSXAudit-Q2oW46" }, { "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_SDDCManagerRoot-bFulOq" }, { "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_SDDCManagerVCF-Ec3gES" }, { "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_SDDCManagerAdmin-JMTAAb" } ] }`
	AttrCredentials() awscdk.IResolvable
	// The Amazon Resource Name (ARN) that is associated with the environment.
	//
	// For example: `arn:aws:evs:us-east-1:000000000000:environment/env-1234567890` .
	AttrEnvironmentArn() *string
	// The unique ID for the environment.
	//
	// For example: `env-1234567890` .
	AttrEnvironmentId() *string
	// The state of an environment.
	//
	// For example: `CREATED` .
	AttrEnvironmentState() *string
	// The date and time that the environment was modified.
	//
	// For example: `1749081600.000` .
	AttrModifiedAt() *string
	// A detailed description of the `environmentState` of an environment.
	//
	// For example: `Environment successfully created` .
	AttrStateDetails() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The connectivity configuration for the environment.
	ConnectivityInfo() interface{}
	SetConnectivityInfo(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	Env() *interfaces.ResourceEnvironment
	// The name of the environment.
	EnvironmentName() *string
	SetEnvironmentName(val *string)
	// A reference to a Environment resource.
	EnvironmentRef() *interfacesawsevs.EnvironmentReference
	// Required for environment resource creation.
	Hosts() interface{}
	SetHosts(val interface{})
	// > Amazon EVS is in public preview release and is subject to change.
	InitialVlans() interface{}
	SetInitialVlans(val interface{})
	// The AWS KMS key ID that AWS Secrets Manager uses to encrypt secrets that are associated with the environment.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// The license information that Amazon EVS requires to create an environment.
	LicenseInfo() interface{}
	SetLicenseInfo(val interface{})
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The security groups that allow traffic between the Amazon EVS control plane and your VPC for service access.
	ServiceAccessSecurityGroups() interface{}
	SetServiceAccessSecurityGroups(val interface{})
	// The subnet that is used to establish connectivity between the Amazon EVS control plane and VPC.
	ServiceAccessSubnetId() *string
	SetServiceAccessSubnetId(val *string)
	// The Broadcom Site ID that is associated with your Amazon EVS environment.
	SiteId() *string
	SetSiteId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata that assists with categorization and organization.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// Customer confirmation that the customer has purchased and will continue to maintain the required number of VCF software licenses to cover all physical processor cores in the Amazon EVS environment.
	TermsAccepted() interface{}
	SetTermsAccepted(val interface{})
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
	// The DNS hostnames to be used by the VCF management appliances in your environment.
	VcfHostnames() interface{}
	SetVcfHostnames(val interface{})
	// The VCF version of the environment.
	VcfVersion() *string
	SetVcfVersion(val *string)
	// The VPC associated with the environment.
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

// The jsii proxy struct for CfnEnvironment
type jsiiProxy_CfnEnvironment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsevsIEnvironmentRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnEnvironment) AttrChecks() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrCredentials() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrEnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrEnvironmentState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrStateDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) ConnectivityInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectivityInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) EnvironmentRef() *interfacesawsevs.EnvironmentReference {
	var returns *interfacesawsevs.EnvironmentReference
	_jsii_.Get(
		j,
		"environmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) Hosts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) InitialVlans() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialVlans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) LicenseInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) ServiceAccessSecurityGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAccessSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) ServiceAccessSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) SiteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) TermsAccepted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"termsAccepted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) VcfHostnames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vcfHostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) VcfVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcfVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new `AWS::EVS::Environment`.
func NewCfnEnvironment(scope constructs.Construct, id *string, props *CfnEnvironmentProps) CfnEnvironment {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_evs.CfnEnvironment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EVS::Environment`.
func NewCfnEnvironment_Override(c CfnEnvironment, scope constructs.Construct, id *string, props *CfnEnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_evs.CfnEnvironment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetConnectivityInfo(val interface{}) {
	if err := j.validateSetConnectivityInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectivityInfo",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetEnvironmentName(val *string) {
	_jsii_.Set(
		j,
		"environmentName",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetHosts(val interface{}) {
	if err := j.validateSetHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hosts",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetInitialVlans(val interface{}) {
	if err := j.validateSetInitialVlansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialVlans",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetLicenseInfo(val interface{}) {
	if err := j.validateSetLicenseInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseInfo",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetServiceAccessSecurityGroups(val interface{}) {
	if err := j.validateSetServiceAccessSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetServiceAccessSubnetId(val *string) {
	if err := j.validateSetServiceAccessSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessSubnetId",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetSiteId(val *string) {
	if err := j.validateSetSiteIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"siteId",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetTermsAccepted(val interface{}) {
	if err := j.validateSetTermsAcceptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"termsAccepted",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetVcfHostnames(val interface{}) {
	if err := j.validateSetVcfHostnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcfHostnames",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetVcfVersion(val *string) {
	if err := j.validateSetVcfVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcfVersion",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func CfnEnvironment_ArnForEnvironment(resource interfacesawsevs.IEnvironmentRef) *string {
	_init_.Initialize()

	if err := validateCfnEnvironment_ArnForEnvironmentParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evs.CfnEnvironment",
		"arnForEnvironment",
		[]interface{}{resource},
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
func CfnEnvironment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironment_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evs.CfnEnvironment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnEnvironment_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironment_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evs.CfnEnvironment",
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
func CfnEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evs.CfnEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_evs.CfnEnvironment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironment) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEnvironment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEnvironment) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnEnvironment) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnEnvironment) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEnvironment) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironment) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEnvironment) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEnvironment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnEnvironment) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnEnvironment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironment) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

