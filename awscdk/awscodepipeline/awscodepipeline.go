package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v3"
)

// Low-level class for generic CodePipeline Actions implementing the {@link IAction} interface.
//
// Contains some common logic that can be re-used by all {@link IAction} implementations.
// If you're writing your own Action class,
// feel free to extend this class.
// Experimental.
type Action interface {
	IAction
	ActionProperties() *ActionProperties
	ProvidedActionProperties() *ActionProperties
	Bind(scope awscdk.Construct, stage IStage, options *ActionBindOptions) *ActionConfig
	Bound(scope awscdk.Construct, stage IStage, options *ActionBindOptions) *ActionConfig
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for Action
type jsiiProxy_Action struct {
	jsiiProxy_IAction
}

func (j *jsiiProxy_Action) ActionProperties() *ActionProperties {
	var returns *ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Action) ProvidedActionProperties() *ActionProperties {
	var returns *ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


// Experimental.
func NewAction_Override(a Action) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline.Action",
		nil, // no parameters
		a,
	)
}

// The callback invoked when this Action is added to a Pipeline.
// Experimental.
func (a *jsiiProxy_Action) Bind(scope awscdk.Construct, stage IStage, options *ActionBindOptions) *ActionConfig {
	var returns *ActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

// This is a renamed version of the {@link IAction.bind} method.
// Experimental.
func (a *jsiiProxy_Action) Bound(scope awscdk.Construct, stage IStage, options *ActionBindOptions) *ActionConfig {
	var returns *ActionConfig

	_jsii_.Invoke(
		a,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

// Creates an Event that will be triggered whenever the state of this Action changes.
// Experimental.
func (a *jsiiProxy_Action) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		a,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_Action) VariableExpression(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

// Specifies the constraints on the number of input and output artifacts an action can have.
//
// The constraints for each action type are documented on the
// {@link https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html Pipeline Structure Reference} page.
//
// TODO: EXAMPLE
//
// Experimental.
type ActionArtifactBounds struct {
	// Experimental.
	MaxInputs *float64 `json:"maxInputs" yaml:"maxInputs"`
	// Experimental.
	MaxOutputs *float64 `json:"maxOutputs" yaml:"maxOutputs"`
	// Experimental.
	MinInputs *float64 `json:"minInputs" yaml:"minInputs"`
	// Experimental.
	MinOutputs *float64 `json:"minOutputs" yaml:"minOutputs"`
}

// TODO: EXAMPLE
//
// Experimental.
type ActionBindOptions struct {
	// Experimental.
	Bucket awss3.IBucket `json:"bucket" yaml:"bucket"`
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// TODO: EXAMPLE
//
// Experimental.
type ActionCategory string

const (
	ActionCategory_SOURCE ActionCategory = "SOURCE"
	ActionCategory_BUILD ActionCategory = "BUILD"
	ActionCategory_TEST ActionCategory = "TEST"
	ActionCategory_APPROVAL ActionCategory = "APPROVAL"
	ActionCategory_DEPLOY ActionCategory = "DEPLOY"
	ActionCategory_INVOKE ActionCategory = "INVOKE"
)

// TODO: EXAMPLE
//
// Experimental.
type ActionConfig struct {
	// Experimental.
	Configuration interface{} `json:"configuration" yaml:"configuration"`
}

// TODO: EXAMPLE
//
// Experimental.
type ActionProperties struct {
	// Experimental.
	ActionName *string `json:"actionName" yaml:"actionName"`
	// Experimental.
	ArtifactBounds *ActionArtifactBounds `json:"artifactBounds" yaml:"artifactBounds"`
	// The category of the action.
	//
	// The category defines which action type the owner
	// (the entity that performs the action) performs.
	// Experimental.
	Category ActionCategory `json:"category" yaml:"category"`
	// The service provider that the action calls.
	// Experimental.
	Provider *string `json:"provider" yaml:"provider"`
	// The account the Action is supposed to live in.
	//
	// For Actions backed by resources,
	// this is inferred from the Stack {@link resource} is part of.
	// However, some Actions, like the CloudFormation ones,
	// are not backed by any resource, and they still might want to be cross-account.
	// In general, a concrete Action class should specify either {@link resource},
	// or {@link account} - but not both.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// Experimental.
	Inputs *[]Artifact `json:"inputs" yaml:"inputs"`
	// Experimental.
	Outputs *[]Artifact `json:"outputs" yaml:"outputs"`
	// Experimental.
	Owner *string `json:"owner" yaml:"owner"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the {@link PipelineProps#crossRegionReplicationBuckets} property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// The optional resource that is backing this Action.
	//
	// This is used for automatically handling Actions backed by
	// resources from a different account and/or region.
	// Experimental.
	Resource awscdk.IResource `json:"resource" yaml:"resource"`
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// The order in which AWS CodePipeline runs this action. For more information, see the AWS CodePipeline User Guide.
	//
	// https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements
	// Experimental.
	RunOrder *float64 `json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace" yaml:"variablesNamespace"`
	// Experimental.
	Version *string `json:"version" yaml:"version"`
}

// An output artifact of an action.
//
// Artifacts can be used as input by some actions.
//
// TODO: EXAMPLE
//
// Experimental.
type Artifact interface {
	ArtifactName() *string
	BucketName() *string
	ObjectKey() *string
	S3Location() *awss3.Location
	Url() *string
	AtPath(fileName *string) ArtifactPath
	GetMetadata(key *string) interface{}
	GetParam(jsonFile *string, keyName *string) *string
	SetMetadata(key *string, value interface{})
	ToString() *string
}

// The jsii proxy struct for Artifact
type jsiiProxy_Artifact struct {
	_ byte // padding
}

func (j *jsiiProxy_Artifact) ArtifactName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Artifact) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Artifact) ObjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Artifact) S3Location() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Artifact) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Experimental.
func NewArtifact(artifactName *string) Artifact {
	_init_.Initialize()

	j := jsiiProxy_Artifact{}

	_jsii_.Create(
		"monocdk.aws_codepipeline.Artifact",
		[]interface{}{artifactName},
		&j,
	)

	return &j
}

// Experimental.
func NewArtifact_Override(a Artifact, artifactName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline.Artifact",
		[]interface{}{artifactName},
		a,
	)
}

// A static factory method used to create instances of the Artifact class.
//
// Mainly meant to be used from `decdk`.
// Experimental.
func Artifact_Artifact(name *string) Artifact {
	_init_.Initialize()

	var returns Artifact

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.Artifact",
		"artifact",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns an ArtifactPath for a file within this artifact.
//
// CfnOutput is in the form "<artifact-name>::<file-name>"
// Experimental.
func (a *jsiiProxy_Artifact) AtPath(fileName *string) ArtifactPath {
	var returns ArtifactPath

	_jsii_.Invoke(
		a,
		"atPath",
		[]interface{}{fileName},
		&returns,
	)

	return returns
}

// Retrieve the metadata stored in this artifact under the given key.
//
// If there is no metadata stored under the given key,
// null will be returned.
// Experimental.
func (a *jsiiProxy_Artifact) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Returns a token for a value inside a JSON file within this artifact.
// Experimental.
func (a *jsiiProxy_Artifact) GetParam(jsonFile *string, keyName *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getParam",
		[]interface{}{jsonFile, keyName},
		&returns,
	)

	return returns
}

