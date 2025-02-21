package awswisdom

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon Q in Connect message template.
//
// The name of the message template has to be unique for each knowledge base. The channel subtype of the message template is immutable and cannot be modified after creation. After the message template is created, you can use the `$LATEST` qualifier to reference the created message template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMessageTemplate := awscdk.Aws_wisdom.NewCfnMessageTemplate(this, jsii.String("MyCfnMessageTemplate"), &CfnMessageTemplateProps{
//   	ChannelSubtype: jsii.String("channelSubtype"),
//   	Content: &ContentProperty{
//   		EmailMessageTemplateContent: &EmailMessageTemplateContentProperty{
//   			Body: &EmailMessageTemplateContentBodyProperty{
//   				Html: &MessageTemplateBodyContentProviderProperty{
//   					Content: jsii.String("content"),
//   				},
//   				PlainText: &MessageTemplateBodyContentProviderProperty{
//   					Content: jsii.String("content"),
//   				},
//   			},
//   			Headers: []interface{}{
//   				&EmailMessageTemplateHeaderProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Subject: jsii.String("subject"),
//   		},
//   		SmsMessageTemplateContent: &SmsMessageTemplateContentProperty{
//   			Body: &SmsMessageTemplateContentBodyProperty{
//   				PlainText: &MessageTemplateBodyContentProviderProperty{
//   					Content: jsii.String("content"),
//   				},
//   			},
//   		},
//   	},
//   	KnowledgeBaseArn: jsii.String("knowledgeBaseArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DefaultAttributes: &MessageTemplateAttributesProperty{
//   		AgentAttributes: &AgentAttributesProperty{
//   			FirstName: jsii.String("firstName"),
//   			LastName: jsii.String("lastName"),
//   		},
//   		CustomAttributes: map[string]*string{
//   			"customAttributesKey": jsii.String("customAttributes"),
//   		},
//   		CustomerProfileAttributes: &CustomerProfileAttributesProperty{
//   			AccountNumber: jsii.String("accountNumber"),
//   			AdditionalInformation: jsii.String("additionalInformation"),
//   			Address1: jsii.String("address1"),
//   			Address2: jsii.String("address2"),
//   			Address3: jsii.String("address3"),
//   			Address4: jsii.String("address4"),
//   			BillingAddress1: jsii.String("billingAddress1"),
//   			BillingAddress2: jsii.String("billingAddress2"),
//   			BillingAddress3: jsii.String("billingAddress3"),
//   			BillingAddress4: jsii.String("billingAddress4"),
//   			BillingCity: jsii.String("billingCity"),
//   			BillingCountry: jsii.String("billingCountry"),
//   			BillingCounty: jsii.String("billingCounty"),
//   			BillingPostalCode: jsii.String("billingPostalCode"),
//   			BillingProvince: jsii.String("billingProvince"),
//   			BillingState: jsii.String("billingState"),
//   			BirthDate: jsii.String("birthDate"),
//   			BusinessEmailAddress: jsii.String("businessEmailAddress"),
//   			BusinessName: jsii.String("businessName"),
//   			BusinessPhoneNumber: jsii.String("businessPhoneNumber"),
//   			City: jsii.String("city"),
//   			Country: jsii.String("country"),
//   			County: jsii.String("county"),
//   			Custom: map[string]*string{
//   				"customKey": jsii.String("custom"),
//   			},
//   			EmailAddress: jsii.String("emailAddress"),
//   			FirstName: jsii.String("firstName"),
//   			Gender: jsii.String("gender"),
//   			HomePhoneNumber: jsii.String("homePhoneNumber"),
//   			LastName: jsii.String("lastName"),
//   			MailingAddress1: jsii.String("mailingAddress1"),
//   			MailingAddress2: jsii.String("mailingAddress2"),
//   			MailingAddress3: jsii.String("mailingAddress3"),
//   			MailingAddress4: jsii.String("mailingAddress4"),
//   			MailingCity: jsii.String("mailingCity"),
//   			MailingCountry: jsii.String("mailingCountry"),
//   			MailingCounty: jsii.String("mailingCounty"),
//   			MailingPostalCode: jsii.String("mailingPostalCode"),
//   			MailingProvince: jsii.String("mailingProvince"),
//   			MailingState: jsii.String("mailingState"),
//   			MiddleName: jsii.String("middleName"),
//   			MobilePhoneNumber: jsii.String("mobilePhoneNumber"),
//   			PartyType: jsii.String("partyType"),
//   			PhoneNumber: jsii.String("phoneNumber"),
//   			PostalCode: jsii.String("postalCode"),
//   			ProfileArn: jsii.String("profileArn"),
//   			ProfileId: jsii.String("profileId"),
//   			Province: jsii.String("province"),
//   			ShippingAddress1: jsii.String("shippingAddress1"),
//   			ShippingAddress2: jsii.String("shippingAddress2"),
//   			ShippingAddress3: jsii.String("shippingAddress3"),
//   			ShippingAddress4: jsii.String("shippingAddress4"),
//   			ShippingCity: jsii.String("shippingCity"),
//   			ShippingCountry: jsii.String("shippingCountry"),
//   			ShippingCounty: jsii.String("shippingCounty"),
//   			ShippingPostalCode: jsii.String("shippingPostalCode"),
//   			ShippingProvince: jsii.String("shippingProvince"),
//   			ShippingState: jsii.String("shippingState"),
//   			State: jsii.String("state"),
//   		},
//   		SystemAttributes: &SystemAttributesProperty{
//   			CustomerEndpoint: &SystemEndpointAttributesProperty{
//   				Address: jsii.String("address"),
//   			},
//   			Name: jsii.String("name"),
//   			SystemEndpoint: &SystemEndpointAttributesProperty{
//   				Address: jsii.String("address"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	GroupingConfiguration: &GroupingConfigurationProperty{
//   		Criteria: jsii.String("criteria"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Language: jsii.String("language"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html
//
type CfnMessageTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggableV2
	// The Amazon Resource Name (ARN) of the message template.
	AttrMessageTemplateArn() *string
	// The checksum value of the message template content that is referenced by the `$LATEST` qualifier.
	//
	// It can be returned in `MessageTemplateData` or `ExtendedMessageTemplateData` . It’s calculated by content, language, `defaultAttributes` and `Attachments` of the message template.
	AttrMessageTemplateContentSha256() *string
	// The identifier of the message template.
	AttrMessageTemplateId() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The channel subtype this message template applies to.
	ChannelSubtype() *string
	SetChannelSubtype(val *string)
	// The content of the message template.
	Content() interface{}
	SetContent(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An object that specifies the default values to use for variables in the message template.
	DefaultAttributes() interface{}
	SetDefaultAttributes(val interface{})
	// The description of the message template.
	Description() *string
	SetDescription(val *string)
	// The configuration information of the external data source.
	GroupingConfiguration() interface{}
	SetGroupingConfiguration(val interface{})
	// The Amazon Resource Name (ARN) of the knowledge base.
	KnowledgeBaseArn() *string
	SetKnowledgeBaseArn(val *string)
	// The language code value for the language in which the quick response is written.
	Language() *string
	SetLanguage(val *string)
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
	// The name of the message template.
	Name() *string
	SetName(val *string)
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

// The jsii proxy struct for CfnMessageTemplate
type jsiiProxy_CfnMessageTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnMessageTemplate) AttrMessageTemplateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMessageTemplateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) AttrMessageTemplateContentSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMessageTemplateContentSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) AttrMessageTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMessageTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) ChannelSubtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelSubtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) Content() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) DefaultAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) GroupingConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) KnowledgeBaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) Language() *string {
	var returns *string
	_jsii_.Get(
		j,
		"language",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplate) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnMessageTemplate(scope constructs.Construct, id *string, props *CfnMessageTemplateProps) CfnMessageTemplate {
	_init_.Initialize()

	if err := validateNewCfnMessageTemplateParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMessageTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wisdom.CfnMessageTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnMessageTemplate_Override(c CfnMessageTemplate, scope constructs.Construct, id *string, props *CfnMessageTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wisdom.CfnMessageTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMessageTemplate)SetChannelSubtype(val *string) {
	if err := j.validateSetChannelSubtypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelSubtype",
		val,
	)
}

func (j *jsiiProxy_CfnMessageTemplate)SetContent(val interface{}) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_CfnMessageTemplate)SetDefaultAttributes(val interface{}) {
	if err := j.validateSetDefaultAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnMessageTemplate)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnMessageTemplate)SetGroupingConfiguration(val interface{}) {
	if err := j.validateSetGroupingConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupingConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnMessageTemplate)SetKnowledgeBaseArn(val *string) {
	if err := j.validateSetKnowledgeBaseArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"knowledgeBaseArn",
		val,
	)
}

func (j *jsiiProxy_CfnMessageTemplate)SetLanguage(val *string) {
	_jsii_.Set(
		j,
		"language",
		val,
	)
}

func (j *jsiiProxy_CfnMessageTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnMessageTemplate)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMessageTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMessageTemplate_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnMessageTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnMessageTemplate_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMessageTemplate_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnMessageTemplate",
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
func CfnMessageTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMessageTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wisdom.CfnMessageTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMessageTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wisdom.CfnMessageTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMessageTemplate) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnMessageTemplate) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnMessageTemplate) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMessageTemplate) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMessageTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnMessageTemplate) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnMessageTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMessageTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMessageTemplate) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

