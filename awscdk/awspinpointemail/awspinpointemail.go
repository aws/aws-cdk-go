package awspinpointemail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpointemail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::PinpointEmail::ConfigurationSet`.
//
// Create a configuration set. *Configuration sets* are groups of rules that you can apply to the emails you send using Amazon Pinpoint. You apply a configuration set to an email by including a reference to the configuration set in the headers of the email. When you apply a configuration set to an email, all of the rules in that configuration set are applied to the email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSet := awscdk.Aws_pinpointemail.NewCfnConfigurationSet(this, jsii.String("MyCfnConfigurationSet"), &cfnConfigurationSetProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	deliveryOptions: &deliveryOptionsProperty{
//   		sendingPoolName: jsii.String("sendingPoolName"),
//   	},
//   	reputationOptions: &reputationOptionsProperty{
//   		reputationMetricsEnabled: jsii.Boolean(false),
//   	},
//   	sendingOptions: &sendingOptionsProperty{
//   		sendingEnabled: jsii.Boolean(false),
//   	},
//   	tags: []tagsProperty{
//   		&tagsProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	trackingOptions: &trackingOptionsProperty{
//   		customRedirectDomain: jsii.String("customRedirectDomain"),
//   	},
//   })
//
type CfnConfigurationSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.
	DeliveryOptions() interface{}
	SetDeliveryOptions(val interface{})
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
	// The name of the configuration set.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// An object that defines whether or not Amazon Pinpoint collects reputation metrics for the emails that you send that use the configuration set.
	ReputationOptions() interface{}
	SetReputationOptions(val interface{})
	// An object that defines whether or not Amazon Pinpoint can send email that you send using the configuration set.
	SendingOptions() interface{}
	SetSendingOptions(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An object that defines the tags (keys and values) that you want to associate with the configuration set.
	Tags() *[]*CfnConfigurationSet_TagsProperty
	SetTags(val *[]*CfnConfigurationSet_TagsProperty)
	// An object that defines the open and click tracking options for emails that you send using the configuration set.
	TrackingOptions() interface{}
	SetTrackingOptions(val interface{})
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

// The jsii proxy struct for CfnConfigurationSet
type jsiiProxy_CfnConfigurationSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConfigurationSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) DeliveryOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) ReputationOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reputationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) SendingOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) Tags() *[]*CfnConfigurationSet_TagsProperty {
	var returns *[]*CfnConfigurationSet_TagsProperty
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) TrackingOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::PinpointEmail::ConfigurationSet`.
func NewCfnConfigurationSet(scope constructs.Construct, id *string, props *CfnConfigurationSetProps) CfnConfigurationSet {
	_init_.Initialize()

	j := jsiiProxy_CfnConfigurationSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::PinpointEmail::ConfigurationSet`.
func NewCfnConfigurationSet_Override(c CfnConfigurationSet, scope constructs.Construct, id *string, props *CfnConfigurationSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConfigurationSet) SetDeliveryOptions(val interface{}) {
	_jsii_.Set(
		j,
		"deliveryOptions",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationSet) SetReputationOptions(val interface{}) {
	_jsii_.Set(
		j,
		"reputationOptions",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationSet) SetSendingOptions(val interface{}) {
	_jsii_.Set(
		j,
		"sendingOptions",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationSet) SetTags(val *[]*CfnConfigurationSet_TagsProperty) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationSet) SetTrackingOptions(val interface{}) {
	_jsii_.Set(
		j,
		"trackingOptions",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConfigurationSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConfigurationSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSet",
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
func CfnConfigurationSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationSet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConfigurationSet) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConfigurationSet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConfigurationSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConfigurationSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConfigurationSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConfigurationSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConfigurationSet) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSet) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConfigurationSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConfigurationSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Used to associate a configuration set with a dedicated IP pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryOptionsProperty := &deliveryOptionsProperty{
//   	sendingPoolName: jsii.String("sendingPoolName"),
//   }
//
type CfnConfigurationSet_DeliveryOptionsProperty struct {
	// The name of the dedicated IP pool that you want to associate with the configuration set.
	SendingPoolName *string `field:"optional" json:"sendingPoolName" yaml:"sendingPoolName"`
}

// Enable or disable collection of reputation metrics for emails that you send using this configuration set in the current AWS Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reputationOptionsProperty := &reputationOptionsProperty{
//   	reputationMetricsEnabled: jsii.Boolean(false),
//   }
//
type CfnConfigurationSet_ReputationOptionsProperty struct {
	// If `true` , tracking of reputation metrics is enabled for the configuration set.
	//
	// If `false` , tracking of reputation metrics is disabled for the configuration set.
	ReputationMetricsEnabled interface{} `field:"optional" json:"reputationMetricsEnabled" yaml:"reputationMetricsEnabled"`
}

// Used to enable or disable email sending for messages that use this configuration set in the current AWS Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sendingOptionsProperty := &sendingOptionsProperty{
//   	sendingEnabled: jsii.Boolean(false),
//   }
//
type CfnConfigurationSet_SendingOptionsProperty struct {
	// If `true` , email sending is enabled for the configuration set.
	//
	// If `false` , email sending is disabled for the configuration set.
	SendingEnabled interface{} `field:"optional" json:"sendingEnabled" yaml:"sendingEnabled"`
}

// An object that defines the tags (keys and values) that you want to associate with the configuration set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsProperty := &tagsProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnConfigurationSet_TagsProperty struct {
	// One part of a key-value pair that defines a tag.
	//
	// The maximum length of a tag key is 128 characters. The minimum length is 1 character.
	//
	// If you specify tags for the configuration set, then this value is required.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The optional part of a key-value pair that defines a tag.
	//
	// The maximum length of a tag value is 256 characters. The minimum length is 0 characters. If you don’t want a resource to have a specific tag value, don’t specify a value for this parameter. Amazon Pinpoint will set the value to an empty string.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// An object that defines the tracking options for a configuration set.
//
// When you use Amazon Pinpoint to send an email, it contains an invisible image that's used to track when recipients open your email. If your email contains links, those links are changed slightly in order to track when recipients click them.
//
// These images and links include references to a domain operated by AWS . You can optionally configure Amazon Pinpoint to use a domain that you operate for these images and links.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trackingOptionsProperty := &trackingOptionsProperty{
//   	customRedirectDomain: jsii.String("customRedirectDomain"),
//   }
//
type CfnConfigurationSet_TrackingOptionsProperty struct {
	// The domain that you want to use for tracking open and click events.
	CustomRedirectDomain *string `field:"optional" json:"customRedirectDomain" yaml:"customRedirectDomain"`
}

// A CloudFormation `AWS::PinpointEmail::ConfigurationSetEventDestination`.
//
// Create an event destination. In Amazon Pinpoint, *events* include message sends, deliveries, opens, clicks, bounces, and complaints. *Event destinations* are places that you can send information about these events to. For example, you can send event data to Amazon SNS to receive notifications when you receive bounces or complaints, or you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for long-term storage.
//
// A single configuration set can include more than one event destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSetEventDestination := awscdk.Aws_pinpointemail.NewCfnConfigurationSetEventDestination(this, jsii.String("MyCfnConfigurationSetEventDestination"), &cfnConfigurationSetEventDestinationProps{
//   	configurationSetName: jsii.String("configurationSetName"),
//   	eventDestinationName: jsii.String("eventDestinationName"),
//
//   	// the properties below are optional
//   	eventDestination: &eventDestinationProperty{
//   		matchingEventTypes: []*string{
//   			jsii.String("matchingEventTypes"),
//   		},
//
//   		// the properties below are optional
//   		cloudWatchDestination: &cloudWatchDestinationProperty{
//   			dimensionConfigurations: []interface{}{
//   				&dimensionConfigurationProperty{
//   					defaultDimensionValue: jsii.String("defaultDimensionValue"),
//   					dimensionName: jsii.String("dimensionName"),
//   					dimensionValueSource: jsii.String("dimensionValueSource"),
//   				},
//   			},
//   		},
//   		enabled: jsii.Boolean(false),
//   		kinesisFirehoseDestination: &kinesisFirehoseDestinationProperty{
//   			deliveryStreamArn: jsii.String("deliveryStreamArn"),
//   			iamRoleArn: jsii.String("iamRoleArn"),
//   		},
//   		pinpointDestination: &pinpointDestinationProperty{
//   			applicationArn: jsii.String("applicationArn"),
//   		},
//   		snsDestination: &snsDestinationProperty{
//   			topicArn: jsii.String("topicArn"),
//   		},
//   	},
//   })
//
type CfnConfigurationSetEventDestination interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The name of the configuration set that contains the event destination that you want to modify.
	ConfigurationSetName() *string
	SetConfigurationSetName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An object that defines the event destination.
	EventDestination() interface{}
	SetEventDestination(val interface{})
	// The name of the event destination that you want to modify.
	EventDestinationName() *string
	SetEventDestinationName(val *string)
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

// The jsii proxy struct for CfnConfigurationSetEventDestination
type jsiiProxy_CfnConfigurationSetEventDestination struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) ConfigurationSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) EventDestination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) EventDestinationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDestinationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::PinpointEmail::ConfigurationSetEventDestination`.
func NewCfnConfigurationSetEventDestination(scope constructs.Construct, id *string, props *CfnConfigurationSetEventDestinationProps) CfnConfigurationSetEventDestination {
	_init_.Initialize()

	j := jsiiProxy_CfnConfigurationSetEventDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::PinpointEmail::ConfigurationSetEventDestination`.
func NewCfnConfigurationSetEventDestination_Override(c CfnConfigurationSetEventDestination, scope constructs.Construct, id *string, props *CfnConfigurationSetEventDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) SetConfigurationSetName(val *string) {
	_jsii_.Set(
		j,
		"configurationSetName",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) SetEventDestination(val interface{}) {
	_jsii_.Set(
		j,
		"eventDestination",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationSetEventDestination) SetEventDestinationName(val *string) {
	_jsii_.Set(
		j,
		"eventDestinationName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConfigurationSetEventDestination_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConfigurationSetEventDestination_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination",
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
func CfnConfigurationSetEventDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationSetEventDestination_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfigurationSetEventDestination) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object that defines an Amazon CloudWatch destination for email events.
//
// You can use Amazon CloudWatch to monitor and gain insights on your email sending metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchDestinationProperty := &cloudWatchDestinationProperty{
//   	dimensionConfigurations: []interface{}{
//   		&dimensionConfigurationProperty{
//   			defaultDimensionValue: jsii.String("defaultDimensionValue"),
//   			dimensionName: jsii.String("dimensionName"),
//   			dimensionValueSource: jsii.String("dimensionValueSource"),
//   		},
//   	},
//   }
//
type CfnConfigurationSetEventDestination_CloudWatchDestinationProperty struct {
	// An array of objects that define the dimensions to use when you send email events to Amazon CloudWatch.
	DimensionConfigurations interface{} `field:"optional" json:"dimensionConfigurations" yaml:"dimensionConfigurations"`
}

// An array of objects that define the dimensions to use when you send email events to Amazon CloudWatch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dimensionConfigurationProperty := &dimensionConfigurationProperty{
//   	defaultDimensionValue: jsii.String("defaultDimensionValue"),
//   	dimensionName: jsii.String("dimensionName"),
//   	dimensionValueSource: jsii.String("dimensionValueSource"),
//   }
//
type CfnConfigurationSetEventDestination_DimensionConfigurationProperty struct {
	// The default value of the dimension that is published to Amazon CloudWatch if you don't provide the value of the dimension when you send an email.
	//
	// This value has to meet the following criteria:
	//
	// - It can only contain ASCII letters (a–z, A–Z), numbers (0–9), underscores (_), or dashes (-).
	// - It can contain no more than 256 characters.
	DefaultDimensionValue *string `field:"required" json:"defaultDimensionValue" yaml:"defaultDimensionValue"`
	// The name of an Amazon CloudWatch dimension associated with an email sending metric.
	//
	// The name has to meet the following criteria:
	//
	// - It can only contain ASCII letters (a–z, A–Z), numbers (0–9), underscores (_), or dashes (-).
	// - It can contain no more than 256 characters.
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// The location where Amazon Pinpoint finds the value of a dimension to publish to Amazon CloudWatch.
	//
	// Acceptable values: `MESSAGE_TAG` , `EMAIL_HEADER` , and `LINK_TAG` .
	//
	// If you want Amazon Pinpoint to use the message tags that you specify using an `X-SES-MESSAGE-TAGS` header or a parameter to the `SendEmail` API, choose `MESSAGE_TAG` . If you want Amazon Pinpoint to use your own email headers, choose `EMAIL_HEADER` . If you want Amazon Pinpoint to use tags that are specified in your links, choose `LINK_TAG` .
	DimensionValueSource *string `field:"required" json:"dimensionValueSource" yaml:"dimensionValueSource"`
}

// In Amazon Pinpoint, *events* include message sends, deliveries, opens, clicks, bounces, and complaints.
//
// *Event destinations* are places that you can send information about these events to. For example, you can send event data to Amazon SNS to receive notifications when you receive bounces or complaints, or you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for long-term storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventDestinationProperty := &eventDestinationProperty{
//   	matchingEventTypes: []*string{
//   		jsii.String("matchingEventTypes"),
//   	},
//
//   	// the properties below are optional
//   	cloudWatchDestination: &cloudWatchDestinationProperty{
//   		dimensionConfigurations: []interface{}{
//   			&dimensionConfigurationProperty{
//   				defaultDimensionValue: jsii.String("defaultDimensionValue"),
//   				dimensionName: jsii.String("dimensionName"),
//   				dimensionValueSource: jsii.String("dimensionValueSource"),
//   			},
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   	kinesisFirehoseDestination: &kinesisFirehoseDestinationProperty{
//   		deliveryStreamArn: jsii.String("deliveryStreamArn"),
//   		iamRoleArn: jsii.String("iamRoleArn"),
//   	},
//   	pinpointDestination: &pinpointDestinationProperty{
//   		applicationArn: jsii.String("applicationArn"),
//   	},
//   	snsDestination: &snsDestinationProperty{
//   		topicArn: jsii.String("topicArn"),
//   	},
//   }
//
type CfnConfigurationSetEventDestination_EventDestinationProperty struct {
	// The types of events that Amazon Pinpoint sends to the specified event destinations.
	//
	// Acceptable values: `SEND` , `REJECT` , `BOUNCE` , `COMPLAINT` , `DELIVERY` , `OPEN` , `CLICK` , and `RENDERING_FAILURE` .
	MatchingEventTypes *[]*string `field:"required" json:"matchingEventTypes" yaml:"matchingEventTypes"`
	// An object that defines an Amazon CloudWatch destination for email events.
	//
	// You can use Amazon CloudWatch to monitor and gain insights on your email sending metrics.
	CloudWatchDestination interface{} `field:"optional" json:"cloudWatchDestination" yaml:"cloudWatchDestination"`
	// If `true` , the event destination is enabled.
	//
	// When the event destination is enabled, the specified event types are sent to the destinations in this `EventDestinationDefinition` .
	//
	// If `false` , the event destination is disabled. When the event destination is disabled, events aren't sent to the specified destinations.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// An object that defines an Amazon Kinesis Data Firehose destination for email events.
	//
	// You can use Amazon Kinesis Data Firehose to stream data to other services, such as Amazon S3 and Amazon Redshift.
	KinesisFirehoseDestination interface{} `field:"optional" json:"kinesisFirehoseDestination" yaml:"kinesisFirehoseDestination"`
	// An object that defines a Amazon Pinpoint destination for email events.
	//
	// You can use Amazon Pinpoint events to create attributes in Amazon Pinpoint projects. You can use these attributes to create segments for your campaigns.
	PinpointDestination interface{} `field:"optional" json:"pinpointDestination" yaml:"pinpointDestination"`
	// An object that defines an Amazon SNS destination for email events.
	//
	// You can use Amazon SNS to send notification when certain email events occur.
	SnsDestination interface{} `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}

// An object that defines an Amazon Kinesis Data Firehose destination for email events.
//
// You can use Amazon Kinesis Data Firehose to stream data to other services, such as Amazon S3 and Amazon Redshift.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisFirehoseDestinationProperty := &kinesisFirehoseDestinationProperty{
//   	deliveryStreamArn: jsii.String("deliveryStreamArn"),
//   	iamRoleArn: jsii.String("iamRoleArn"),
//   }
//
type CfnConfigurationSetEventDestination_KinesisFirehoseDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Kinesis Data Firehose stream that Amazon Pinpoint sends email events to.
	DeliveryStreamArn *string `field:"required" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// The Amazon Resource Name (ARN) of the IAM role that Amazon Pinpoint uses when sending email events to the Amazon Kinesis Data Firehose stream.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
}

// An object that defines a Amazon Pinpoint destination for email events.
//
// You can use Amazon Pinpoint events to create attributes in Amazon Pinpoint projects. You can use these attributes to create segments for your campaigns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pinpointDestinationProperty := &pinpointDestinationProperty{
//   	applicationArn: jsii.String("applicationArn"),
//   }
//
type CfnConfigurationSetEventDestination_PinpointDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Pinpoint project that you want to send email events to.
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
}

// An object that defines an Amazon SNS destination for email events.
//
// You can use Amazon SNS to send notification when certain email events occur.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsDestinationProperty := &snsDestinationProperty{
//   	topicArn: jsii.String("topicArn"),
//   }
//
type CfnConfigurationSetEventDestination_SnsDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic that you want to publish email events to.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

// Properties for defining a `CfnConfigurationSetEventDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSetEventDestinationProps := &cfnConfigurationSetEventDestinationProps{
//   	configurationSetName: jsii.String("configurationSetName"),
//   	eventDestinationName: jsii.String("eventDestinationName"),
//
//   	// the properties below are optional
//   	eventDestination: &eventDestinationProperty{
//   		matchingEventTypes: []*string{
//   			jsii.String("matchingEventTypes"),
//   		},
//
//   		// the properties below are optional
//   		cloudWatchDestination: &cloudWatchDestinationProperty{
//   			dimensionConfigurations: []interface{}{
//   				&dimensionConfigurationProperty{
//   					defaultDimensionValue: jsii.String("defaultDimensionValue"),
//   					dimensionName: jsii.String("dimensionName"),
//   					dimensionValueSource: jsii.String("dimensionValueSource"),
//   				},
//   			},
//   		},
//   		enabled: jsii.Boolean(false),
//   		kinesisFirehoseDestination: &kinesisFirehoseDestinationProperty{
//   			deliveryStreamArn: jsii.String("deliveryStreamArn"),
//   			iamRoleArn: jsii.String("iamRoleArn"),
//   		},
//   		pinpointDestination: &pinpointDestinationProperty{
//   			applicationArn: jsii.String("applicationArn"),
//   		},
//   		snsDestination: &snsDestinationProperty{
//   			topicArn: jsii.String("topicArn"),
//   		},
//   	},
//   }
//
type CfnConfigurationSetEventDestinationProps struct {
	// The name of the configuration set that contains the event destination that you want to modify.
	ConfigurationSetName *string `field:"required" json:"configurationSetName" yaml:"configurationSetName"`
	// The name of the event destination that you want to modify.
	EventDestinationName *string `field:"required" json:"eventDestinationName" yaml:"eventDestinationName"`
	// An object that defines the event destination.
	EventDestination interface{} `field:"optional" json:"eventDestination" yaml:"eventDestination"`
}

// Properties for defining a `CfnConfigurationSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSetProps := &cfnConfigurationSetProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	deliveryOptions: &deliveryOptionsProperty{
//   		sendingPoolName: jsii.String("sendingPoolName"),
//   	},
//   	reputationOptions: &reputationOptionsProperty{
//   		reputationMetricsEnabled: jsii.Boolean(false),
//   	},
//   	sendingOptions: &sendingOptionsProperty{
//   		sendingEnabled: jsii.Boolean(false),
//   	},
//   	tags: []tagsProperty{
//   		&tagsProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	trackingOptions: &trackingOptionsProperty{
//   		customRedirectDomain: jsii.String("customRedirectDomain"),
//   	},
//   }
//
type CfnConfigurationSetProps struct {
	// The name of the configuration set.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.
	DeliveryOptions interface{} `field:"optional" json:"deliveryOptions" yaml:"deliveryOptions"`
	// An object that defines whether or not Amazon Pinpoint collects reputation metrics for the emails that you send that use the configuration set.
	ReputationOptions interface{} `field:"optional" json:"reputationOptions" yaml:"reputationOptions"`
	// An object that defines whether or not Amazon Pinpoint can send email that you send using the configuration set.
	SendingOptions interface{} `field:"optional" json:"sendingOptions" yaml:"sendingOptions"`
	// An object that defines the tags (keys and values) that you want to associate with the configuration set.
	Tags *[]*CfnConfigurationSet_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
	// An object that defines the open and click tracking options for emails that you send using the configuration set.
	TrackingOptions interface{} `field:"optional" json:"trackingOptions" yaml:"trackingOptions"`
}

// A CloudFormation `AWS::PinpointEmail::DedicatedIpPool`.
//
// A request to create a new dedicated IP pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDedicatedIpPool := awscdk.Aws_pinpointemail.NewCfnDedicatedIpPool(this, jsii.String("MyCfnDedicatedIpPool"), &cfnDedicatedIpPoolProps{
//   	poolName: jsii.String("poolName"),
//   	tags: []tagsProperty{
//   		&tagsProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDedicatedIpPool interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The name of the dedicated IP pool.
	PoolName() *string
	SetPoolName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An object that defines the tags (keys and values) that you want to associate with the dedicated IP pool.
	Tags() *[]*CfnDedicatedIpPool_TagsProperty
	SetTags(val *[]*CfnDedicatedIpPool_TagsProperty)
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

// The jsii proxy struct for CfnDedicatedIpPool
type jsiiProxy_CfnDedicatedIpPool struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDedicatedIpPool) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDedicatedIpPool) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDedicatedIpPool) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDedicatedIpPool) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDedicatedIpPool) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDedicatedIpPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDedicatedIpPool) PoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDedicatedIpPool) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDedicatedIpPool) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDedicatedIpPool) Tags() *[]*CfnDedicatedIpPool_TagsProperty {
	var returns *[]*CfnDedicatedIpPool_TagsProperty
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDedicatedIpPool) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::PinpointEmail::DedicatedIpPool`.
func NewCfnDedicatedIpPool(scope constructs.Construct, id *string, props *CfnDedicatedIpPoolProps) CfnDedicatedIpPool {
	_init_.Initialize()

	j := jsiiProxy_CfnDedicatedIpPool{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpointemail.CfnDedicatedIpPool",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::PinpointEmail::DedicatedIpPool`.