// Add arbitrary extra payload to the artifact under a given key.
//
// This can be used by CodePipeline actions to communicate data between themselves.
// If metadata was already present under the given key,
// it will be overwritten with the new value.
// Experimental.
func (a *jsiiProxy_Artifact) SetMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"setMetadata",
		[]interface{}{key, value},
	)
}

// Experimental.
func (a *jsiiProxy_Artifact) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A specific file within an output artifact.
//
// The most common use case for this is specifying the template file
// for a CloudFormation action.
//
// TODO: EXAMPLE
//
// Experimental.
type ArtifactPath interface {
	Artifact() Artifact
	FileName() *string
	Location() *string
}

// The jsii proxy struct for ArtifactPath
type jsiiProxy_ArtifactPath struct {
	_ byte // padding
}

func (j *jsiiProxy_ArtifactPath) Artifact() Artifact {
	var returns Artifact
	_jsii_.Get(
		j,
		"artifact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactPath) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactPath) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}


// Experimental.
func NewArtifactPath(artifact Artifact, fileName *string) ArtifactPath {
	_init_.Initialize()

	j := jsiiProxy_ArtifactPath{}

	_jsii_.Create(
		"monocdk.aws_codepipeline.ArtifactPath",
		[]interface{}{artifact, fileName},
		&j,
	)

	return &j
}

// Experimental.
func NewArtifactPath_Override(a ArtifactPath, artifact Artifact, fileName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline.ArtifactPath",
		[]interface{}{artifact, fileName},
		a,
	)
}

// Experimental.
func ArtifactPath_ArtifactPath(artifactName *string, fileName *string) ArtifactPath {
	_init_.Initialize()

	var returns ArtifactPath

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.ArtifactPath",
		"artifactPath",
		[]interface{}{artifactName, fileName},
		&returns,
	)

	return returns
}

