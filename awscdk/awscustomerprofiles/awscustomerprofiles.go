package awscustomerprofiles

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::CustomerProfiles::Domain`.
//
// Specifies an Amazon Connect Customer Profiles Domain.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   cfnDomain := customerprofiles.NewCfnDomain(this, jsii.String("MyCfnDomain"), &cfnDomainProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	deadLetterQueueUrl: jsii.String("deadLetterQueueUrl"),
//   	defaultEncryptionKey: jsii.String("defaultEncryptionKey"),
//   	defaultExpirationDays: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDomain interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The timestamp of when the domain was created.
	AttrCreatedAt() *string
	// The timestamp of when the domain was most recently edited.
	AttrLastUpdatedAt() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The URL of the SQS dead letter queue, which is used for reporting errors associated with ingesting data from third party applications.
	//
	// You must set up a policy on the DeadLetterQueue for the SendMessage operation to enable Amazon Connect Customer Profiles to send messages to the DeadLetterQueue.
	DeadLetterQueueUrl() *string
	SetDeadLetterQueueUrl(val *string)
	// The default encryption key, which is an AWS managed key, is used when no specific type of encryption key is specified.
	//
	// It is used to encrypt all data before it is placed in permanent or semi-permanent storage.
	DefaultEncryptionKey() *string
	SetDefaultEncryptionKey(val *string)
	// The default number of days until the data within the domain expires.
	DefaultExpirationDays() *float64
	SetDefaultExpirationDays(val *float64)
	// The unique name of the domain.
	DomainName() *string
	SetDomainName(val *string)
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags used to organize, track, or control access for this resource.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

func (j *jsiiProxy_CfnDomain) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrLastUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedAt",
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

func (j *jsiiProxy_CfnDomain) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DeadLetterQueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deadLetterQueueUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DefaultEncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DefaultExpirationDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultExpirationDays",
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

func (j *jsiiProxy_CfnDomain) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
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

func (j *jsiiProxy_CfnDomain) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
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


// Create a new `AWS::CustomerProfiles::Domain`.
func NewCfnDomain(scope constructs.Construct, id *string, props *CfnDomainProps) CfnDomain {
	_init_.Initialize()

	j := jsiiProxy_CfnDomain{}

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CustomerProfiles::Domain`.
func NewCfnDomain_Override(c CfnDomain, scope constructs.Construct, id *string, props *CfnDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomain) SetDeadLetterQueueUrl(val *string) {
	_jsii_.Set(
		j,
		"deadLetterQueueUrl",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDefaultEncryptionKey(val *string) {
	_jsii_.Set(
		j,
		"defaultEncryptionKey",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDefaultExpirationDays(val *float64) {
	_jsii_.Set(
		j,
		"defaultExpirationDays",
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
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
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
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
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
		"aws-cdk-lib.aws_customerprofiles.CfnDomain",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomain) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDomain) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomain) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDomain) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDomain) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDomain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnDomain) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDomain`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   cfnDomainProps := &cfnDomainProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	deadLetterQueueUrl: jsii.String("deadLetterQueueUrl"),
//   	defaultEncryptionKey: jsii.String("defaultEncryptionKey"),
//   	defaultExpirationDays: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDomainProps struct {
	// The unique name of the domain.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// The URL of the SQS dead letter queue, which is used for reporting errors associated with ingesting data from third party applications.
	//
	// You must set up a policy on the DeadLetterQueue for the SendMessage operation to enable Amazon Connect Customer Profiles to send messages to the DeadLetterQueue.
	DeadLetterQueueUrl *string `json:"deadLetterQueueUrl" yaml:"deadLetterQueueUrl"`
	// The default encryption key, which is an AWS managed key, is used when no specific type of encryption key is specified.
	//
	// It is used to encrypt all data before it is placed in permanent or semi-permanent storage.
	DefaultEncryptionKey *string `json:"defaultEncryptionKey" yaml:"defaultEncryptionKey"`
	// The default number of days until the data within the domain expires.
	DefaultExpirationDays *float64 `json:"defaultExpirationDays" yaml:"defaultExpirationDays"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::CustomerProfiles::Integration`.
//
// Specifies an Amazon Connect Customer Profiles Integration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   cfnIntegration := customerprofiles.NewCfnIntegration(this, jsii.String("MyCfnIntegration"), &cfnIntegrationProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	flowDefinition: &flowDefinitionProperty{
//   		flowName: jsii.String("flowName"),
//   		kmsArn: jsii.String("kmsArn"),
//   		sourceFlowConfig: &sourceFlowConfigProperty{
//   			connectorType: jsii.String("connectorType"),
//   			sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   				marketo: &marketoSourcePropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   				s3: &s3SourcePropertiesProperty{
//   					bucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				salesforce: &salesforceSourcePropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					enableDynamicFieldUpdate: jsii.Boolean(false),
//   					includeDeletedRecords: jsii.Boolean(false),
//   				},
//   				serviceNow: &serviceNowSourcePropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   				zendesk: &zendeskSourcePropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			connectorProfileName: jsii.String("connectorProfileName"),
//   			incrementalPullConfig: &incrementalPullConfigProperty{
//   				datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   			},
//   		},
//   		tasks: []interface{}{
//   			&taskProperty{
//   				sourceFields: []*string{
//   					jsii.String("sourceFields"),
//   				},
//   				taskType: jsii.String("taskType"),
//
//   				// the properties below are optional
//   				connectorOperator: &connectorOperatorProperty{
//   					marketo: jsii.String("marketo"),
//   					s3: jsii.String("s3"),
//   					salesforce: jsii.String("salesforce"),
//   					serviceNow: jsii.String("serviceNow"),
//   					zendesk: jsii.String("zendesk"),
//   				},
//   				destinationField: jsii.String("destinationField"),
//   				taskProperties: []interface{}{
//   					&taskPropertiesMapProperty{
//   						operatorPropertyKey: jsii.String("operatorPropertyKey"),
//   						property: jsii.String("property"),
//   					},
//   				},
//   			},
//   		},
//   		triggerConfig: &triggerConfigProperty{
//   			triggerType: jsii.String("triggerType"),
//
//   			// the properties below are optional
//   			triggerProperties: &triggerPropertiesProperty{
//   				scheduled: &scheduledTriggerPropertiesProperty{
//   					scheduleExpression: jsii.String("scheduleExpression"),
//
//   					// the properties below are optional
//   					dataPullMode: jsii.String("dataPullMode"),
//   					firstExecutionFrom: jsii.Number(123),
//   					scheduleEndTime: jsii.Number(123),
//   					scheduleOffset: jsii.Number(123),
//   					scheduleStartTime: jsii.Number(123),
//   					timezone: jsii.String("timezone"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   	},
//   	objectTypeName: jsii.String("objectTypeName"),
//   	objectTypeNames: []interface{}{
//   		&objectTypeMappingProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	uri: jsii.String("uri"),
//   })
//
type CfnIntegration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The timestamp of when the integration was created.
	AttrCreatedAt() *string
	// The timestamp of when the integration was most recently edited.
	AttrLastUpdatedAt() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The unique name of the domain.
	DomainName() *string
	SetDomainName(val *string)
	// The configuration that controls how Customer Profiles retrieves data from the source.
	FlowDefinition() interface{}
	SetFlowDefinition(val interface{})
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
	// The name of the profile object type mapping to use.
	ObjectTypeName() *string
	SetObjectTypeName(val *string)
	// The object type mapping.
	ObjectTypeNames() interface{}
	SetObjectTypeNames(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags used to organize, track, or control access for this resource.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The URI of the S3 bucket or any other type of data source.
	Uri() *string
	SetUri(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnIntegration
type jsiiProxy_CfnIntegration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIntegration) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) AttrLastUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) FlowDefinition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) ObjectTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) ObjectTypeNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectTypeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}


// Create a new `AWS::CustomerProfiles::Integration`.
func NewCfnIntegration(scope constructs.Construct, id *string, props *CfnIntegrationProps) CfnIntegration {
	_init_.Initialize()

	j := jsiiProxy_CfnIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CustomerProfiles::Integration`.
