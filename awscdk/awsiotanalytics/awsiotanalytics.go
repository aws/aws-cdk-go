package awsiotanalytics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::IoTAnalytics::Channel`.
//
// The AWS::IoTAnalytics::Channel resource collects data from an MQTT topic and archives the raw, unprocessed messages before publishing the data to a pipeline. For more information, see [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *AWS IoT Analytics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serviceManagedS3 interface{}
//
//   cfnChannel := awscdk.Aws_iotanalytics.NewCfnChannel(this, jsii.String("MyCfnChannel"), &cfnChannelProps{
//   	channelName: jsii.String("channelName"),
//   	channelStorage: &channelStorageProperty{
//   		customerManagedS3: &customerManagedS3Property{
//   			bucket: jsii.String("bucket"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			keyPrefix: jsii.String("keyPrefix"),
//   		},
//   		serviceManagedS3: serviceManagedS3,
//   	},
//   	retentionPeriod: &retentionPeriodProperty{
//   		numberOfDays: jsii.Number(123),
//   		unlimited: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The name of the channel.
	ChannelName() *string
	SetChannelName(val *string)
	// Where channel data is stored.
	ChannelStorage() interface{}
	SetChannelStorage(val interface{})
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// How long, in days, message data is kept for the channel.
	RetentionPeriod() interface{}
	SetRetentionPeriod(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata which can be used to manage the channel.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
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

// The jsii proxy struct for CfnChannel
type jsiiProxy_CfnChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnChannel) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) ChannelStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) RetentionPeriod() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTAnalytics::Channel`.
func NewCfnChannel(scope constructs.Construct, id *string, props *CfnChannelProps) CfnChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Channel`.
func NewCfnChannel_Override(c CfnChannel, scope constructs.Construct, id *string, props *CfnChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnChannel) SetChannelName(val *string) {
	_jsii_.Set(
		j,
		"channelName",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetChannelStorage(val interface{}) {
	_jsii_.Set(
		j,
		"channelStorage",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetRetentionPeriod(val interface{}) {
	_jsii_.Set(
		j,
		"retentionPeriod",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
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
func CfnChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Where channel data is stored.
//
// You may choose one of `serviceManagedS3` , `customerManagedS3` storage. If not specified, the default is `serviceManagedS3` . This can't be changed after creation of the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serviceManagedS3 interface{}
//
//   channelStorageProperty := &channelStorageProperty{
//   	customerManagedS3: &customerManagedS3Property{
//   		bucket: jsii.String("bucket"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		keyPrefix: jsii.String("keyPrefix"),
//   	},
//   	serviceManagedS3: serviceManagedS3,
//   }
//
type CfnChannel_ChannelStorageProperty struct {
	// Used to store channel data in an S3 bucket that you manage.
	//
	// If customer managed storage is selected, the `retentionPeriod` parameter is ignored. You can't change the choice of S3 storage after the data store is created.
	CustomerManagedS3 interface{} `field:"optional" json:"customerManagedS3" yaml:"customerManagedS3"`
	// Used to store channel data in an S3 bucket managed by AWS IoT Analytics .
	//
	// You can't change the choice of S3 storage after the data store is created.
	ServiceManagedS3 interface{} `field:"optional" json:"serviceManagedS3" yaml:"serviceManagedS3"`
}

// Used to store channel data in an S3 bucket that you manage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerManagedS3Property := &customerManagedS3Property{
//   	bucket: jsii.String("bucket"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	keyPrefix: jsii.String("keyPrefix"),
//   }
//
type CfnChannel_CustomerManagedS3Property struct {
	// The name of the S3 bucket in which channel data is stored.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3 resources.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// (Optional) The prefix used to create the keys of the channel data objects.
	//
	// Each object in an S3 bucket has a key that is its unique identifier within the bucket (each object in a bucket has exactly one key). The prefix must end with a forward slash (/).
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

// How long, in days, message data is kept.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retentionPeriodProperty := &retentionPeriodProperty{
//   	numberOfDays: jsii.Number(123),
//   	unlimited: jsii.Boolean(false),
//   }
//
type CfnChannel_RetentionPeriodProperty struct {
	// The number of days that message data is kept.
	//
	// The `unlimited` parameter must be false.
	NumberOfDays *float64 `field:"optional" json:"numberOfDays" yaml:"numberOfDays"`
	// If true, message data is kept indefinitely.
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}

// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serviceManagedS3 interface{}
//
//   cfnChannelProps := &cfnChannelProps{
//   	channelName: jsii.String("channelName"),
//   	channelStorage: &channelStorageProperty{
//   		customerManagedS3: &customerManagedS3Property{
//   			bucket: jsii.String("bucket"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			keyPrefix: jsii.String("keyPrefix"),
//   		},
//   		serviceManagedS3: serviceManagedS3,
//   	},
//   	retentionPeriod: &retentionPeriodProperty{
//   		numberOfDays: jsii.Number(123),
//   		unlimited: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnChannelProps struct {
	// The name of the channel.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Where channel data is stored.
	ChannelStorage interface{} `field:"optional" json:"channelStorage" yaml:"channelStorage"`
	// How long, in days, message data is kept for the channel.
	RetentionPeriod interface{} `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Metadata which can be used to manage the channel.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTAnalytics::Dataset`.
//
// The AWS::IoTAnalytics::Dataset resource stores data retrieved from a data store by applying a `queryAction` (an SQL query) or a `containerAction` (executing a containerized application). The data set can be populated manually by calling `CreateDatasetContent` or automatically according to a `trigger` you specify. For more information, see [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *AWS IoT Analytics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataset := awscdk.Aws_iotanalytics.NewCfnDataset(this, jsii.String("MyCfnDataset"), &cfnDatasetProps{
//   	actions: []interface{}{
//   		&actionProperty{
//   			actionName: jsii.String("actionName"),
//
//   			// the properties below are optional
//   			containerAction: &containerActionProperty{
//   				executionRoleArn: jsii.String("executionRoleArn"),
//   				image: jsii.String("image"),
//   				resourceConfiguration: &resourceConfigurationProperty{
//   					computeType: jsii.String("computeType"),
//   					volumeSizeInGb: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				variables: []interface{}{
//   					&variableProperty{
//   						variableName: jsii.String("variableName"),
//
//   						// the properties below are optional
//   						datasetContentVersionValue: &datasetContentVersionValueProperty{
//   							datasetName: jsii.String("datasetName"),
//   						},
//   						doubleValue: jsii.Number(123),
//   						outputFileUriValue: &outputFileUriValueProperty{
//   							fileName: jsii.String("fileName"),
//   						},
//   						stringValue: jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   			queryAction: &queryActionProperty{
//   				sqlQuery: jsii.String("sqlQuery"),
//
//   				// the properties below are optional
//   				filters: []interface{}{
//   					&filterProperty{
//   						deltaTime: &deltaTimeProperty{
//   							offsetSeconds: jsii.Number(123),
//   							timeExpression: jsii.String("timeExpression"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	contentDeliveryRules: []interface{}{
//   		&datasetContentDeliveryRuleProperty{
//   			destination: &datasetContentDeliveryRuleDestinationProperty{
//   				iotEventsDestinationConfiguration: &iotEventsDestinationConfigurationProperty{
//   					inputName: jsii.String("inputName"),
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				s3DestinationConfiguration: &s3DestinationConfigurationProperty{
//   					bucket: jsii.String("bucket"),
//   					key: jsii.String("key"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					glueConfiguration: &glueConfigurationProperty{
//   						databaseName: jsii.String("databaseName"),
//   						tableName: jsii.String("tableName"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			entryName: jsii.String("entryName"),
//   		},
//   	},
//   	datasetName: jsii.String("datasetName"),
//   	lateDataRules: []interface{}{
//   		&lateDataRuleProperty{
//   			ruleConfiguration: &lateDataRuleConfigurationProperty{
//   				deltaTimeSessionWindowConfiguration: &deltaTimeSessionWindowConfigurationProperty{
//   					timeoutInMinutes: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			ruleName: jsii.String("ruleName"),
//   		},
//   	},
//   	retentionPeriod: &retentionPeriodProperty{
//   		numberOfDays: jsii.Number(123),
//   		unlimited: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	triggers: []interface{}{
//   		&triggerProperty{
//   			schedule: &scheduleProperty{
//   				scheduleExpression: jsii.String("scheduleExpression"),
//   			},
//   			triggeringDataset: &triggeringDatasetProperty{
//   				datasetName: jsii.String("datasetName"),
//   			},
//   		},
//   	},
//   	versioningConfiguration: &versioningConfigurationProperty{
//   		maxVersions: jsii.Number(123),
//   		unlimited: jsii.Boolean(false),
//   	},
//   })
//
type CfnDataset interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The `DatasetAction` objects that automatically create the dataset contents.
	Actions() interface{}
	SetActions(val interface{})
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// When dataset contents are created they are delivered to destinations specified here.
	ContentDeliveryRules() interface{}
	SetContentDeliveryRules(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the dataset.
	DatasetName() *string
	SetDatasetName(val *string)
	// A list of data rules that send notifications to CloudWatch, when data arrives late.
	//
	// To specify `lateDataRules` , the dataset must use a [DeltaTimer](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_DeltaTime.html) filter.
	LateDataRules() interface{}
	SetLateDataRules(val interface{})
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
	// Optional.
	//
	// How long, in days, message data is kept for the dataset.
	RetentionPeriod() interface{}
	SetRetentionPeriod(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata which can be used to manage the data set.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// The `DatasetTrigger` objects that specify when the dataset is automatically updated.
	Triggers() interface{}
	SetTriggers(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Optional.
	//
	// How many versions of dataset contents are kept. If not specified or set to null, only the latest version plus the latest succeeded version (if they are different) are kept for the time period specified by the `retentionPeriod` parameter. For more information, see [Keeping Multiple Versions of AWS IoT Analytics datasets](https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions) in the *AWS IoT Analytics User Guide* .
	VersioningConfiguration() interface{}
	SetVersioningConfiguration(val interface{})
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

// The jsii proxy struct for CfnDataset
type jsiiProxy_CfnDataset struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataset) Actions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) ContentDeliveryRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentDeliveryRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) DatasetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) LateDataRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lateDataRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) RetentionPeriod() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Triggers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) VersioningConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioningConfiguration",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTAnalytics::Dataset`.
func NewCfnDataset(scope constructs.Construct, id *string, props *CfnDatasetProps) CfnDataset {
	_init_.Initialize()

	j := jsiiProxy_CfnDataset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Dataset`.
func NewCfnDataset_Override(c CfnDataset, scope constructs.Construct, id *string, props *CfnDatasetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataset) SetActions(val interface{}) {
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetContentDeliveryRules(val interface{}) {
	_jsii_.Set(
		j,
		"contentDeliveryRules",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetDatasetName(val *string) {
	_jsii_.Set(
		j,
		"datasetName",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetLateDataRules(val interface{}) {
	_jsii_.Set(
		j,
		"lateDataRules",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetRetentionPeriod(val interface{}) {
	_jsii_.Set(
		j,
		"retentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetTriggers(val interface{}) {
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetVersioningConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"versioningConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataset_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataset_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
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
func CfnDataset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataset_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataset) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataset) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataset) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataset) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataset) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataset) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataset) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataset) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataset) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataset) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information needed to run the "containerAction" to produce data set contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	actionName: jsii.String("actionName"),
//
//   	// the properties below are optional
//   	containerAction: &containerActionProperty{
//   		executionRoleArn: jsii.String("executionRoleArn"),
//   		image: jsii.String("image"),
//   		resourceConfiguration: &resourceConfigurationProperty{
//   			computeType: jsii.String("computeType"),
//   			volumeSizeInGb: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		variables: []interface{}{
//   			&variableProperty{
//   				variableName: jsii.String("variableName"),
//
//   				// the properties below are optional
//   				datasetContentVersionValue: &datasetContentVersionValueProperty{
//   					datasetName: jsii.String("datasetName"),
//   				},
//   				doubleValue: jsii.Number(123),
//   				outputFileUriValue: &outputFileUriValueProperty{
//   					fileName: jsii.String("fileName"),
//   				},
//   				stringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	queryAction: &queryActionProperty{
//   		sqlQuery: jsii.String("sqlQuery"),
//
//   		// the properties below are optional
//   		filters: []interface{}{
//   			&filterProperty{
//   				deltaTime: &deltaTimeProperty{
//   					offsetSeconds: jsii.Number(123),
//   					timeExpression: jsii.String("timeExpression"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDataset_ActionProperty struct {
	// The name of the data set action by which data set contents are automatically created.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// Information which allows the system to run a containerized application in order to create the data set contents.
	//
	// The application must be in a Docker container along with any needed support libraries.
	ContainerAction interface{} `field:"optional" json:"containerAction" yaml:"containerAction"`
	// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
	QueryAction interface{} `field:"optional" json:"queryAction" yaml:"queryAction"`
}

// Information needed to run the "containerAction" to produce data set contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerActionProperty := &containerActionProperty{
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	image: jsii.String("image"),
//   	resourceConfiguration: &resourceConfigurationProperty{
//   		computeType: jsii.String("computeType"),
//   		volumeSizeInGb: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	variables: []interface{}{
//   		&variableProperty{
//   			variableName: jsii.String("variableName"),
//
//   			// the properties below are optional
//   			datasetContentVersionValue: &datasetContentVersionValueProperty{
//   				datasetName: jsii.String("datasetName"),
//   			},
//   			doubleValue: jsii.Number(123),
//   			outputFileUriValue: &outputFileUriValueProperty{
//   				fileName: jsii.String("fileName"),
//   			},
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   }
//
type CfnDataset_ContainerActionProperty struct {
	// The ARN of the role which gives permission to the system to access needed resources in order to run the "containerAction".
	//
	// This includes, at minimum, permission to retrieve the data set contents which are the input to the containerized application.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The ARN of the Docker container stored in your account.
	//
	// The Docker container contains an application and needed support libraries and is used to generate data set contents.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Configuration of the resource which executes the "containerAction".
	ResourceConfiguration interface{} `field:"required" json:"resourceConfiguration" yaml:"resourceConfiguration"`
	// The values of variables used within the context of the execution of the containerized application (basically, parameters passed to the application).
	//
	// Each variable must have a name and a value given by one of "stringValue", "datasetContentVersionValue", or "outputFileUriValue".
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

// The destination to which dataset contents are delivered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetContentDeliveryRuleDestinationProperty := &datasetContentDeliveryRuleDestinationProperty{
//   	iotEventsDestinationConfiguration: &iotEventsDestinationConfigurationProperty{
//   		inputName: jsii.String("inputName"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	s3DestinationConfiguration: &s3DestinationConfigurationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		glueConfiguration: &glueConfigurationProperty{
//   			databaseName: jsii.String("databaseName"),
//   			tableName: jsii.String("tableName"),
//   		},
//   	},
//   }
//
type CfnDataset_DatasetContentDeliveryRuleDestinationProperty struct {
	// Configuration information for delivery of dataset contents to AWS IoT Events .
	IotEventsDestinationConfiguration interface{} `field:"optional" json:"iotEventsDestinationConfiguration" yaml:"iotEventsDestinationConfiguration"`
	// Configuration information for delivery of dataset contents to Amazon S3.
	S3DestinationConfiguration interface{} `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
}