// A CloudFormation `AWS::CodePipeline::CustomActionType`.
//
// The `AWS::CodePipeline::CustomActionType` resource creates a custom action for activities that aren't included in the CodePipeline default actions, such as running an internally developed build process or a test suite. You can use these custom actions in the stage of a pipeline. For more information, see [Create and Add a Custom Action in AWS CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html) in the *AWS CodePipeline User Guide* .
//
// TODO: EXAMPLE
//
type CfnCustomActionType interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Category() *string
	SetCategory(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConfigurationProperties() interface{}
	SetConfigurationProperties(val interface{})
	CreationStack() *[]*string
	InputArtifactDetails() interface{}
	SetInputArtifactDetails(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
	OutputArtifactDetails() interface{}
	SetOutputArtifactDetails(val interface{})
	Provider() *string
	SetProvider(val *string)
	Ref() *string
	Settings() interface{}
	SetSettings(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	Version() *string
	SetVersion(val *string)
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCustomActionType
type jsiiProxy_CfnCustomActionType struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCustomActionType) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) ConfigurationProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) InputArtifactDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputArtifactDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) OutputArtifactDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputArtifactDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) Settings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionType) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new `AWS::CodePipeline::CustomActionType`.
func NewCfnCustomActionType(scope awscdk.Construct, id *string, props *CfnCustomActionTypeProps) CfnCustomActionType {
	_init_.Initialize()

	j := jsiiProxy_CfnCustomActionType{}

	_jsii_.Create(
		"monocdk.aws_codepipeline.CfnCustomActionType",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodePipeline::CustomActionType`.
func NewCfnCustomActionType_Override(c CfnCustomActionType, scope awscdk.Construct, id *string, props *CfnCustomActionTypeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline.CfnCustomActionType",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCustomActionType) SetCategory(val *string) {
	_jsii_.Set(
		j,
		"category",
		val,
	)
}

func (j *jsiiProxy_CfnCustomActionType) SetConfigurationProperties(val interface{}) {
	_jsii_.Set(
		j,
		"configurationProperties",
		val,
	)
}

func (j *jsiiProxy_CfnCustomActionType) SetInputArtifactDetails(val interface{}) {
	_jsii_.Set(
		j,
		"inputArtifactDetails",
		val,
	)
}

func (j *jsiiProxy_CfnCustomActionType) SetOutputArtifactDetails(val interface{}) {
	_jsii_.Set(
		j,
		"outputArtifactDetails",
		val,
	)
}

func (j *jsiiProxy_CfnCustomActionType) SetProvider(val *string) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CfnCustomActionType) SetSettings(val interface{}) {
	_jsii_.Set(
		j,
		"settings",
		val,
	)
}

func (j *jsiiProxy_CfnCustomActionType) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
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
func CfnCustomActionType_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.CfnCustomActionType",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCustomActionType_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.CfnCustomActionType",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCustomActionType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.CfnCustomActionType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomActionType_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codepipeline.CfnCustomActionType",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnCustomActionType) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnCustomActionType) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnCustomActionType) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnCustomActionType) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnCustomActionType) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCustomActionType) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCustomActionType) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnCustomActionType) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCustomActionType) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnCustomActionType) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Returns information about the details of an artifact.
//
// TODO: EXAMPLE
//
type CfnCustomActionType_ArtifactDetailsProperty struct {
	// The maximum number of artifacts allowed for the action type.
	MaximumCount *float64 `json:"maximumCount" yaml:"maximumCount"`
	// The minimum number of artifacts allowed for the action type.
	MinimumCount *float64 `json:"minimumCount" yaml:"minimumCount"`
}

// The configuration properties for the custom action.
//
// > You can refer to a name in the configuration properties of the custom action within the URL templates by following the format of {Config:name}, as long as the configuration property is both required and not secret. For more information, see [Create a Custom Action for a Pipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html) .
//
// TODO: EXAMPLE
//
type CfnCustomActionType_ConfigurationPropertiesProperty struct {
	// Whether the configuration property is a key.
	Key interface{} `json:"key" yaml:"key"`
	// The name of the action configuration property.
	Name *string `json:"name" yaml:"name"`
	// Whether the configuration property is a required value.
	Required interface{} `json:"required" yaml:"required"`
	// Whether the configuration property is secret.
	//
	// Secrets are hidden from all calls except for `GetJobDetails` , `GetThirdPartyJobDetails` , `PollForJobs` , and `PollForThirdPartyJobs` .
	//
	// When updating a pipeline, passing * * * * * without changing any other values of the action preserves the previous value of the secret.
	Secret interface{} `json:"secret" yaml:"secret"`
	// The description of the action configuration property that is displayed to users.
	Description *string `json:"description" yaml:"description"`
	// Indicates that the property is used with `PollForJobs` .
	//
	// When creating a custom action, an action can have up to one queryable property. If it has one, that property must be both required and not secret.
	//
	// If you create a pipeline with a custom action type, and that custom action contains a queryable property, the value for that configuration property is subject to other restrictions. The value must be less than or equal to twenty (20) characters. The value can contain only alphanumeric characters, underscores, and hyphens.
	Queryable interface{} `json:"queryable" yaml:"queryable"`
	// The type of the configuration property.
	Type *string `json:"type" yaml:"type"`
}

// `Settings` is a property of the `AWS::CodePipeline::CustomActionType` resource that provides URLs that users can access to view information about the CodePipeline custom action.
//
// TODO: EXAMPLE
//
type CfnCustomActionType_SettingsProperty struct {
	// The URL returned to the CodePipeline console that provides a deep link to the resources of the external system, such as the configuration page for a CodeDeploy deployment group.
	//
	// This link is provided as part of the action display in the pipeline.
	EntityUrlTemplate *string `json:"entityUrlTemplate" yaml:"entityUrlTemplate"`
	// The URL returned to the CodePipeline console that contains a link to the top-level landing page for the external system, such as the console page for CodeDeploy.
	//
	// This link is shown on the pipeline view page in the CodePipeline console and provides a link to the execution entity of the external action.
	ExecutionUrlTemplate *string `json:"executionUrlTemplate" yaml:"executionUrlTemplate"`
	// The URL returned to the CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action.
	RevisionUrlTemplate *string `json:"revisionUrlTemplate" yaml:"revisionUrlTemplate"`
	// The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
	ThirdPartyConfigurationUrl *string `json:"thirdPartyConfigurationUrl" yaml:"thirdPartyConfigurationUrl"`
}

// Properties for defining a `CfnCustomActionType`.
//
// TODO: EXAMPLE
//
type CfnCustomActionTypeProps struct {
	// The category of the custom action, such as a build action or a test action.
	Category *string `json:"category" yaml:"category"`
	// The details of the input artifact for the action, such as its commit ID.
	InputArtifactDetails interface{} `json:"inputArtifactDetails" yaml:"inputArtifactDetails"`
	// The details of the output artifact of the action, such as its commit ID.
	OutputArtifactDetails interface{} `json:"outputArtifactDetails" yaml:"outputArtifactDetails"`
	// The provider of the service used in the custom action, such as CodeDeploy.
	Provider *string `json:"provider" yaml:"provider"`
	// The version identifier of the custom action.
	Version *string `json:"version" yaml:"version"`
	// The configuration properties for the custom action.
	//
	// > You can refer to a name in the configuration properties of the custom action within the URL templates by following the format of {Config:name}, as long as the configuration property is both required and not secret. For more information, see [Create a Custom Action for a Pipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html) .
	ConfigurationProperties interface{} `json:"configurationProperties" yaml:"configurationProperties"`
	// URLs that provide users information about this custom action.
	Settings interface{} `json:"settings" yaml:"settings"`
	// The tags for the custom action.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::CodePipeline::Pipeline`.
//
// The `AWS::CodePipeline::Pipeline` resource creates a CodePipeline pipeline that describes how software changes go through a release process. For more information, see [What Is CodePipeline?](https://docs.aws.amazon.com/codepipeline/latest/userguide/welcome.html) in the *AWS CodePipeline User Guide* .
//
// TODO: EXAMPLE
//
type CfnPipeline interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ArtifactStore() interface{}
	SetArtifactStore(val interface{})
	ArtifactStores() interface{}
	SetArtifactStores(val interface{})
	AttrVersion() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DisableInboundStageTransitions() interface{}
	SetDisableInboundStageTransitions(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	RestartExecutionOnUpdate() interface{}
	SetRestartExecutionOnUpdate(val interface{})
	RoleArn() *string
	SetRoleArn(val *string)
	Stack() awscdk.Stack
	Stages() interface{}
	SetStages(val interface{})
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnPipeline
type jsiiProxy_CfnPipeline struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPipeline) ArtifactStore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"artifactStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) ArtifactStores() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"artifactStores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) AttrVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVersion",
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

func (j *jsiiProxy_CfnPipeline) DisableInboundStageTransitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableInboundStageTransitions",
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

func (j *jsiiProxy_CfnPipeline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
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

func (j *jsiiProxy_CfnPipeline) RestartExecutionOnUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartExecutionOnUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
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

func (j *jsiiProxy_CfnPipeline) Stages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stages",
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


// Create a new `AWS::CodePipeline::Pipeline`.
func NewCfnPipeline(scope awscdk.Construct, id *string, props *CfnPipelineProps) CfnPipeline {
	_init_.Initialize()

	j := jsiiProxy_CfnPipeline{}

	_jsii_.Create(
		"monocdk.aws_codepipeline.CfnPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodePipeline::Pipeline`.
func NewCfnPipeline_Override(c CfnPipeline, scope awscdk.Construct, id *string, props *CfnPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline.CfnPipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPipeline) SetArtifactStore(val interface{}) {
	_jsii_.Set(
		j,
		"artifactStore",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetArtifactStores(val interface{}) {
	_jsii_.Set(
		j,
		"artifactStores",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetDisableInboundStageTransitions(val interface{}) {
	_jsii_.Set(
		j,
		"disableInboundStageTransitions",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetRestartExecutionOnUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"restartExecutionOnUpdate",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetStages(val interface{}) {
	_jsii_.Set(
		j,
		"stages",
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
func CfnPipeline_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.CfnPipeline",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPipeline_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.CfnPipeline",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.CfnPipeline",
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
		"monocdk.aws_codepipeline.CfnPipeline",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnPipeline) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnPipeline) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnPipeline) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnPipeline) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnPipeline) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnPipeline) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnPipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
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

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnPipeline) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnPipeline) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnPipeline) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnPipeline) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnPipeline) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnPipeline) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnPipeline) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnPipeline) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnPipeline) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Represents information about an action declaration.
//
// TODO: EXAMPLE
//
type CfnPipeline_ActionDeclarationProperty struct {
	// Specifies the action type and the provider of the action.
	ActionTypeId interface{} `json:"actionTypeId" yaml:"actionTypeId"`
	// The action declaration's name.
	Name *string `json:"name" yaml:"name"`
	// The action's configuration.
	//
	// These are key-value pairs that specify input values for an action. For more information, see [Action Structure Requirements in CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements) . For the list of configuration properties for the AWS CloudFormation action type in CodePipeline, see [Configuration Properties Reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-action-reference.html) in the *AWS CloudFormation User Guide* . For template snippets with examples, see [Using Parameter Override Functions with CodePipeline Pipelines](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-parameter-override-functions.html) in the *AWS CloudFormation User Guide* .
	//
	// The values can be represented in either JSON or YAML format. For example, the JSON configuration item format is as follows:
	//
	// *JSON:*
	//
	// `"Configuration" : { Key : Value },`
	Configuration interface{} `json:"configuration" yaml:"configuration"`
	// The name or ID of the artifact consumed by the action, such as a test or build artifact.
	//
	// > For a CodeBuild action with multiple input artifacts, one of your input sources must be designated the PrimarySource. For more information, see the [CodeBuild action reference page](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeBuild.html) in the *AWS CodePipeline User Guide* .
	InputArtifacts interface{} `json:"inputArtifacts" yaml:"inputArtifacts"`
	// The variable namespace associated with the action.
	//
	// All variables produced as output by this action fall under this namespace.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// The name or ID of the result of the action declaration, such as a test or build artifact.
	OutputArtifacts interface{} `json:"outputArtifacts" yaml:"outputArtifacts"`
	// The action declaration's AWS Region, such as us-east-1.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the IAM service role that performs the declared action.
	//
	// This is assumed through the roleArn for the pipeline.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The order in which actions are run.
	RunOrder *float64 `json:"runOrder" yaml:"runOrder"`
}