func NewCfnIntegration_Override(c CfnIntegration, scope constructs.Construct, id *string, props *CfnIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIntegration) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetFlowDefinition(val interface{}) {
	_jsii_.Set(
		j,
		"flowDefinition",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetObjectTypeName(val *string) {
	_jsii_.Set(
		j,
		"objectTypeName",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetObjectTypeNames(val interface{}) {
	_jsii_.Set(
		j,
		"objectTypeNames",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetUri(val *string) {
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnIntegration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnIntegration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIntegration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_customerprofiles.CfnIntegration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIntegration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnIntegration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIntegration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnIntegration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnIntegration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnIntegration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnIntegration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnIntegration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnIntegration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIntegration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The operation to be performed on the provided source fields.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   connectorOperatorProperty := &connectorOperatorProperty{
//   	marketo: jsii.String("marketo"),
//   	s3: jsii.String("s3"),
//   	salesforce: jsii.String("salesforce"),
//   	serviceNow: jsii.String("serviceNow"),
//   	zendesk: jsii.String("zendesk"),
//   }
//
type CfnIntegration_ConnectorOperatorProperty struct {
	// The operation to be performed on the provided Marketo source fields.
	Marketo *string `json:"marketo" yaml:"marketo"`
	// The operation to be performed on the provided Amazon S3 source fields.
	S3 *string `json:"s3" yaml:"s3"`
	// The operation to be performed on the provided Salesforce source fields.
	Salesforce *string `json:"salesforce" yaml:"salesforce"`
	// The operation to be performed on the provided ServiceNow source fields.
	ServiceNow *string `json:"serviceNow" yaml:"serviceNow"`
	// The operation to be performed on the provided Zendesk source fields.
	Zendesk *string `json:"zendesk" yaml:"zendesk"`
}

// The configurations that control how Customer Profiles retrieves data from the source, Amazon AppFlow.
//
// Customer Profiles uses this information to create an AppFlow flow on behalf of customers.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   flowDefinitionProperty := &flowDefinitionProperty{
//   	flowName: jsii.String("flowName"),
//   	kmsArn: jsii.String("kmsArn"),
//   	sourceFlowConfig: &sourceFlowConfigProperty{
//   		connectorType: jsii.String("connectorType"),
//   		sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   			marketo: &marketoSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			s3: &s3SourcePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//
//   				// the properties below are optional
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   			},
//   			salesforce: &salesforceSourcePropertiesProperty{
//   				object: jsii.String("object"),
//
//   				// the properties below are optional
//   				enableDynamicFieldUpdate: jsii.Boolean(false),
//   				includeDeletedRecords: jsii.Boolean(false),
//   			},
//   			serviceNow: &serviceNowSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			zendesk: &zendeskSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		connectorProfileName: jsii.String("connectorProfileName"),
//   		incrementalPullConfig: &incrementalPullConfigProperty{
//   			datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   		},
//   	},
//   	tasks: []interface{}{
//   		&taskProperty{
//   			sourceFields: []*string{
//   				jsii.String("sourceFields"),
//   			},
//   			taskType: jsii.String("taskType"),
//
//   			// the properties below are optional
//   			connectorOperator: &connectorOperatorProperty{
//   				marketo: jsii.String("marketo"),
//   				s3: jsii.String("s3"),
//   				salesforce: jsii.String("salesforce"),
//   				serviceNow: jsii.String("serviceNow"),
//   				zendesk: jsii.String("zendesk"),
//   			},
//   			destinationField: jsii.String("destinationField"),
//   			taskProperties: []interface{}{
//   				&taskPropertiesMapProperty{
//   					operatorPropertyKey: jsii.String("operatorPropertyKey"),
//   					property: jsii.String("property"),
//   				},
//   			},
//   		},
//   	},
//   	triggerConfig: &triggerConfigProperty{
//   		triggerType: jsii.String("triggerType"),
//
//   		// the properties below are optional
//   		triggerProperties: &triggerPropertiesProperty{
//   			scheduled: &scheduledTriggerPropertiesProperty{
//   				scheduleExpression: jsii.String("scheduleExpression"),
//
//   				// the properties below are optional
//   				dataPullMode: jsii.String("dataPullMode"),
//   				firstExecutionFrom: jsii.Number(123),
//   				scheduleEndTime: jsii.Number(123),
//   				scheduleOffset: jsii.Number(123),
//   				scheduleStartTime: jsii.Number(123),
//   				timezone: jsii.String("timezone"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnIntegration_FlowDefinitionProperty struct {
	// The specified name of the flow.
	//
	// Use underscores (_) or hyphens (-) only. Spaces are not allowed.
	FlowName *string `json:"flowName" yaml:"flowName"`
	// The Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key you provide for encryption.
	KmsArn *string `json:"kmsArn" yaml:"kmsArn"`
	// The configuration that controls how Customer Profiles retrieves data from the source.
	SourceFlowConfig interface{} `json:"sourceFlowConfig" yaml:"sourceFlowConfig"`
	// A list of tasks that Customer Profiles performs while transferring the data in the flow run.
	Tasks interface{} `json:"tasks" yaml:"tasks"`
	// The trigger settings that determine how and when the flow runs.
	TriggerConfig interface{} `json:"triggerConfig" yaml:"triggerConfig"`
	// A description of the flow you want to create.
	Description *string `json:"description" yaml:"description"`
}

// Specifies the configuration used when importing incremental records from the source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   incrementalPullConfigProperty := &incrementalPullConfigProperty{
//   	datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   }
//
type CfnIntegration_IncrementalPullConfigProperty struct {
	// A field that specifies the date time or timestamp field as the criteria to use when importing incremental records from the source.
	DatetimeTypeFieldName *string `json:"datetimeTypeFieldName" yaml:"datetimeTypeFieldName"`
}

// The properties that are applied when Marketo is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   marketoSourcePropertiesProperty := &marketoSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnIntegration_MarketoSourcePropertiesProperty struct {
	// The object specified in the Marketo flow source.
	Object *string `json:"object" yaml:"object"`
}

// A map in which each key is an event type from an external application such as Segment or Shopify, and each value is an `ObjectTypeName` (template) used to ingest the event.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   objectTypeMappingProperty := &objectTypeMappingProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnIntegration_ObjectTypeMappingProperty struct {
	// The key.
	Key *string `json:"key" yaml:"key"`
	// The value.
	Value *string `json:"value" yaml:"value"`
}

// The properties that are applied when Amazon S3 is being used as the flow source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   s3SourcePropertiesProperty := &s3SourcePropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   }
//
type CfnIntegration_S3SourcePropertiesProperty struct {
	// The Amazon S3 bucket name where the source files are stored.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The object key for the Amazon S3 bucket in which the source files are stored.
	BucketPrefix *string `json:"bucketPrefix" yaml:"bucketPrefix"`
}

// The properties that are applied when Salesforce is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   salesforceSourcePropertiesProperty := &salesforceSourcePropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	enableDynamicFieldUpdate: jsii.Boolean(false),
//   	includeDeletedRecords: jsii.Boolean(false),
//   }
//
type CfnIntegration_SalesforceSourcePropertiesProperty struct {
	// The object specified in the Salesforce flow source.
	Object *string `json:"object" yaml:"object"`
	// The flag that enables dynamic fetching of new (recently added) fields in the Salesforce objects while running a flow.
	EnableDynamicFieldUpdate interface{} `json:"enableDynamicFieldUpdate" yaml:"enableDynamicFieldUpdate"`
	// Indicates whether Amazon AppFlow includes deleted files in the flow run.
	IncludeDeletedRecords interface{} `json:"includeDeletedRecords" yaml:"includeDeletedRecords"`
}

// Specifies the configuration details of a scheduled-trigger flow that you define.
//
// Currently, these settings only apply to the scheduled-trigger type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   scheduledTriggerPropertiesProperty := &scheduledTriggerPropertiesProperty{
//   	scheduleExpression: jsii.String("scheduleExpression"),
//
//   	// the properties below are optional
//   	dataPullMode: jsii.String("dataPullMode"),
//   	firstExecutionFrom: jsii.Number(123),
//   	scheduleEndTime: jsii.Number(123),
//   	scheduleOffset: jsii.Number(123),
//   	scheduleStartTime: jsii.Number(123),
//   	timezone: jsii.String("timezone"),
//   }
//
type CfnIntegration_ScheduledTriggerPropertiesProperty struct {
	// The scheduling expression that determines the rate at which the schedule will run, for example rate (5 minutes).
	ScheduleExpression *string `json:"scheduleExpression" yaml:"scheduleExpression"`
	// Specifies whether a scheduled flow has an incremental data transfer or a complete data transfer for each flow run.
	DataPullMode *string `json:"dataPullMode" yaml:"dataPullMode"`
	// Specifies the date range for the records to import from the connector in the first flow run.
	FirstExecutionFrom *float64 `json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// Specifies the scheduled end time for a scheduled-trigger flow.
	ScheduleEndTime *float64 `json:"scheduleEndTime" yaml:"scheduleEndTime"`
	// Specifies the optional offset that is added to the time interval for a schedule-triggered flow.
	ScheduleOffset *float64 `json:"scheduleOffset" yaml:"scheduleOffset"`
	// Specifies the scheduled start time for a scheduled-trigger flow.
	ScheduleStartTime *float64 `json:"scheduleStartTime" yaml:"scheduleStartTime"`
	// Specifies the time zone used when referring to the date and time of a scheduled-triggered flow, such as America/New_York.
	Timezone *string `json:"timezone" yaml:"timezone"`
}

// The properties that are applied when ServiceNow is being used as a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   serviceNowSourcePropertiesProperty := &serviceNowSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnIntegration_ServiceNowSourcePropertiesProperty struct {
	// The object specified in the ServiceNow flow source.
	Object *string `json:"object" yaml:"object"`
}

// Specifies the information that is required to query a particular Amazon AppFlow connector.
//
// Customer Profiles supports Salesforce, Zendesk, Marketo, ServiceNow and Amazon S3.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   sourceConnectorPropertiesProperty := &sourceConnectorPropertiesProperty{
//   	marketo: &marketoSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	s3: &s3SourcePropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   	},
//   	salesforce: &salesforceSourcePropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		enableDynamicFieldUpdate: jsii.Boolean(false),
//   		includeDeletedRecords: jsii.Boolean(false),
//   	},
//   	serviceNow: &serviceNowSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	zendesk: &zendeskSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   }
//
type CfnIntegration_SourceConnectorPropertiesProperty struct {
	// The properties that are applied when Marketo is being used as a source.
	Marketo interface{} `json:"marketo" yaml:"marketo"`
	// The properties that are applied when Amazon S3 is being used as the flow source.
	S3 interface{} `json:"s3" yaml:"s3"`
	// The properties that are applied when Salesforce is being used as a source.
	Salesforce interface{} `json:"salesforce" yaml:"salesforce"`
	// The properties that are applied when ServiceNow is being used as a source.
	ServiceNow interface{} `json:"serviceNow" yaml:"serviceNow"`
	// The properties that are applied when using Zendesk as a flow source.
	Zendesk interface{} `json:"zendesk" yaml:"zendesk"`
}

// The configuration that controls how Customer Profiles retrieves data from the source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   sourceFlowConfigProperty := &sourceFlowConfigProperty{
//   	connectorType: jsii.String("connectorType"),
//   	sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   		marketo: &marketoSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		s3: &s3SourcePropertiesProperty{
//   			bucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   		},
//   		salesforce: &salesforceSourcePropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			enableDynamicFieldUpdate: jsii.Boolean(false),
//   			includeDeletedRecords: jsii.Boolean(false),
//   		},
//   		serviceNow: &serviceNowSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		zendesk: &zendeskSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	connectorProfileName: jsii.String("connectorProfileName"),
//   	incrementalPullConfig: &incrementalPullConfigProperty{
//   		datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   	},
//   }
//
type CfnIntegration_SourceFlowConfigProperty struct {
	// The type of connector, such as Salesforce, Marketo, and so on.
	ConnectorType *string `json:"connectorType" yaml:"connectorType"`
	// Specifies the information that is required to query a particular source connector.
	SourceConnectorProperties interface{} `json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
	// The name of the Amazon AppFlow connector profile.
	//
	// This name must be unique for each connector profile in the AWS account .
	ConnectorProfileName *string `json:"connectorProfileName" yaml:"connectorProfileName"`
	// Defines the configuration for a scheduled incremental data pull.
	//
	// If a valid configuration is provided, the fields specified in the configuration are used when querying for the incremental data pull.
	IncrementalPullConfig interface{} `json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

// A map used to store task-related information.
//
// The execution service looks for particular information based on the `TaskType` .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   taskPropertiesMapProperty := &taskPropertiesMapProperty{
//   	operatorPropertyKey: jsii.String("operatorPropertyKey"),
//   	property: jsii.String("property"),
//   }
//
type CfnIntegration_TaskPropertiesMapProperty struct {
	// The task property key.
	OperatorPropertyKey *string `json:"operatorPropertyKey" yaml:"operatorPropertyKey"`
	// The task property value.
	Property *string `json:"property" yaml:"property"`
}

// The `Task` property type specifies the class for modeling different type of tasks.
//
// Task implementation varies based on the TaskType.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   taskProperty := &taskProperty{
//   	sourceFields: []*string{
//   		jsii.String("sourceFields"),
//   	},
//   	taskType: jsii.String("taskType"),
//
//   	// the properties below are optional
//   	connectorOperator: &connectorOperatorProperty{
//   		marketo: jsii.String("marketo"),
//   		s3: jsii.String("s3"),
//   		salesforce: jsii.String("salesforce"),
//   		serviceNow: jsii.String("serviceNow"),
//   		zendesk: jsii.String("zendesk"),
//   	},
//   	destinationField: jsii.String("destinationField"),
//   	taskProperties: []interface{}{
//   		&taskPropertiesMapProperty{
//   			operatorPropertyKey: jsii.String("operatorPropertyKey"),
//   			property: jsii.String("property"),
//   		},
//   	},
//   }
//
type CfnIntegration_TaskProperty struct {
	// The source fields to which a particular task is applied.
	SourceFields *[]*string `json:"sourceFields" yaml:"sourceFields"`
	// Specifies the particular task implementation that Amazon AppFlow performs.
	TaskType *string `json:"taskType" yaml:"taskType"`
	// The operation to be performed on the provided source fields.
	ConnectorOperator interface{} `json:"connectorOperator" yaml:"connectorOperator"`
	// A field in a destination connector, or a field value against which Amazon AppFlow validates a source field.
	DestinationField *string `json:"destinationField" yaml:"destinationField"`
	// A map used to store task-related information.
	//
	// The service looks for particular information based on the TaskType.
	TaskProperties interface{} `json:"taskProperties" yaml:"taskProperties"`
}

// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   triggerConfigProperty := &triggerConfigProperty{
//   	triggerType: jsii.String("triggerType"),
//
//   	// the properties below are optional
//   	triggerProperties: &triggerPropertiesProperty{
//   		scheduled: &scheduledTriggerPropertiesProperty{
//   			scheduleExpression: jsii.String("scheduleExpression"),
//
//   			// the properties below are optional
//   			dataPullMode: jsii.String("dataPullMode"),
//   			firstExecutionFrom: jsii.Number(123),
//   			scheduleEndTime: jsii.Number(123),
//   			scheduleOffset: jsii.Number(123),
//   			scheduleStartTime: jsii.Number(123),
//   			timezone: jsii.String("timezone"),
//   		},
//   	},
//   }
//
type CfnIntegration_TriggerConfigProperty struct {
	// Specifies the type of flow trigger.
	//
	// It can be OnDemand, Scheduled, or Event.
	TriggerType *string `json:"triggerType" yaml:"triggerType"`
	// Specifies the configuration details of a schedule-triggered flow that you define.
	//
	// Currently, these settings only apply to the Scheduled trigger type.
	TriggerProperties interface{} `json:"triggerProperties" yaml:"triggerProperties"`
}

// Specifies the configuration details that control the trigger for a flow.
//
// Currently, these settings only apply to the Scheduled trigger type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   triggerPropertiesProperty := &triggerPropertiesProperty{
//   	scheduled: &scheduledTriggerPropertiesProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//
//   		// the properties below are optional
//   		dataPullMode: jsii.String("dataPullMode"),
//   		firstExecutionFrom: jsii.Number(123),
//   		scheduleEndTime: jsii.Number(123),
//   		scheduleOffset: jsii.Number(123),
//   		scheduleStartTime: jsii.Number(123),
//   		timezone: jsii.String("timezone"),
//   	},
//   }
//
type CfnIntegration_TriggerPropertiesProperty struct {
	// Specifies the configuration details of a schedule-triggered flow that you define.
	Scheduled interface{} `json:"scheduled" yaml:"scheduled"`
}

// The properties that are applied when using Zendesk as a flow source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   zendeskSourcePropertiesProperty := &zendeskSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnIntegration_ZendeskSourcePropertiesProperty struct {
	// The object specified in the Zendesk flow source.
	Object *string `json:"object" yaml:"object"`
}

// Properties for defining a `CfnIntegration`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   cfnIntegrationProps := &cfnIntegrationProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	flowDefinition: &flowDefinitionProperty{
//   		flowName: jsii.String("flowName"),
//   		kmsArn: jsii.String("kmsArn"),
//   		sourceFlowConfig: &sourceFlowConfigProperty{
//   			connectorType: jsii.String("connectorType"),
//   			sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   				marketo: &marketoSourcePropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   				s3: &s3SourcePropertiesProperty{
//   					bucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				salesforce: &salesforceSourcePropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					enableDynamicFieldUpdate: jsii.Boolean(false),
//   					includeDeletedRecords: jsii.Boolean(false),
//   				},
//   				serviceNow: &serviceNowSourcePropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   				zendesk: &zendeskSourcePropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			connectorProfileName: jsii.String("connectorProfileName"),
//   			incrementalPullConfig: &incrementalPullConfigProperty{
//   				datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   			},
//   		},
//   		tasks: []interface{}{
//   			&taskProperty{
//   				sourceFields: []*string{
//   					jsii.String("sourceFields"),
//   				},
//   				taskType: jsii.String("taskType"),
//
//   				// the properties below are optional
//   				connectorOperator: &connectorOperatorProperty{
//   					marketo: jsii.String("marketo"),
//   					s3: jsii.String("s3"),
//   					salesforce: jsii.String("salesforce"),
//   					serviceNow: jsii.String("serviceNow"),
//   					zendesk: jsii.String("zendesk"),
//   				},
//   				destinationField: jsii.String("destinationField"),
//   				taskProperties: []interface{}{
//   					&taskPropertiesMapProperty{
//   						operatorPropertyKey: jsii.String("operatorPropertyKey"),
//   						property: jsii.String("property"),
//   					},
//   				},
//   			},
//   		},
//   		triggerConfig: &triggerConfigProperty{
//   			triggerType: jsii.String("triggerType"),
//
//   			// the properties below are optional
//   			triggerProperties: &triggerPropertiesProperty{
//   				scheduled: &scheduledTriggerPropertiesProperty{
//   					scheduleExpression: jsii.String("scheduleExpression"),
//
//   					// the properties below are optional
//   					dataPullMode: jsii.String("dataPullMode"),
//   					firstExecutionFrom: jsii.Number(123),
//   					scheduleEndTime: jsii.Number(123),
//   					scheduleOffset: jsii.Number(123),
//   					scheduleStartTime: jsii.Number(123),
//   					timezone: jsii.String("timezone"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   	},
//   	objectTypeName: jsii.String("objectTypeName"),
//   	objectTypeNames: []interface{}{
//   		&objectTypeMappingProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	uri: jsii.String("uri"),
//   }
//
type CfnIntegrationProps struct {
	// The unique name of the domain.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// The configuration that controls how Customer Profiles retrieves data from the source.
	FlowDefinition interface{} `json:"flowDefinition" yaml:"flowDefinition"`
	// The name of the profile object type mapping to use.
	ObjectTypeName *string `json:"objectTypeName" yaml:"objectTypeName"`
	// The object type mapping.
	ObjectTypeNames interface{} `json:"objectTypeNames" yaml:"objectTypeNames"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The URI of the S3 bucket or any other type of data source.
	Uri *string `json:"uri" yaml:"uri"`
}

// A CloudFormation `AWS::CustomerProfiles::ObjectType`.
//
// Specifies an Amazon Connect Customer Profiles Object Type Mapping.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   cfnObjectType := customerprofiles.NewCfnObjectType(this, jsii.String("MyCfnObjectType"), &cfnObjectTypeProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	allowProfileCreation: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	encryptionKey: jsii.String("encryptionKey"),
//   	expirationDays: jsii.Number(123),
//   	fields: []interface{}{
//   		&fieldMapProperty{
//   			name: jsii.String("name"),
//   			objectTypeField: &objectTypeFieldProperty{
//   				contentType: jsii.String("contentType"),
//   				source: jsii.String("source"),
//   				target: jsii.String("target"),
//   			},
//   		},
//   	},
//   	keys: []interface{}{
//   		&keyMapProperty{
//   			name: jsii.String("name"),
//   			objectTypeKeyList: []interface{}{
//   				&objectTypeKeyProperty{
//   					fieldNames: []*string{
//   						jsii.String("fieldNames"),
//   					},
//   					standardIdentifiers: []*string{
//   						jsii.String("standardIdentifiers"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	objectTypeName: jsii.String("objectTypeName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	templateId: jsii.String("templateId"),
//   })
//
type CfnObjectType interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Indicates whether a profile should be created when data is received if one doesn’t exist for an object of this type.
	//
	// The default is `FALSE` . If the AllowProfileCreation flag is set to `FALSE` , then the service tries to fetch a standard profile and associate this object with the profile. If it is set to `TRUE` , and if no match is found, then the service creates a new standard profile.
	AllowProfileCreation() interface{}
	SetAllowProfileCreation(val interface{})
	// The timestamp of when the object type was created.
	AttrCreatedAt() *string
	// The timestamp of when the object type was most recently edited.
	AttrLastUpdatedAt() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the profile object type mapping.
	Description() *string
	SetDescription(val *string)
	// The unique name of the domain.
	DomainName() *string
	SetDomainName(val *string)
	// The customer-provided key to encrypt the profile object that will be created in this profile object type mapping.
	//
	// If not specified the system will use the encryption key of the domain.
	EncryptionKey() *string
	SetEncryptionKey(val *string)
	// The number of days until the data of this type expires.
	ExpirationDays() *float64
	SetExpirationDays(val *float64)
	// A list of field definitions for the object type mapping.
	Fields() interface{}
	SetFields(val interface{})
	// A list of keys that can be used to map data to the profile or search for the profile.
	Keys() interface{}
	SetKeys(val interface{})
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
	// The name of the profile object type.
	ObjectTypeName() *string
	SetObjectTypeName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags used to organize, track, or control access for this resource.
	Tags() awscdk.TagManager
	// A unique identifier for the template mapping.
	//
	// This can be used instead of specifying the Keys and Fields properties directly.
	TemplateId() *string
	SetTemplateId(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnObjectType
type jsiiProxy_CfnObjectType struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnObjectType) AllowProfileCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowProfileCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) AttrLastUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) EncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) ExpirationDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Fields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Keys() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) ObjectTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) TemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectType) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CustomerProfiles::ObjectType`.
func NewCfnObjectType(scope constructs.Construct, id *string, props *CfnObjectTypeProps) CfnObjectType {
	_init_.Initialize()

	j := jsiiProxy_CfnObjectType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CustomerProfiles::ObjectType`.