func NewCfnDedicatedIpPool_Override(c CfnDedicatedIpPool, scope constructs.Construct, id *string, props *CfnDedicatedIpPoolProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpointemail.CfnDedicatedIpPool",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDedicatedIpPool) SetPoolName(val *string) {
	_jsii_.Set(
		j,
		"poolName",
		val,
	)
}

func (j *jsiiProxy_CfnDedicatedIpPool) SetTags(val *[]*CfnDedicatedIpPool_TagsProperty) {
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
func CfnDedicatedIpPool_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnDedicatedIpPool",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDedicatedIpPool_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnDedicatedIpPool",
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
func CfnDedicatedIpPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnDedicatedIpPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDedicatedIpPool_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpointemail.CfnDedicatedIpPool",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDedicatedIpPool) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDedicatedIpPool) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDedicatedIpPool) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDedicatedIpPool) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDedicatedIpPool) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDedicatedIpPool) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDedicatedIpPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDedicatedIpPool) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDedicatedIpPool) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDedicatedIpPool) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDedicatedIpPool) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDedicatedIpPool) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDedicatedIpPool) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDedicatedIpPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDedicatedIpPool) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object that defines the tags (keys and values) that you want to associate with the dedicated IP pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsProperty := &tagsProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnDedicatedIpPool_TagsProperty struct {
	// One part of a key-value pair that defines a tag.
	//
	// The maximum length of a tag key is 128 characters. The minimum length is 1 character.
	//
	// If you specify tags for the dedicated IP pool, then this value is required.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The optional part of a key-value pair that defines a tag.
	//
	// The maximum length of a tag value is 256 characters. The minimum length is 0 characters. If you don’t want a resource to have a specific tag value, don’t specify a value for this parameter. Amazon Pinpoint will set the value to an empty string.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Properties for defining a `CfnDedicatedIpPool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDedicatedIpPoolProps := &cfnDedicatedIpPoolProps{
//   	poolName: jsii.String("poolName"),
//   	tags: []tagsProperty{
//   		&tagsProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDedicatedIpPoolProps struct {
	// The name of the dedicated IP pool.
	PoolName *string `field:"optional" json:"poolName" yaml:"poolName"`
	// An object that defines the tags (keys and values) that you want to associate with the dedicated IP pool.
	Tags *[]*CfnDedicatedIpPool_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::PinpointEmail::Identity`.
//
// Specifies an identity to use for sending email through Amazon Pinpoint. In Amazon Pinpoint, an *identity* is an email address or domain that you use when you send email. Before you can use Amazon Pinpoint to send an email from an identity, you first have to verify it. By verifying an identity, you demonstrate that you're the owner of the address or domain, and that you've given Amazon Pinpoint permission to send email from that identity.
//
// When you verify an email address, Amazon Pinpoint sends an email to the address. Your email address is verified as soon as you follow the link in the verification email.
//
// When you verify a domain, this operation provides a set of DKIM tokens, which you can convert into CNAME tokens. You add these CNAME tokens to the DNS configuration for your domain. Your domain is verified when Amazon Pinpoint detects these records in the DNS configuration for your domain. It usually takes around 72 hours to complete the domain verification process.
//
// > When you use CloudFormation to specify an identity, CloudFormation might indicate that the identity was created successfully. However, you have to verify the identity before you can use it to send email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdentity := awscdk.Aws_pinpointemail.NewCfnIdentity(this, jsii.String("MyCfnIdentity"), &cfnIdentityProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dkimSigningEnabled: jsii.Boolean(false),
//   	feedbackForwardingEnabled: jsii.Boolean(false),
//   	mailFromAttributes: &mailFromAttributesProperty{
//   		behaviorOnMxFailure: jsii.String("behaviorOnMxFailure"),
//   		mailFromDomain: jsii.String("mailFromDomain"),
//   	},
//   	tags: []tagsProperty{
//   		&tagsProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnIdentity interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The host name for the first token that you have to add to the DNS configuration for your domain.
	//
	// For more information, see [Verifying a Domain](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html#channels-email-manage-verify-domain) in the Amazon Pinpoint User Guide.
	AttrIdentityDnsRecordName1() *string
	// The host name for the second token that you have to add to the DNS configuration for your domain.
	AttrIdentityDnsRecordName2() *string
	// The host name for the third token that you have to add to the DNS configuration for your domain.
	AttrIdentityDnsRecordName3() *string
	// The record value for the first token that you have to add to the DNS configuration for your domain.
	AttrIdentityDnsRecordValue1() *string
	// The record value for the second token that you have to add to the DNS configuration for your domain.
	AttrIdentityDnsRecordValue2() *string
	// The record value for the third token that you have to add to the DNS configuration for your domain.
	AttrIdentityDnsRecordValue3() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// For domain identities, this attribute is used to enable or disable DomainKeys Identified Mail (DKIM) signing for the domain.
	//
	// If the value is `true` , then the messages that you send from the domain are signed using both the DKIM keys for your domain, as well as the keys for the `amazonses.com` domain. If the value is `false` , then the messages that you send are only signed using the DKIM keys for the `amazonses.com` domain.
	DkimSigningEnabled() interface{}
	SetDkimSigningEnabled(val interface{})
	// Used to enable or disable feedback forwarding for an identity.
	//
	// This setting determines what happens when an identity is used to send an email that results in a bounce or complaint event.
	//
	// When you enable feedback forwarding, Amazon Pinpoint sends you email notifications when bounce or complaint events occur. Amazon Pinpoint sends this notification to the address that you specified in the Return-Path header of the original email.
	//
	// When you disable feedback forwarding, Amazon Pinpoint sends notifications through other mechanisms, such as by notifying an Amazon SNS topic. You're required to have a method of tracking bounces and complaints. If you haven't set up another mechanism for receiving bounce or complaint notifications, Amazon Pinpoint sends an email notification when these events occur (even if this setting is disabled).
	FeedbackForwardingEnabled() interface{}
	SetFeedbackForwardingEnabled(val interface{})
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
	// Used to enable or disable the custom Mail-From domain configuration for an email identity.
	MailFromAttributes() interface{}
	SetMailFromAttributes(val interface{})
	// The address or domain of the identity, such as *sender@example.com* or *example.co.uk* .
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
	// An object that defines the tags (keys and values) that you want to associate with the email identity.
	Tags() *[]*CfnIdentity_TagsProperty
	SetTags(val *[]*CfnIdentity_TagsProperty)
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

// The jsii proxy struct for CfnIdentity
type jsiiProxy_CfnIdentity struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIdentity) AttrIdentityDnsRecordName1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIdentityDnsRecordName1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) AttrIdentityDnsRecordName2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIdentityDnsRecordName2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) AttrIdentityDnsRecordName3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIdentityDnsRecordName3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) AttrIdentityDnsRecordValue1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIdentityDnsRecordValue1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) AttrIdentityDnsRecordValue2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIdentityDnsRecordValue2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) AttrIdentityDnsRecordValue3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIdentityDnsRecordValue3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) DkimSigningEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dkimSigningEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) FeedbackForwardingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"feedbackForwardingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) MailFromAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mailFromAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) Tags() *[]*CfnIdentity_TagsProperty {
	var returns *[]*CfnIdentity_TagsProperty
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentity) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::PinpointEmail::Identity`.
func NewCfnIdentity(scope constructs.Construct, id *string, props *CfnIdentityProps) CfnIdentity {
	_init_.Initialize()

	j := jsiiProxy_CfnIdentity{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpointemail.CfnIdentity",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::PinpointEmail::Identity`.
func NewCfnIdentity_Override(c CfnIdentity, scope constructs.Construct, id *string, props *CfnIdentityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpointemail.CfnIdentity",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIdentity) SetDkimSigningEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"dkimSigningEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnIdentity) SetFeedbackForwardingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"feedbackForwardingEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnIdentity) SetMailFromAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"mailFromAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnIdentity) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnIdentity) SetTags(val *[]*CfnIdentity_TagsProperty) {
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
func CfnIdentity_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnIdentity",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnIdentity_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnIdentity",
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
func CfnIdentity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpointemail.CfnIdentity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdentity_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpointemail.CfnIdentity",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdentity) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnIdentity) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIdentity) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnIdentity) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnIdentity) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnIdentity) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnIdentity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnIdentity) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentity) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentity) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnIdentity) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIdentity) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentity) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIdentity) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A list of attributes that are associated with a MAIL FROM domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailFromAttributesProperty := &mailFromAttributesProperty{
//   	behaviorOnMxFailure: jsii.String("behaviorOnMxFailure"),
//   	mailFromDomain: jsii.String("mailFromDomain"),
//   }
//
type CfnIdentity_MailFromAttributesProperty struct {
	// The action that Amazon Pinpoint to takes if it can't read the required MX record for a custom MAIL FROM domain.
	//
	// When you set this value to `UseDefaultValue` , Amazon Pinpoint uses *amazonses.com* as the MAIL FROM domain. When you set this value to `RejectMessage` , Amazon Pinpoint returns a `MailFromDomainNotVerified` error, and doesn't attempt to deliver the email.
	//
	// These behaviors are taken when the custom MAIL FROM domain configuration is in the `Pending` , `Failed` , and `TemporaryFailure` states.
	BehaviorOnMxFailure *string `field:"optional" json:"behaviorOnMxFailure" yaml:"behaviorOnMxFailure"`
	// The name of a domain that an email identity uses as a custom MAIL FROM domain.
	MailFromDomain *string `field:"optional" json:"mailFromDomain" yaml:"mailFromDomain"`
}

// An object that defines the tags (keys and values) that you want to associate with the identity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsProperty := &tagsProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnIdentity_TagsProperty struct {
	// One part of a key-value pair that defines a tag.
	//
	// The maximum length of a tag key is 128 characters. The minimum length is 1 character.
	//
	// If you specify tags for the identity, then this value is required.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The optional part of a key-value pair that defines a tag.
	//
	// The maximum length of a tag value is 256 characters. The minimum length is 0 characters. If you don’t want a resource to have a specific tag value, don’t specify a value for this parameter. Amazon Pinpoint will set the value to an empty string.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Properties for defining a `CfnIdentity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdentityProps := &cfnIdentityProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dkimSigningEnabled: jsii.Boolean(false),
//   	feedbackForwardingEnabled: jsii.Boolean(false),
//   	mailFromAttributes: &mailFromAttributesProperty{
//   		behaviorOnMxFailure: jsii.String("behaviorOnMxFailure"),
//   		mailFromDomain: jsii.String("mailFromDomain"),
//   	},
//   	tags: []tagsProperty{
//   		&tagsProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIdentityProps struct {
	// The address or domain of the identity, such as *sender@example.com* or *example.co.uk* .
	Name *string `field:"required" json:"name" yaml:"name"`
	// For domain identities, this attribute is used to enable or disable DomainKeys Identified Mail (DKIM) signing for the domain.
	//
	// If the value is `true` , then the messages that you send from the domain are signed using both the DKIM keys for your domain, as well as the keys for the `amazonses.com` domain. If the value is `false` , then the messages that you send are only signed using the DKIM keys for the `amazonses.com` domain.
	DkimSigningEnabled interface{} `field:"optional" json:"dkimSigningEnabled" yaml:"dkimSigningEnabled"`
	// Used to enable or disable feedback forwarding for an identity.
	//
	// This setting determines what happens when an identity is used to send an email that results in a bounce or complaint event.
	//
	// When you enable feedback forwarding, Amazon Pinpoint sends you email notifications when bounce or complaint events occur. Amazon Pinpoint sends this notification to the address that you specified in the Return-Path header of the original email.
	//
	// When you disable feedback forwarding, Amazon Pinpoint sends notifications through other mechanisms, such as by notifying an Amazon SNS topic. You're required to have a method of tracking bounces and complaints. If you haven't set up another mechanism for receiving bounce or complaint notifications, Amazon Pinpoint sends an email notification when these events occur (even if this setting is disabled).
	FeedbackForwardingEnabled interface{} `field:"optional" json:"feedbackForwardingEnabled" yaml:"feedbackForwardingEnabled"`
	// Used to enable or disable the custom Mail-From domain configuration for an email identity.
	MailFromAttributes interface{} `field:"optional" json:"mailFromAttributes" yaml:"mailFromAttributes"`
	// An object that defines the tags (keys and values) that you want to associate with the email identity.
	Tags *[]*CfnIdentity_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