// When dataset contents are created, they are delivered to destination specified here.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetContentDeliveryRuleProperty := &datasetContentDeliveryRuleProperty{
//   	destination: &datasetContentDeliveryRuleDestinationProperty{
//   		iotEventsDestinationConfiguration: &iotEventsDestinationConfigurationProperty{
//   			inputName: jsii.String("inputName"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		s3DestinationConfiguration: &s3DestinationConfigurationProperty{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			glueConfiguration: &glueConfigurationProperty{
//   				databaseName: jsii.String("databaseName"),
//   				tableName: jsii.String("tableName"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	entryName: jsii.String("entryName"),
//   }
//
type CfnDataset_DatasetContentDeliveryRuleProperty struct {
	// The destination to which dataset contents are delivered.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// The name of the dataset content delivery rules entry.
	EntryName *string `field:"optional" json:"entryName" yaml:"entryName"`
}

// The dataset whose latest contents are used as input to the notebook or application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetContentVersionValueProperty := &datasetContentVersionValueProperty{
//   	datasetName: jsii.String("datasetName"),
//   }
//
type CfnDataset_DatasetContentVersionValueProperty struct {
	// The name of the dataset whose latest contents are used as input to the notebook or application.
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
}

// Used to limit data to that which has arrived since the last execution of the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deltaTimeProperty := &deltaTimeProperty{
//   	offsetSeconds: jsii.Number(123),
//   	timeExpression: jsii.String("timeExpression"),
//   }
//
type CfnDataset_DeltaTimeProperty struct {
	// The number of seconds of estimated in-flight lag time of message data.
	//
	// When you create dataset contents using message data from a specified timeframe, some message data might still be in flight when processing begins, and so do not arrive in time to be processed. Use this field to make allowances for the in flight time of your message data, so that data not processed from a previous timeframe is included with the next timeframe. Otherwise, missed message data would be excluded from processing during the next timeframe too, because its timestamp places it within the previous timeframe.
	OffsetSeconds *float64 `field:"required" json:"offsetSeconds" yaml:"offsetSeconds"`
	// An expression by which the time of the message data might be determined.
	//
	// This can be the name of a timestamp field or a SQL expression that is used to derive the time the message data was generated.
	TimeExpression *string `field:"required" json:"timeExpression" yaml:"timeExpression"`
}

