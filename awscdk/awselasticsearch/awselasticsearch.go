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
//
// TODO: EXAMPLE
//
type AdvancedSecurityOptions struct {
	// ARN for the master user.
	//
	// Only specify this or masterUserName, but not both.
	MasterUserArn *string `json:"masterUserArn"`
	// Username for the master user.
	//
	// Only specify this or masterUserArn, but not both.
	MasterUserName *string `json:"masterUserName"`
	// Password for the master user.
	//
	// You can use `SecretValue.plainText` to specify a password in plain text or
	// use `secretsmanager.Secret.fromSecretAttributes` to reference a secret in
	// Secrets Manager.
	MasterUserPassword awscdk.SecretValue `json:"masterUserPassword"`
}

// Configures the capacity of the cluster such as the instance type and the number of instances.
//
// TODO: EXAMPLE
//
type CapacityConfig struct {
	// The instance type for your data nodes, such as `m3.medium.elasticsearch`. For valid values, see [Supported Instance Types](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html) in the Amazon Elasticsearch Service Developer Guide.
	DataNodeInstanceType *string `json:"dataNodeInstanceType"`
	// The number of data nodes (instances) to use in the Amazon ES domain.
	DataNodes *float64 `json:"dataNodes"`
	// The hardware configuration of the computer that hosts the dedicated master node, such as `m3.medium.elasticsearch`. For valid values, see [Supported Instance Types] (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html) in the Amazon Elasticsearch Service Developer Guide.
	MasterNodeInstanceType *string `json:"masterNodeInstanceType"`
	// The number of instances to use for the master node.
	MasterNodes *float64 `json:"masterNodes"`
	// The instance type for your UltraWarm node, such as `ultrawarm1.medium.elasticsearch`. For valid values, see [UltraWarm Storage Limits] (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-limits.html#limits-ultrawarm) in the Amazon Elasticsearch Service Developer Guide.
	WarmInstanceType *string `json:"warmInstanceType"`
	// The number of UltraWarm nodes (instances) to use in the Amazon ES domain.
	WarmNodes *float64 `json:"warmNodes"`
}

// A CloudFormation `AWS::Elasticsearch::Domain`.
//
// The AWS::Elasticsearch::Domain resource creates an Amazon OpenSearch Service (successor to Amazon Elasticsearch Service) domain.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
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
func (c *jsiiProxy_CfnDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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
func (c *jsiiProxy_CfnDomain) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies options for fine-grained access control.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDomain_AdvancedSecurityOptionsInputProperty struct {
	// True to enable fine-grained access control.
	//
	// You must also enable encryption of data at rest and node-to-node encryption.
	Enabled interface{} `json:"enabled"`
	// True to enable the internal user database.
	InternalUserDatabaseEnabled interface{} `json:"internalUserDatabaseEnabled"`
	// Specifies information about the master user.
	MasterUserOptions interface{} `json:"masterUserOptions"`
}

// Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDomain_CognitoOptionsProperty struct {
	// Whether to enable or disable Amazon Cognito authentication for OpenSearch Dashboards.
	//
	// See [Amazon Cognito authentication for OpenSearch Dashboards](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html) .
	Enabled interface{} `json:"enabled"`
	// The Amazon Cognito identity pool ID that you want OpenSearch Service to use for OpenSearch Dashboards authentication.
	IdentityPoolId *string `json:"identityPoolId"`
	// The `AmazonESCognitoAccess` role that allows OpenSearch Service to configure your user pool and identity pool.
	RoleArn *string `json:"roleArn"`
	// The Amazon Cognito user pool ID that you want OpenSearch Service to use for OpenSearch Dashboards authentication.
	UserPoolId *string `json:"userPoolId"`
}

// Specifies options for cold storage. For more information, see [Cold storage for Amazon Elasticsearch Service](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/cold-storage.html) .
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDomain_ColdStorageOptionsProperty struct {
	// Whether to enable or disable cold storage on the domain.
	//
	// You must enable UltraWarm storage in order to enable cold storage.
	Enabled interface{} `json:"enabled"`
}

// Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDomain_DomainEndpointOptionsProperty struct {
	// The fully qualified URL for your custom endpoint.
	CustomEndpoint *string `json:"customEndpoint"`
	// The AWS Certificate Manager ARN for your domain's SSL/TLS certificate.
	CustomEndpointCertificateArn *string `json:"customEndpointCertificateArn"`
	// True to enable a custom endpoint for the domain.
	//
	// If enabled, you must also provide values for `CustomEndpoint` and `CustomEndpointCertificateArn` .
	CustomEndpointEnabled interface{} `json:"customEndpointEnabled"`
	// True to require that all traffic to the domain arrive over HTTPS.
	EnforceHttps interface{} `json:"enforceHttps"`
	// The minimum TLS version required for traffic to the domain. Valid values are TLS 1.0 (default) or 1.2:.
	//
	// - `Policy-Min-TLS-1-0-2019-07`
	// - `Policy-Min-TLS-1-2-2019-07`
	TlsSecurityPolicy *string `json:"tlsSecurityPolicy"`
}

// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the OpenSearch Service domain.
//
// For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDomain_EBSOptionsProperty struct {
	// Specifies whether Amazon EBS volumes are attached to data nodes in the OpenSearch Service domain.
	EbsEnabled interface{} `json:"ebsEnabled"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	//
	// This property applies only to the Provisioned IOPS (SSD) EBS volume type.
	Iops *float64 `json:"iops"`
	// The size (in GiB) of the EBS volume for each data node.
	//
	// The minimum and maximum size of an EBS volume depends on the EBS volume type and the instance type to which it is attached. For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
	VolumeSize *float64 `json:"volumeSize"`
	// The EBS volume type to use with the OpenSearch Service domain, such as standard, gp2, or io1.
	//
	// For more information about each type, see [Amazon EBS volume types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the *Amazon EC2 User Guide for Linux Instances* .
	VolumeType *string `json:"volumeType"`
}

// The cluster configuration for the OpenSearch Service domain.
//
// You can specify options such as the instance type and the number of instances. For more information, see [Creating and managing Amazon OpenSearch Service domains](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html) in the *Amazon OpenSearch Service Developer Guide* .
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDomain_ElasticsearchClusterConfigProperty struct {
	// Specifies cold storage options for the domain.
	ColdStorageOptions interface{} `json:"coldStorageOptions"`
	// The number of instances to use for the master node.
	//
	// If you specify this property, you must specify true for the DedicatedMasterEnabled property.
	DedicatedMasterCount *float64 `json:"dedicatedMasterCount"`
	// Indicates whether to use a dedicated master node for the OpenSearch Service domain.
	//
	// A dedicated master node is a cluster node that performs cluster management tasks, but doesn't hold data or respond to data upload requests. Dedicated master nodes offload cluster management tasks to increase the stability of your search clusters. See [Dedicated master nodes in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-dedicatedmasternodes.html) .
	DedicatedMasterEnabled interface{} `json:"dedicatedMasterEnabled"`
	// The hardware configuration of the computer that hosts the dedicated master node, such as `m3.medium.elasticsearch` . If you specify this property, you must specify true for the `DedicatedMasterEnabled` property. For valid values, see [Supported instance types in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) .
	DedicatedMasterType *string `json:"dedicatedMasterType"`
	// The number of data nodes (instances) to use in the OpenSearch Service domain.
	InstanceCount *float64 `json:"instanceCount"`
	// The instance type for your data nodes, such as `m3.medium.elasticsearch` . For valid values, see [Supported instance types in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) .
	InstanceType *string `json:"instanceType"`
	// The number of warm nodes in the cluster.
	WarmCount *float64 `json:"warmCount"`
	// Whether to enable warm storage for the cluster.
	WarmEnabled interface{} `json:"warmEnabled"`
	// The instance type for the cluster's warm nodes.
	WarmType *string `json:"warmType"`
	// Specifies zone awareness configuration options.
	//
	// Only use if `ZoneAwarenessEnabled` is `true` .
	ZoneAwarenessConfig interface{} `json:"zoneAwarenessConfig"`
	// Indicates whether to enable zone awareness for the OpenSearch Service domain.
	//
	// When you enable zone awareness, OpenSearch Service allocates the nodes and replica index shards that belong to a cluster across two Availability Zones (AZs) in the same region to prevent data loss and minimize downtime in the event of node or data center failure. Don't enable zone awareness if your cluster has no replica index shards or is a single-node cluster. For more information, see [Configuring a multi-AZ domain in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html) .
	ZoneAwarenessEnabled interface{} `json:"zoneAwarenessEnabled"`
}

// Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service key to use.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDomain_EncryptionAtRestOptionsProperty struct {
	// Specify `true` to enable encryption at rest.
	Enabled interface{} `json:"enabled"`
	// The KMS key ID.
	//
	// Takes the form `1a2a3a4-1a2a-3a4a-5a6a-1a2a3a4a5a6a` .
	KmsKeyId *string `json:"kmsKeyId"`
}

// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// Specifies whether the OpenSearch Service domain publishes the Elasticsearch application, search slow logs, or index slow logs to Amazon CloudWatch. Each option must be an object of name `SEARCH_SLOW_LOGS` , `ES_APPLICATION_LOGS` , `INDEX_SLOW_LOGS` , or `AUDIT_LOGS` depending on the type of logs you want to publish.
//
// If you enable a slow log, you still have to enable the *collection* of slow logs using the Configuration API. To learn more, see [Enabling log publishing ( AWS CLI)](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createdomain-configure-slow-logs.html#createdomain-configure-slow-logs-cli) .
//
// TODO: EXAMPLE
//
type CfnDomain_LogPublishingOptionProperty struct {
	// Specifies the CloudWatch log group to publish to.
	CloudWatchLogsLogGroupArn *string `json:"cloudWatchLogsLogGroupArn"`
	// If `true` , enables the publishing of logs to CloudWatch.
	//
	// Default: `false` .
	Enabled interface{} `json:"enabled"`
}

// Specifies information about the master user.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDomain_MasterUserOptionsProperty struct {
	// ARN for the master user.
	//
	// Only specify if `InternalUserDatabaseEnabled` is false in `AdvancedSecurityOptions` .
	MasterUserArn *string `json:"masterUserArn"`
	// Username for the master user.
	//
	// Only specify if `InternalUserDatabaseEnabled` is true in `AdvancedSecurityOptions` .
	MasterUserName *string `json:"masterUserName"`
	// Password for the master user.
	//
	// Only specify if `InternalUserDatabaseEnabled` is true in `AdvancedSecurityOptions` .
	MasterUserPassword *string `json:"masterUserPassword"`
}

// Specifies whether node-to-node encryption is enabled.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDomain_NodeToNodeEncryptionOptionsProperty struct {
	// Specifies whether node-to-node encryption is enabled, as a Boolean.
	Enabled interface{} `json:"enabled"`
}

// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// *DEPRECATED* . For domains running Elasticsearch 5.3 and later, OpenSearch Service takes hourly automated snapshots, making this setting irrelevant. For domains running earlier versions of Elasticsearch, OpenSearch Service takes daily automated snapshots.
//
// The automated snapshot configuration for the OpenSearch Service domain indices.
//
// TODO: EXAMPLE
//
type CfnDomain_SnapshotOptionsProperty struct {
	// The hour in UTC during which the service takes an automated daily snapshot of the indices in the OpenSearch Service domain.
	//
	// For example, if you specify 0, OpenSearch Service takes an automated snapshot everyday between midnight and 1 am. You can specify a value between 0 and 23.
	AutomatedSnapshotStartHour *float64 `json:"automatedSnapshotStartHour"`
}

// The virtual private cloud (VPC) configuration for the OpenSearch Service domain.
//
// For more information, see [Launching your Amazon OpenSearch Service domains using a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the *Amazon OpenSearch Service Developer Guide* .
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDomain_VPCOptionsProperty struct {
	// The list of security group IDs that are associated with the VPC endpoints for the domain.
	//
	// If you don't provide a security group ID, OpenSearch Service uses the default security group for the VPC. To learn more, see [Security groups for your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html) in the *Amazon VPC User Guide* .
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Provide one subnet ID for each Availability Zone that your domain uses.
	//
	// For example, you must specify three subnet IDs for a three Availability Zone domain. To learn more, see [VPCs and subnets](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in the *Amazon VPC User Guide* .
	SubnetIds *[]*string `json:"subnetIds"`
}

// Specifies zone awareness configuration options. Only use if `ZoneAwarenessEnabled` is `true` .
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDomain_ZoneAwarenessConfigProperty struct {
	// If you enabled multiple Availability Zones (AZs), the number of AZs that you want the domain to use.
	//
	// Valid values are `2` and `3` . Default is 2.
	AvailabilityZoneCount *float64 `json:"availabilityZoneCount"`
}

// Properties for defining a `CfnDomain`.
//
// TODO: EXAMPLE
//
type CfnDomainProps struct {
	// An AWS Identity and Access Management ( IAM ) policy document that specifies who can access the OpenSearch Service domain and their permissions.
	//
	// For more information, see [Configuring access policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ac.html#ac-creating) in the *Amazon OpenSearch Service Developer Guid* e.
	AccessPolicies interface{} `json:"accessPolicies"`
	// Additional options to specify for the OpenSearch Service domain.
	//
	// For more information, see [Advanced cluster parameters](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options) in the *Amazon OpenSearch Service Developer Guide* .
	AdvancedOptions interface{} `json:"advancedOptions"`
	// Specifies options for fine-grained access control.
	AdvancedSecurityOptions interface{} `json:"advancedSecurityOptions"`
	// Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
	CognitoOptions interface{} `json:"cognitoOptions"`
	// Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.
	DomainEndpointOptions interface{} `json:"domainEndpointOptions"`
	// A name for the OpenSearch Service domain.
	//
	// For valid values, see the [DomainName](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/configuration-api.html#configuration-api-datatypes-domainname) data type in the *Amazon OpenSearch Service Developer Guide* . If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the domain name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DomainName *string `json:"domainName"`
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the OpenSearch Service domain.
	//
	// For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
	EbsOptions interface{} `json:"ebsOptions"`
	// ElasticsearchClusterConfig is a property of the AWS::Elasticsearch::Domain resource that configures the cluster of an Amazon OpenSearch Service domain.
	ElasticsearchClusterConfig interface{} `json:"elasticsearchClusterConfig"`
	// The version of Elasticsearch to use, such as 2.3. If not specified, 1.5 is used as the default. For information about the versions that OpenSearch Service supports, see [Supported versions of OpenSearch and Elasticsearch](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/what-is.html#choosing-version) in the *Amazon OpenSearch Service Developer Guide* .
	//
	// If you set the [EnableVersionUpgrade](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeopensearchdomain) update policy to `true` , you can update `ElasticsearchVersion` without interruption. When `EnableVersionUpgrade` is set to `false` , or is not specified, updating `ElasticsearchVersion` results in [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	ElasticsearchVersion *string `json:"elasticsearchVersion"`
	// Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service key to use.
	//
	// See [Encryption of data at rest for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html) .
	EncryptionAtRestOptions interface{} `json:"encryptionAtRestOptions"`
	// An object with one or more of the following keys: `SEARCH_SLOW_LOGS` , `ES_APPLICATION_LOGS` , `INDEX_SLOW_LOGS` , `AUDIT_LOGS` , depending on the types of logs you want to publish.
	//
	// Each key needs a valid `LogPublishingOption` value.
	LogPublishingOptions interface{} `json:"logPublishingOptions"`
	// Specifies whether node-to-node encryption is enabled.
	//
	// See [Node-to-node encryption for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html) .
	NodeToNodeEncryptionOptions interface{} `json:"nodeToNodeEncryptionOptions"`
	// *DEPRECATED* .
	//
	// The automated snapshot configuration for the OpenSearch Service domain indices.
	SnapshotOptions interface{} `json:"snapshotOptions"`
	// An arbitrary set of tags (key–value pairs) to associate with the OpenSearch Service domain.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// The virtual private cloud (VPC) configuration for the OpenSearch Service domain.
	//
	// For more information, see [Launching your Amazon OpenSearch Service domains within a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the *Amazon OpenSearch Service Developer Guide* .
	VpcOptions interface{} `json:"vpcOptions"`
}

// Configures Amazon ES to use Amazon Cognito authentication for Kibana.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html
//
type CognitoOptions struct {
	// The Amazon Cognito identity pool ID that you want Amazon ES to use for Kibana authentication.
	IdentityPoolId *string `json:"identityPoolId"`
	// A role that allows Amazon ES to configure your user pool and identity pool.
	//
	// It must have the `AmazonESCognitoAccess` policy attached to it.
	// See: https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html#es-cognito-auth-prereq
	//
	Role awsiam.IRole `json:"role"`
	// The Amazon Cognito user pool ID that you want Amazon ES to use for Kibana authentication.
	UserPoolId *string `json:"userPoolId"`
}

// Configures a custom domain endpoint for the ES domain.
//
// TODO: EXAMPLE
//
type CustomEndpointOptions struct {
	// The custom domain name to assign.
	DomainName *string `json:"domainName"`
	// The certificate to use.
	Certificate awscertificatemanager.ICertificate `json:"certificate"`
	// The hosted zone in Route53 to create the CNAME record in.
	HostedZone awsroute53.IHostedZone `json:"hostedZone"`
}

// Provides an Elasticsearch domain.
//
// TODO: EXAMPLE
//
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

func NewDomain_Override(d Domain, scope constructs.Construct, id *string, props *DomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticsearch.Domain",
		[]interface{}{scope, id, props},
		d,
	)
}

// Creates a Domain construct that represents an external domain.
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
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_Domain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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
//
// TODO: EXAMPLE
//
type DomainAttributes struct {
	// The ARN of the Elasticsearch domain.
	DomainArn *string `json:"domainArn"`
	// The domain endpoint of the Elasticsearch domain.
	DomainEndpoint *string `json:"domainEndpoint"`
}

// Properties for an AWS Elasticsearch Domain.
//
// TODO: EXAMPLE
//
type DomainProps struct {
	// The Elasticsearch version that your domain will leverage.
	Version ElasticsearchVersion `json:"version"`
	// Domain Access policies.
	AccessPolicies *[]awsiam.PolicyStatement `json:"accessPolicies"`
	// Additional options to specify for the Amazon ES domain.
	// See: https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options
	//
	AdvancedOptions *map[string]*string `json:"advancedOptions"`
	// The hour in UTC during which the service takes an automated daily snapshot of the indices in the Amazon ES domain.
	//
	// Only applies for Elasticsearch
	// versions below 5.3.
	AutomatedSnapshotStartHour *float64 `json:"automatedSnapshotStartHour"`
	// The cluster capacity configuration for the Amazon ES domain.
	Capacity *CapacityConfig `json:"capacity"`
	// Configures Amazon ES to use Amazon Cognito authentication for Kibana.
	CognitoKibanaAuth *CognitoOptions `json:"cognitoKibanaAuth"`
	// To configure a custom domain configure these options.
	//
	// If you specify a Route53 hosted zone it will create a CNAME record and use DNS validation for the certificate
	CustomEndpoint *CustomEndpointOptions `json:"customEndpoint"`
	// Enforces a particular physical domain name.
	DomainName *string `json:"domainName"`
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the Amazon ES domain.
	//
	// For more information, see
	// [Configuring EBS-based Storage]
	// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
	// in the Amazon Elasticsearch Service Developer Guide.
	Ebs *EbsOptions `json:"ebs"`
	// To upgrade an Amazon ES domain to a new version of Elasticsearch rather than replacing the entire domain resource, use the EnableVersionUpgrade update policy.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeelasticsearchdomain
	//
	EnableVersionUpgrade *bool `json:"enableVersionUpgrade"`
	// Encryption at rest options for the cluster.
	EncryptionAtRest *EncryptionAtRestOptions `json:"encryptionAtRest"`
	// True to require that all traffic to the domain arrive over HTTPS.
	EnforceHttps *bool `json:"enforceHttps"`
	// Specifies options for fine-grained access control.
	//
	// Requires Elasticsearch version 6.7 or later. Enabling fine-grained access control
	// also requires encryption of data at rest and node-to-node encryption, along with
	// enforced HTTPS.
	FineGrainedAccessControl *AdvancedSecurityOptions `json:"fineGrainedAccessControl"`
	// Configuration log publishing configuration options.
	Logging *LoggingOptions `json:"logging"`
	// Specify true to enable node to node encryption.
	//
	// Requires Elasticsearch version 6.0 or later.
	NodeToNodeEncryption *bool `json:"nodeToNodeEncryption"`
	// Policy to apply when the domain is removed from the stack.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// The list of security groups that are associated with the VPC endpoints for the domain.
	//
	// Only used if `vpc` is specified.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html
	//
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The minimum TLS version required for traffic to the domain.
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
	UseUnsignedBasicAuth *bool `json:"useUnsignedBasicAuth"`
	// Place the domain inside this VPC.
	// See: https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html
	//
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
	VpcSubnets *[]*awsec2.SubnetSelection `json:"vpcSubnets"`
	// The cluster zone awareness configuration for the Amazon ES domain.
	ZoneAwareness *ZoneAwarenessConfig `json:"zoneAwareness"`
}

// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the Amazon ES domain.
//
// For more information, see
// [Configuring EBS-based Storage]
// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
// in the Amazon Elasticsearch Service Developer Guide.
//
// TODO: EXAMPLE
//
type EbsOptions struct {
	// Specifies whether Amazon EBS volumes are attached to data nodes in the Amazon ES domain.
	Enabled *bool `json:"enabled"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	//
	// This property applies only to the Provisioned IOPS (SSD) EBS
	// volume type.
	Iops *float64 `json:"iops"`
	// The size (in GiB) of the EBS volume for each data node.
	//
	// The minimum and
	// maximum size of an EBS volume depends on the EBS volume type and the
	// instance type to which it is attached.  For more information, see
	// [Configuring EBS-based Storage]
	// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
	// in the Amazon Elasticsearch Service Developer Guide.
	VolumeSize *float64 `json:"volumeSize"`
	// The EBS volume type to use with the Amazon ES domain, such as standard, gp2, io1.
	//
	// For more information, see[Configuring EBS-based Storage]
	// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
	// in the Amazon Elasticsearch Service Developer Guide.
	VolumeType awsec2.EbsDeviceVolumeType `json:"volumeType"`
}