func NewCfnObjectType_Override(c CfnObjectType, scope constructs.Construct, id *string, props *CfnObjectTypeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnObjectType) SetAllowProfileCreation(val interface{}) {
	_jsii_.Set(
		j,
		"allowProfileCreation",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetEncryptionKey(val *string) {
	_jsii_.Set(
		j,
		"encryptionKey",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetExpirationDays(val *float64) {
	_jsii_.Set(
		j,
		"expirationDays",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetFields(val interface{}) {
	_jsii_.Set(
		j,
		"fields",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetKeys(val interface{}) {
	_jsii_.Set(
		j,
		"keys",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetObjectTypeName(val *string) {
	_jsii_.Set(
		j,
		"objectTypeName",
		val,
	)
}

func (j *jsiiProxy_CfnObjectType) SetTemplateId(val *string) {
	_jsii_.Set(
		j,
		"templateId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnObjectType_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnObjectType_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnObjectType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnObjectType_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_customerprofiles.CfnObjectType",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnObjectType) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnObjectType) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnObjectType) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnObjectType) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnObjectType) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnObjectType) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnObjectType) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnObjectType) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObjectType) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObjectType) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnObjectType) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnObjectType) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObjectType) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObjectType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObjectType) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A map of the name and ObjectType field.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   fieldMapProperty := &fieldMapProperty{
//   	name: jsii.String("name"),
//   	objectTypeField: &objectTypeFieldProperty{
//   		contentType: jsii.String("contentType"),
//   		source: jsii.String("source"),
//   		target: jsii.String("target"),
//   	},
//   }
//
type CfnObjectType_FieldMapProperty struct {
	// Name of the field.
	Name *string `json:"name" yaml:"name"`
	// Represents a field in a ProfileObjectType.
	ObjectTypeField interface{} `json:"objectTypeField" yaml:"objectTypeField"`
}

// A unique key map that can be used to map data to the profile.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   keyMapProperty := &keyMapProperty{
//   	name: jsii.String("name"),
//   	objectTypeKeyList: []interface{}{
//   		&objectTypeKeyProperty{
//   			fieldNames: []*string{
//   				jsii.String("fieldNames"),
//   			},
//   			standardIdentifiers: []*string{
//   				jsii.String("standardIdentifiers"),
//   			},
//   		},
//   	},
//   }
//
type CfnObjectType_KeyMapProperty struct {
	// Name of the key.
	Name *string `json:"name" yaml:"name"`
	// A list of ObjectTypeKey.
	ObjectTypeKeyList interface{} `json:"objectTypeKeyList" yaml:"objectTypeKeyList"`
}

// Represents a field in a ProfileObjectType.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   objectTypeFieldProperty := &objectTypeFieldProperty{
//   	contentType: jsii.String("contentType"),
//   	source: jsii.String("source"),
//   	target: jsii.String("target"),
//   }
//
type CfnObjectType_ObjectTypeFieldProperty struct {
	// The content type of the field.
	//
	// Used for determining equality when searching.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// A field of a ProfileObject.
	//
	// For example: _source.FirstName, where “_source” is a ProfileObjectType of a Zendesk user and “FirstName” is a field in that ObjectType.
	Source *string `json:"source" yaml:"source"`
	// The location of the data in the standard ProfileObject model.
	//
	// For example: _profile.Address.PostalCode.
	Target *string `json:"target" yaml:"target"`
}

// An object that defines the Key element of a ProfileObject.
//
// A Key is a special element that can be used to search for a customer profile.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   objectTypeKeyProperty := &objectTypeKeyProperty{
//   	fieldNames: []*string{
//   		jsii.String("fieldNames"),
//   	},
//   	standardIdentifiers: []*string{
//   		jsii.String("standardIdentifiers"),
//   	},
//   }
//
type CfnObjectType_ObjectTypeKeyProperty struct {
	// The reference for the key name of the fields map.
	FieldNames *[]*string `json:"fieldNames" yaml:"fieldNames"`
	// The types of keys that a ProfileObject can have.
	//
	// Each ProfileObject can have only 1 UNIQUE key but multiple PROFILE keys. PROFILE means that this key can be used to tie an object to a PROFILE. UNIQUE means that it can be used to uniquely identify an object. If a key a is marked as SECONDARY, it will be used to search for profiles after all other PROFILE keys have been searched. A LOOKUP_ONLY key is only used to match a profile but is not persisted to be used for searching of the profile. A NEW_ONLY key is only used if the profile does not already exist before the object is ingested, otherwise it is only used for matching objects to profiles.
	StandardIdentifiers *[]*string `json:"standardIdentifiers" yaml:"standardIdentifiers"`
}

// Properties for defining a `CfnObjectType`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import customerprofiles "github.com/aws/aws-cdk-go/awscdk/aws_customerprofiles"
//   cfnObjectTypeProps := &cfnObjectTypeProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	allowProfileCreation: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	encryptionKey: jsii.String("encryptionKey"),
//   	expirationDays: jsii.Number(123),
//   	fields: []interface{}{
//   		&fieldMapProperty{
//   			name: jsii.String("name"),
//   			objectTypeField: &objectTypeFieldProperty{
//   				contentType: jsii.String("contentType"),
//   				source: jsii.String("source"),
//   				target: jsii.String("target"),
//   			},
//   		},
//   	},
//   	keys: []interface{}{
//   		&keyMapProperty{
//   			name: jsii.String("name"),
//   			objectTypeKeyList: []interface{}{
//   				&objectTypeKeyProperty{
//   					fieldNames: []*string{
//   						jsii.String("fieldNames"),
//   					},
//   					standardIdentifiers: []*string{
//   						jsii.String("standardIdentifiers"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	objectTypeName: jsii.String("objectTypeName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	templateId: jsii.String("templateId"),
//   }
//
type CfnObjectTypeProps struct {
	// The unique name of the domain.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Indicates whether a profile should be created when data is received if one doesn’t exist for an object of this type.
	//
	// The default is `FALSE` . If the AllowProfileCreation flag is set to `FALSE` , then the service tries to fetch a standard profile and associate this object with the profile. If it is set to `TRUE` , and if no match is found, then the service creates a new standard profile.
	AllowProfileCreation interface{} `json:"allowProfileCreation" yaml:"allowProfileCreation"`
	// The description of the profile object type mapping.
	Description *string `json:"description" yaml:"description"`
	// The customer-provided key to encrypt the profile object that will be created in this profile object type mapping.
	//
	// If not specified the system will use the encryption key of the domain.
	EncryptionKey *string `json:"encryptionKey" yaml:"encryptionKey"`
	// The number of days until the data of this type expires.
	ExpirationDays *float64 `json:"expirationDays" yaml:"expirationDays"`
	// A list of field definitions for the object type mapping.
	Fields interface{} `json:"fields" yaml:"fields"`
	// A list of keys that can be used to map data to the profile or search for the profile.
	Keys interface{} `json:"keys" yaml:"keys"`
	// The name of the profile object type.
	ObjectTypeName *string `json:"objectTypeName" yaml:"objectTypeName"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// A unique identifier for the template mapping.
	//
	// This can be used instead of specifying the Keys and Fields properties directly.
	TemplateId *string `json:"templateId" yaml:"templateId"`
}

