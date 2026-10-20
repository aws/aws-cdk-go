package awsworkspaces

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspaces/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspaces"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::WorkSpaces::Directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDirectory := awscdk.Aws_workspaces.NewCfnDirectory(this, jsii.String("MyCfnDirectory"), &CfnDirectoryProps{
//   	ActiveDirectoryConfig: &ActiveDirectoryConfigProperty{
//   		DomainName: jsii.String("domainName"),
//   		ServiceAccountSecretArn: jsii.String("serviceAccountSecretArn"),
//   	},
//   	CertificateBasedAuthProperties: &CertificateBasedAuthPropertiesProperty{
//   		CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   		Status: jsii.String("status"),
//   	},
//   	EnableSelfService: jsii.Boolean(false),
//   	EndpointEncryptionMode: jsii.String("endpointEncryptionMode"),
//   	IdcInstanceArn: jsii.String("idcInstanceArn"),
//   	IpGroupIds: []*string{
//   		jsii.String("ipGroupIds"),
//   	},
//   	MicrosoftEntraConfig: &MicrosoftEntraConfigProperty{
//   		ApplicationConfigSecretArn: jsii.String("applicationConfigSecretArn"),
//   		TenantId: jsii.String("tenantId"),
//   	},
//   	SamlProperties: &SamlPropertiesProperty{
//   		RelayStateParameterName: jsii.String("relayStateParameterName"),
//   		Status: jsii.String("status"),
//   		UserAccessUrl: jsii.String("userAccessUrl"),
//   	},
//   	SelfservicePermissions: &SelfservicePermissionsProperty{
//   		ChangeComputeType: jsii.String("changeComputeType"),
//   		IncreaseVolumeSize: jsii.String("increaseVolumeSize"),
//   		RebuildWorkspace: jsii.String("rebuildWorkspace"),
//   		RestartWorkspace: jsii.String("restartWorkspace"),
//   		SwitchRunningMode: jsii.String("switchRunningMode"),
//   	},
//   	StreamingProperties: &StreamingPropertiesProperty{
//   		GlobalAccelerator: &GlobalAcceleratorForDirectoryProperty{
//   			Mode: jsii.String("mode"),
//
//   			// the properties below are optional
//   			PreferredProtocol: jsii.String("preferredProtocol"),
//   		},
//   		StorageConnectors: []interface{}{
//   			&StorageConnectorProperty{
//   				ConnectorType: jsii.String("connectorType"),
//   				Status: jsii.String("status"),
//   			},
//   		},
//   		StreamingExperiencePreferredProtocol: jsii.String("streamingExperiencePreferredProtocol"),
//   		UserSettings: []interface{}{
//   			&UserSettingProperty{
//   				Action: jsii.String("action"),
//   				Permission: jsii.String("permission"),
//
//   				// the properties below are optional
//   				MaximumLength: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tenancy: jsii.String("tenancy"),
//   	UserIdentityType: jsii.String("userIdentityType"),
//   	WorkspaceAccessProperties: &WorkspaceAccessPropertiesProperty{
//   		AccessEndpointConfig: &AccessEndpointConfigProperty{
//   			AccessEndpoints: []interface{}{
//   				&AccessEndpointProperty{
//   					AccessEndpointType: jsii.String("accessEndpointType"),
//   					VpcEndpointId: jsii.String("vpcEndpointId"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			InternetFallbackProtocols: []*string{
//   				jsii.String("internetFallbackProtocols"),
//   			},
//   		},
//   		DeviceTypeAndroid: jsii.String("deviceTypeAndroid"),
//   		DeviceTypeChromeOs: jsii.String("deviceTypeChromeOs"),
//   		DeviceTypeIos: jsii.String("deviceTypeIos"),
//   		DeviceTypeLinux: jsii.String("deviceTypeLinux"),
//   		DeviceTypeOsx: jsii.String("deviceTypeOsx"),
//   		DeviceTypeWeb: jsii.String("deviceTypeWeb"),
//   		DeviceTypeWindows: jsii.String("deviceTypeWindows"),
//   		DeviceTypeWorkSpacesThinClient: jsii.String("deviceTypeWorkSpacesThinClient"),
//   		DeviceTypeZeroClient: jsii.String("deviceTypeZeroClient"),
//   	},
//   	WorkspaceCreationProperties: &DefaultWorkspaceCreationPropertiesProperty{
//   		CustomSecurityGroupId: jsii.String("customSecurityGroupId"),
//   		DefaultOu: jsii.String("defaultOu"),
//   		EnableInternetAccess: jsii.Boolean(false),
//   		EnableMaintenanceMode: jsii.Boolean(false),
//   		InstanceIamRoleArn: jsii.String("instanceIamRoleArn"),
//   		UserEnabledAsLocalAdministrator: jsii.Boolean(false),
//   	},
//   	WorkspaceDirectoryDescription: jsii.String("workspaceDirectoryDescription"),
//   	WorkspaceDirectoryName: jsii.String("workspaceDirectoryName"),
//   	WorkspaceType: jsii.String("workspaceType"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html
//
type CfnDirectory interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsworkspaces.IDirectoryRef
	awscdk.ITaggableV2
	// Information about the Active Directory config.
	ActiveDirectoryConfig() interface{}
	SetActiveDirectoryConfig(val interface{})
	// The directory alias.
	AttrAlias() *string
	// The Amazon Resource Name (ARN) of the WorkSpace directory.
	AttrArn() *string
	// The user name for the service account.
	AttrCustomerUserName() *string
	// The directory identifier.
	AttrDirectoryId() *string
	// The name of the directory.
	AttrDirectoryName() *string
	// The directory type.
	AttrDirectoryType() *string
	// The IP addresses of the DNS servers for the directory.
	AttrDnsIpAddresses() *[]*string
	// The IPv6 addresses of the DNS servers for the directory.
	AttrDnsIpv6Addresses() *[]*string
	// The identifier of the IAM role.
	AttrIamRoleId() *string
	// Specifies the configurations of the identity center.
	AttrIdcConfig() awscdk.IResolvable
	// The registration code for the directory.
	AttrRegistrationCode() *string
	// The state of the directory's registration with Amazon WorkSpaces.
	AttrState() *string
	// The identifier of the security group that is assigned to new WorkSpaces.
	AttrWorkspaceSecurityGroupId() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Describes the properties of the certificate-based authentication you want to use with your WorkSpaces.
	CertificateBasedAuthProperties() interface{}
	SetCertificateBasedAuthProperties(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A reference to a Directory resource.
	DirectoryRef() *interfacesawsworkspaces.DirectoryReference
	// Indicates whether self-service capabilities are enabled or disabled.
	EnableSelfService() interface{}
	SetEnableSelfService(val interface{})
	// Endpoint encryption mode that allows you to configure the specified directory between Standard TLS and FIPS 140-2 validated mode.
	EndpointEncryptionMode() *string
	SetEndpointEncryptionMode(val *string)
	Env() *interfaces.ResourceEnvironment
	// The Amazon Resource Name (ARN) of the identity center instance.
	IdcInstanceArn() *string
	SetIdcInstanceArn(val *string)
	// The identifiers of the IP access control groups associated with the directory.
	IpGroupIds() *[]*string
	SetIpGroupIds(val *[]*string)
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
	// Specifies the configurations of the Microsoft Entra.
	MicrosoftEntraConfig() interface{}
	SetMicrosoftEntraConfig(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Describes the enablement status, user access URL, and relay state parameter name that are used for configuring federation with an SAML 2.0 identity provider.
	SamlProperties() interface{}
	SetSamlProperties(val interface{})
	// Describes the self-service permissions for a directory.
	SelfservicePermissions() interface{}
	SetSelfservicePermissions(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Describes the streaming properties.
	StreamingProperties() interface{}
	SetStreamingProperties(val interface{})
	// The identifiers of the subnets for your virtual private cloud (VPC).
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// The tags associated with the directory.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// Indicates whether your WorkSpace directory is dedicated or shared.
	Tenancy() *string
	SetTenancy(val *string)
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
	// The type of identity management the user is using.
	UserIdentityType() *string
	SetUserIdentityType(val *string)
	// The device types and operating systems that can be used to access a WorkSpace.
	WorkspaceAccessProperties() interface{}
	SetWorkspaceAccessProperties(val interface{})
	// The default values that are used to create WorkSpaces.
	WorkspaceCreationProperties() interface{}
	SetWorkspaceCreationProperties(val interface{})
	// Description of the directory to register.
	WorkspaceDirectoryDescription() *string
	SetWorkspaceDirectoryDescription(val *string)
	// The name of the directory to register.
	WorkspaceDirectoryName() *string
	SetWorkspaceDirectoryName(val *string)
	// Indicates whether the directory's WorkSpace type is personal or pools.
	WorkspaceType() *string
	SetWorkspaceType(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This method has been renamed to `addResourceDependency` to more clearly
	// set it apart from `construct.node.addDependency`. See the documentation
	// of that function for more details.
	// Deprecated: Use `addResourceDependency` instead.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method has been renamed to `addResourceDependency`, which makes it
	// more clear that this method operates at a different level from the
	// construct-level `construct.node.addDependency()` mechanism.
	// Deprecated: Use `addResourceDependency` instead.
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
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method only adds dependencies between L1 resources. If you are
	// looking for a generic construct-to-construct dependency mechanism that works
	// for all constructs including L2s, use `construct.node.addDependency` instead.
	AddResourceDependency(target awscdk.CfnResource, reason *string)
	// Sets the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// strength instead of the global default from the consuming stack's context.
	ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength)
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
	CfnPropertyName(cdkPropertyName *string) *string
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
	// Retrieves an array of resources and stacks this resource depends on.
	//
	// For resources depended on directly, returns the `CfnResource` object. For
	// dependencies on other stacks, returns the `Stack` object. The order of the
	// array is not guaranteed.
	ObtainDependencies() *[]interface{}
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	// Deprecated: Use `removeResourceDependency` instead.
	RemoveDependency(target awscdk.CfnResource)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveResourceDependency(target awscdk.CfnResource)
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnDirectory
type jsiiProxy_CfnDirectory struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsworkspacesIDirectoryRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnDirectory) ActiveDirectoryConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrCustomerUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCustomerUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrDirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDirectoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrDirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDirectoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrDirectoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDirectoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrDnsIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDnsIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrDnsIpv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDnsIpv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrIamRoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIamRoleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrIdcConfig() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIdcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrRegistrationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRegistrationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) AttrWorkspaceSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWorkspaceSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) CertificateBasedAuthProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateBasedAuthProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) DirectoryRef() *interfacesawsworkspaces.DirectoryReference {
	var returns *interfacesawsworkspaces.DirectoryReference
	_jsii_.Get(
		j,
		"directoryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) EnableSelfService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSelfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) EndpointEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) IdcInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) IpGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) MicrosoftEntraConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftEntraConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) SamlProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) SelfservicePermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfservicePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) StreamingProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) UserIdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdentityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) WorkspaceAccessProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceAccessProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) WorkspaceCreationProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceCreationProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) WorkspaceDirectoryDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceDirectoryDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) WorkspaceDirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceDirectoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectory) WorkspaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceType",
		&returns,
	)
	return returns
}