// A structure that contains the configuration information of a delta time session window.
//
// [`DeltaTime`](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_DeltaTime.html) specifies a time interval. You can use `DeltaTime` to create dataset contents with data that has arrived in the data store since the last execution. For an example of `DeltaTime` , see [Creating a SQL dataset with a delta window (CLI)](https://docs.aws.amazon.com/iotanalytics/latest/userguide/automate-create-dataset.html#automate-example6) in the *AWS IoT Analytics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deltaTimeSessionWindowConfigurationProperty := &deltaTimeSessionWindowConfigurationProperty{
//   	timeoutInMinutes: jsii.Number(123),
//   }
//
type CfnDataset_DeltaTimeSessionWindowConfigurationProperty struct {
	// A time interval.
	//
	// You can use `timeoutInMinutes` so that AWS IoT Analytics can batch up late data notifications that have been generated since the last execution. AWS IoT Analytics sends one batch of notifications to Amazon CloudWatch Events at one time.
	//
	// For more information about how to write a timestamp expression, see [Date and Time Functions and Operators](https://docs.aws.amazon.com/https://prestodb.io/docs/0.172/functions/datetime.html) , in the *Presto 0.172 Documentation* .
	TimeoutInMinutes *float64 `field:"required" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

// Information which is used to filter message data, to segregate it according to the time frame in which it arrives.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &filterProperty{
//   	deltaTime: &deltaTimeProperty{
//   		offsetSeconds: jsii.Number(123),
//   		timeExpression: jsii.String("timeExpression"),
//   	},
//   }
//
type CfnDataset_FilterProperty struct {
	// Used to limit data to that which has arrived since the last execution of the action.
	DeltaTime interface{} `field:"optional" json:"deltaTime" yaml:"deltaTime"`
}

// Configuration information for coordination with AWS Glue , a fully managed extract, transform and load (ETL) service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   glueConfigurationProperty := &glueConfigurationProperty{
//   	databaseName: jsii.String("databaseName"),
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnDataset_GlueConfigurationProperty struct {
	// The name of the database in your AWS Glue Data Catalog in which the table is located.
	//
	// An AWS Glue Data Catalog database contains metadata tables.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The name of the table in your AWS Glue Data Catalog that is used to perform the ETL operations.
	//
	// An AWS Glue Data Catalog table contains partitioned data and descriptions of data sources and targets.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

// Configuration information for delivery of dataset contents to AWS IoT Events .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotEventsDestinationConfigurationProperty := &iotEventsDestinationConfigurationProperty{
//   	inputName: jsii.String("inputName"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnDataset_IotEventsDestinationConfigurationProperty struct {
	// The name of the AWS IoT Events input to which dataset contents are delivered.
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
	// The ARN of the role that grants AWS IoT Analytics permission to deliver dataset contents to an AWS IoT Events input.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

// The information needed to configure a delta time session window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lateDataRuleConfigurationProperty := &lateDataRuleConfigurationProperty{
//   	deltaTimeSessionWindowConfiguration: &deltaTimeSessionWindowConfigurationProperty{
//   		timeoutInMinutes: jsii.Number(123),
//   	},
//   }
//
type CfnDataset_LateDataRuleConfigurationProperty struct {
	// The information needed to configure a delta time session window.
	DeltaTimeSessionWindowConfiguration interface{} `field:"optional" json:"deltaTimeSessionWindowConfiguration" yaml:"deltaTimeSessionWindowConfiguration"`
}

// A structure that contains the name and configuration information of a late data rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lateDataRuleProperty := &lateDataRuleProperty{
//   	ruleConfiguration: &lateDataRuleConfigurationProperty{
//   		deltaTimeSessionWindowConfiguration: &deltaTimeSessionWindowConfigurationProperty{
//   			timeoutInMinutes: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ruleName: jsii.String("ruleName"),
//   }
//
type CfnDataset_LateDataRuleProperty struct {
	// The information needed to configure the late data rule.
	RuleConfiguration interface{} `field:"required" json:"ruleConfiguration" yaml:"ruleConfiguration"`
	// The name of the late data rule.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

// The value of the variable as a structure that specifies an output file URI.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputFileUriValueProperty := &outputFileUriValueProperty{
//   	fileName: jsii.String("fileName"),
//   }
//
type CfnDataset_OutputFileUriValueProperty struct {
	// The URI of the location where dataset contents are stored, usually the URI of a file in an S3 bucket.
	FileName *string `field:"required" json:"fileName" yaml:"fileName"`
}

// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryActionProperty := &queryActionProperty{
//   	sqlQuery: jsii.String("sqlQuery"),
//
//   	// the properties below are optional
//   	filters: []interface{}{
//   		&filterProperty{
//   			deltaTime: &deltaTimeProperty{
//   				offsetSeconds: jsii.Number(123),
//   				timeExpression: jsii.String("timeExpression"),
//   			},
//   		},
//   	},
//   }
//
type CfnDataset_QueryActionProperty struct {
	// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
	SqlQuery *string `field:"required" json:"sqlQuery" yaml:"sqlQuery"`
	// Pre-filters applied to message data.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

// The configuration of the resource used to execute the `containerAction` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceConfigurationProperty := &resourceConfigurationProperty{
//   	computeType: jsii.String("computeType"),
//   	volumeSizeInGb: jsii.Number(123),
//   }
//
type CfnDataset_ResourceConfigurationProperty struct {
	// The type of the compute resource used to execute the `containerAction` .
	//
	// Possible values are: `ACU_1` (vCPU=4, memory=16 GiB) or `ACU_2` (vCPU=8, memory=32 GiB).
	ComputeType *string `field:"required" json:"computeType" yaml:"computeType"`
	// The size, in GB, of the persistent storage available to the resource instance used to execute the `containerAction` (min: 1, max: 50).
	VolumeSizeInGb *float64 `field:"required" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

// How long, in days, message data is kept.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retentionPeriodProperty := &retentionPeriodProperty{
//   	numberOfDays: jsii.Number(123),
//   	unlimited: jsii.Boolean(false),
//   }
//
type CfnDataset_RetentionPeriodProperty struct {
	// The number of days that message data is kept.
	//
	// The `unlimited` parameter must be false.
	NumberOfDays *float64 `field:"optional" json:"numberOfDays" yaml:"numberOfDays"`
	// If true, message data is kept indefinitely.
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}

// Configuration information for delivery of dataset contents to Amazon Simple Storage Service (Amazon S3).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationConfigurationProperty := &s3DestinationConfigurationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	glueConfiguration: &glueConfigurationProperty{
//   		databaseName: jsii.String("databaseName"),
//   		tableName: jsii.String("tableName"),
//   	},
//   }
//
type CfnDataset_S3DestinationConfigurationProperty struct {
	// The name of the S3 bucket to which dataset contents are delivered.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The key of the dataset contents object in an S3 bucket.
	//
	// Each object has a key that is a unique identifier. Each object has exactly one key.
	//
	// You can create a unique key with the following options:
	//
	// - Use `!{iotanalytics:scheduleTime}` to insert the time of a scheduled SQL query run.
	// - Use `!{iotanalytics:versionId}` to insert a unique hash that identifies a dataset content.
	// - Use `!{iotanalytics:creationTime}` to insert the creation time of a dataset content.
	//
	// The following example creates a unique key for a CSV file: `dataset/mydataset/!{iotanalytics:scheduleTime}/!{iotanalytics:versionId}.csv`
	//
	// > If you don't use `!{iotanalytics:versionId}` to specify the key, you might get duplicate keys. For example, you might have two dataset contents with the same `scheduleTime` but different `versionId` s. This means that one dataset content overwrites the other.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3 and AWS Glue resources.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Configuration information for coordination with AWS Glue , a fully managed extract, transform and load (ETL) service.
	GlueConfiguration interface{} `field:"optional" json:"glueConfiguration" yaml:"glueConfiguration"`
}

// The schedule for when to trigger an update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleProperty := &scheduleProperty{
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   }
//
type CfnDataset_ScheduleProperty struct {
	// The expression that defines when to trigger an update.
	//
	// For more information, see [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) in the Amazon CloudWatch documentation.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
}

