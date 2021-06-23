package awselasticsearch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticsearch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies options for fine-grained access control.
// Experimental.
type AdvancedSecurityOptions struct {
	// ARN for the master user.
	//
	// Only specify this or masterUserName, but not both.
	// Experimental.
	MasterUserArn *string `json:"masterUserArn"`
	// Username for the master user.
	//
	// Only specify this or masterUserArn, but not both.
	// Experimental.
	MasterUserName *string `json:"masterUserName"`
	// Password for the master user.
	//
	// You can use `SecretValue.plainText` to specify a password in plain text or
	// use `secretsmanager.Secret.fromSecretAttributes` to reference a secret in
	// Secrets Manager.
	// Experimental.
	MasterUserPassword awscdk.SecretValue `json:"masterUserPassword"`
}

// Configures the capacity of the cluster such as the instance type and the number of instances.
// Experimental.
type CapacityConfig struct {
	// The instance type for your data nodes, such as `m3.medium.elasticsearch`. For valid values, see [Supported Instance Types](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html) in the Amazon Elasticsearch Service Developer Guide.
	// Experimental.
	DataNodeInstanceType *string `json:"dataNodeInstanceType"`
	// The number of data nodes (instances) to use in the Amazon ES domain.
	// Experimental.
	DataNodes *float64 `json:"dataNodes"`
	// The hardware configuration of the computer that hosts the dedicated master node, such as `m3.medium.elasticsearch`. For valid values, see [Supported Instance Types] (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html) in the Amazon Elasticsearch Service Developer Guide.
	// Experimental.
	MasterNodeInstanceType *string `json:"masterNodeInstanceType"`
	// The number of instances to use for the master node.
	// Experimental.
	MasterNodes *float64 `json:"masterNodes"`
	// The instance type for your UltraWarm node, such as `ultrawarm1.medium.elasticsearch`. For valid values, see [UltraWarm Storage Limits] (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-limits.html#limits-ultrawarm) in the Amazon Elasticsearch Service Developer Guide.
	// Experimental.
	WarmInstanceType *string `json:"warmInstanceType"`
	// The number of UltraWarm nodes (instances) to use in the Amazon ES domain.
	// Experimental.
	WarmNodes *float64 `json:"warmNodes"`
}

// A CloudFormation `AWS::Elasticsearch::Domain`.
type CfnDomain interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AccessPolicies() interface{}
	SetAccessPolicies(val interface{})
	AdvancedOptions() interface{}
	SetAdvancedOptions(val interface{})
	AdvancedSecurityOptions() interface{}
	SetAdvancedSecurityOptions(val interface{})
	AttrArn() *string
	AttrDomainEndpoint() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CognitoOptions() interface{}
	SetCognitoOptions(val interface{})
	CreationStack() *[]*string
	DomainEndpointOptions() interface{}
	SetDomainEndpointOptions(val interface{})
	DomainName() *string
	SetDomainName(val *string)
	EbsOptions() interface{}
	SetEbsOptions(val interface{})
	ElasticsearchClusterConfig() interface{}
	SetElasticsearchClusterConfig(val interface{})
	ElasticsearchVersion() *string
	SetElasticsearchVersion(val *string)
	EncryptionAtRestOptions() interface{}
	SetEncryptionAtRestOptions(val interface{})
	LogicalId() *string
	LogPublishingOptions() interface{}
	SetLogPublishingOptions(val interface{})
	Node() constructs.Node
	NodeToNodeEncryptionOptions() interface{}
	SetNodeToNodeEncryptionOptions(val interface{})
	Ref() *string
	SnapshotOptions() interface{}
	SetSnapshotOptions(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VpcOptions() interface{}
	SetVpcOptions(val interface{})
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDomain
type jsiiProxy_CfnDomain struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnDomain) ElasticsearchClusterConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) ElasticsearchVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchVersion",
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

func (j *jsiiProxy_CfnDomain) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
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

func (j *jsiiProxy_CfnDomain) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
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


