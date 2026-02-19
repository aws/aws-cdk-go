package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiot"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the resource definition of AWS IoT Command.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCommand := awscdk.Aws_iot.NewCfnCommand(this, jsii.String("MyCfnCommand"), &CfnCommandProps{
//   	CommandId: jsii.String("commandId"),
//
//   	// the properties below are optional
//   	CreatedAt: jsii.String("createdAt"),
//   	Deprecated: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	LastUpdatedAt: jsii.String("lastUpdatedAt"),
//   	MandatoryParameters: []interface{}{
//   		&CommandParameterProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			DefaultValue: &CommandParameterValueProperty{
//   				B: jsii.Boolean(false),
//   				Bin: jsii.String("bin"),
//   				D: jsii.Number(123),
//   				I: jsii.Number(123),
//   				L: jsii.String("l"),
//   				S: jsii.String("s"),
//   				Ul: jsii.String("ul"),
//   			},
//   			Description: jsii.String("description"),
//   			Type: jsii.String("type"),
//   			Value: &CommandParameterValueProperty{
//   				B: jsii.Boolean(false),
//   				Bin: jsii.String("bin"),
//   				D: jsii.Number(123),
//   				I: jsii.Number(123),
//   				L: jsii.String("l"),
//   				S: jsii.String("s"),
//   				Ul: jsii.String("ul"),
//   			},
//   			ValueConditions: []interface{}{
//   				&CommandParameterValueConditionProperty{
//   					ComparisonOperator: jsii.String("comparisonOperator"),
//   					Operand: &CommandParameterValueComparisonOperandProperty{
//   						Number: jsii.String("number"),
//   						NumberRange: &CommandParameterValueNumberRangeProperty{
//   							Max: jsii.String("max"),
//   							Min: jsii.String("min"),
//   						},
//   						Numbers: []*string{
//   							jsii.String("numbers"),
//   						},
//   						String: jsii.String("string"),
//   						Strings: []*string{
//   							jsii.String("strings"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   	Payload: &CommandPayloadProperty{
//   		Content: jsii.String("content"),
//   		ContentType: jsii.String("contentType"),
//   	},
//   	PayloadTemplate: jsii.String("payloadTemplate"),
//   	PendingDeletion: jsii.Boolean(false),
//   	Preprocessor: &CommandPreprocessorProperty{
//   		AwsJsonSubstitution: &AwsJsonSubstitutionCommandPreprocessorConfigProperty{
//   			OutputFormat: jsii.String("outputFormat"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html
//
type CfnCommand interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsiot.ICommandRef
	awscdk.ITaggableV2
	// The Amazon Resource Name (ARN) of the command.
	AttrCommandArn() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The unique identifier of the command.
	CommandId() *string
	SetCommandId(val *string)
	// A reference to a Command resource.
	CommandRef() *interfacesawsiot.CommandReference
	// The timestamp, when the command was created.
	CreatedAt() *string
	SetCreatedAt(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Indicates whether the command has been deprecated.
	Deprecated() interface{}
	SetDeprecated(val interface{})
	// The description of the command parameter.
	Description() *string
	SetDescription(val *string)
	// The display name of the command.
	DisplayName() *string
	SetDisplayName(val *string)
	Env() *interfaces.ResourceEnvironment
	// The timestamp, when the command was last updated.
	LastUpdatedAt() *string
	SetLastUpdatedAt(val *string)
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
	MandatoryParameters() interface{}
	SetMandatoryParameters(val interface{})
	// The namespace to which the command belongs.
	Namespace() *string
	SetNamespace(val *string)
	// The tree node.
	Node() constructs.Node
	Payload() interface{}
	SetPayload(val interface{})
	// The payload template associated with the command.
	PayloadTemplate() *string
	SetPayloadTemplate(val *string)
	// Indicates whether the command is pending deletion.
	PendingDeletion() interface{}
	SetPendingDeletion(val interface{})
	Preprocessor() interface{}
	SetPreprocessor(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The customer role associated with the command.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to be associated with the command.
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

// The jsii proxy struct for CfnCommand
type jsiiProxy_CfnCommand struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsiotICommandRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnCommand) AttrCommandArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCommandArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) CommandId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commandId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) CommandRef() *interfacesawsiot.CommandReference {
	var returns *interfacesawsiot.CommandReference
	_jsii_.Get(
		j,
		"commandRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) Deprecated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deprecated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) LastUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) MandatoryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mandatoryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) Payload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) PayloadTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) PendingDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pendingDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) Preprocessor() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preprocessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommand) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::Command`.
func NewCfnCommand(scope constructs.Construct, id *string, props *CfnCommandProps) CfnCommand {
	_init_.Initialize()

	if err := validateNewCfnCommandParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCommand{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnCommand",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::Command`.
func NewCfnCommand_Override(c CfnCommand, scope constructs.Construct, id *string, props *CfnCommandProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnCommand",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCommand)SetCommandId(val *string) {
	if err := j.validateSetCommandIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commandId",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetCreatedAt(val *string) {
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetDeprecated(val interface{}) {
	if err := j.validateSetDeprecatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deprecated",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetLastUpdatedAt(val *string) {
	_jsii_.Set(
		j,
		"lastUpdatedAt",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetMandatoryParameters(val interface{}) {
	if err := j.validateSetMandatoryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mandatoryParameters",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetPayload(val interface{}) {
	if err := j.validateSetPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payload",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetPayloadTemplate(val *string) {
	_jsii_.Set(
		j,
		"payloadTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetPendingDeletion(val interface{}) {
	if err := j.validateSetPendingDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pendingDeletion",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetPreprocessor(val interface{}) {
	if err := j.validateSetPreprocessorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preprocessor",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnCommand)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func CfnCommand_ArnForCommand(resource interfacesawsiot.ICommandRef) *string {
	_init_.Initialize()

	if err := validateCfnCommand_ArnForCommandParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCommand",
		"arnForCommand",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Creates a new ICommandRef from an ARN.
func CfnCommand_FromCommandArn(scope constructs.Construct, id *string, arn *string) interfacesawsiot.ICommandRef {
	_init_.Initialize()

	if err := validateCfnCommand_FromCommandArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns interfacesawsiot.ICommandRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCommand",
		"fromCommandArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

// Creates a new ICommandRef from a commandId.
func CfnCommand_FromCommandId(scope constructs.Construct, id *string, commandId *string) interfacesawsiot.ICommandRef {
	_init_.Initialize()

	if err := validateCfnCommand_FromCommandIdParameters(scope, id, commandId); err != nil {
		panic(err)
	}
	var returns interfacesawsiot.ICommandRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCommand",
		"fromCommandId",
		[]interface{}{scope, id, commandId},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnCommand.
func CfnCommand_IsCfnCommand(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCommand_IsCfnCommandParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCommand",
		"isCfnCommand",
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
func CfnCommand_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCommand_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCommand",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCommand_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCommand_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCommand",
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
func CfnCommand_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCommand_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCommand",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCommand_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnCommand",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCommand) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCommand) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCommand) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCommand) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCommand) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCommand) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCommand) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCommand) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCommand) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCommand) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCommand) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCommand) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCommand) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCommand) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCommand) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCommand) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCommand) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCommand) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCommand) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCommand) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnCommand) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

