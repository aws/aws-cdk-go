package awsopensearchservice

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::OpenSearchService::Domain resource creates an Amazon OpenSearch Service domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessPolicies interface{}
//
//   cfnDomain := awscdk.Aws_opensearchservice.NewCfnDomain(this, jsii.String("MyCfnDomain"), &CfnDomainProps{
//   	AccessPolicies: accessPolicies,
//   	AdvancedOptions: map[string]*string{
//   		"advancedOptionsKey": jsii.String("advancedOptions"),
//   	},
//   	AdvancedSecurityOptions: &AdvancedSecurityOptionsInputProperty{
//   		AnonymousAuthDisableDate: jsii.String("anonymousAuthDisableDate"),
//   		AnonymousAuthEnabled: jsii.Boolean(false),
//   		Enabled: jsii.Boolean(false),
//   		InternalUserDatabaseEnabled: jsii.Boolean(false),
//   		JwtOptions: &JWTOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			PublicKey: jsii.String("publicKey"),
//   			RolesKey: jsii.String("rolesKey"),
//   			SubjectKey: jsii.String("subjectKey"),
//   		},
//   		MasterUserOptions: &MasterUserOptionsProperty{
//   			MasterUserArn: jsii.String("masterUserArn"),
//   			MasterUserName: jsii.String("masterUserName"),
//   			MasterUserPassword: jsii.String("masterUserPassword"),
//   		},
//   		SamlOptions: &SAMLOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			Idp: &IdpProperty{
//   				EntityId: jsii.String("entityId"),
//   				MetadataContent: jsii.String("metadataContent"),
//   			},
//   			MasterBackendRole: jsii.String("masterBackendRole"),
//   			MasterUserName: jsii.String("masterUserName"),
//   			RolesKey: jsii.String("rolesKey"),
//   			SessionTimeoutMinutes: jsii.Number(123),
//   			SubjectKey: jsii.String("subjectKey"),
//   		},
//   	},
//   	ClusterConfig: &ClusterConfigProperty{
//   		ColdStorageOptions: &ColdStorageOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		DedicatedMasterCount: jsii.Number(123),
//   		DedicatedMasterEnabled: jsii.Boolean(false),
//   		DedicatedMasterType: jsii.String("dedicatedMasterType"),
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//   		MultiAzWithStandbyEnabled: jsii.Boolean(false),
//   		NodeOptions: []interface{}{
//   			&NodeOptionProperty{
//   				NodeConfig: &NodeConfigProperty{
//   					Count: jsii.Number(123),
//   					Enabled: jsii.Boolean(false),
//   					Type: jsii.String("type"),
//   				},
//   				NodeType: jsii.String("nodeType"),
//   			},
//   		},
//   		WarmCount: jsii.Number(123),
//   		WarmEnabled: jsii.Boolean(false),
//   		WarmType: jsii.String("warmType"),
//   		ZoneAwarenessConfig: &ZoneAwarenessConfigProperty{
//   			AvailabilityZoneCount: jsii.Number(123),
//   		},
//   		ZoneAwarenessEnabled: jsii.Boolean(false),
//   	},
//   	CognitoOptions: &CognitoOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		IdentityPoolId: jsii.String("identityPoolId"),
//   		RoleArn: jsii.String("roleArn"),
//   		UserPoolId: jsii.String("userPoolId"),
//   	},
//   	DomainArn: jsii.String("domainArn"),
//   	DomainEndpointOptions: &DomainEndpointOptionsProperty{
//   		CustomEndpoint: jsii.String("customEndpoint"),
//   		CustomEndpointCertificateArn: jsii.String("customEndpointCertificateArn"),
//   		CustomEndpointEnabled: jsii.Boolean(false),
//   		EnforceHttps: jsii.Boolean(false),
//   		TlsSecurityPolicy: jsii.String("tlsSecurityPolicy"),
//   	},
//   	DomainName: jsii.String("domainName"),
//   	EbsOptions: &EBSOptionsProperty{
//   		EbsEnabled: jsii.Boolean(false),
//   		Iops: jsii.Number(123),
//   		Throughput: jsii.Number(123),
//   		VolumeSize: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	EncryptionAtRestOptions: &EncryptionAtRestOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	EngineVersion: jsii.String("engineVersion"),
//   	IdentityCenterOptions: &IdentityCenterOptionsProperty{
//   		EnabledApiAccess: jsii.Boolean(false),
//   		IdentityCenterApplicationArn: jsii.String("identityCenterApplicationArn"),
//   		IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   		IdentityStoreId: jsii.String("identityStoreId"),
//   		RolesKey: jsii.String("rolesKey"),
//   		SubjectKey: jsii.String("subjectKey"),
//   	},
//   	IpAddressType: jsii.String("ipAddressType"),
//   	LogPublishingOptions: map[string]interface{}{
//   		"logPublishingOptionsKey": &LogPublishingOptionProperty{
//   			"cloudWatchLogsLogGroupArn": jsii.String("cloudWatchLogsLogGroupArn"),
//   			"enabled": jsii.Boolean(false),
//   		},
//   	},
//   	NodeToNodeEncryptionOptions: &NodeToNodeEncryptionOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	OffPeakWindowOptions: &OffPeakWindowOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		OffPeakWindow: &OffPeakWindowProperty{
//   			WindowStartTime: &WindowStartTimeProperty{
//   				Hours: jsii.Number(123),
//   				Minutes: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SkipShardMigrationWait: jsii.Boolean(false),
//   	SnapshotOptions: &SnapshotOptionsProperty{
//   		AutomatedSnapshotStartHour: jsii.Number(123),
//   	},
//   	SoftwareUpdateOptions: &SoftwareUpdateOptionsProperty{
//   		AutoSoftwareUpdateEnabled: jsii.Boolean(false),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcOptions: &VPCOptionsProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html
//
type CfnDomain interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// An AWS Identity and Access Management ( IAM ) policy document that specifies who can access the OpenSearch Service domain and their permissions.
	AccessPolicies() interface{}
	SetAccessPolicies(val interface{})
	// Additional options to specify for the OpenSearch Service domain.
	AdvancedOptions() interface{}
	SetAdvancedOptions(val interface{})
	// Specifies options for fine-grained access control and SAML authentication.
	AdvancedSecurityOptions() interface{}
	SetAdvancedSecurityOptions(val interface{})
	// Date and time when the migration period will be disabled.
	//
	// Only necessary when [enabling fine-grained access control on an existing domain](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html#fgac-enabling-existing) .
	AttrAdvancedSecurityOptionsAnonymousAuthDisableDate() *string
	// The Amazon Resource Name (ARN) of the CloudFormation stack.
	AttrArn() *string
	// The domain-specific endpoint used for requests to the OpenSearch APIs, such as `search-mystack-1ab2cdefghij-ab1c2deckoyb3hofw7wpqa3cm.us-west-1.es.amazonaws.com` .
	AttrDomainEndpoint() *string
	AttrDomainEndpoints() awscdk.IResolvable
	// If `IPAddressType` to set to `dualstack` , a version 2 domain endpoint is provisioned.
	//
	// This endpoint functions like a normal endpoint, except that it works with both IPv4 and IPv6 IP addresses. Normal endpoints work only with IPv4 IP addresses.
	AttrDomainEndpointV2() *string
	// The resource ID.
	//
	// For example, `123456789012/my-domain` .
	AttrId() *string
	// The ARN of the IAM Identity Center application that integrates with Amazon OpenSearch Service.
	AttrIdentityCenterOptionsIdentityCenterApplicationArn() *string
	// The identifier of the IAM Identity Store.
	AttrIdentityCenterOptionsIdentityStoreId() *string
	AttrServiceSoftwareOptions() awscdk.IResolvable
	AttrServiceSoftwareOptionsAutomatedUpdateDate() *string
	AttrServiceSoftwareOptionsCancellable() awscdk.IResolvable
	AttrServiceSoftwareOptionsCurrentVersion() *string
	AttrServiceSoftwareOptionsDescription() *string
	AttrServiceSoftwareOptionsNewVersion() *string
	AttrServiceSoftwareOptionsOptionalDeployment() awscdk.IResolvable
	AttrServiceSoftwareOptionsUpdateAvailable() awscdk.IResolvable
	AttrServiceSoftwareOptionsUpdateStatus() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Container for the cluster configuration of a domain.
	ClusterConfig() interface{}
	SetClusterConfig(val interface{})
	// Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
	CognitoOptions() interface{}
	SetCognitoOptions(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	DomainArn() *string
	SetDomainArn(val *string)
	// Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.
	DomainEndpointOptions() interface{}
	SetDomainEndpointOptions(val interface{})
	// A name for the OpenSearch Service domain.
	DomainName() *string
	SetDomainName(val *string)
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the OpenSearch Service domain.
	EbsOptions() interface{}
	SetEbsOptions(val interface{})
	// Whether the domain should encrypt data at rest, and if so, the AWS KMS key to use.
	EncryptionAtRestOptions() interface{}
	SetEncryptionAtRestOptions(val interface{})
	// The version of OpenSearch to use.
	EngineVersion() *string
	SetEngineVersion(val *string)
	// Configuration options for controlling IAM Identity Center integration within a domain.
	IdentityCenterOptions() interface{}
	SetIdentityCenterOptions(val interface{})
	// Choose either dual stack or IPv4 as your IP address type.
	IpAddressType() *string
	SetIpAddressType(val *string)
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
	// An object with one or more of the following keys: `SEARCH_SLOW_LOGS` , `ES_APPLICATION_LOGS` , `INDEX_SLOW_LOGS` , `AUDIT_LOGS` , depending on the types of logs you want to publish.
	LogPublishingOptions() interface{}
	SetLogPublishingOptions(val interface{})
	// The tree node.
	Node() constructs.Node
	// Specifies whether node-to-node encryption is enabled.
	NodeToNodeEncryptionOptions() interface{}
	SetNodeToNodeEncryptionOptions(val interface{})
	// Options for a domain's off-peak window, during which OpenSearch Service can perform mandatory configuration changes on the domain.
	OffPeakWindowOptions() interface{}
	SetOffPeakWindowOptions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	SkipShardMigrationWait() interface{}
	SetSkipShardMigrationWait(val interface{})
	// *DEPRECATED* .
	SnapshotOptions() interface{}
	SetSnapshotOptions(val interface{})
	// Service software update options for the domain.
	SoftwareUpdateOptions() interface{}
	SetSoftwareUpdateOptions(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// An arbitrary set of tags (key–value pairs) to associate with the OpenSearch Service domain.
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
	// The virtual private cloud (VPC) configuration for the OpenSearch Service domain.
	VpcOptions() interface{}
	SetVpcOptions(val interface{})
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

// The jsii proxy struct for CfnDomain
type jsiiProxy_CfnDomain struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnDomain) AccessPolicies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AdvancedOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AdvancedSecurityOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrAdvancedSecurityOptionsAnonymousAuthDisableDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAdvancedSecurityOptionsAnonymousAuthDisableDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrDomainEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrDomainEndpoints() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrDomainEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrDomainEndpointV2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainEndpointV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrIdentityCenterOptionsIdentityCenterApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIdentityCenterOptionsIdentityCenterApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrIdentityCenterOptionsIdentityStoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIdentityCenterOptionsIdentityStoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrServiceSoftwareOptions() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrServiceSoftwareOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrServiceSoftwareOptionsAutomatedUpdateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceSoftwareOptionsAutomatedUpdateDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrServiceSoftwareOptionsCancellable() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrServiceSoftwareOptionsCancellable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrServiceSoftwareOptionsCurrentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceSoftwareOptionsCurrentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrServiceSoftwareOptionsDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceSoftwareOptionsDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrServiceSoftwareOptionsNewVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceSoftwareOptionsNewVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrServiceSoftwareOptionsOptionalDeployment() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrServiceSoftwareOptionsOptionalDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrServiceSoftwareOptionsUpdateAvailable() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrServiceSoftwareOptionsUpdateAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrServiceSoftwareOptionsUpdateStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceSoftwareOptionsUpdateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) ClusterConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CognitoOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cognitoOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DomainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DomainEndpointOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainEndpointOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) EbsOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) EncryptionAtRestOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtRestOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) IdentityCenterOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityCenterOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) LogPublishingOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) NodeToNodeEncryptionOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeToNodeEncryptionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) OffPeakWindowOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"offPeakWindowOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) SkipShardMigrationWait() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipShardMigrationWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) SnapshotOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) SoftwareUpdateOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"softwareUpdateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) VpcOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcOptions",
		&returns,
	)
	return returns
}