// The "DatasetTrigger" that specifies when the data set is automatically updated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerProperty := &triggerProperty{
//   	schedule: &scheduleProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	triggeringDataset: &triggeringDatasetProperty{
//   		datasetName: jsii.String("datasetName"),
//   	},
//   }
//
type CfnDataset_TriggerProperty struct {
	// The "Schedule" when the trigger is initiated.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Information about the data set whose content generation triggers the new data set content generation.
	TriggeringDataset interface{} `field:"optional" json:"triggeringDataset" yaml:"triggeringDataset"`
}

// Information about the dataset whose content generation triggers the new dataset content generation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggeringDatasetProperty := &triggeringDatasetProperty{
//   	datasetName: jsii.String("datasetName"),
//   }
//
type CfnDataset_TriggeringDatasetProperty struct {
	// The name of the data set whose content generation triggers the new data set content generation.
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
}

// An instance of a variable to be passed to the `containerAction` execution.
//
// Each variable must have a name and a value given by one of `stringValue` , `datasetContentVersionValue` , or `outputFileUriValue` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variableProperty := &variableProperty{
//   	variableName: jsii.String("variableName"),
//
//   	// the properties below are optional
//   	datasetContentVersionValue: &datasetContentVersionValueProperty{
//   		datasetName: jsii.String("datasetName"),
//   	},
//   	doubleValue: jsii.Number(123),
//   	outputFileUriValue: &outputFileUriValueProperty{
//   		fileName: jsii.String("fileName"),
//   	},
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnDataset_VariableProperty struct {
	// The name of the variable.
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
	// The value of the variable as a structure that specifies a dataset content version.
	DatasetContentVersionValue interface{} `field:"optional" json:"datasetContentVersionValue" yaml:"datasetContentVersionValue"`
	// The value of the variable as a double (numeric).
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The value of the variable as a structure that specifies an output file URI.
	OutputFileUriValue interface{} `field:"optional" json:"outputFileUriValue" yaml:"outputFileUriValue"`
	// The value of the variable as a string.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

// Information about the versioning of dataset contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   versioningConfigurationProperty := &versioningConfigurationProperty{
//   	maxVersions: jsii.Number(123),
//   	unlimited: jsii.Boolean(false),
//   }
//
type CfnDataset_VersioningConfigurationProperty struct {
	// How many versions of dataset contents are kept.
	//
	// The `unlimited` parameter must be `false` .
	MaxVersions *float64 `field:"optional" json:"maxVersions" yaml:"maxVersions"`
	// If true, unlimited versions of dataset contents are kept.
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}

// Properties for defining a `CfnDataset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetProps := &cfnDatasetProps{
//   	actions: []interface{}{
//   		&actionProperty{
//   			actionName: jsii.String("actionName"),
//
//   			// the properties below are optional
//   			containerAction: &containerActionProperty{
//   				executionRoleArn: jsii.String("executionRoleArn"),
//   				image: jsii.String("image"),
//   				resourceConfiguration: &resourceConfigurationProperty{
//   					computeType: jsii.String("computeType"),
//   					volumeSizeInGb: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				variables: []interface{}{
//   					&variableProperty{
//   						variableName: jsii.String("variableName"),
//
//   						// the properties below are optional
//   						datasetContentVersionValue: &datasetContentVersionValueProperty{
//   							datasetName: jsii.String("datasetName"),
//   						},
//   						doubleValue: jsii.Number(123),
//   						outputFileUriValue: &outputFileUriValueProperty{
//   							fileName: jsii.String("fileName"),
//   						},
//   						stringValue: jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   			queryAction: &queryActionProperty{
//   				sqlQuery: jsii.String("sqlQuery"),
//
//   				// the properties below are optional
//   				filters: []interface{}{
//   					&filterProperty{
//   						deltaTime: &deltaTimeProperty{
//   							offsetSeconds: jsii.Number(123),
//   							timeExpression: jsii.String("timeExpression"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	contentDeliveryRules: []interface{}{
//   		&datasetContentDeliveryRuleProperty{
//   			destination: &datasetContentDeliveryRuleDestinationProperty{
//   				iotEventsDestinationConfiguration: &iotEventsDestinationConfigurationProperty{
//   					inputName: jsii.String("inputName"),
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				s3DestinationConfiguration: &s3DestinationConfigurationProperty{
//   					bucket: jsii.String("bucket"),
//   					key: jsii.String("key"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					glueConfiguration: &glueConfigurationProperty{
//   						databaseName: jsii.String("databaseName"),
//   						tableName: jsii.String("tableName"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			entryName: jsii.String("entryName"),
//   		},
//   	},
//   	datasetName: jsii.String("datasetName"),
//   	lateDataRules: []interface{}{
//   		&lateDataRuleProperty{
//   			ruleConfiguration: &lateDataRuleConfigurationProperty{
//   				deltaTimeSessionWindowConfiguration: &deltaTimeSessionWindowConfigurationProperty{
//   					timeoutInMinutes: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			ruleName: jsii.String("ruleName"),
//   		},
//   	},
//   	retentionPeriod: &retentionPeriodProperty{
//   		numberOfDays: jsii.Number(123),
//   		unlimited: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	triggers: []interface{}{
//   		&triggerProperty{
//   			schedule: &scheduleProperty{
//   				scheduleExpression: jsii.String("scheduleExpression"),
//   			},
//   			triggeringDataset: &triggeringDatasetProperty{
//   				datasetName: jsii.String("datasetName"),
//   			},
//   		},
//   	},
//   	versioningConfiguration: &versioningConfigurationProperty{
//   		maxVersions: jsii.Number(123),
//   		unlimited: jsii.Boolean(false),
//   	},
//   }
//
type CfnDatasetProps struct {
	// The `DatasetAction` objects that automatically create the dataset contents.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// When dataset contents are created they are delivered to destinations specified here.
	ContentDeliveryRules interface{} `field:"optional" json:"contentDeliveryRules" yaml:"contentDeliveryRules"`
	// The name of the dataset.
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// A list of data rules that send notifications to CloudWatch, when data arrives late.
	//
	// To specify `lateDataRules` , the dataset must use a [DeltaTimer](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_DeltaTime.html) filter.
	LateDataRules interface{} `field:"optional" json:"lateDataRules" yaml:"lateDataRules"`
	// Optional.
	//
	// How long, in days, message data is kept for the dataset.
	RetentionPeriod interface{} `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Metadata which can be used to manage the data set.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The `DatasetTrigger` objects that specify when the dataset is automatically updated.
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
	// Optional.
	//
	// How many versions of dataset contents are kept. If not specified or set to null, only the latest version plus the latest succeeded version (if they are different) are kept for the time period specified by the `retentionPeriod` parameter. For more information, see [Keeping Multiple Versions of AWS IoT Analytics datasets](https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions) in the *AWS IoT Analytics User Guide* .
	VersioningConfiguration interface{} `field:"optional" json:"versioningConfiguration" yaml:"versioningConfiguration"`
}

// A CloudFormation `AWS::IoTAnalytics::Datastore`.
//
// AWS::IoTAnalytics::Datastore resource is a repository for messages. For more information, see [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *AWS IoT Analytics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var jsonConfiguration interface{}
//   var serviceManagedS3 interface{}
//
//   cfnDatastore := awscdk.Aws_iotanalytics.NewCfnDatastore(this, jsii.String("MyCfnDatastore"), &cfnDatastoreProps{
//   	datastoreName: jsii.String("datastoreName"),
//   	datastorePartitions: &datastorePartitionsProperty{
//   		partitions: []interface{}{
//   			&datastorePartitionProperty{
//   				partition: &partitionProperty{
//   					attributeName: jsii.String("attributeName"),
//   				},
//   				timestampPartition: &timestampPartitionProperty{
//   					attributeName: jsii.String("attributeName"),
//
//   					// the properties below are optional
//   					timestampFormat: jsii.String("timestampFormat"),
//   				},
//   			},
//   		},
//   	},
//   	datastoreStorage: &datastoreStorageProperty{
//   		customerManagedS3: &customerManagedS3Property{
//   			bucket: jsii.String("bucket"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			keyPrefix: jsii.String("keyPrefix"),
//   		},
//   		iotSiteWiseMultiLayerStorage: &iotSiteWiseMultiLayerStorageProperty{
//   			customerManagedS3Storage: &customerManagedS3StorageProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				keyPrefix: jsii.String("keyPrefix"),
//   			},
//   		},
//   		serviceManagedS3: serviceManagedS3,
//   	},
//   	fileFormatConfiguration: &fileFormatConfigurationProperty{
//   		jsonConfiguration: jsonConfiguration,
//   		parquetConfiguration: &parquetConfigurationProperty{
//   			schemaDefinition: &schemaDefinitionProperty{
//   				columns: []interface{}{
//   					&columnProperty{
//   						name: jsii.String("name"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	retentionPeriod: &retentionPeriodProperty{
//   		numberOfDays: jsii.Number(123),
//   		unlimited: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDatastore interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the data store.
	DatastoreName() *string
	SetDatastoreName(val *string)
	// Information about the partition dimensions in a data store.
	DatastorePartitions() interface{}
	SetDatastorePartitions(val interface{})
	// Where data store data is stored.
	DatastoreStorage() interface{}
	SetDatastoreStorage(val interface{})
	// Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
	//
	// The default file format is JSON. You can specify only one format.
	//
	// You can't change the file format after you create the data store.
	FileFormatConfiguration() interface{}
	SetFileFormatConfiguration(val interface{})
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
	// How long, in days, message data is kept for the data store.
	//
	// When `customerManagedS3` storage is selected, this parameter is ignored.
	RetentionPeriod() interface{}
	SetRetentionPeriod(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata which can be used to manage the data store.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
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

// The jsii proxy struct for CfnDatastore
type jsiiProxy_CfnDatastore struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDatastore) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) DatastoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) DatastorePartitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datastorePartitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) DatastoreStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datastoreStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) FileFormatConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileFormatConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) RetentionPeriod() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTAnalytics::Datastore`.
func NewCfnDatastore(scope constructs.Construct, id *string, props *CfnDatastoreProps) CfnDatastore {
	_init_.Initialize()

	j := jsiiProxy_CfnDatastore{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Datastore`.
func NewCfnDatastore_Override(c CfnDatastore, scope constructs.Construct, id *string, props *CfnDatastoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDatastore) SetDatastoreName(val *string) {
	_jsii_.Set(
		j,
		"datastoreName",
		val,
	)
}

func (j *jsiiProxy_CfnDatastore) SetDatastorePartitions(val interface{}) {
	_jsii_.Set(
		j,
		"datastorePartitions",
		val,
	)
}

func (j *jsiiProxy_CfnDatastore) SetDatastoreStorage(val interface{}) {
	_jsii_.Set(
		j,
		"datastoreStorage",
		val,
	)
}

func (j *jsiiProxy_CfnDatastore) SetFileFormatConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"fileFormatConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDatastore) SetRetentionPeriod(val interface{}) {
	_jsii_.Set(
		j,
		"retentionPeriod",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDatastore_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDatastore_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
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
func CfnDatastore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatastore_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDatastore) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDatastore) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDatastore) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDatastore) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDatastore) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDatastore) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDatastore) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDatastore) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatastore) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatastore) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDatastore) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDatastore) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatastore) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatastore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatastore) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains information about a column that stores your data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnProperty := &columnProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//   }
//
type CfnDatastore_ColumnProperty struct {
	// The name of the column.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of data.
	//
	// For more information about the supported data types, see [Common data types](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-common.html) in the *AWS Glue Developer Guide* .
	Type *string `field:"required" json:"type" yaml:"type"`
}

// S3-customer-managed;
//
// When you choose customer-managed storage, the `retentionPeriod` parameter is ignored. You can't change the choice of Amazon S3 storage after your data store is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerManagedS3Property := &customerManagedS3Property{
//   	bucket: jsii.String("bucket"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	keyPrefix: jsii.String("keyPrefix"),
//   }
//
type CfnDatastore_CustomerManagedS3Property struct {
	// The name of the Amazon S3 bucket where your data is stored.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3 resources.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// (Optional) The prefix used to create the keys of the data store data objects.
	//
	// Each object in an Amazon S3 bucket has a key that is its unique identifier in the bucket. Each object in a bucket has exactly one key. The prefix must end with a forward slash (/).
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

// Amazon S3 -customer-managed;
//
// When you choose customer-managed storage, the `retentionPeriod` parameter is ignored. You can't change the choice of Amazon S3 storage after your data store is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerManagedS3StorageProperty := &customerManagedS3StorageProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	keyPrefix: jsii.String("keyPrefix"),
//   }
//
type CfnDatastore_CustomerManagedS3StorageProperty struct {
	// The name of the Amazon S3 bucket where your data is stored.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// (Optional) The prefix used to create the keys of the data store data objects.
	//
	// Each object in an Amazon S3 bucket has a key that is its unique identifier in the bucket. Each object in a bucket has exactly one key. The prefix must end with a forward slash (/).
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

// A single dimension to partition a data store.
//
// The dimension must be an `AttributePartition` or a `TimestampPartition` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datastorePartitionProperty := &datastorePartitionProperty{
//   	partition: &partitionProperty{
//   		attributeName: jsii.String("attributeName"),
//   	},
//   	timestampPartition: &timestampPartitionProperty{
//   		attributeName: jsii.String("attributeName"),
//
//   		// the properties below are optional
//   		timestampFormat: jsii.String("timestampFormat"),
//   	},
//   }
//
type CfnDatastore_DatastorePartitionProperty struct {
	// A partition dimension defined by an attribute.
	Partition interface{} `field:"optional" json:"partition" yaml:"partition"`
	// A partition dimension defined by a timestamp attribute.
	TimestampPartition interface{} `field:"optional" json:"timestampPartition" yaml:"timestampPartition"`
}

// Information about the partition dimensions in a data store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datastorePartitionsProperty := &datastorePartitionsProperty{
//   	partitions: []interface{}{
//   		&datastorePartitionProperty{
//   			partition: &partitionProperty{
//   				attributeName: jsii.String("attributeName"),
//   			},
//   			timestampPartition: &timestampPartitionProperty{
//   				attributeName: jsii.String("attributeName"),
//
//   				// the properties below are optional
//   				timestampFormat: jsii.String("timestampFormat"),
//   			},
//   		},
//   	},
//   }
//
type CfnDatastore_DatastorePartitionsProperty struct {
	// A list of partition dimensions in a data store.
	Partitions interface{} `field:"optional" json:"partitions" yaml:"partitions"`
}

// Where data store data is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serviceManagedS3 interface{}
//
//   datastoreStorageProperty := &datastoreStorageProperty{
//   	customerManagedS3: &customerManagedS3Property{
//   		bucket: jsii.String("bucket"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		keyPrefix: jsii.String("keyPrefix"),
//   	},
//   	iotSiteWiseMultiLayerStorage: &iotSiteWiseMultiLayerStorageProperty{
//   		customerManagedS3Storage: &customerManagedS3StorageProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			keyPrefix: jsii.String("keyPrefix"),
//   		},
//   	},
//   	serviceManagedS3: serviceManagedS3,
//   }
//
type CfnDatastore_DatastoreStorageProperty struct {
	// Use this to store data store data in an S3 bucket that you manage.
	//
	// The choice of service-managed or customer-managed S3 storage cannot be changed after creation of the data store.
	CustomerManagedS3 interface{} `field:"optional" json:"customerManagedS3" yaml:"customerManagedS3"`
	// Use this to store data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
	//
	// You can't change the choice of Amazon S3 storage after your data store is created.
	IotSiteWiseMultiLayerStorage interface{} `field:"optional" json:"iotSiteWiseMultiLayerStorage" yaml:"iotSiteWiseMultiLayerStorage"`
	// Use this to store data store data in an S3 bucket managed by the AWS IoT Analytics service.
	//
	// The choice of service-managed or customer-managed S3 storage cannot be changed after creation of the data store.
	ServiceManagedS3 interface{} `field:"optional" json:"serviceManagedS3" yaml:"serviceManagedS3"`
}

// Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
//
// The default file format is JSON. You can specify only one format.
//
// You can't change the file format after you create the data store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var jsonConfiguration interface{}
//
//   fileFormatConfigurationProperty := &fileFormatConfigurationProperty{
//   	jsonConfiguration: jsonConfiguration,
//   	parquetConfiguration: &parquetConfigurationProperty{
//   		schemaDefinition: &schemaDefinitionProperty{
//   			columns: []interface{}{
//   				&columnProperty{
//   					name: jsii.String("name"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDatastore_FileFormatConfigurationProperty struct {
	// Contains the configuration information of the JSON format.
	JsonConfiguration interface{} `field:"optional" json:"jsonConfiguration" yaml:"jsonConfiguration"`
	// Contains the configuration information of the Parquet format.
	ParquetConfiguration interface{} `field:"optional" json:"parquetConfiguration" yaml:"parquetConfiguration"`
}

// Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
//
// You can't change the choice of Amazon S3 storage after your data store is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotSiteWiseMultiLayerStorageProperty := &iotSiteWiseMultiLayerStorageProperty{
//   	customerManagedS3Storage: &customerManagedS3StorageProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		keyPrefix: jsii.String("keyPrefix"),
//   	},
//   }
//
type CfnDatastore_IotSiteWiseMultiLayerStorageProperty struct {
	// Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
	CustomerManagedS3Storage interface{} `field:"optional" json:"customerManagedS3Storage" yaml:"customerManagedS3Storage"`
}

// Contains the configuration information of the Parquet format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parquetConfigurationProperty := &parquetConfigurationProperty{
//   	schemaDefinition: &schemaDefinitionProperty{
//   		columns: []interface{}{
//   			&columnProperty{
//   				name: jsii.String("name"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
type CfnDatastore_ParquetConfigurationProperty struct {
	// Information needed to define a schema.
	SchemaDefinition interface{} `field:"optional" json:"schemaDefinition" yaml:"schemaDefinition"`
}

// A single dimension to partition a data store.
//
// The dimension must be an `AttributePartition` or a `TimestampPartition` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partitionProperty := &partitionProperty{
//   	attributeName: jsii.String("attributeName"),
//   }
//
type CfnDatastore_PartitionProperty struct {
	// The name of the attribute that defines a partition dimension.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
}

// How long, in days, message data is kept.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retentionPeriodProperty := &retentionPeriodProperty{
//   	numberOfDays: jsii.Number(123),
//   	unlimited: jsii.Boolean(false),
//   }
//
type CfnDatastore_RetentionPeriodProperty struct {
	// The number of days that message data is kept.
	//
	// The `unlimited` parameter must be false.
	NumberOfDays *float64 `field:"optional" json:"numberOfDays" yaml:"numberOfDays"`
	// If true, message data is kept indefinitely.
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}

// Information needed to define a schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaDefinitionProperty := &schemaDefinitionProperty{
//   	columns: []interface{}{
//   		&columnProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnDatastore_SchemaDefinitionProperty struct {
	// Specifies one or more columns that store your data.
	//
	// Each schema can have up to 100 columns. Each column can have up to 100 nested types.
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
}

// A partition dimension defined by a timestamp attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timestampPartitionProperty := &timestampPartitionProperty{
//   	attributeName: jsii.String("attributeName"),
//
//   	// the properties below are optional
//   	timestampFormat: jsii.String("timestampFormat"),
//   }
//
type CfnDatastore_TimestampPartitionProperty struct {
	// The attribute name of the partition defined by a timestamp.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// The timestamp format of a partition defined by a timestamp.
	//
	// The default format is seconds since epoch (January 1, 1970 at midnight UTC time).
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
}

// Properties for defining a `CfnDatastore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var jsonConfiguration interface{}
//   var serviceManagedS3 interface{}
//
//   cfnDatastoreProps := &cfnDatastoreProps{
//   	datastoreName: jsii.String("datastoreName"),
//   	datastorePartitions: &datastorePartitionsProperty{
//   		partitions: []interface{}{
//   			&datastorePartitionProperty{
//   				partition: &partitionProperty{
//   					attributeName: jsii.String("attributeName"),
//   				},
//   				timestampPartition: &timestampPartitionProperty{
//   					attributeName: jsii.String("attributeName"),
//
//   					// the properties below are optional
//   					timestampFormat: jsii.String("timestampFormat"),
//   				},
//   			},
//   		},
//   	},
//   	datastoreStorage: &datastoreStorageProperty{
//   		customerManagedS3: &customerManagedS3Property{
//   			bucket: jsii.String("bucket"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			keyPrefix: jsii.String("keyPrefix"),
//   		},
//   		iotSiteWiseMultiLayerStorage: &iotSiteWiseMultiLayerStorageProperty{
//   			customerManagedS3Storage: &customerManagedS3StorageProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				keyPrefix: jsii.String("keyPrefix"),
//   			},
//   		},
//   		serviceManagedS3: serviceManagedS3,
//   	},
//   	fileFormatConfiguration: &fileFormatConfigurationProperty{
//   		jsonConfiguration: jsonConfiguration,
//   		parquetConfiguration: &parquetConfigurationProperty{
//   			schemaDefinition: &schemaDefinitionProperty{
//   				columns: []interface{}{
//   					&columnProperty{
//   						name: jsii.String("name"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	retentionPeriod: &retentionPeriodProperty{
//   		numberOfDays: jsii.Number(123),
//   		unlimited: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDatastoreProps struct {
	// The name of the data store.
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// Information about the partition dimensions in a data store.
	DatastorePartitions interface{} `field:"optional" json:"datastorePartitions" yaml:"datastorePartitions"`
	// Where data store data is stored.
	DatastoreStorage interface{} `field:"optional" json:"datastoreStorage" yaml:"datastoreStorage"`
	// Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
	//
	// The default file format is JSON. You can specify only one format.
	//
	// You can't change the file format after you create the data store.
	FileFormatConfiguration interface{} `field:"optional" json:"fileFormatConfiguration" yaml:"fileFormatConfiguration"`
	// How long, in days, message data is kept for the data store.
	//
	// When `customerManagedS3` storage is selected, this parameter is ignored.
	RetentionPeriod interface{} `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Metadata which can be used to manage the data store.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTAnalytics::Pipeline`.
//
// The AWS::IoTAnalytics::Pipeline resource consumes messages from one or more channels and allows you to process the messages before storing them in a data store. You must specify both a `channel` and a `datastore` activity and, optionally, as many as 23 additional activities in the `pipelineActivities` array. For more information, see [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *AWS IoT Analytics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipeline := awscdk.Aws_iotanalytics.NewCfnPipeline(this, jsii.String("MyCfnPipeline"), &cfnPipelineProps{
//   	pipelineActivities: []interface{}{
//   		&activityProperty{
//   			addAttributes: &addAttributesProperty{
//   				attributes: map[string]*string{
//   					"attributesKey": jsii.String("attributes"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			channel: &channelProperty{
//   				channelName: jsii.String("channelName"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			datastore: &datastoreProperty{
//   				datastoreName: jsii.String("datastoreName"),
//   				name: jsii.String("name"),
//   			},
//   			deviceRegistryEnrich: &deviceRegistryEnrichProperty{
//   				attribute: jsii.String("attribute"),
//   				name: jsii.String("name"),
//   				roleArn: jsii.String("roleArn"),
//   				thingName: jsii.String("thingName"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			deviceShadowEnrich: &deviceShadowEnrichProperty{
//   				attribute: jsii.String("attribute"),
//   				name: jsii.String("name"),
//   				roleArn: jsii.String("roleArn"),
//   				thingName: jsii.String("thingName"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			filter: &filterProperty{
//   				filter: jsii.String("filter"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			lambda: &lambdaProperty{
//   				batchSize: jsii.Number(123),
//   				lambdaName: jsii.String("lambdaName"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			math: &mathProperty{
//   				attribute: jsii.String("attribute"),
//   				math: jsii.String("math"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			removeAttributes: &removeAttributesProperty{
//   				attributes: []*string{
//   					jsii.String("attributes"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			selectAttributes: &selectAttributesProperty{
//   				attributes: []*string{
//   					jsii.String("attributes"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	pipelineName: jsii.String("pipelineName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPipeline interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
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
	// A list of "PipelineActivity" objects.
	//
	// Activities perform transformations on your messages, such as removing, renaming or adding message attributes; filtering messages based on attribute values; invoking your Lambda functions on messages for advanced processing; or performing mathematical transformations to normalize device data.
	//
	// The list can be 2-25 *PipelineActivity* objects and must contain both a `channel` and a `datastore` activity. Each entry in the list must contain only one activity, for example:
	//
	// `pipelineActivities = [ { "channel": { ... } }, { "lambda": { ... } }, ... ]`
	PipelineActivities() interface{}
	SetPipelineActivities(val interface{})
	// The name of the pipeline.
	PipelineName() *string
	SetPipelineName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata which can be used to manage the pipeline.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
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

// The jsii proxy struct for CfnPipeline
type jsiiProxy_CfnPipeline struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPipeline) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) PipelineActivities() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineActivities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) PipelineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTAnalytics::Pipeline`.
func NewCfnPipeline(scope constructs.Construct, id *string, props *CfnPipelineProps) CfnPipeline {
	_init_.Initialize()

	j := jsiiProxy_CfnPipeline{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Pipeline`.
func NewCfnPipeline_Override(c CfnPipeline, scope constructs.Construct, id *string, props *CfnPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPipeline) SetPipelineActivities(val interface{}) {
	_jsii_.Set(
		j,
		"pipelineActivities",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetPipelineName(val *string) {
	_jsii_.Set(
		j,
		"pipelineName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPipeline_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPipeline_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
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
func CfnPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipeline_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipeline) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPipeline) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPipeline) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPipeline) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPipeline) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPipeline) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPipeline) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPipeline) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPipeline) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An activity that performs a transformation on a message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activityProperty := &activityProperty{
//   	addAttributes: &addAttributesProperty{
//   		attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	channel: &channelProperty{
//   		channelName: jsii.String("channelName"),
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	datastore: &datastoreProperty{
//   		datastoreName: jsii.String("datastoreName"),
//   		name: jsii.String("name"),
//   	},
//   	deviceRegistryEnrich: &deviceRegistryEnrichProperty{
//   		attribute: jsii.String("attribute"),
//   		name: jsii.String("name"),
//   		roleArn: jsii.String("roleArn"),
//   		thingName: jsii.String("thingName"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	deviceShadowEnrich: &deviceShadowEnrichProperty{
//   		attribute: jsii.String("attribute"),
//   		name: jsii.String("name"),
//   		roleArn: jsii.String("roleArn"),
//   		thingName: jsii.String("thingName"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	filter: &filterProperty{
//   		filter: jsii.String("filter"),
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	lambda: &lambdaProperty{
//   		batchSize: jsii.Number(123),
//   		lambdaName: jsii.String("lambdaName"),
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	math: &mathProperty{
//   		attribute: jsii.String("attribute"),
//   		math: jsii.String("math"),
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	removeAttributes: &removeAttributesProperty{
//   		attributes: []*string{
//   			jsii.String("attributes"),
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	selectAttributes: &selectAttributesProperty{
//   		attributes: []*string{
//   			jsii.String("attributes"),
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   }
//
type CfnPipeline_ActivityProperty struct {
	// Adds other attributes based on existing attributes in the message.
	AddAttributes interface{} `field:"optional" json:"addAttributes" yaml:"addAttributes"`
	// Determines the source of the messages to be processed.
	Channel interface{} `field:"optional" json:"channel" yaml:"channel"`
	// Specifies where to store the processed message data.
	Datastore interface{} `field:"optional" json:"datastore" yaml:"datastore"`
	// Adds data from the AWS IoT device registry to your message.
	DeviceRegistryEnrich interface{} `field:"optional" json:"deviceRegistryEnrich" yaml:"deviceRegistryEnrich"`
	// Adds information from the AWS IoT Device Shadows service to a message.
	DeviceShadowEnrich interface{} `field:"optional" json:"deviceShadowEnrich" yaml:"deviceShadowEnrich"`
	// Filters a message based on its attributes.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Runs a Lambda function to modify the message.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Computes an arithmetic expression using the message's attributes and adds it to the message.
	Math interface{} `field:"optional" json:"math" yaml:"math"`
	// Removes attributes from a message.
	RemoveAttributes interface{} `field:"optional" json:"removeAttributes" yaml:"removeAttributes"`
	// Creates a new message using only the specified attributes from the original message.
	SelectAttributes interface{} `field:"optional" json:"selectAttributes" yaml:"selectAttributes"`
}

// An activity that adds other attributes based on existing attributes in the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addAttributesProperty := &addAttributesProperty{
//   	attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_AddAttributesProperty struct {
	// A list of 1-50 "AttributeNameMapping" objects that map an existing attribute to a new attribute.
	//
	// > The existing attributes remain in the message, so if you want to remove the originals, use "RemoveAttributeActivity".
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the 'addAttributes' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

// Determines the source of the messages to be processed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelProperty := &channelProperty{
//   	channelName: jsii.String("channelName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_ChannelProperty struct {
	// The name of the channel from which the messages are processed.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The name of the 'channel' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

// The datastore activity that specifies where to store the processed data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datastoreProperty := &datastoreProperty{
//   	datastoreName: jsii.String("datastoreName"),
//   	name: jsii.String("name"),
//   }
//
type CfnPipeline_DatastoreProperty struct {
	// The name of the data store where processed messages are stored.
	DatastoreName *string `field:"required" json:"datastoreName" yaml:"datastoreName"`
	// The name of the datastore activity.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// An activity that adds data from the AWS IoT device registry to your message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceRegistryEnrichProperty := &deviceRegistryEnrichProperty{
//   	attribute: jsii.String("attribute"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	thingName: jsii.String("thingName"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_DeviceRegistryEnrichProperty struct {
	// The name of the attribute that is added to the message.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The name of the 'deviceRegistryEnrich' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the role that allows access to the device's registry information.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the IoT device whose registry information is added to the message.
	ThingName *string `field:"required" json:"thingName" yaml:"thingName"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

// An activity that adds information from the AWS IoT Device Shadows service to a message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceShadowEnrichProperty := &deviceShadowEnrichProperty{
//   	attribute: jsii.String("attribute"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	thingName: jsii.String("thingName"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_DeviceShadowEnrichProperty struct {
	// The name of the attribute that is added to the message.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The name of the 'deviceShadowEnrich' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the role that allows access to the device's shadow.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the IoT device whose shadow information is added to the message.
	ThingName *string `field:"required" json:"thingName" yaml:"thingName"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

// An activity that filters a message based on its attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &filterProperty{
//   	filter: jsii.String("filter"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_FilterProperty struct {
	// An expression that looks like an SQL WHERE clause that must return a Boolean value.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// The name of the 'filter' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

// An activity that runs a Lambda function to modify the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaProperty := &lambdaProperty{
//   	batchSize: jsii.Number(123),
//   	lambdaName: jsii.String("lambdaName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_LambdaProperty struct {
	// The number of messages passed to the Lambda function for processing.
	//
	// The AWS Lambda function must be able to process all of these messages within five minutes, which is the maximum timeout duration for Lambda functions.
	BatchSize *float64 `field:"required" json:"batchSize" yaml:"batchSize"`
	// The name of the Lambda function that is run on the message.
	LambdaName *string `field:"required" json:"lambdaName" yaml:"lambdaName"`
	// The name of the 'lambda' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

// An activity that computes an arithmetic expression using the message's attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mathProperty := &mathProperty{
//   	attribute: jsii.String("attribute"),
//   	math: jsii.String("math"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_MathProperty struct {
	// The name of the attribute that contains the result of the math operation.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// An expression that uses one or more existing attributes and must return an integer value.
	Math *string `field:"required" json:"math" yaml:"math"`
	// The name of the 'math' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

// An activity that removes attributes from a message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   removeAttributesProperty := &removeAttributesProperty{
//   	attributes: []*string{
//   		jsii.String("attributes"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_RemoveAttributesProperty struct {
	// A list of 1-50 attributes to remove from the message.
	Attributes *[]*string `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the 'removeAttributes' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

// Creates a new message using only the specified attributes from the original message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selectAttributesProperty := &selectAttributesProperty{
//   	attributes: []*string{
//   		jsii.String("attributes"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_SelectAttributesProperty struct {
	// A list of the attributes to select from the message.
	Attributes *[]*string `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the 'selectAttributes' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

// Properties for defining a `CfnPipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipelineProps := &cfnPipelineProps{
//   	pipelineActivities: []interface{}{
//   		&activityProperty{
//   			addAttributes: &addAttributesProperty{
//   				attributes: map[string]*string{
//   					"attributesKey": jsii.String("attributes"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			channel: &channelProperty{
//   				channelName: jsii.String("channelName"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			datastore: &datastoreProperty{
//   				datastoreName: jsii.String("datastoreName"),
//   				name: jsii.String("name"),
//   			},
//   			deviceRegistryEnrich: &deviceRegistryEnrichProperty{
//   				attribute: jsii.String("attribute"),
//   				name: jsii.String("name"),
//   				roleArn: jsii.String("roleArn"),
//   				thingName: jsii.String("thingName"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			deviceShadowEnrich: &deviceShadowEnrichProperty{
//   				attribute: jsii.String("attribute"),
//   				name: jsii.String("name"),
//   				roleArn: jsii.String("roleArn"),
//   				thingName: jsii.String("thingName"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			filter: &filterProperty{
//   				filter: jsii.String("filter"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			lambda: &lambdaProperty{
//   				batchSize: jsii.Number(123),
//   				lambdaName: jsii.String("lambdaName"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			math: &mathProperty{
//   				attribute: jsii.String("attribute"),
//   				math: jsii.String("math"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			removeAttributes: &removeAttributesProperty{
//   				attributes: []*string{
//   					jsii.String("attributes"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			selectAttributes: &selectAttributesProperty{
//   				attributes: []*string{
//   					jsii.String("attributes"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	pipelineName: jsii.String("pipelineName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipelineProps struct {
	// A list of "PipelineActivity" objects.
	//
	// Activities perform transformations on your messages, such as removing, renaming or adding message attributes; filtering messages based on attribute values; invoking your Lambda functions on messages for advanced processing; or performing mathematical transformations to normalize device data.
	//
	// The list can be 2-25 *PipelineActivity* objects and must contain both a `channel` and a `datastore` activity. Each entry in the list must contain only one activity, for example:
	//
	// `pipelineActivities = [ { "channel": { ... } }, { "lambda": { ... } }, ... ]`
	PipelineActivities interface{} `field:"required" json:"pipelineActivities" yaml:"pipelineActivities"`
	// The name of the pipeline.
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// Metadata which can be used to manage the pipeline.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

