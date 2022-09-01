package awsopensearchservice

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::OpenSearchService::Domain`.
//
// The AWS::OpenSearchService::Domain resource creates an Amazon OpenSearch Service domain.
//
// > The `AWS::OpenSearchService::Domain` resource replaces the legacy [AWS::Elasticsearch::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html) resource. While the Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and legacy Elasticsearch engines. For instructions to upgrade domains defined within CloudFormation from Elasticsearch to OpenSearch, see [Remarks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#aws-resource-opensearchservice-domain--remarks) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessPolicies interface{}
//
//   cfnDomain := awscdk.Aws_opensearchservice.NewCfnDomain(this, jsii.String("MyCfnDomain"), &cfnDomainProps{
//   	accessPolicies: accessPolicies,
//   	advancedOptions: map[string]*string{
//   		"advancedOptionsKey": jsii.String("advancedOptions"),
//   	},
//   	advancedSecurityOptions: &advancedSecurityOptionsInputProperty{
//   		enabled: jsii.Boolean(false),
//   		internalUserDatabaseEnabled: jsii.Boolean(false),
//   		masterUserOptions: &masterUserOptionsProperty{
//   			masterUserArn: jsii.String("masterUserArn"),
//   			masterUserName: jsii.String("masterUserName"),
//   			masterUserPassword: jsii.String("masterUserPassword"),
//   		},
//   	},
//   	clusterConfig: &clusterConfigProperty{
//   		dedicatedMasterCount: jsii.Number(123),
//   		dedicatedMasterEnabled: jsii.Boolean(false),
//   		dedicatedMasterType: jsii.String("dedicatedMasterType"),
//   		instanceCount: jsii.Number(123),
//   		instanceType: jsii.String("instanceType"),
//   		warmCount: jsii.Number(123),
//   		warmEnabled: jsii.Boolean(false),
//   		warmType: jsii.String("warmType"),
//   		zoneAwarenessConfig: &zoneAwarenessConfigProperty{
//   			availabilityZoneCount: jsii.Number(123),
//   		},
//   		zoneAwarenessEnabled: jsii.Boolean(false),
//   	},
//   	cognitoOptions: &cognitoOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		identityPoolId: jsii.String("identityPoolId"),
//   		roleArn: jsii.String("roleArn"),
//   		userPoolId: jsii.String("userPoolId"),
//   	},
//   	domainEndpointOptions: &domainEndpointOptionsProperty{
//   		customEndpoint: jsii.String("customEndpoint"),
//   		customEndpointCertificateArn: jsii.String("customEndpointCertificateArn"),
//   		customEndpointEnabled: jsii.Boolean(false),
//   		enforceHttps: jsii.Boolean(false),
//   		tlsSecurityPolicy: jsii.String("tlsSecurityPolicy"),
//   	},
//   	domainName: jsii.String("domainName"),
//   	ebsOptions: &eBSOptionsProperty{
//   		ebsEnabled: jsii.Boolean(false),
//   		iops: jsii.Number(123),
//   		throughput: jsii.Number(123),
//   		volumeSize: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//   	},
//   	encryptionAtRestOptions: &encryptionAtRestOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	engineVersion: jsii.String("engineVersion"),
//   	logPublishingOptions: map[string]interface{}{
//   		"logPublishingOptionsKey": &LogPublishingOptionProperty{
//   			"cloudWatchLogsLogGroupArn": jsii.String("cloudWatchLogsLogGroupArn"),
//   			"enabled": jsii.Boolean(false),
//   		},
//   	},
//   	nodeToNodeEncryptionOptions: &nodeToNodeEncryptionOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	snapshotOptions: &snapshotOptionsProperty{
//   		automatedSnapshotStartHour: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcOptions: &vPCOptionsProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   })
//
type CfnDomain interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// An AWS Identity and Access Management ( IAM ) policy document that specifies who can access the OpenSearch Service domain and their permissions.
	//
	// For more information, see [Configuring access policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ac.html#ac-creating) in the *Amazon OpenSearch Service Developer Guide* .
	AccessPolicies() interface{}
	SetAccessPolicies(val interface{})
	// Additional options to specify for the OpenSearch Service domain.
	//
	// For more information, see [AdvancedOptions](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/configuration-api.html#configuration-api-datatypes-advancedoptions) in the OpenSearch Service configuration API reference.
	AdvancedOptions() interface{}
	SetAdvancedOptions(val interface{})
	// Specifies options for fine-grained access control.
	AdvancedSecurityOptions() interface{}
	SetAdvancedSecurityOptions(val interface{})
	// The Amazon Resource Name (ARN) of the domain, such as `arn:aws:es:us-west-2:123456789012:domain/mystack-1ab2cdefghij` .
	AttrArn() *string
	// The domain-specific endpoint used for requests to the OpenSearch APIs, such as `search-mystack-1ab2cdefghij-ab1c2deckoyb3hofw7wpqa3cm.us-west-1.es.amazonaws.com` .
	AttrDomainEndpoint() *string
	// The resource ID.
	//
	// For example, `123456789012/my-domain` .
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `ClusterConfig` is a property of the AWS::OpenSearchService::Domain resource that configures an Amazon OpenSearch Service cluster.
	ClusterConfig() interface{}
	SetClusterConfig(val interface{})
	// Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
	CognitoOptions() interface{}
	SetCognitoOptions(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.
	DomainEndpointOptions() interface{}
	SetDomainEndpointOptions(val interface{})
	// A name for the OpenSearch Service domain.
	//
	// For valid values, see the [DomainName](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/configuration-api.html#configuration-api-datatypes-domainname) data type in the *Amazon OpenSearch Service Developer Guide* . If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the domain name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// Required when creating a new domain.
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DomainName() *string
	SetDomainName(val *string)
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the OpenSearch Service domain.
	//
	// For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
	EbsOptions() interface{}
	SetEbsOptions(val interface{})
	// Whether the domain should encrypt data at rest, and if so, the AWS KMS key to use.
	//
	// See [Encryption of data at rest for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html) .
	EncryptionAtRestOptions() interface{}
	SetEncryptionAtRestOptions(val interface{})
	// The version of OpenSearch to use.
	//
	// The value must be in the format `OpenSearch_X.Y` or `Elasticsearch_X.Y` . If not specified, the latest version of OpenSearch is used. For information about the versions that OpenSearch Service supports, see [Supported versions of OpenSearch and Elasticsearch](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/what-is.html#choosing-version) in the *Amazon OpenSearch Service Developer Guide* .
	//
	// If you set the [EnableVersionUpgrade](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeopensearchdomain) update policy to `true` , you can update `EngineVersion` without interruption. When `EnableVersionUpgrade` is set to `false` , or is not specified, updating `EngineVersion` results in [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	EngineVersion() *string
	SetEngineVersion(val *string)
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
	//
	// Each key needs a valid `LogPublishingOption` value. For the full syntax, see the [examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#aws-resource-opensearchservice-domain--examples) .
	LogPublishingOptions() interface{}
	SetLogPublishingOptions(val interface{})
	// The tree node.
	Node() constructs.Node
	// Specifies whether node-to-node encryption is enabled.
	//
	// See [Node-to-node encryption for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html) .
	NodeToNodeEncryptionOptions() interface{}
	SetNodeToNodeEncryptionOptions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// *DEPRECATED* .
	//
	// The automated snapshot configuration for the OpenSearch Service domain indices.
	SnapshotOptions() interface{}
	SetSnapshotOptions(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An arbitrary set of tags (key–value pairs) to associate with the OpenSearch Service domain.
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
	// The virtual private cloud (VPC) configuration for the OpenSearch Service domain.
	//
	// For more information, see [Launching your Amazon OpenSearch Service domains within a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the *Amazon OpenSearch Service Developer Guide* .
	VpcOptions() interface{}
	SetVpcOptions(val interface{})
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

func (j *jsiiProxy_CfnDomain) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
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


// Create a new `AWS::OpenSearchService::Domain`.
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

// Create a new `AWS::OpenSearchService::Domain`.
func NewCfnDomain_Override(c CfnDomain, scope constructs.Construct, id *string, props *CfnDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomain)SetAccessPolicies(val interface{}) {
	if err := j.validateSetAccessPoliciesParameters(val); err != nil {
		panic(err)
	}
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

// Check whether the given construct is a CfnResource.
func CfnDomain_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnDomain_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain",
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

func (c *jsiiProxy_CfnDomain) GetAtt(attributeName *string) awscdk.Reference {
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