func NewCfnDomain(scope constructs.Construct, id *string, props *CfnDomainProps) CfnDomain {
	_init_.Initialize()

	if err := validateNewCfnDomainParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDomain{}

	_jsii_.Create(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnDomain_Override(c CfnDomain, scope constructs.Construct, id *string, props *CfnDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomain)SetAccessPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"accessPolicies",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetAdvancedOptions(val interface{}) {
	if err := j.validateSetAdvancedOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetAdvancedSecurityOptions(val interface{}) {
	if err := j.validateSetAdvancedSecurityOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedSecurityOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetClusterConfig(val interface{}) {
	if err := j.validateSetClusterConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetCognitoOptions(val interface{}) {
	if err := j.validateSetCognitoOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cognitoOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetDomainArn(val *string) {
	_jsii_.Set(
		j,
		"domainArn",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetDomainEndpointOptions(val interface{}) {
	if err := j.validateSetDomainEndpointOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainEndpointOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetEbsOptions(val interface{}) {
	if err := j.validateSetEbsOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetEncryptionAtRestOptions(val interface{}) {
	if err := j.validateSetEncryptionAtRestOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAtRestOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetIdentityCenterOptions(val interface{}) {
	if err := j.validateSetIdentityCenterOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityCenterOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetIpAddressType(val *string) {
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetLogPublishingOptions(val interface{}) {
	if err := j.validateSetLogPublishingOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logPublishingOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetNodeToNodeEncryptionOptions(val interface{}) {
	if err := j.validateSetNodeToNodeEncryptionOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeToNodeEncryptionOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetOffPeakWindowOptions(val interface{}) {
	if err := j.validateSetOffPeakWindowOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offPeakWindowOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetSkipShardMigrationWait(val interface{}) {
	if err := j.validateSetSkipShardMigrationWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipShardMigrationWait",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetSnapshotOptions(val interface{}) {
	if err := j.validateSetSnapshotOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetSoftwareUpdateOptions(val interface{}) {
	if err := j.validateSetSoftwareUpdateOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"softwareUpdateOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnDomain)SetVpcOptions(val interface{}) {
	if err := j.validateSetVpcOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcOptions",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDomain_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomain_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDomain_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomain_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain",
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
func CfnDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomain_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomain) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDomain) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomain) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomain) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDomain) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDomain) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDomain) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDomain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDomain) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnDomain) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnDomain) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDomain) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDomain) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomain) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnDomain) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDomain) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