// Represents information about an action type.
//
// TODO: EXAMPLE
//
type CfnPipeline_ActionTypeIdProperty struct {
	// A category defines what kind of action can be taken in the stage, and constrains the provider type for the action.
	//
	// Valid categories are limited to one of the values below.
	//
	// - `Source`
	// - `Build`
	// - `Test`
	// - `Deploy`
	// - `Invoke`
	// - `Approval`
	Category *string `json:"category" yaml:"category"`
	// The creator of the action being called.
	//
	// There are three valid values for the `Owner` field in the action category section within your pipeline structure: `AWS` , `ThirdParty` , and `Custom` . For more information, see [Valid Action Types and Providers in CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#actions-valid-providers) .
	Owner *string `json:"owner" yaml:"owner"`
	// The provider of the service being called by the action.
	//
	// Valid providers are determined by the action category. For example, an action in the Deploy category type might have a provider of CodeDeploy, which would be specified as `CodeDeploy` . For more information, see [Valid Action Types and Providers in CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#actions-valid-providers) .
	Provider *string `json:"provider" yaml:"provider"`
	// A string that describes the action version.
	Version *string `json:"version" yaml:"version"`
}

// A mapping of `artifactStore` objects and their corresponding AWS Regions.
//
// There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.
//
// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
//
// TODO: EXAMPLE
//
type CfnPipeline_ArtifactStoreMapProperty struct {
	// Represents information about the S3 bucket where artifacts are stored for the pipeline.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	ArtifactStore interface{} `json:"artifactStore" yaml:"artifactStore"`
	// The action declaration's AWS Region, such as us-east-1.
	Region *string `json:"region" yaml:"region"`
}

// The S3 bucket where artifacts for the pipeline are stored.
//
// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
//
// TODO: EXAMPLE
//
type CfnPipeline_ArtifactStoreProperty struct {
	// The S3 bucket used for storing the artifacts for a pipeline.
	//
	// You can specify the name of an S3 bucket but not a folder in the bucket. A folder to contain the pipeline artifacts is created for you based on the name of the pipeline. You can use any S3 bucket in the same AWS Region as the pipeline to store your pipeline artifacts.
	Location *string `json:"location" yaml:"location"`
	// The type of the artifact store, such as S3.
	Type *string `json:"type" yaml:"type"`
	// The encryption key used to encrypt the data in the artifact store, such as an AWS Key Management Service ( AWS KMS) key.
	//
	// If this is undefined, the default key for Amazon S3 is used. To see an example artifact store encryption key field, see the example structure here: [AWS::CodePipeline::Pipeline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html) .
	EncryptionKey interface{} `json:"encryptionKey" yaml:"encryptionKey"`
}

// Reserved for future use.
//
// TODO: EXAMPLE
//
type CfnPipeline_BlockerDeclarationProperty struct {
	// Reserved for future use.
	Name *string `json:"name" yaml:"name"`
	// Reserved for future use.
	Type *string `json:"type" yaml:"type"`
}

// Represents information about the key used to encrypt data in the artifact store, such as an AWS Key Management Service ( AWS KMS) key.
//
// `EncryptionKey` is a property of the [ArtifactStore](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html) property type.
//
// TODO: EXAMPLE
//
type CfnPipeline_EncryptionKeyProperty struct {
	// The ID used to identify the key.
	//
	// For an AWS KMS key, you can use the key ID, the key ARN, or the alias ARN.
	//
	// > Aliases are recognized only in the account that created the AWS KMS key. For cross-account actions, you can only use the key ID or key ARN to identify the key.
	Id *string `json:"id" yaml:"id"`
	// The type of encryption key, such as an AWS KMS key.
	//
	// When creating or updating a pipeline, the value must be set to 'KMS'.
	Type *string `json:"type" yaml:"type"`
}

// Represents information about an artifact to be worked on, such as a test or build artifact.
//
// TODO: EXAMPLE
//
type CfnPipeline_InputArtifactProperty struct {
	// The name of the artifact to be worked on (for example, "My App").
	//
	// The input artifact of an action must exactly match the output artifact declared in a preceding action, but the input artifact does not have to be the next action in strict sequence from the action that provided the output artifact. Actions in parallel can declare different output artifacts, which are in turn consumed by different following actions.
	Name *string `json:"name" yaml:"name"`
}

// Represents information about the output of an action.
//
// TODO: EXAMPLE
//
type CfnPipeline_OutputArtifactProperty struct {
	// The name of the output of an artifact, such as "My App".
	//
	// The output artifact name must exactly match the input artifact declared for a downstream action. However, the downstream action's input artifact does not have to be the next action in strict sequence from the action that provided the output artifact. Actions in parallel can declare different output artifacts, which are in turn consumed by different following actions.
	//
	// Output artifact names must be unique within a pipeline.
	Name *string `json:"name" yaml:"name"`
}

// Represents information about a stage and its definition.
//
// TODO: EXAMPLE
//
type CfnPipeline_StageDeclarationProperty struct {
	// The actions included in a stage.
	Actions interface{} `json:"actions" yaml:"actions"`
	// The name of the stage.
	Name *string `json:"name" yaml:"name"`
	// Reserved for future use.
	Blockers interface{} `json:"blockers" yaml:"blockers"`
}

// The name of the pipeline in which you want to disable the flow of artifacts from one stage to another.
//
// TODO: EXAMPLE
//
type CfnPipeline_StageTransitionProperty struct {
	// The reason given to the user that a stage is disabled, such as waiting for manual approval or manual tests.
	//
	// This message is displayed in the pipeline console UI.
	Reason *string `json:"reason" yaml:"reason"`
	// The name of the stage where you want to disable the inbound or outbound transition of artifacts.
	StageName *string `json:"stageName" yaml:"stageName"`
}

// Properties for defining a `CfnPipeline`.
//
// TODO: EXAMPLE
//
type CfnPipelineProps struct {
	// The Amazon Resource Name (ARN) for CodePipeline to use to either perform actions with no `actionRoleArn` , or to use to assume roles for actions with an `actionRoleArn` .
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Represents information about a stage and its definition.
	Stages interface{} `json:"stages" yaml:"stages"`
	// The S3 bucket where artifacts for the pipeline are stored.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	ArtifactStore interface{} `json:"artifactStore" yaml:"artifactStore"`
	// A mapping of `artifactStore` objects and their corresponding AWS Regions.
	//
	// There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	ArtifactStores interface{} `json:"artifactStores" yaml:"artifactStores"`
	// Represents the input of a `DisableStageTransition` action.
	DisableInboundStageTransitions interface{} `json:"disableInboundStageTransitions" yaml:"disableInboundStageTransitions"`
	// The name of the pipeline.
	Name *string `json:"name" yaml:"name"`
	// Indicates whether to rerun the CodePipeline pipeline after you update it.
	RestartExecutionOnUpdate interface{} `json:"restartExecutionOnUpdate" yaml:"restartExecutionOnUpdate"`
	// Specifies the tags applied to the pipeline.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::CodePipeline::Webhook`.
//
// The `AWS::CodePipeline::Webhook` resource creates and registers your webhook. After the webhook is created and registered, it triggers your pipeline to start every time an external event occurs. For more information, see [Configure Your GitHub Pipelines to Use Webhooks for Change Detection](https://docs.aws.amazon.com/codepipeline/latest/userguide/pipelines-webhooks-migration.html) in the *AWS CodePipeline User Guide* .
//
// We strongly recommend that you use AWS Secrets Manager to store your credentials. If you use Secrets Manager, you must have already configured and stored your secret parameters in Secrets Manager. For more information, see [Using Dynamic References to Specify Template Values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) .
//
// > When passing secret parameters, do not enter the value directly into the template. The value is rendered as plaintext and is therefore readable. For security reasons, do not use plaintext in your AWS CloudFormation template to store your credentials.
//
// TODO: EXAMPLE
//
type CfnWebhook interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrUrl() *string
	Authentication() *string
	SetAuthentication(val *string)
	AuthenticationConfiguration() interface{}
	SetAuthenticationConfiguration(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Filters() interface{}
	SetFilters(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	RegisterWithThirdParty() interface{}
	SetRegisterWithThirdParty(val interface{})
	Stack() awscdk.Stack
	TargetAction() *string
	SetTargetAction(val *string)
	TargetPipeline() *string
	SetTargetPipeline(val *string)
	TargetPipelineVersion() *float64
	SetTargetPipelineVersion(val *float64)
	UpdatedProperites() *map[string]interface{}
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnWebhook
type jsiiProxy_CfnWebhook struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWebhook) AttrUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) Authentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) AuthenticationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) Filters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) RegisterWithThirdParty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registerWithThirdParty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) TargetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) TargetPipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) TargetPipelineVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPipelineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebhook) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CodePipeline::Webhook`.
func NewCfnWebhook(scope awscdk.Construct, id *string, props *CfnWebhookProps) CfnWebhook {
	_init_.Initialize()

	j := jsiiProxy_CfnWebhook{}

	_jsii_.Create(
		"monocdk.aws_codepipeline.CfnWebhook",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodePipeline::Webhook`.