// Create a new `AWS::WorkSpaces::Directory`.
func NewCfnDirectory(scope constructs.Construct, id *string, props *CfnDirectoryProps) CfnDirectory {
	_init_.Initialize()

	if err := validateNewCfnDirectoryParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDirectory{}

	_jsii_.Create(
		"aws-cdk-lib.aws_workspaces.CfnDirectory",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WorkSpaces::Directory`.
func NewCfnDirectory_Override(c CfnDirectory, scope constructs.Construct, id *string, props *CfnDirectoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_workspaces.CfnDirectory",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDirectory)SetActiveDirectoryConfig(val interface{}) {
	if err := j.validateSetActiveDirectoryConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDirectoryConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetCertificateBasedAuthProperties(val interface{}) {
	if err := j.validateSetCertificateBasedAuthPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateBasedAuthProperties",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetEnableSelfService(val interface{}) {
	if err := j.validateSetEnableSelfServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSelfService",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetEndpointEncryptionMode(val *string) {
	_jsii_.Set(
		j,
		"endpointEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetIdcInstanceArn(val *string) {
	_jsii_.Set(
		j,
		"idcInstanceArn",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetIpGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"ipGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetMicrosoftEntraConfig(val interface{}) {
	if err := j.validateSetMicrosoftEntraConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftEntraConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetSamlProperties(val interface{}) {
	if err := j.validateSetSamlPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlProperties",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetSelfservicePermissions(val interface{}) {
	if err := j.validateSetSelfservicePermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfservicePermissions",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetStreamingProperties(val interface{}) {
	if err := j.validateSetStreamingPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingProperties",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetTenancy(val *string) {
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetUserIdentityType(val *string) {
	_jsii_.Set(
		j,
		"userIdentityType",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetWorkspaceAccessProperties(val interface{}) {
	if err := j.validateSetWorkspaceAccessPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceAccessProperties",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetWorkspaceCreationProperties(val interface{}) {
	if err := j.validateSetWorkspaceCreationPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceCreationProperties",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetWorkspaceDirectoryDescription(val *string) {
	_jsii_.Set(
		j,
		"workspaceDirectoryDescription",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetWorkspaceDirectoryName(val *string) {
	_jsii_.Set(
		j,
		"workspaceDirectoryName",
		val,
	)
}

func (j *jsiiProxy_CfnDirectory)SetWorkspaceType(val *string) {
	_jsii_.Set(
		j,
		"workspaceType",
		val,
	)
}

func CfnDirectory_ArnForDirectory(resource interfacesawsworkspaces.IDirectoryRef) *string {
	_init_.Initialize()

	if err := validateCfnDirectory_ArnForDirectoryParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspaces.CfnDirectory",
		"arnForDirectory",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnDirectory.
func CfnDirectory_IsCfnDirectory(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDirectory_IsCfnDirectoryParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspaces.CfnDirectory",
		"isCfnDirectory",
		[]interface{}{x},
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
func CfnDirectory_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDirectory_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspaces.CfnDirectory",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDirectory_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDirectory_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspaces.CfnDirectory",
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
func CfnDirectory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDirectory_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspaces.CfnDirectory",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDirectory_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_workspaces.CfnDirectory",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDirectory) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDirectory) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDirectory) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDirectory) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDirectory) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDirectory) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDirectory) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDirectory) AddResourceDependency(target awscdk.CfnResource, reason *string) {
	if err := c.validateAddResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addResourceDependency",
		[]interface{}{target, reason},
	)
}

func (c *jsiiProxy_CfnDirectory) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnDirectory) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDirectory) CfnPropertyName(cdkPropertyName *string) *string {
	if err := c.validateCfnPropertyNameParameters(cdkPropertyName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"cfnPropertyName",
		[]interface{}{cdkPropertyName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDirectory) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnDirectory) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnDirectory) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDirectory) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDirectory) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDirectory) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDirectory) RemoveResourceDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeResourceDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDirectory) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnDirectory) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDirectory) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDirectory) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDirectory) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnDirectory) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

