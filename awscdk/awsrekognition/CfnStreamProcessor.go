package awsrekognition

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrekognition/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Rekognition::StreamProcessor` type creates a stream processor used to detect and recognize faces or to detect connected home labels in a streaming video.
//
// Amazon Rekognition Video is a consumer of live video from Amazon Kinesis Video Streams. There are two different settings for stream processors in Amazon Rekognition, one for detecting faces and one for connected home features.
//
// If you are creating a stream processor for detecting faces, you provide a Kinesis video stream (input) and a Kinesis data stream (output). You also specify the face recognition criteria in FaceSearchSettings. For example, the collection containing faces that you want to recognize.
//
// If you are creating a stream processor for detection of connected home labels, you provide a Kinesis video stream for input, and for output an Amazon S3 bucket and an Amazon SNS topic. You can also provide a KMS key ID to encrypt the data sent to your Amazon S3 bucket. You specify what you want to detect in ConnectedHomeSettings, such as people, packages, and pets.
//
// You can also specify where in the frame you want Amazon Rekognition to monitor with BoundingBoxRegionsOfInterest and PolygonRegionsOfInterest. The Name is used to manage the stream processor and it is the identifier for the stream processor. The `AWS::Rekognition::StreamProcessor` resource creates a stream processor in the same Region where you create the Amazon CloudFormation stack.
//
// For more information, see [CreateStreamProcessor](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateStreamProcessor) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var polygonRegionsOfInterest interface{}
//
//   cfnStreamProcessor := awscdk.Aws_rekognition.NewCfnStreamProcessor(this, jsii.String("MyCfnStreamProcessor"), &CfnStreamProcessorProps{
//   	KinesisVideoStream: &KinesisVideoStreamProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	BoundingBoxRegionsOfInterest: []interface{}{
//   		&BoundingBoxProperty{
//   			Height: jsii.Number(123),
//   			Left: jsii.Number(123),
//   			Top: jsii.Number(123),
//   			Width: jsii.Number(123),
//   		},
//   	},
//   	ConnectedHomeSettings: &ConnectedHomeSettingsProperty{
//   		Labels: []*string{
//   			jsii.String("labels"),
//   		},
//
//   		// the properties below are optional
//   		MinConfidence: jsii.Number(123),
//   	},
//   	DataSharingPreference: &DataSharingPreferenceProperty{
//   		OptIn: jsii.Boolean(false),
//   	},
//   	FaceSearchSettings: &FaceSearchSettingsProperty{
//   		CollectionId: jsii.String("collectionId"),
//
//   		// the properties below are optional
//   		FaceMatchThreshold: jsii.Number(123),
//   	},
//   	KinesisDataStream: &KinesisDataStreamProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	NotificationChannel: &NotificationChannelProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	PolygonRegionsOfInterest: polygonRegionsOfInterest,
//   	S3Destination: &S3DestinationProperty{
//   		BucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rekognition-streamprocessor.html
//
type CfnStreamProcessor interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// Amazon Resource Name for the newly created stream processor.
	AttrArn() *string
	// Current status of the Amazon Rekognition stream processor.
	AttrStatus() *string
	// Detailed status message about the stream processor.
	AttrStatusMessage() *string
	// List of BoundingBox objects, each of which denotes a region of interest on screen.
	BoundingBoxRegionsOfInterest() interface{}
	SetBoundingBoxRegionsOfInterest(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Connected home settings to use on a streaming video.
	ConnectedHomeSettings() interface{}
	SetConnectedHomeSettings(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Allows you to opt in or opt out to share data with Rekognition to improve model performance.
	DataSharingPreference() interface{}
	SetDataSharingPreference(val interface{})
	// The input parameters used to recognize faces in a streaming video analyzed by an Amazon Rekognition stream processor.
	FaceSearchSettings() interface{}
	SetFaceSearchSettings(val interface{})
	// Amazon Rekognition's Video Stream Processor takes a Kinesis video stream as input.
	KinesisDataStream() interface{}
	SetKinesisDataStream(val interface{})
	// The Kinesis video stream that provides the source of the streaming video for an Amazon Rekognition Video stream processor.
	KinesisVideoStream() interface{}
	SetKinesisVideoStream(val interface{})
	// The identifier for your Amazon Key Management Service key (Amazon KMS key).
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	// The Name attribute specifies the name of the stream processor and it must be within the constraints described in the Name section of [StreamProcessor](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StreamProcessor) . If you don't specify a name, Amazon CloudFormation generates a unique ID and uses that ID for the stream processor name.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The Amazon Simple Notification Service topic to which Amazon Rekognition publishes the object detection results and completion status of a video analysis operation.
	NotificationChannel() interface{}
	SetNotificationChannel(val interface{})
	// A set of ordered lists of [Point](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_Point) objects. Each entry of the set contains a polygon denoting a region of interest on the screen. Each polygon is an ordered list of [Point](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_Point) objects. For more information, see the Polygon field of [RegionOfInterest](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_RegionOfInterest) .
	PolygonRegionsOfInterest() interface{}
	SetPolygonRegionsOfInterest(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ARN of the IAM role that allows access to the stream processor.
	RoleArn() *string
	SetRoleArn(val *string)
	// The Amazon S3 bucket location to which Amazon Rekognition publishes the detailed inference results of a video analysis operation.
	S3Destination() interface{}
	SetS3Destination(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// A set of tags (key-value pairs) that you want to attach to the stream processor.
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

// The jsii proxy struct for CfnStreamProcessor
type jsiiProxy_CfnStreamProcessor struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnStreamProcessor) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) AttrStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) BoundingBoxRegionsOfInterest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boundingBoxRegionsOfInterest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) ConnectedHomeSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectedHomeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) DataSharingPreference() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSharingPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) FaceSearchSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"faceSearchSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) KinesisDataStream() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisDataStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) KinesisVideoStream() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisVideoStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) NotificationChannel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) PolygonRegionsOfInterest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"polygonRegionsOfInterest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) S3Destination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamProcessor) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnStreamProcessor(scope constructs.Construct, id *string, props *CfnStreamProcessorProps) CfnStreamProcessor {
	_init_.Initialize()

	if err := validateNewCfnStreamProcessorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStreamProcessor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rekognition.CfnStreamProcessor",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnStreamProcessor_Override(c CfnStreamProcessor, scope constructs.Construct, id *string, props *CfnStreamProcessorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rekognition.CfnStreamProcessor",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetBoundingBoxRegionsOfInterest(val interface{}) {
	if err := j.validateSetBoundingBoxRegionsOfInterestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundingBoxRegionsOfInterest",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetConnectedHomeSettings(val interface{}) {
	if err := j.validateSetConnectedHomeSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectedHomeSettings",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetDataSharingPreference(val interface{}) {
	if err := j.validateSetDataSharingPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSharingPreference",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetFaceSearchSettings(val interface{}) {
	if err := j.validateSetFaceSearchSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faceSearchSettings",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetKinesisDataStream(val interface{}) {
	if err := j.validateSetKinesisDataStreamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kinesisDataStream",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetKinesisVideoStream(val interface{}) {
	if err := j.validateSetKinesisVideoStreamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kinesisVideoStream",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetNotificationChannel(val interface{}) {
	if err := j.validateSetNotificationChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationChannel",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetPolygonRegionsOfInterest(val interface{}) {
	_jsii_.Set(
		j,
		"polygonRegionsOfInterest",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetS3Destination(val interface{}) {
	if err := j.validateSetS3DestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Destination",
		val,
	)
}

func (j *jsiiProxy_CfnStreamProcessor)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStreamProcessor_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStreamProcessor_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rekognition.CfnStreamProcessor",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStreamProcessor_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnStreamProcessor_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rekognition.CfnStreamProcessor",
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
func CfnStreamProcessor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStreamProcessor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rekognition.CfnStreamProcessor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStreamProcessor_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rekognition.CfnStreamProcessor",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStreamProcessor) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnStreamProcessor) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnStreamProcessor) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStreamProcessor) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStreamProcessor) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnStreamProcessor) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnStreamProcessor) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStreamProcessor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStreamProcessor) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