func NewCfnWebhook_Override(c CfnWebhook, scope awscdk.Construct, id *string, props *CfnWebhookProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline.CfnWebhook",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWebhook) SetAuthentication(val *string) {
	_jsii_.Set(
		j,
		"authentication",
		val,
	)
}

func (j *jsiiProxy_CfnWebhook) SetAuthenticationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"authenticationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnWebhook) SetFilters(val interface{}) {
	_jsii_.Set(
		j,
		"filters",
		val,
	)
}

func (j *jsiiProxy_CfnWebhook) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnWebhook) SetRegisterWithThirdParty(val interface{}) {
	_jsii_.Set(
		j,
		"registerWithThirdParty",
		val,
	)
}

func (j *jsiiProxy_CfnWebhook) SetTargetAction(val *string) {
	_jsii_.Set(
		j,
		"targetAction",
		val,
	)
}

func (j *jsiiProxy_CfnWebhook) SetTargetPipeline(val *string) {
	_jsii_.Set(
		j,
		"targetPipeline",
		val,
	)
}

func (j *jsiiProxy_CfnWebhook) SetTargetPipelineVersion(val *float64) {
	_jsii_.Set(
		j,
		"targetPipelineVersion",
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
func CfnWebhook_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.CfnWebhook",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnWebhook_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.CfnWebhook",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnWebhook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.CfnWebhook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebhook_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codepipeline.CfnWebhook",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnWebhook) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnWebhook) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnWebhook) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnWebhook) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnWebhook) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnWebhook) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnWebhook) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnWebhook) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnWebhook) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnWebhook) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnWebhook) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnWebhook) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnWebhook) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnWebhook) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnWebhook) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnWebhook) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnWebhook) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnWebhook) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnWebhook) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnWebhook) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnWebhook) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The authentication applied to incoming webhook trigger requests.
//
// TODO: EXAMPLE
//
type CfnWebhook_WebhookAuthConfigurationProperty struct {
	// The property used to configure acceptance of webhooks in an IP address range.
	//
	// For IP, only the `AllowedIPRange` property must be set. This property must be set to a valid CIDR range.
	AllowedIpRange *string `json:"allowedIpRange" yaml:"allowedIpRange"`
	// The property used to configure GitHub authentication.
	//
	// For GITHUB_HMAC, only the `SecretToken` property must be set.
	SecretToken *string `json:"secretToken" yaml:"secretToken"`
}

// The event criteria that specify when a webhook notification is sent to your URL.
//
// TODO: EXAMPLE
//
type CfnWebhook_WebhookFilterRuleProperty struct {
	// A JsonPath expression that is applied to the body/payload of the webhook.
	//
	// The value selected by the JsonPath expression must match the value specified in the `MatchEquals` field. Otherwise, the request is ignored. For more information, see [Java JsonPath implementation](https://docs.aws.amazon.com/https://github.com/json-path/JsonPath) in GitHub.
	JsonPath *string `json:"jsonPath" yaml:"jsonPath"`
	// The value selected by the `JsonPath` expression must match what is supplied in the `MatchEquals` field.
	//
	// Otherwise, the request is ignored. Properties from the target action configuration can be included as placeholders in this value by surrounding the action configuration key with curly brackets. For example, if the value supplied here is "refs/heads/{Branch}" and the target action has an action configuration property called "Branch" with a value of "main", the `MatchEquals` value is evaluated as "refs/heads/main". For a list of action configuration properties for built-in action types, see [Pipeline Structure Reference Action Requirements](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements) .
	MatchEquals *string `json:"matchEquals" yaml:"matchEquals"`
}

// Properties for defining a `CfnWebhook`.
//
// TODO: EXAMPLE
//
type CfnWebhookProps struct {
	// Supported options are GITHUB_HMAC, IP, and UNAUTHENTICATED.
	//
	// - For information about the authentication scheme implemented by GITHUB_HMAC, see [Securing your webhooks](https://docs.aws.amazon.com/https://developer.github.com/webhooks/securing/) on the GitHub Developer website.
	// - IP rejects webhooks trigger requests unless they originate from an IP address in the IP range whitelisted in the authentication configuration.
	// - UNAUTHENTICATED accepts all webhook trigger requests regardless of origin.
	Authentication *string `json:"authentication" yaml:"authentication"`
	// Properties that configure the authentication applied to incoming webhook trigger requests.
	//
	// The required properties depend on the authentication type. For GITHUB_HMAC, only the `SecretToken` property must be set. For IP, only the `AllowedIPRange` property must be set to a valid CIDR range. For UNAUTHENTICATED, no properties can be set.
	AuthenticationConfiguration interface{} `json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// A list of rules applied to the body/payload sent in the POST request to a webhook URL.
	//
	// All defined rules must pass for the request to be accepted and the pipeline started.
	Filters interface{} `json:"filters" yaml:"filters"`
	// The name of the action in a pipeline you want to connect to the webhook.
	//
	// The action must be from the source (first) stage of the pipeline.
	TargetAction *string `json:"targetAction" yaml:"targetAction"`
	// The name of the pipeline you want to connect to the webhook.
	TargetPipeline *string `json:"targetPipeline" yaml:"targetPipeline"`
	// The version number of the pipeline to be connected to the trigger request.
	//
	// Required: Yes
	//
	// Type: Integer
	//
	// Update requires: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)
	TargetPipelineVersion *float64 `json:"targetPipelineVersion" yaml:"targetPipelineVersion"`
	// The name of the webhook.
	Name *string `json:"name" yaml:"name"`
	// Configures a connection between the webhook that was created and the external tool with events to be detected.
	RegisterWithThirdParty interface{} `json:"registerWithThirdParty" yaml:"registerWithThirdParty"`
}

// Common properties shared by all Actions.
//
// TODO: EXAMPLE
//
// Experimental.
type CommonActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace" yaml:"variablesNamespace"`
}

// Common properties shared by all Actions whose {@link ActionProperties.owner} field is 'AWS' (or unset, as 'AWS' is the default).
//
// TODO: EXAMPLE
//
// Experimental.
type CommonAwsActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// An interface representing resources generated in order to support the cross-region capabilities of CodePipeline.
//
// You get instances of this interface from the {@link Pipeline#crossRegionSupport} property.
//
// TODO: EXAMPLE
//
// Experimental.
type CrossRegionSupport struct {
	// The replication Bucket used by CodePipeline to operate in this region.
	//
	// Belongs to {@link stack}.
	// Experimental.
	ReplicationBucket awss3.IBucket `json:"replicationBucket" yaml:"replicationBucket"`
	// The Stack that has been created to house the replication Bucket required for this  region.
	// Experimental.
	Stack awscdk.Stack `json:"stack" yaml:"stack"`
}

// The creation attributes used for defining a configuration property of a custom Action.
//
// TODO: EXAMPLE
//
// Experimental.
type CustomActionProperty struct {
	// The name of the property.
	//
	// You use this name in the `configuration` attribute when defining your custom Action class.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Whether this property is required.
	// Experimental.
	Required *bool `json:"required" yaml:"required"`
	// The description of the property.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Whether this property is a key.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-key
	//
	// Experimental.
	Key *bool `json:"key" yaml:"key"`
	// Whether this property is queryable.
	//
	// Note that only a single property of a custom Action can be queryable.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-queryable
	//
	// Experimental.
	Queryable *bool `json:"queryable" yaml:"queryable"`
	// Whether this property is secret, like a password, or access key.
	// Experimental.
	Secret *bool `json:"secret" yaml:"secret"`
	// The type of the property, like 'String', 'Number', or 'Boolean'.
	// Experimental.
	Type *string `json:"type" yaml:"type"`
}

// The resource representing registering a custom Action with CodePipeline.
//
// For the Action to be usable, it has to be registered for every region and every account it's used in.
// In addition to this class, you should most likely also provide your clients a class
// representing your custom Action, extending the Action class,
// and taking the `actionProperties` as properly typed, construction properties.
//
// TODO: EXAMPLE
//
// Experimental.
type CustomActionRegistration interface {
	awscdk.Construct
	Node() awscdk.ConstructNode
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for CustomActionRegistration
type jsiiProxy_CustomActionRegistration struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_CustomActionRegistration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomActionRegistration(scope constructs.Construct, id *string, props *CustomActionRegistrationProps) CustomActionRegistration {
	_init_.Initialize()

	j := jsiiProxy_CustomActionRegistration{}

	_jsii_.Create(
		"monocdk.aws_codepipeline.CustomActionRegistration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomActionRegistration_Override(c CustomActionRegistration, scope constructs.Construct, id *string, props *CustomActionRegistrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline.CustomActionRegistration",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func CustomActionRegistration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.CustomActionRegistration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CustomActionRegistration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CustomActionRegistration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CustomActionRegistration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CustomActionRegistration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CustomActionRegistration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CustomActionRegistration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CustomActionRegistration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of registering a custom Action.
//
// TODO: EXAMPLE
//
// Experimental.
type CustomActionRegistrationProps struct {
	// The artifact bounds of the Action.
	// Experimental.
	ArtifactBounds *ActionArtifactBounds `json:"artifactBounds" yaml:"artifactBounds"`
	// The category of the Action.
	// Experimental.
	Category ActionCategory `json:"category" yaml:"category"`
	// The provider of the Action.
	//
	// For example, `'MyCustomActionProvider'`
	// Experimental.
	Provider *string `json:"provider" yaml:"provider"`
	// The properties used for customizing the instance of your Action.
	// Experimental.
	ActionProperties *[]*CustomActionProperty `json:"actionProperties" yaml:"actionProperties"`
	// The URL shown for the entire Action in the Pipeline UI.
	// Experimental.
	EntityUrl *string `json:"entityUrl" yaml:"entityUrl"`
	// The URL shown for a particular execution of an Action in the Pipeline UI.
	// Experimental.
	ExecutionUrl *string `json:"executionUrl" yaml:"executionUrl"`
	// The version of your Action.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
}

// The CodePipeline variables that are global, not bound to a specific action.
//
// This class defines a bunch of static fields that represent the different variables.
// These can be used can be used in any action configuration.
//
// TODO: EXAMPLE
//
// Experimental.
type GlobalVariables interface {
}

// The jsii proxy struct for GlobalVariables
type jsiiProxy_GlobalVariables struct {
	_ byte // padding
}

// Experimental.
func NewGlobalVariables() GlobalVariables {
	_init_.Initialize()

	j := jsiiProxy_GlobalVariables{}

	_jsii_.Create(
		"monocdk.aws_codepipeline.GlobalVariables",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewGlobalVariables_Override(g GlobalVariables) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline.GlobalVariables",
		nil, // no parameters
		g,
	)
}

func GlobalVariables_ExecutionId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codepipeline.GlobalVariables",
		"executionId",
		&returns,
	)
	return returns
}

// A Pipeline Action.
//
// If you want to implement this interface,
// consider extending the {@link Action} class,
// which contains some common logic.
// Experimental.
type IAction interface {
	// The callback invoked when this Action is added to a Pipeline.
	// Experimental.
	Bind(scope awscdk.Construct, stage IStage, options *ActionBindOptions) *ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	// Experimental.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	// Experimental.
	ActionProperties() *ActionProperties
}

// The jsii proxy for IAction
type jsiiProxy_IAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IAction) Bind(scope awscdk.Construct, stage IStage, options *ActionBindOptions) *ActionConfig {
	var returns *ActionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAction) ActionProperties() *ActionProperties {
	var returns *ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

// The abstract view of an AWS CodePipeline as required and used by Actions.
//
// It extends {@link events.IRuleTarget},
// so this interface can be used as a Target for CloudWatch Events.
// Experimental.
type IPipeline interface {
	awscodestarnotifications.INotificationRuleSource
	awscdk.IResource
	// Defines a CodeStar notification rule triggered when the pipeline events emitted by you specified, it very similar to `onEvent` API.
	//
	// You can also use the methods `notifyOnExecutionStateChange`, `notifyOnAnyStageStateChange`,
	// `notifyOnAnyActionStateChange` and `notifyOnAnyManualApprovalStateChange`
	// to define rules for these specific event emitted.
	//
	// Returns: CodeStar notification rule associated with this build project.
	// Experimental.
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Action execution" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	// Experimental.
	NotifyOnAnyActionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Manual approval" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	// Experimental.
	NotifyOnAnyManualApprovalStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Stage execution" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	// Experimental.
	NotifyOnAnyStageStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Pipeline execution" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	// Experimental.
	NotifyOnExecutionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an event rule triggered by this CodePipeline.
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Define an event rule triggered by the "CodePipeline Pipeline Execution State Change" event emitted from this pipeline.
	// Experimental.
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The ARN of the Pipeline.
	// Experimental.
	PipelineArn() *string
	// The name of the Pipeline.
	// Experimental.
	PipelineName() *string
}

// The jsii proxy for IPipeline
type jsiiProxy_IPipeline struct {
	internal.Type__awscodestarnotificationsINotificationRuleSource
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPipeline) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOn",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) NotifyOnAnyActionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnAnyActionStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) NotifyOnAnyManualApprovalStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnAnyManualApprovalStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) NotifyOnAnyStageStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnAnyStageStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) NotifyOnExecutionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnExecutionStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IPipeline) BindAsNotificationRuleSource(scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
	var returns *awscodestarnotifications.NotificationRuleSourceConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleSource",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPipeline) PipelineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) PipelineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// The abstract interface of a Pipeline Stage that is used by Actions.
// Experimental.
type IStage interface {
	// Experimental.
	AddAction(action IAction)
	// Experimental.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// The actions belonging to this stage.
	// Experimental.
	Actions() *[]IAction
	// Experimental.
	Pipeline() IPipeline
	// The physical, human-readable name of this Pipeline Stage.
	// Experimental.
	StageName() *string
}

// The jsii proxy for IStage
type jsiiProxy_IStage struct {
	_ byte // padding
}

func (i *jsiiProxy_IStage) AddAction(action IAction) {
	_jsii_.InvokeVoid(
		i,
		"addAction",
		[]interface{}{action},
	)
}

func (i *jsiiProxy_IStage) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IStage) Actions() *[]IAction {
	var returns *[]IAction
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStage) Pipeline() IPipeline {
	var returns IPipeline
	_jsii_.Get(
		j,
		"pipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

// An AWS CodePipeline pipeline with its associated IAM role and S3 bucket.
//
// TODO: EXAMPLE
//
// Experimental.
type Pipeline interface {
	awscdk.Resource
	IPipeline
	ArtifactBucket() awss3.IBucket
	CrossRegionSupport() *map[string]*CrossRegionSupport
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	PipelineArn() *string
	PipelineName() *string
	PipelineVersion() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	StageCount() *float64
	Stages() *[]IStage
	AddStage(props *StageOptions) IStage
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) awscodestarnotifications.INotificationRule
	NotifyOnAnyActionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	NotifyOnAnyManualApprovalStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	NotifyOnAnyStageStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	NotifyOnExecutionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnPrepare()
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Stage(stageName *string) IStage
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for Pipeline
type jsiiProxy_Pipeline struct {
	internal.Type__awscdkResource
	jsiiProxy_IPipeline
}

func (j *jsiiProxy_Pipeline) ArtifactBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"artifactBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) CrossRegionSupport() *map[string]*CrossRegionSupport {
	var returns *map[string]*CrossRegionSupport
	_jsii_.Get(
		j,
		"crossRegionSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) PipelineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) PipelineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) PipelineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) StageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Stages() *[]IStage {
	var returns *[]IStage
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}


// Experimental.
func NewPipeline(scope constructs.Construct, id *string, props *PipelineProps) Pipeline {
	_init_.Initialize()

	j := jsiiProxy_Pipeline{}

	_jsii_.Create(
		"monocdk.aws_codepipeline.Pipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPipeline_Override(p Pipeline, scope constructs.Construct, id *string, props *PipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline.Pipeline",
		[]interface{}{scope, id, props},
		p,
	)
}

// Import a pipeline into this app.
// Experimental.
func Pipeline_FromPipelineArn(scope constructs.Construct, id *string, pipelineArn *string) IPipeline {
	_init_.Initialize()

	var returns IPipeline

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.Pipeline",
		"fromPipelineArn",
		[]interface{}{scope, id, pipelineArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Pipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.Pipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Pipeline_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.Pipeline",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Creates a new Stage, and adds it to this Pipeline.
//
// Returns: the newly created Stage
// Experimental.
func (p *jsiiProxy_Pipeline) AddStage(props *StageOptions) IStage {
	var returns IStage

	_jsii_.Invoke(
		p,
		"addStage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Adds a statement to the pipeline role.
// Experimental.
func (p *jsiiProxy_Pipeline) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		p,
		"addToRolePolicy",
		[]interface{}{statement},
	)
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
// Experimental.
func (p *jsiiProxy_Pipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Returns a source configuration for notification rule.
// Experimental.
func (p *jsiiProxy_Pipeline) BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
	var returns *awscodestarnotifications.NotificationRuleSourceConfig

	_jsii_.Invoke(
		p,
		"bindAsNotificationRuleSource",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Experimental.
func (p *jsiiProxy_Pipeline) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_Pipeline) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_Pipeline) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Defines a CodeStar notification rule triggered when the pipeline events emitted by you specified, it very similar to `onEvent` API.
//
// You can also use the methods `notifyOnExecutionStateChange`, `notifyOnAnyStageStateChange`,
// `notifyOnAnyActionStateChange` and `notifyOnAnyManualApprovalStateChange`
// to define rules for these specific event emitted.
// Experimental.
func (p *jsiiProxy_Pipeline) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOn",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

// Define an notification rule triggered by the set of the "Action execution" events emitted from this pipeline.
// Experimental.
func (p *jsiiProxy_Pipeline) NotifyOnAnyActionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnAnyActionStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

// Define an notification rule triggered by the set of the "Manual approval" events emitted from this pipeline.
// Experimental.
func (p *jsiiProxy_Pipeline) NotifyOnAnyManualApprovalStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnAnyManualApprovalStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

// Define an notification rule triggered by the set of the "Stage execution" events emitted from this pipeline.
// Experimental.
func (p *jsiiProxy_Pipeline) NotifyOnAnyStageStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnAnyStageStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

// Define an notification rule triggered by the set of the "Pipeline execution" events emitted from this pipeline.
// Experimental.
func (p *jsiiProxy_Pipeline) NotifyOnExecutionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnExecutionStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

// Defines an event rule triggered by this CodePipeline.
// Experimental.
func (p *jsiiProxy_Pipeline) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (p *jsiiProxy_Pipeline) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

// Defines an event rule triggered by the "CodePipeline Pipeline Execution State Change" event emitted from this pipeline.
// Experimental.
func (p *jsiiProxy_Pipeline) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (p *jsiiProxy_Pipeline) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (p *jsiiProxy_Pipeline) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (p *jsiiProxy_Pipeline) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

// Access one of the pipeline's stages by stage name.
// Experimental.
func (p *jsiiProxy_Pipeline) Stage(stageName *string) IStage {
	var returns IStage

	_jsii_.Invoke(
		p,
		"stage",
		[]interface{}{stageName},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (p *jsiiProxy_Pipeline) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (p *jsiiProxy_Pipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the pipeline structure.
//
// Validation happens according to the rules documented at
//
// https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#pipeline-requirements
// Experimental.
func (p *jsiiProxy_Pipeline) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The list of event types for AWS Codepipeline Pipeline.
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
//
// Experimental.
type PipelineNotificationEvents string

const (
	PipelineNotificationEvents_PIPELINE_EXECUTION_FAILED PipelineNotificationEvents = "PIPELINE_EXECUTION_FAILED"
	PipelineNotificationEvents_PIPELINE_EXECUTION_CANCELED PipelineNotificationEvents = "PIPELINE_EXECUTION_CANCELED"
	PipelineNotificationEvents_PIPELINE_EXECUTION_STARTED PipelineNotificationEvents = "PIPELINE_EXECUTION_STARTED"
	PipelineNotificationEvents_PIPELINE_EXECUTION_RESUMED PipelineNotificationEvents = "PIPELINE_EXECUTION_RESUMED"
	PipelineNotificationEvents_PIPELINE_EXECUTION_SUCCEEDED PipelineNotificationEvents = "PIPELINE_EXECUTION_SUCCEEDED"
	PipelineNotificationEvents_PIPELINE_EXECUTION_SUPERSEDED PipelineNotificationEvents = "PIPELINE_EXECUTION_SUPERSEDED"
	PipelineNotificationEvents_STAGE_EXECUTION_STARTED PipelineNotificationEvents = "STAGE_EXECUTION_STARTED"
	PipelineNotificationEvents_STAGE_EXECUTION_SUCCEEDED PipelineNotificationEvents = "STAGE_EXECUTION_SUCCEEDED"
	PipelineNotificationEvents_STAGE_EXECUTION_RESUMED PipelineNotificationEvents = "STAGE_EXECUTION_RESUMED"
	PipelineNotificationEvents_STAGE_EXECUTION_CANCELED PipelineNotificationEvents = "STAGE_EXECUTION_CANCELED"
	PipelineNotificationEvents_STAGE_EXECUTION_FAILED PipelineNotificationEvents = "STAGE_EXECUTION_FAILED"
	PipelineNotificationEvents_ACTION_EXECUTION_SUCCEEDED PipelineNotificationEvents = "ACTION_EXECUTION_SUCCEEDED"
	PipelineNotificationEvents_ACTION_EXECUTION_FAILED PipelineNotificationEvents = "ACTION_EXECUTION_FAILED"
	PipelineNotificationEvents_ACTION_EXECUTION_CANCELED PipelineNotificationEvents = "ACTION_EXECUTION_CANCELED"
	PipelineNotificationEvents_ACTION_EXECUTION_STARTED PipelineNotificationEvents = "ACTION_EXECUTION_STARTED"
	PipelineNotificationEvents_MANUAL_APPROVAL_FAILED PipelineNotificationEvents = "MANUAL_APPROVAL_FAILED"
	PipelineNotificationEvents_MANUAL_APPROVAL_NEEDED PipelineNotificationEvents = "MANUAL_APPROVAL_NEEDED"
	PipelineNotificationEvents_MANUAL_APPROVAL_SUCCEEDED PipelineNotificationEvents = "MANUAL_APPROVAL_SUCCEEDED"
)

// Additional options to pass to the notification rule.
//
// TODO: EXAMPLE
//
// Experimental.
type PipelineNotifyOnOptions struct {
	// The level of detail to include in the notifications for this resource.
	//
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	// Experimental.
	DetailType awscodestarnotifications.DetailType `json:"detailType" yaml:"detailType"`
	// The status of the notification rule.
	//
	// If the enabled is set to DISABLED, notifications aren't sent for the notification rule.
	// Experimental.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account.
	// Experimental.
	NotificationRuleName *string `json:"notificationRuleName" yaml:"notificationRuleName"`
	// A list of event types associated with this notification rule for CodePipeline Pipeline.
	//
	// For a complete list of event types and IDs, see Notification concepts in the Developer Tools Console User Guide.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#concepts-api
	//
	// Experimental.
	Events *[]PipelineNotificationEvents `json:"events" yaml:"events"`
}

// TODO: EXAMPLE
//
// Experimental.
type PipelineProps struct {
	// The S3 bucket used by this Pipeline to store artifacts.
	// Experimental.
	ArtifactBucket awss3.IBucket `json:"artifactBucket" yaml:"artifactBucket"`
	// Create KMS keys for cross-account deployments.
	//
	// This controls whether the pipeline is enabled for cross-account deployments.
	//
	// By default cross-account deployments are enabled, but this feature requires
	// that KMS Customer Master Keys are created which have a cost of $1/month.
	//
	// If you do not need cross-account deployments, you can set this to `false` to
	// not create those keys and save on that cost (the artifact bucket will be
	// encrypted with an AWS-managed key). However, cross-account deployments will
	// no longer be possible.
	// Experimental.
	CrossAccountKeys *bool `json:"crossAccountKeys" yaml:"crossAccountKeys"`
	// A map of region to S3 bucket name used for cross-region CodePipeline.
	//
	// For every Action that you specify targeting a different region than the Pipeline itself,
	// if you don't provide an explicit Bucket for that region using this property,
	// the construct will automatically create a Stack containing an S3 Bucket in that region.
	// Experimental.
	CrossRegionReplicationBuckets *map[string]awss3.IBucket `json:"crossRegionReplicationBuckets" yaml:"crossRegionReplicationBuckets"`
	// Enable KMS key rotation for the generated KMS keys.
	//
	// By default KMS key rotation is disabled, but will add an additional $1/month
	// for each year the key exists when enabled.
	// Experimental.
	EnableKeyRotation *bool `json:"enableKeyRotation" yaml:"enableKeyRotation"`
	// Name of the pipeline.
	// Experimental.
	PipelineName *string `json:"pipelineName" yaml:"pipelineName"`
	// Indicates whether to rerun the AWS CodePipeline pipeline after you update it.
	// Experimental.
	RestartExecutionOnUpdate *bool `json:"restartExecutionOnUpdate" yaml:"restartExecutionOnUpdate"`
	// Reuse the same cross region support stack for all pipelines in the App.
	// Experimental.
	ReuseCrossRegionSupportStacks *bool `json:"reuseCrossRegionSupportStacks" yaml:"reuseCrossRegionSupportStacks"`
	// The IAM role to be assumed by this Pipeline.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// The list of Stages, in order, to create this Pipeline with.
	//
	// You can always add more Stages later by calling {@link Pipeline#addStage}.
	// Experimental.
	Stages *[]*StageProps `json:"stages" yaml:"stages"`
}

// TODO: EXAMPLE
//
// Experimental.
type StageOptions struct {
	// The physical, human-readable name to assign to this Pipeline Stage.
	// Experimental.
	StageName *string `json:"stageName" yaml:"stageName"`
	// The list of Actions to create this Stage with.
	//
	// You can always add more Actions later by calling {@link IStage#addAction}.
	// Experimental.
	Actions *[]IAction `json:"actions" yaml:"actions"`
	// Experimental.
	Placement *StagePlacement `json:"placement" yaml:"placement"`
}

// Allows you to control where to place a new Stage when it's added to the Pipeline.
//
// Note that you can provide only one of the below properties -
// specifying more than one will result in a validation error.
//
// TODO: EXAMPLE
//
// See: #justAfter
//
// Experimental.
type StagePlacement struct {
	// Inserts the new Stage as a child of the given Stage (changing its current child Stage, if it had one).
	// Experimental.
	JustAfter IStage `json:"justAfter" yaml:"justAfter"`
	// Inserts the new Stage as a parent of the given Stage (changing its current parent Stage, if it had one).
	// Experimental.
	RightBefore IStage `json:"rightBefore" yaml:"rightBefore"`
}

// Construction properties of a Pipeline Stage.
//
// TODO: EXAMPLE
//
// Experimental.
type StageProps struct {
	// The physical, human-readable name to assign to this Pipeline Stage.
	// Experimental.
	StageName *string `json:"stageName" yaml:"stageName"`
	// The list of Actions to create this Stage with.
	//
	// You can always add more Actions later by calling {@link IStage#addAction}.
	// Experimental.
	Actions *[]IAction `json:"actions" yaml:"actions"`
}

