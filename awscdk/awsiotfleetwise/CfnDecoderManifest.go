package awsiotfleetwise

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotfleetwise"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates the decoder manifest associated with a model manifest. To create a decoder manifest, the following must be true:.
//
// - Every signal decoder has a unique name.
// - Each signal decoder is associated with a network interface.
// - Each network interface has a unique ID.
// - The signal decoders are specified in the model manifest.
//
// For more information, see [Decoder manifests](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/decoder-manifests.html) in the *AWS IoT FleetWise Developer Guide* .
//
// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDecoderManifest := awscdk.Aws_iotfleetwise.NewCfnDecoderManifest(this, jsii.String("MyCfnDecoderManifest"), &CfnDecoderManifestProps{
//   	ModelManifestArn: jsii.String("modelManifestArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DefaultForUnmappedSignals: jsii.String("defaultForUnmappedSignals"),
//   	Description: jsii.String("description"),
//   	NetworkInterfaces: []interface{}{
//   		&NetworkInterfacesItemsProperty{
//   			InterfaceId: jsii.String("interfaceId"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			CanInterface: &CanInterfaceProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				ProtocolName: jsii.String("protocolName"),
//   				ProtocolVersion: jsii.String("protocolVersion"),
//   			},
//   			ObdInterface: &ObdInterfaceProperty{
//   				Name: jsii.String("name"),
//   				RequestMessageId: jsii.String("requestMessageId"),
//
//   				// the properties below are optional
//   				DtcRequestIntervalSeconds: jsii.String("dtcRequestIntervalSeconds"),
//   				HasTransmissionEcu: jsii.String("hasTransmissionEcu"),
//   				ObdStandard: jsii.String("obdStandard"),
//   				PidRequestIntervalSeconds: jsii.String("pidRequestIntervalSeconds"),
//   				UseExtendedIds: jsii.String("useExtendedIds"),
//   			},
//   		},
//   	},
//   	SignalDecoders: []interface{}{
//   		&SignalDecodersItemsProperty{
//   			FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   			InterfaceId: jsii.String("interfaceId"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			CanSignal: &CanSignalProperty{
//   				Factor: jsii.String("factor"),
//   				IsBigEndian: jsii.String("isBigEndian"),
//   				IsSigned: jsii.String("isSigned"),
//   				Length: jsii.String("length"),
//   				MessageId: jsii.String("messageId"),
//   				Offset: jsii.String("offset"),
//   				StartBit: jsii.String("startBit"),
//
//   				// the properties below are optional
//   				Name: jsii.String("name"),
//   				SignalValueType: jsii.String("signalValueType"),
//   			},
//   			ObdSignal: &ObdSignalProperty{
//   				ByteLength: jsii.String("byteLength"),
//   				Offset: jsii.String("offset"),
//   				Pid: jsii.String("pid"),
//   				PidResponseLength: jsii.String("pidResponseLength"),
//   				Scaling: jsii.String("scaling"),
//   				ServiceMode: jsii.String("serviceMode"),
//   				StartByte: jsii.String("startByte"),
//
//   				// the properties below are optional
//   				BitMaskLength: jsii.String("bitMaskLength"),
//   				BitRightShift: jsii.String("bitRightShift"),
//   				IsSigned: jsii.String("isSigned"),
//   				SignalValueType: jsii.String("signalValueType"),
//   			},
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-decodermanifest.html
//
type CfnDecoderManifest interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsiotfleetwise.IDecoderManifestRef
	awscdk.ITaggable
	// The Amazon Resource Name (ARN) of the decoder manifest.
	AttrArn() *string
	// The time the decoder manifest was created in seconds since epoch (January 1, 1970 at midnight UTC time).
	AttrCreationTime() *string
	// The time the decoder manifest was last updated in seconds since epoch (January 1, 1970 at midnight UTC time).
	AttrLastModificationTime() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A reference to a DecoderManifest resource.
	DecoderManifestRef() *interfacesawsiotfleetwise.DecoderManifestReference
	// Use default decoders for all unmapped signals in the model.
	DefaultForUnmappedSignals() *string
	SetDefaultForUnmappedSignals(val *string)
	// A brief description of the decoder manifest.
	Description() *string
	SetDescription(val *string)
	Env() *interfaces.ResourceEnvironment
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
	// The Amazon Resource Name (ARN) of a vehicle model (model manifest) associated with the decoder manifest.
	ModelManifestArn() *string
	SetModelManifestArn(val *string)
	// The name of the decoder manifest.
	Name() *string
	SetName(val *string)
	// A list of information about available network interfaces.
	NetworkInterfaces() interface{}
	SetNetworkInterfaces(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A list of information about signal decoders.
	SignalDecoders() interface{}
	SetSignalDecoders(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The state of the decoder manifest.
	Status() *string
	SetStatus(val *string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Metadata that can be used to manage the decoder manifest.
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnDecoderManifest
type jsiiProxy_CfnDecoderManifest struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsiotfleetwiseIDecoderManifestRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnDecoderManifest) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) AttrLastModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) DecoderManifestRef() *interfacesawsiotfleetwise.DecoderManifestReference {
	var returns *interfacesawsiotfleetwise.DecoderManifestReference
	_jsii_.Get(
		j,
		"decoderManifestRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) DefaultForUnmappedSignals() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultForUnmappedSignals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) ModelManifestArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelManifestArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) NetworkInterfaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) SignalDecoders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signalDecoders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifest) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTFleetWise::DecoderManifest`.
func NewCfnDecoderManifest(scope constructs.Construct, id *string, props *CfnDecoderManifestProps) CfnDecoderManifest {
	_init_.Initialize()

	if err := validateNewCfnDecoderManifestParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDecoderManifest{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotfleetwise.CfnDecoderManifest",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTFleetWise::DecoderManifest`.
func NewCfnDecoderManifest_Override(c CfnDecoderManifest, scope constructs.Construct, id *string, props *CfnDecoderManifestProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotfleetwise.CfnDecoderManifest",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDecoderManifest)SetDefaultForUnmappedSignals(val *string) {
	_jsii_.Set(
		j,
		"defaultForUnmappedSignals",
		val,
	)
}

func (j *jsiiProxy_CfnDecoderManifest)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDecoderManifest)SetModelManifestArn(val *string) {
	if err := j.validateSetModelManifestArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelManifestArn",
		val,
	)
}

func (j *jsiiProxy_CfnDecoderManifest)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDecoderManifest)SetNetworkInterfaces(val interface{}) {
	if err := j.validateSetNetworkInterfacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaces",
		val,
	)
}