// Elasticsearch version.
//
// TODO: EXAMPLE
//
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
//
// TODO: EXAMPLE
//
type EncryptionAtRestOptions struct {
	// Specify true to enable encryption at rest.
	Enabled *bool `json:"enabled"`
	// Supply if using KMS key for encryption at rest.
	KmsKey awskms.IKey `json:"kmsKey"`
}

// An interface that represents an Elasticsearch domain - either created with the CDK, or an existing one.
type IDomain interface {
	awscdk.IResource
	// Grant read permissions for an index in this domain to an IAM principal (Role/Group/User).
	GrantIndexRead(index *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read/write permissions for an index in this domain to an IAM principal (Role/Group/User).
	GrantIndexReadWrite(index *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for an index in this domain to an IAM principal (Role/Group/User).
	GrantIndexWrite(index *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read permissions for a specific path in this domain to an IAM principal (Role/Group/User).
	GrantPathRead(path *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read/write permissions for a specific path in this domain to an IAM principal (Role/Group/User).
	GrantPathReadWrite(path *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for a specific path in this domain to an IAM principal (Role/Group/User).
	GrantPathWrite(path *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read permissions for this domain and its contents to an IAM principal (Role/Group/User).
	GrantRead(identity awsiam.IGrantable) awsiam.Grant
	// Grant read/write permissions for this domain and its contents to an IAM principal (Role/Group/User).
	GrantReadWrite(identity awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for this domain and its contents to an IAM principal (Role/Group/User).
	GrantWrite(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Domain.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for automated snapshot failures.
	MetricAutomatedSnapshotFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the cluster blocking index writes.
	MetricClusterIndexWritesBlocked(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time the cluster status is red.
	MetricClusterStatusRed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time the cluster status is yellow.
	MetricClusterStatusYellow(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for CPU utilization.
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the storage space of nodes in the cluster.
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for indexing latency.
	MetricIndexingLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for JVM memory pressure.
	MetricJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for KMS key errors.
	MetricKMSKeyError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for KMS key being inaccessible.
	MetricKMSKeyInaccessible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for master CPU utilization.
	MetricMasterCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for master JVM memory pressure.
	MetricMasterJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of nodes.
	MetricNodes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for number of searchable documents.
	MetricSearchableDocuments(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for search latency.
	MetricSearchLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Arn of the Elasticsearch domain.
	DomainArn() *string
	// Endpoint of the Elasticsearch domain.
	DomainEndpoint() *string
	// Domain name of the Elasticsearch domain.
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
//
// TODO: EXAMPLE
//
type LoggingOptions struct {
	// Specify if Elasticsearch application logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later.
	AppLogEnabled *bool `json:"appLogEnabled"`
	// Log Elasticsearch application logs to this log group.
	AppLogGroup awslogs.ILogGroup `json:"appLogGroup"`
	// Specify if Elasticsearch audit logging should be set up.
	//
	// Requires Elasticsearch version 6.7 or later and fine grained access control to be enabled.
	AuditLogEnabled *bool `json:"auditLogEnabled"`
	// Log Elasticsearch audit logs to this log group.
	AuditLogGroup awslogs.ILogGroup `json:"auditLogGroup"`
	// Specify if slow index logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later.
	SlowIndexLogEnabled *bool `json:"slowIndexLogEnabled"`
	// Log slow indices to this log group.
	SlowIndexLogGroup awslogs.ILogGroup `json:"slowIndexLogGroup"`
	// Specify if slow search logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later.
	SlowSearchLogEnabled *bool `json:"slowSearchLogEnabled"`
	// Log slow searches to this log group.
	SlowSearchLogGroup awslogs.ILogGroup `json:"slowSearchLogGroup"`
}

// The minimum TLS version required for traffic to the domain.
type TLSSecurityPolicy string

const (
	TLSSecurityPolicy_TLS_1_0 TLSSecurityPolicy = "TLS_1_0"
	TLSSecurityPolicy_TLS_1_2 TLSSecurityPolicy = "TLS_1_2"
)

// Specifies zone awareness configuration options.
//
// TODO: EXAMPLE
//
type ZoneAwarenessConfig struct {
	// If you enabled multiple Availability Zones (AZs), the number of AZs that you want the domain to use.
	//
	// Valid values are 2 and 3.
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
	Enabled *bool `json:"enabled"`
}