// Create a new `AWS::Elasticsearch::Domain`.
func NewCfnDomain(scope constructs.Construct, id *string, props *CfnDomainProps) CfnDomain {
	_init_.Initialize()

	j := jsiiProxy_CfnDomain{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Elasticsearch::Domain`.
func NewCfnDomain_Override(c CfnDomain, scope constructs.Construct, id *string, props *CfnDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomain) SetAccessPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"accessPolicies",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetAdvancedOptions(val interface{}) {
	_jsii_.Set(
		j,
		"advancedOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetAdvancedSecurityOptions(val interface{}) {
	_jsii_.Set(
		j,
		"advancedSecurityOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetCognitoOptions(val interface{}) {
	_jsii_.Set(
		j,
		"cognitoOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDomainEndpointOptions(val interface{}) {
	_jsii_.Set(
		j,
		"domainEndpointOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetEbsOptions(val interface{}) {
	_jsii_.Set(
		j,
		"ebsOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetElasticsearchClusterConfig(val interface{}) {
	_jsii_.Set(
		j,
		"elasticsearchClusterConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetElasticsearchVersion(val *string) {
	_jsii_.Set(
		j,
		"elasticsearchVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetEncryptionAtRestOptions(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionAtRestOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetLogPublishingOptions(val interface{}) {
	_jsii_.Set(
		j,
		"logPublishingOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetNodeToNodeEncryptionOptions(val interface{}) {
	_jsii_.Set(
		j,
		"nodeToNodeEncryptionOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetSnapshotOptions(val interface{}) {
	_jsii_.Set(
		j,
		"snapshotOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetVpcOptions(val interface{}) {
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
// Experimental.
func CfnDomain_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDomain_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
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
		"aws-cdk-lib.aws_elasticsearch.CfnDomain",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDomain) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnDomain) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnDomain) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDomain) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnDomain) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnDomain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnDomain) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnDomain) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDomain) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnDomain) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDomain) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDomain) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDomain_AdvancedSecurityOptionsInputProperty struct {
	// `CfnDomain.AdvancedSecurityOptionsInputProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnDomain.AdvancedSecurityOptionsInputProperty.InternalUserDatabaseEnabled`.
	InternalUserDatabaseEnabled interface{} `json:"internalUserDatabaseEnabled"`
	// `CfnDomain.AdvancedSecurityOptionsInputProperty.MasterUserOptions`.
	MasterUserOptions interface{} `json:"masterUserOptions"`
}

type CfnDomain_CognitoOptionsProperty struct {
	// `CfnDomain.CognitoOptionsProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnDomain.CognitoOptionsProperty.IdentityPoolId`.
	IdentityPoolId *string `json:"identityPoolId"`
	// `CfnDomain.CognitoOptionsProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnDomain.CognitoOptionsProperty.UserPoolId`.
	UserPoolId *string `json:"userPoolId"`
}

type CfnDomain_DomainEndpointOptionsProperty struct {
	// `CfnDomain.DomainEndpointOptionsProperty.CustomEndpoint`.
	CustomEndpoint *string `json:"customEndpoint"`
	// `CfnDomain.DomainEndpointOptionsProperty.CustomEndpointCertificateArn`.
	CustomEndpointCertificateArn *string `json:"customEndpointCertificateArn"`
	// `CfnDomain.DomainEndpointOptionsProperty.CustomEndpointEnabled`.
	CustomEndpointEnabled interface{} `json:"customEndpointEnabled"`
	// `CfnDomain.DomainEndpointOptionsProperty.EnforceHTTPS`.
	EnforceHttps interface{} `json:"enforceHttps"`
	// `CfnDomain.DomainEndpointOptionsProperty.TLSSecurityPolicy`.
	TlsSecurityPolicy *string `json:"tlsSecurityPolicy"`
}

type CfnDomain_EBSOptionsProperty struct {
	// `CfnDomain.EBSOptionsProperty.EBSEnabled`.
	EbsEnabled interface{} `json:"ebsEnabled"`
	// `CfnDomain.EBSOptionsProperty.Iops`.
	Iops *float64 `json:"iops"`
	// `CfnDomain.EBSOptionsProperty.VolumeSize`.
	VolumeSize *float64 `json:"volumeSize"`
	// `CfnDomain.EBSOptionsProperty.VolumeType`.
	VolumeType *string `json:"volumeType"`
}

type CfnDomain_ElasticsearchClusterConfigProperty struct {
	// `CfnDomain.ElasticsearchClusterConfigProperty.DedicatedMasterCount`.
	DedicatedMasterCount *float64 `json:"dedicatedMasterCount"`
	// `CfnDomain.ElasticsearchClusterConfigProperty.DedicatedMasterEnabled`.
	DedicatedMasterEnabled interface{} `json:"dedicatedMasterEnabled"`
	// `CfnDomain.ElasticsearchClusterConfigProperty.DedicatedMasterType`.
	DedicatedMasterType *string `json:"dedicatedMasterType"`
	// `CfnDomain.ElasticsearchClusterConfigProperty.InstanceCount`.
	InstanceCount *float64 `json:"instanceCount"`
	// `CfnDomain.ElasticsearchClusterConfigProperty.InstanceType`.
	InstanceType *string `json:"instanceType"`
	// `CfnDomain.ElasticsearchClusterConfigProperty.WarmCount`.
	WarmCount *float64 `json:"warmCount"`
	// `CfnDomain.ElasticsearchClusterConfigProperty.WarmEnabled`.
	WarmEnabled interface{} `json:"warmEnabled"`
	// `CfnDomain.ElasticsearchClusterConfigProperty.WarmType`.
	WarmType *string `json:"warmType"`
	// `CfnDomain.ElasticsearchClusterConfigProperty.ZoneAwarenessConfig`.
	ZoneAwarenessConfig interface{} `json:"zoneAwarenessConfig"`
	// `CfnDomain.ElasticsearchClusterConfigProperty.ZoneAwarenessEnabled`.
	ZoneAwarenessEnabled interface{} `json:"zoneAwarenessEnabled"`
}

type CfnDomain_EncryptionAtRestOptionsProperty struct {
	// `CfnDomain.EncryptionAtRestOptionsProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnDomain.EncryptionAtRestOptionsProperty.KmsKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
}

type CfnDomain_LogPublishingOptionProperty struct {
	// `CfnDomain.LogPublishingOptionProperty.CloudWatchLogsLogGroupArn`.
	CloudWatchLogsLogGroupArn *string `json:"cloudWatchLogsLogGroupArn"`
	// `CfnDomain.LogPublishingOptionProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

type CfnDomain_MasterUserOptionsProperty struct {
	// `CfnDomain.MasterUserOptionsProperty.MasterUserARN`.
	MasterUserArn *string `json:"masterUserArn"`
	// `CfnDomain.MasterUserOptionsProperty.MasterUserName`.
	MasterUserName *string `json:"masterUserName"`
	// `CfnDomain.MasterUserOptionsProperty.MasterUserPassword`.
	MasterUserPassword *string `json:"masterUserPassword"`
}

type CfnDomain_NodeToNodeEncryptionOptionsProperty struct {
	// `CfnDomain.NodeToNodeEncryptionOptionsProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

type CfnDomain_SnapshotOptionsProperty struct {
	// `CfnDomain.SnapshotOptionsProperty.AutomatedSnapshotStartHour`.
	AutomatedSnapshotStartHour *float64 `json:"automatedSnapshotStartHour"`
}

type CfnDomain_VPCOptionsProperty struct {
	// `CfnDomain.VPCOptionsProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `CfnDomain.VPCOptionsProperty.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
}

type CfnDomain_ZoneAwarenessConfigProperty struct {
	// `CfnDomain.ZoneAwarenessConfigProperty.AvailabilityZoneCount`.
	AvailabilityZoneCount *float64 `json:"availabilityZoneCount"`
}

// Properties for defining a `AWS::Elasticsearch::Domain`.
type CfnDomainProps struct {
	// `AWS::Elasticsearch::Domain.AccessPolicies`.
	AccessPolicies interface{} `json:"accessPolicies"`
	// `AWS::Elasticsearch::Domain.AdvancedOptions`.
	AdvancedOptions interface{} `json:"advancedOptions"`
	// `AWS::Elasticsearch::Domain.AdvancedSecurityOptions`.
	AdvancedSecurityOptions interface{} `json:"advancedSecurityOptions"`
	// `AWS::Elasticsearch::Domain.CognitoOptions`.
	CognitoOptions interface{} `json:"cognitoOptions"`
	// `AWS::Elasticsearch::Domain.DomainEndpointOptions`.
	DomainEndpointOptions interface{} `json:"domainEndpointOptions"`
	// `AWS::Elasticsearch::Domain.DomainName`.
	DomainName *string `json:"domainName"`
	// `AWS::Elasticsearch::Domain.EBSOptions`.
	EbsOptions interface{} `json:"ebsOptions"`
	// `AWS::Elasticsearch::Domain.ElasticsearchClusterConfig`.
	ElasticsearchClusterConfig interface{} `json:"elasticsearchClusterConfig"`
	// `AWS::Elasticsearch::Domain.ElasticsearchVersion`.
	ElasticsearchVersion *string `json:"elasticsearchVersion"`
	// `AWS::Elasticsearch::Domain.EncryptionAtRestOptions`.
	EncryptionAtRestOptions interface{} `json:"encryptionAtRestOptions"`
	// `AWS::Elasticsearch::Domain.LogPublishingOptions`.
	LogPublishingOptions interface{} `json:"logPublishingOptions"`
	// `AWS::Elasticsearch::Domain.NodeToNodeEncryptionOptions`.
	NodeToNodeEncryptionOptions interface{} `json:"nodeToNodeEncryptionOptions"`
	// `AWS::Elasticsearch::Domain.SnapshotOptions`.
	SnapshotOptions interface{} `json:"snapshotOptions"`
	// `AWS::Elasticsearch::Domain.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::Elasticsearch::Domain.VPCOptions`.
	VpcOptions interface{} `json:"vpcOptions"`
}

// Configures Amazon ES to use Amazon Cognito authentication for Kibana.
// See: https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html
//
// Experimental.
type CognitoOptions struct {
	// The Amazon Cognito identity pool ID that you want Amazon ES to use for Kibana authentication.
	// Experimental.
	IdentityPoolId *string `json:"identityPoolId"`
	// A role that allows Amazon ES to configure your user pool and identity pool.
	//
	// It must have the `AmazonESCognitoAccess` policy attached to it.
	// See: https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html#es-cognito-auth-prereq
	//
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The Amazon Cognito user pool ID that you want Amazon ES to use for Kibana authentication.
	// Experimental.
	UserPoolId *string `json:"userPoolId"`
}

// Configures a custom domain endpoint for the ES domain.
// Experimental.
type CustomEndpointOptions struct {
	// The custom domain name to assign.
	// Experimental.
	DomainName *string `json:"domainName"`
	// The certificate to use.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `json:"certificate"`
	// The hosted zone in Route53 to create the CNAME record in.
	// Experimental.
	HostedZone awsroute53.IHostedZone `json:"hostedZone"`
}

// Provides an Elasticsearch domain.
// Experimental.
type Domain interface {
	awscdk.Resource
	awsec2.IConnectable
	IDomain
	AppLogGroup() awslogs.ILogGroup
	AuditLogGroup() awslogs.ILogGroup
	Connections() awsec2.Connections
	DomainArn() *string
	DomainEndpoint() *string
	DomainName() *string
	Env() *awscdk.ResourceEnvironment
	MasterUserPassword() awscdk.SecretValue
	Node() constructs.Node
	PhysicalName() *string
	SlowIndexLogGroup() awslogs.ILogGroup
	SlowSearchLogGroup() awslogs.ILogGroup
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantIndexRead(index *string, identity awsiam.IGrantable) awsiam.Grant
	GrantIndexReadWrite(index *string, identity awsiam.IGrantable) awsiam.Grant
	GrantIndexWrite(index *string, identity awsiam.IGrantable) awsiam.Grant
	GrantPathRead(path *string, identity awsiam.IGrantable) awsiam.Grant
	GrantPathReadWrite(path *string, identity awsiam.IGrantable) awsiam.Grant
	GrantPathWrite(path *string, identity awsiam.IGrantable) awsiam.Grant
	GrantRead(identity awsiam.IGrantable) awsiam.Grant
	GrantReadWrite(identity awsiam.IGrantable) awsiam.Grant
	GrantWrite(identity awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricAutomatedSnapshotFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricClusterIndexWritesBlocked(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricClusterStatusRed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricClusterStatusYellow(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIndexingLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricKMSKeyError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricKMSKeyInaccessible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricMasterCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricMasterJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNodes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSearchableDocuments(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSearchLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for Domain
type jsiiProxy_Domain struct {
	internal.Type__awscdkResource
	internal.Type__awsec2IConnectable
	jsiiProxy_IDomain
}

func (j *jsiiProxy_Domain) AppLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"appLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) AuditLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"auditLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) DomainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) DomainEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) MasterUserPassword() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) SlowIndexLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"slowIndexLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) SlowSearchLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"slowSearchLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDomain(scope constructs.Construct, id *string, props *DomainProps) Domain {
	_init_.Initialize()

	j := jsiiProxy_Domain{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticsearch.Domain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDomain_Override(d Domain, scope constructs.Construct, id *string, props *DomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticsearch.Domain",
		[]interface{}{scope, id, props},
		d,
	)
}

// Creates a Domain construct that represents an external domain.
// Experimental.
func Domain_FromDomainAttributes(scope constructs.Construct, id *string, attrs *DomainAttributes) IDomain {
	_init_.Initialize()

	var returns IDomain

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.Domain",
		"fromDomainAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Creates a Domain construct that represents an external domain via domain endpoint.
// Experimental.
func Domain_FromDomainEndpoint(scope constructs.Construct, id *string, domainEndpoint *string) IDomain {
	_init_.Initialize()

	var returns IDomain

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.Domain",
		"fromDomainEndpoint",
		[]interface{}{scope, id, domainEndpoint},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Domain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.Domain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Domain_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.Domain",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_Domain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (d *jsiiProxy_Domain) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (d *jsiiProxy_Domain) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (d *jsiiProxy_Domain) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant read permissions for an index in this domain to an IAM principal (Role/Group/User).
// Experimental.
func (d *jsiiProxy_Domain) GrantIndexRead(index *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantIndexRead",
		[]interface{}{index, identity},
		&returns,
	)

	return returns
}

// Grant read/write permissions for an index in this domain to an IAM principal (Role/Group/User).
// Experimental.
func (d *jsiiProxy_Domain) GrantIndexReadWrite(index *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantIndexReadWrite",
		[]interface{}{index, identity},
		&returns,
	)

	return returns
}

// Grant write permissions for an index in this domain to an IAM principal (Role/Group/User).
// Experimental.
func (d *jsiiProxy_Domain) GrantIndexWrite(index *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantIndexWrite",
		[]interface{}{index, identity},
		&returns,
	)

	return returns
}

// Grant read permissions for a specific path in this domain to an IAM principal (Role/Group/User).
// Experimental.
func (d *jsiiProxy_Domain) GrantPathRead(path *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantPathRead",
		[]interface{}{path, identity},
		&returns,
	)

	return returns
}

// Grant read/write permissions for a specific path in this domain to an IAM principal (Role/Group/User).
// Experimental.
func (d *jsiiProxy_Domain) GrantPathReadWrite(path *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantPathReadWrite",
		[]interface{}{path, identity},
		&returns,
	)

	return returns
}

// Grant write permissions for a specific path in this domain to an IAM principal (Role/Group/User).
// Experimental.
func (d *jsiiProxy_Domain) GrantPathWrite(path *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantPathWrite",
		[]interface{}{path, identity},
		&returns,
	)

	return returns
}

// Grant read permissions for this domain and its contents to an IAM principal (Role/Group/User).
// Experimental.
func (d *jsiiProxy_Domain) GrantRead(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantRead",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

// Grant read/write permissions for this domain and its contents to an IAM principal (Role/Group/User).
// Experimental.
func (d *jsiiProxy_Domain) GrantReadWrite(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantReadWrite",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

// Grant write permissions for this domain and its contents to an IAM principal (Role/Group/User).
// Experimental.
func (d *jsiiProxy_Domain) GrantWrite(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantWrite",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

// Return the given named metric for this Domain.
// Experimental.
func (d *jsiiProxy_Domain) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for automated snapshot failures.
// Experimental.
func (d *jsiiProxy_Domain) MetricAutomatedSnapshotFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricAutomatedSnapshotFailure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the cluster blocking index writes.
// Experimental.
func (d *jsiiProxy_Domain) MetricClusterIndexWritesBlocked(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricClusterIndexWritesBlocked",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the time the cluster status is red.
// Experimental.
func (d *jsiiProxy_Domain) MetricClusterStatusRed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricClusterStatusRed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the time the cluster status is yellow.
// Experimental.
func (d *jsiiProxy_Domain) MetricClusterStatusYellow(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricClusterStatusYellow",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for CPU utilization.
// Experimental.
func (d *jsiiProxy_Domain) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the storage space of nodes in the cluster.
// Experimental.
func (d *jsiiProxy_Domain) MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for indexing latency.
// Experimental.
func (d *jsiiProxy_Domain) MetricIndexingLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricIndexingLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for JVM memory pressure.
// Experimental.
func (d *jsiiProxy_Domain) MetricJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricJVMMemoryPressure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for KMS key errors.
// Experimental.
func (d *jsiiProxy_Domain) MetricKMSKeyError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricKMSKeyError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for KMS key being inaccessible.
// Experimental.
func (d *jsiiProxy_Domain) MetricKMSKeyInaccessible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricKMSKeyInaccessible",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for master CPU utilization.
// Experimental.
func (d *jsiiProxy_Domain) MetricMasterCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricMasterCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for master JVM memory pressure.
// Experimental.
func (d *jsiiProxy_Domain) MetricMasterJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricMasterJVMMemoryPressure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of nodes.
// Experimental.
func (d *jsiiProxy_Domain) MetricNodes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNodes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for number of searchable documents.
// Experimental.
func (d *jsiiProxy_Domain) MetricSearchableDocuments(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSearchableDocuments",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for search latency.
// Experimental.
func (d *jsiiProxy_Domain) MetricSearchLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSearchLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_Domain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Reference to an Elasticsearch domain.
// Experimental.
type DomainAttributes struct {
	// The ARN of the Elasticsearch domain.
	// Experimental.
	DomainArn *string `json:"domainArn"`
	// The domain endpoint of the Elasticsearch domain.
	// Experimental.
	DomainEndpoint *string `json:"domainEndpoint"`
}

// Properties for an AWS Elasticsearch Domain.
// Experimental.
type DomainProps struct {
	// The Elasticsearch version that your domain will leverage.
	// Experimental.
	Version ElasticsearchVersion `json:"version"`
	// Domain Access policies.
	// Experimental.
	AccessPolicies *[]awsiam.PolicyStatement `json:"accessPolicies"`
	// Additional options to specify for the Amazon ES domain.
	// Experimental.
	AdvancedOptions *map[string]*string `json:"advancedOptions"`
	// The hour in UTC during which the service takes an automated daily snapshot of the indices in the Amazon ES domain.
	//
	// Only applies for Elasticsearch
	// versions below 5.3.
	// Experimental.
	AutomatedSnapshotStartHour *float64 `json:"automatedSnapshotStartHour"`
	// The cluster capacity configuration for the Amazon ES domain.
	// Experimental.
	Capacity *CapacityConfig `json:"capacity"`
	// Configures Amazon ES to use Amazon Cognito authentication for Kibana.
	// Experimental.
	CognitoKibanaAuth *CognitoOptions `json:"cognitoKibanaAuth"`
	// To configure a custom domain configure these options.
	//
	// If you specify a Route53 hosted zone it will create a CNAME record and use DNS validation for the certificate
	// Experimental.
	CustomEndpoint *CustomEndpointOptions `json:"customEndpoint"`
	// Enforces a particular physical domain name.
	// Experimental.
	DomainName *string `json:"domainName"`
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the Amazon ES domain.
	//
	// For more information, see
	// [Configuring EBS-based Storage]
	// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
	// in the Amazon Elasticsearch Service Developer Guide.
	// Experimental.
	Ebs *EbsOptions `json:"ebs"`
	// To upgrade an Amazon ES domain to a new version of Elasticsearch rather than replacing the entire domain resource, use the EnableVersionUpgrade update policy.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeelasticsearchdomain
	//
	// Experimental.
	EnableVersionUpgrade *bool `json:"enableVersionUpgrade"`
	// Encryption at rest options for the cluster.
	// Experimental.
	EncryptionAtRest *EncryptionAtRestOptions `json:"encryptionAtRest"`
	// True to require that all traffic to the domain arrive over HTTPS.
	// Experimental.
	EnforceHttps *bool `json:"enforceHttps"`
	// Specifies options for fine-grained access control.
	//
	// Requires Elasticsearch version 6.7 or later. Enabling fine-grained access control
	// also requires encryption of data at rest and node-to-node encryption, along with
	// enforced HTTPS.
	// Experimental.
	FineGrainedAccessControl *AdvancedSecurityOptions `json:"fineGrainedAccessControl"`
	// Configuration log publishing configuration options.
	// Experimental.
	Logging *LoggingOptions `json:"logging"`
	// Specify true to enable node to node encryption.
	//
	// Requires Elasticsearch version 6.0 or later.
	// Experimental.
	NodeToNodeEncryption *bool `json:"nodeToNodeEncryption"`
	// Policy to apply when the domain is removed from the stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// The list of security groups that are associated with the VPC endpoints for the domain.
	//
	// Only used if `vpc` is specified.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The minimum TLS version required for traffic to the domain.
	// Experimental.
	TlsSecurityPolicy TLSSecurityPolicy `json:"tlsSecurityPolicy"`
	// Configures the domain so that unsigned basic auth is enabled.
	//
	// If no master user is provided a default master user
	// with username `admin` and a dynamically generated password stored in KMS is created. The password can be retrieved
	// by getting `masterUserPassword` from the domain instance.
	//
	// Setting this to true will also add an access policy that allows unsigned
	// access, enable node to node encryption, encryption at rest. If conflicting
	// settings are encountered (like disabling encryption at rest) enabling this
	// setting will cause a failure.
	// Experimental.
	UseUnsignedBasicAuth *bool `json:"useUnsignedBasicAuth"`
	// Place the domain inside this VPC.
	// See: https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html
	//
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The specific vpc subnets the domain will be placed in.
	//
	// You must provide one subnet for each Availability Zone
	// that your domain uses. For example, you must specify three subnet IDs for a three Availability Zone
	// domain.
	//
	// Only used if `vpc` is specified.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html
	//
	// Experimental.
	VpcSubnets *[]*awsec2.SubnetSelection `json:"vpcSubnets"`
	// The cluster zone awareness configuration for the Amazon ES domain.
	// Experimental.
	ZoneAwareness *ZoneAwarenessConfig `json:"zoneAwareness"`
}

// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the Amazon ES domain.
//
// For more information, see
// [Configuring EBS-based Storage]
// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
// in the Amazon Elasticsearch Service Developer Guide.
// Experimental.
type EbsOptions struct {
	// Specifies whether Amazon EBS volumes are attached to data nodes in the Amazon ES domain.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	//
	// This property applies only to the Provisioned IOPS (SSD) EBS
	// volume type.
	// Experimental.
	Iops *float64 `json:"iops"`
	// The size (in GiB) of the EBS volume for each data node.
	//
	// The minimum and
	// maximum size of an EBS volume depends on the EBS volume type and the
	// instance type to which it is attached.  For more information, see
	// [Configuring EBS-based Storage]
	// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
	// in the Amazon Elasticsearch Service Developer Guide.
	// Experimental.
	VolumeSize *float64 `json:"volumeSize"`
	// The EBS volume type to use with the Amazon ES domain, such as standard, gp2, io1.
	//
	// For more information, see[Configuring EBS-based Storage]
	// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
	// in the Amazon Elasticsearch Service Developer Guide.
	// Experimental.
	VolumeType awsec2.EbsDeviceVolumeType `json:"volumeType"`
}

// Elasticsearch version.
// Experimental.
type ElasticsearchVersion interface {
	Version() *string
}

// The jsii proxy struct for ElasticsearchVersion
type jsiiProxy_ElasticsearchVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_ElasticsearchVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Custom Elasticsearch version.
// Experimental.
func ElasticsearchVersion_Of(version *string) ElasticsearchVersion {
	_init_.Initialize()

	var returns ElasticsearchVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func ElasticsearchVersion_V1_5() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V1_5",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V2_3() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V2_3",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V5_1() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V5_1",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V5_3() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V5_3",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V5_5() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V5_5",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V5_6() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V5_6",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_0() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_0",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_2() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_2",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_3() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_3",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_4() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_4",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_5() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_5",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_7() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_7",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V6_8() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V6_8",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_1() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_1",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_10() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_10",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_4() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_4",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_7() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_7",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_8() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_8",
		&returns,
	)
	return returns
}

func ElasticsearchVersion_V7_9() ElasticsearchVersion {
	_init_.Initialize()
	var returns ElasticsearchVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticsearch.ElasticsearchVersion",
		"V7_9",
		&returns,
	)
	return returns
}

// Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service (KMS) key to use.
//
// Can only be used to create a new domain,
// not update an existing one. Requires Elasticsearch version 5.1 or later.
// Experimental.
type EncryptionAtRestOptions struct {
	// Specify true to enable encryption at rest.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// Supply if using KMS key for encryption at rest.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey"`
}

// An interface that represents an Elasticsearch domain - either created with the CDK, or an existing one.
// Experimental.
type IDomain interface {
	awscdk.IResource
	// Grant read permissions for an index in this domain to an IAM principal (Role/Group/User).
	// Experimental.
	GrantIndexRead(index *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read/write permissions for an index in this domain to an IAM principal (Role/Group/User).
	// Experimental.
	GrantIndexReadWrite(index *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for an index in this domain to an IAM principal (Role/Group/User).
	// Experimental.
	GrantIndexWrite(index *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read permissions for a specific path in this domain to an IAM principal (Role/Group/User).
	// Experimental.
	GrantPathRead(path *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read/write permissions for a specific path in this domain to an IAM principal (Role/Group/User).
	// Experimental.
	GrantPathReadWrite(path *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for a specific path in this domain to an IAM principal (Role/Group/User).
	// Experimental.
	GrantPathWrite(path *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read permissions for this domain and its contents to an IAM principal (Role/Group/User).
	// Experimental.
	GrantRead(identity awsiam.IGrantable) awsiam.Grant
	// Grant read/write permissions for this domain and its contents to an IAM principal (Role/Group/User).
	// Experimental.
	GrantReadWrite(identity awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for this domain and its contents to an IAM principal (Role/Group/User).
	// Experimental.
	GrantWrite(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Domain.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for automated snapshot failures.
	// Experimental.
	MetricAutomatedSnapshotFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the cluster blocking index writes.
	// Experimental.
	MetricClusterIndexWritesBlocked(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time the cluster status is red.
	// Experimental.
	MetricClusterStatusRed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time the cluster status is yellow.
	// Experimental.
	MetricClusterStatusYellow(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for CPU utilization.
	// Experimental.
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the storage space of nodes in the cluster.
	// Experimental.
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for indexing latency.
	// Experimental.
	MetricIndexingLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for JVM memory pressure.
	// Experimental.
	MetricJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for KMS key errors.
	// Experimental.
	MetricKMSKeyError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for KMS key being inaccessible.
	// Experimental.
	MetricKMSKeyInaccessible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for master CPU utilization.
	// Experimental.
	MetricMasterCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for master JVM memory pressure.
	// Experimental.
	MetricMasterJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of nodes.
	// Experimental.
	MetricNodes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for number of searchable documents.
	// Experimental.
	MetricSearchableDocuments(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for search latency.
	// Experimental.
	MetricSearchLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Arn of the Elasticsearch domain.
	// Experimental.
	DomainArn() *string
	// Endpoint of the Elasticsearch domain.
	// Experimental.
	DomainEndpoint() *string
	// Domain name of the Elasticsearch domain.
	// Experimental.
	DomainName() *string
}

// The jsii proxy for IDomain
type jsiiProxy_IDomain struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDomain) GrantIndexRead(index *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantIndexRead",
		[]interface{}{index, identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) GrantIndexReadWrite(index *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantIndexReadWrite",
		[]interface{}{index, identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) GrantIndexWrite(index *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantIndexWrite",
		[]interface{}{index, identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) GrantPathRead(path *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPathRead",
		[]interface{}{path, identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) GrantPathReadWrite(path *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPathReadWrite",
		[]interface{}{path, identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) GrantPathWrite(path *string, identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPathWrite",
		[]interface{}{path, identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) GrantRead(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) GrantReadWrite(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadWrite",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) GrantWrite(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricAutomatedSnapshotFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricAutomatedSnapshotFailure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricClusterIndexWritesBlocked(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricClusterIndexWritesBlocked",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricClusterStatusRed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricClusterStatusRed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricClusterStatusYellow(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricClusterStatusYellow",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFreeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricIndexingLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIndexingLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricJVMMemoryPressure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricKMSKeyError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricKMSKeyError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricKMSKeyInaccessible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricKMSKeyInaccessible",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricMasterCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMasterCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricMasterJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMasterJVMMemoryPressure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricNodes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNodes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricSearchableDocuments(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSearchableDocuments",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDomain) MetricSearchLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSearchLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDomain) DomainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomain) DomainEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

// Configures log settings for the domain.
// Experimental.
type LoggingOptions struct {
	// Specify if Elasticsearch application logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later.
	// Experimental.
	AppLogEnabled *bool `json:"appLogEnabled"`
	// Log Elasticsearch application logs to this log group.
	// Experimental.
	AppLogGroup awslogs.ILogGroup `json:"appLogGroup"`
	// Specify if Elasticsearch audit logging should be set up.
	//
	// Requires Elasticsearch version 6.7 or later and fine grained access control to be enabled.
	// Experimental.
	AuditLogEnabled *bool `json:"auditLogEnabled"`
	// Log Elasticsearch audit logs to this log group.
	// Experimental.
	AuditLogGroup awslogs.ILogGroup `json:"auditLogGroup"`
	// Specify if slow index logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later.
	// Experimental.
	SlowIndexLogEnabled *bool `json:"slowIndexLogEnabled"`
	// Log slow indices to this log group.
	// Experimental.
	SlowIndexLogGroup awslogs.ILogGroup `json:"slowIndexLogGroup"`
	// Specify if slow search logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later.
	// Experimental.
	SlowSearchLogEnabled *bool `json:"slowSearchLogEnabled"`
	// Log slow searches to this log group.
	// Experimental.
	SlowSearchLogGroup awslogs.ILogGroup `json:"slowSearchLogGroup"`
}

// The minimum TLS version required for traffic to the domain.
// Experimental.
type TLSSecurityPolicy string

const (
	TLSSecurityPolicy_TLS_1_0 TLSSecurityPolicy = "TLS_1_0"
	TLSSecurityPolicy_TLS_1_2 TLSSecurityPolicy = "TLS_1_2"
)

// Specifies zone awareness configuration options.
// Experimental.
type ZoneAwarenessConfig struct {
	// If you enabled multiple Availability Zones (AZs), the number of AZs that you want the domain to use.
	//
	// Valid values are 2 and 3.
	// Experimental.
	AvailabilityZoneCount *float64 `json:"availabilityZoneCount"`
	// Indicates whether to enable zone awareness for the Amazon ES domain.
	//
	// When you enable zone awareness, Amazon ES allocates the nodes and replica
	// index shards that belong to a cluster across two Availability Zones (AZs)
	// in the same region to prevent data loss and minimize downtime in the event
	// of node or data center failure. Don't enable zone awareness if your cluster
	// has no replica index shards or is a single-node cluster. For more information,
	// see [Configuring a Multi-AZ Domain]
	// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-managedomains-multiaz)
	// in the Amazon Elasticsearch Service Developer Guide.
	// Experimental.
	Enabled *bool `json:"enabled"`
}