func (j *jsiiProxy_CfnDecoderManifest)SetSignalDecoders(val interface{}) {
	if err := j.validateSetSignalDecodersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signalDecoders",
		val,
	)
}

func (j *jsiiProxy_CfnDecoderManifest)SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CfnDecoderManifest)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func CfnDecoderManifest_ArnForDecoderManifest(resource interfacesawsiotfleetwise.IDecoderManifestRef) *string {
	_init_.Initialize()

	if err := validateCfnDecoderManifest_ArnForDecoderManifestParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotfleetwise.CfnDecoderManifest",
		"arnForDecoderManifest",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnDecoderManifest.
func CfnDecoderManifest_IsCfnDecoderManifest(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDecoderManifest_IsCfnDecoderManifestParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotfleetwise.CfnDecoderManifest",
		"isCfnDecoderManifest",
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
func CfnDecoderManifest_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDecoderManifest_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotfleetwise.CfnDecoderManifest",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDecoderManifest_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDecoderManifest_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotfleetwise.CfnDecoderManifest",
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
func CfnDecoderManifest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDecoderManifest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotfleetwise.CfnDecoderManifest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDecoderManifest_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotfleetwise.CfnDecoderManifest",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDecoderManifest) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnDecoderManifest) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnDecoderManifest) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDecoderManifest) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDecoderManifest) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnDecoderManifest) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDecoderManifest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDecoderManifest) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnDecoderManifest) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

