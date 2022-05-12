package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Low-level class for generic CodePipeline Actions implementing the {@link IAction} interface.
//
// Contains some common logic that can be re-used by all {@link IAction} implementations.
// If you're writing your own Action class,
// feel free to extend this class.
type Action interface {
	IAction
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *ActionProperties
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
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


func NewAction_Override(a Action) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Action",
		nil, // no parameters
		a,
	)
}

func (a *jsiiProxy_Action) Bind(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig {
	var returns *ActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Action) Bound(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig {
	var returns *ActionConfig

	_jsii_.Invoke(
		a,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

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
// Example:
//   // MyAction is some action type that produces variables, like EcrSourceAction
//   myAction := NewMyAction(&myActionProps{
//   	// ...
//   	actionName: jsii.String("myAction"),
//   })
//   NewOtherAction(&otherActionProps{
//   	// ...
//   	config: myAction.variables.myVariable,
//   	actionName: jsii.String("otherAction"),
//   })
//
type ActionArtifactBounds struct {
	MaxInputs *float64 `field:"required" json:"maxInputs" yaml:"maxInputs"`
	MaxOutputs *float64 `field:"required" json:"maxOutputs" yaml:"maxOutputs"`
	MinInputs *float64 `field:"required" json:"minInputs" yaml:"minInputs"`
	MinOutputs *float64 `field:"required" json:"minOutputs" yaml:"minOutputs"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var role role
//
//   actionBindOptions := &actionBindOptions{
//   	bucket: bucket,
//   	role: role,
//   }
//
type ActionBindOptions struct {
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
}

// Example:
//   // MyAction is some action type that produces variables, like EcrSourceAction
//   myAction := NewMyAction(&myActionProps{
//   	// ...
//   	actionName: jsii.String("myAction"),
//   })
//   NewOtherAction(&otherActionProps{
//   	// ...
//   	config: myAction.variables.myVariable,
//   	actionName: jsii.String("otherAction"),
//   })
//
type ActionCategory string

const (
	ActionCategory_SOURCE ActionCategory = "SOURCE"
	ActionCategory_BUILD ActionCategory = "BUILD"
	ActionCategory_TEST ActionCategory = "TEST"
	ActionCategory_APPROVAL ActionCategory = "APPROVAL"
	ActionCategory_DEPLOY ActionCategory = "DEPLOY"
	ActionCategory_INVOKE ActionCategory = "INVOKE"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   actionConfig := &actionConfig{
//   	configuration: configuration,
//   }
//
type ActionConfig struct {
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var resource resource
//   var role role
//
//   actionProperties := &actionProperties{
//   	actionName: jsii.String("actionName"),
//   	artifactBounds: &actionArtifactBounds{
//   		maxInputs: jsii.Number(123),
//   		maxOutputs: jsii.Number(123),
//   		minInputs: jsii.Number(123),
//   		minOutputs: jsii.Number(123),
//   	},
//   	category: awscdk.Aws_codepipeline.actionCategory_SOURCE,
//   	provider: jsii.String("provider"),
//
//   	// the properties below are optional
//   	account: jsii.String("account"),
//   	inputs: []*artifact{
//   		artifact,
//   	},
//   	outputs: []*artifact{
//   		artifact,
//   	},
//   	owner: jsii.String("owner"),
//   	region: jsii.String("region"),
//   	resource: resource,
//   	role: role,
//   	runOrder: jsii.Number(123),
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   	version: jsii.String("version"),
//   }
//
type ActionProperties struct {
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	ArtifactBounds *ActionArtifactBounds `field:"required" json:"artifactBounds" yaml:"artifactBounds"`
	// The category of the action.
	//
	// The category defines which action type the owner
	// (the entity that performs the action) performs.
	Category ActionCategory `field:"required" json:"category" yaml:"category"`
	// The service provider that the action calls.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// The account the Action is supposed to live in.
	//
	// For Actions backed by resources,
	// this is inferred from the Stack {@link resource} is part of.
	// However, some Actions, like the CloudFormation ones,
	// are not backed by any resource, and they still might want to be cross-account.
	// In general, a concrete Action class should specify either {@link resource},
	// or {@link account} - but not both.
	Account *string `field:"optional" json:"account" yaml:"account"`
	Inputs *[]Artifact `field:"optional" json:"inputs" yaml:"inputs"`
	Outputs *[]Artifact `field:"optional" json:"outputs" yaml:"outputs"`
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the {@link PipelineProps#crossRegionReplicationBuckets} property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The optional resource that is backing this Action.
	//
	// This is used for automatically handling Actions backed by
	// resources from a different account and/or region.
	Resource awscdk.IResource `field:"optional" json:"resource" yaml:"resource"`
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The order in which AWS CodePipeline runs this action. For more information, see the AWS CodePipeline User Guide.
	//
	// https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// An output artifact of an action.
//
// Artifacts can be used as input by some actions.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // later:
//   var project pipelineProject
//   lambdaInvokeAction := codepipeline_actions.NewLambdaInvokeAction(&lambdaInvokeActionProps{
//   	actionName: jsii.String("Lambda"),
//   	lambda: lambda.NewFunction(this, jsii.String("Func"), &functionProps{
//   		runtime: lambda.runtime_NODEJS_12_X(),
//   		handler: jsii.String("index.handler"),
//   		code: lambda.code.fromInline(jsii.String("\n        const AWS = require('aws-sdk');\n\n        exports.handler = async function(event, context) {\n            const codepipeline = new AWS.CodePipeline();\n            await codepipeline.putJobSuccessResult({\n                jobId: event['CodePipeline.job'].id,\n                outputVariables: {\n                    MY_VAR: \"some value\",\n                },\n            }).promise();\n        }\n    ")),
//   	}),
//   	variablesNamespace: jsii.String("MyNamespace"),
//   })
//   sourceOutput := codepipeline.NewArtifact()
//   codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("CodeBuild"),
//   	project: project,
//   	input: sourceOutput,
//   	environmentVariables: map[string]buildEnvironmentVariable{
//   		"MyVar": &buildEnvironmentVariable{
//   			"value": lambdaInvokeAction.variable(jsii.String("MY_VAR")),
//   		},
//   	},
//   })
//
type Artifact interface {
	ArtifactName() *string
	// The artifact attribute for the name of the S3 bucket where the artifact is stored.
	BucketName() *string
	// The artifact attribute for The name of the .zip file that contains the artifact that is generated by AWS CodePipeline, such as 1ABCyZZ.zip.
	ObjectKey() *string
	// Returns the location of the .zip file in S3 that this Artifact represents. Used by Lambda's `CfnParametersCode` when being deployed in a CodePipeline.
	S3Location() *awss3.Location
	// The artifact attribute of the Amazon Simple Storage Service (Amazon S3) URL of the artifact, such as https://s3-us-west-2.amazonaws.com/artifactstorebucket-yivczw8jma0c/test/TemplateSo/1ABCyZZ.zip.
	Url() *string
	// Returns an ArtifactPath for a file within this artifact.
	//
	// CfnOutput is in the form "<artifact-name>::<file-name>".
	AtPath(fileName *string) ArtifactPath
	// Retrieve the metadata stored in this artifact under the given key.
	//
	// If there is no metadata stored under the given key,
	// null will be returned.
	GetMetadata(key *string) interface{}
	// Returns a token for a value inside a JSON file within this artifact.
	GetParam(jsonFile *string, keyName *string) *string
	// Add arbitrary extra payload to the artifact under a given key.
	//
	// This can be used by CodePipeline actions to communicate data between themselves.
	// If metadata was already present under the given key,
	// it will be overwritten with the new value.
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


func NewArtifact(artifactName *string) Artifact {
	_init_.Initialize()

	j := jsiiProxy_Artifact{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Artifact",
		[]interface{}{artifactName},
		&j,
	)

	return &j
}

func NewArtifact_Override(a Artifact, artifactName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Artifact",
		[]interface{}{artifactName},
		a,
	)
}

// A static factory method used to create instances of the Artifact class.
//
// Mainly meant to be used from `decdk`.
func Artifact_Artifact(name *string) Artifact {
	_init_.Initialize()

	var returns Artifact

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.Artifact",
		"artifact",
		[]interface{}{name},
		&returns,
	)

	return returns
}

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

func (a *jsiiProxy_Artifact) SetMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"setMetadata",
		[]interface{}{key, value},
	)
}

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
// Example:
//   // Source stage: read from repository
//   repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &repositoryProps{
//   	repositoryName: jsii.String("template-repo"),
//   })
//   sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
//   source := cpactions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	repository: repo,
//   	output: sourceOutput,
//   	trigger: cpactions.codeCommitTrigger_POLL,
//   })
//   sourceStage := map[string]interface{}{
//   	"stageName": jsii.String("Source"),
//   	"actions": []CodeCommitSourceAction{
//   		source,
//   	},
//   }
//
//   // Deployment stage: create and deploy changeset with manual approval
//   stackName := "OurStack"
//   changeSetName := "StagedChangeSet"
//
//   prodStage := map[string]interface{}{
//   	"stageName": jsii.String("Deploy"),
//   	"actions": []interface{}{
//   		cpactions.NewCloudFormationCreateReplaceChangeSetAction(&CloudFormationCreateReplaceChangeSetActionProps{
//   			"actionName": jsii.String("PrepareChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"adminPermissions": jsii.Boolean(true),
//   			"templatePath": sourceOutput.atPath(jsii.String("template.yaml")),
//   			"runOrder": jsii.Number(1),
//   		}),
//   		cpactions.NewManualApprovalAction(&ManualApprovalActionProps{
//   			"actionName": jsii.String("ApproveChanges"),
//   			"runOrder": jsii.Number(2),
//   		}),
//   		cpactions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
//   			"actionName": jsii.String("ExecuteChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"runOrder": jsii.Number(3),
//   		}),
//   	},
//   }
//
//   codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
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


func NewArtifactPath(artifact Artifact, fileName *string) ArtifactPath {
	_init_.Initialize()

	j := jsiiProxy_ArtifactPath{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.ArtifactPath",
		[]interface{}{artifact, fileName},
		&j,
	)

	return &j
}

func NewArtifactPath_Override(a ArtifactPath, artifact Artifact, fileName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.ArtifactPath",
		[]interface{}{artifact, fileName},
		a,
	)
}

func ArtifactPath_ArtifactPath(artifactName *string, fileName *string) ArtifactPath {
	_init_.Initialize()

	var returns ArtifactPath

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.ArtifactPath",
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomActionType := awscdk.Aws_codepipeline.NewCfnCustomActionType(this, jsii.String("MyCfnCustomActionType"), &cfnCustomActionTypeProps{
//   	category: jsii.String("category"),
//   	inputArtifactDetails: &artifactDetailsProperty{
//   		maximumCount: jsii.Number(123),
//   		minimumCount: jsii.Number(123),
//   	},
//   	outputArtifactDetails: &artifactDetailsProperty{
//   		maximumCount: jsii.Number(123),
//   		minimumCount: jsii.Number(123),
//   	},
//   	provider: jsii.String("provider"),
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	configurationProperties: []interface{}{
//   		&configurationPropertiesProperty{
//   			key: jsii.Boolean(false),
//   			name: jsii.String("name"),
//   			required: jsii.Boolean(false),
//   			secret: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			queryable: jsii.Boolean(false),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	settings: &settingsProperty{
//   		entityUrlTemplate: jsii.String("entityUrlTemplate"),
//   		executionUrlTemplate: jsii.String("executionUrlTemplate"),
//   		revisionUrlTemplate: jsii.String("revisionUrlTemplate"),
//   		thirdPartyConfigurationUrl: jsii.String("thirdPartyConfigurationUrl"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnCustomActionType interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The category of the custom action, such as a build action or a test action.
	Category() *string
	SetCategory(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The configuration properties for the custom action.
	//
	// > You can refer to a name in the configuration properties of the custom action within the URL templates by following the format of {Config:name}, as long as the configuration property is both required and not secret. For more information, see [Create a Custom Action for a Pipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html) .
	ConfigurationProperties() interface{}
	SetConfigurationProperties(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The details of the input artifact for the action, such as its commit ID.
	InputArtifactDetails() interface{}
	SetInputArtifactDetails(val interface{})
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
	// The details of the output artifact of the action, such as its commit ID.
	OutputArtifactDetails() interface{}
	SetOutputArtifactDetails(val interface{})
	// The provider of the service used in the custom action, such as CodeDeploy.
	Provider() *string
	SetProvider(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// URLs that provide users information about this custom action.
	Settings() interface{}
	SetSettings(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags for the custom action.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The version identifier of the custom action.
	Version() *string
	SetVersion(val *string)
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

func (j *jsiiProxy_CfnCustomActionType) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnCustomActionType(scope constructs.Construct, id *string, props *CfnCustomActionTypeProps) CfnCustomActionType {
	_init_.Initialize()

	j := jsiiProxy_CfnCustomActionType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.CfnCustomActionType",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodePipeline::CustomActionType`.
func NewCfnCustomActionType_Override(c CfnCustomActionType, scope constructs.Construct, id *string, props *CfnCustomActionTypeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.CfnCustomActionType",
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
func CfnCustomActionType_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.CfnCustomActionType",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCustomActionType_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.CfnCustomActionType",
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
func CfnCustomActionType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.CfnCustomActionType",
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
		"aws-cdk-lib.aws_codepipeline.CfnCustomActionType",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomActionType) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCustomActionType) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCustomActionType) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCustomActionType) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCustomActionType) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCustomActionType) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCustomActionType) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnCustomActionType) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCustomActionType) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

func (c *jsiiProxy_CfnCustomActionType) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Returns information about the details of an artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactDetailsProperty := &artifactDetailsProperty{
//   	maximumCount: jsii.Number(123),
//   	minimumCount: jsii.Number(123),
//   }
//
type CfnCustomActionType_ArtifactDetailsProperty struct {
	// The maximum number of artifacts allowed for the action type.
	MaximumCount *float64 `field:"required" json:"maximumCount" yaml:"maximumCount"`
	// The minimum number of artifacts allowed for the action type.
	MinimumCount *float64 `field:"required" json:"minimumCount" yaml:"minimumCount"`
}

// The configuration properties for the custom action.
//
// > You can refer to a name in the configuration properties of the custom action within the URL templates by following the format of {Config:name}, as long as the configuration property is both required and not secret. For more information, see [Create a Custom Action for a Pipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationPropertiesProperty := &configurationPropertiesProperty{
//   	key: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	required: jsii.Boolean(false),
//   	secret: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	queryable: jsii.Boolean(false),
//   	type: jsii.String("type"),
//   }
//
type CfnCustomActionType_ConfigurationPropertiesProperty struct {
	// Whether the configuration property is a key.
	Key interface{} `field:"required" json:"key" yaml:"key"`
	// The name of the action configuration property.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether the configuration property is a required value.
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// Whether the configuration property is secret.
	//
	// Secrets are hidden from all calls except for `GetJobDetails` , `GetThirdPartyJobDetails` , `PollForJobs` , and `PollForThirdPartyJobs` .
	//
	// When updating a pipeline, passing * * * * * without changing any other values of the action preserves the previous value of the secret.
	Secret interface{} `field:"required" json:"secret" yaml:"secret"`
	// The description of the action configuration property that is displayed to users.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates that the property is used with `PollForJobs` .
	//
	// When creating a custom action, an action can have up to one queryable property. If it has one, that property must be both required and not secret.
	//
	// If you create a pipeline with a custom action type, and that custom action contains a queryable property, the value for that configuration property is subject to other restrictions. The value must be less than or equal to twenty (20) characters. The value can contain only alphanumeric characters, underscores, and hyphens.
	Queryable interface{} `field:"optional" json:"queryable" yaml:"queryable"`
	// The type of the configuration property.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// `Settings` is a property of the `AWS::CodePipeline::CustomActionType` resource that provides URLs that users can access to view information about the CodePipeline custom action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   settingsProperty := &settingsProperty{
//   	entityUrlTemplate: jsii.String("entityUrlTemplate"),
//   	executionUrlTemplate: jsii.String("executionUrlTemplate"),
//   	revisionUrlTemplate: jsii.String("revisionUrlTemplate"),
//   	thirdPartyConfigurationUrl: jsii.String("thirdPartyConfigurationUrl"),
//   }
//
type CfnCustomActionType_SettingsProperty struct {
	// The URL returned to the CodePipeline console that provides a deep link to the resources of the external system, such as the configuration page for a CodeDeploy deployment group.
	//
	// This link is provided as part of the action display in the pipeline.
	EntityUrlTemplate *string `field:"optional" json:"entityUrlTemplate" yaml:"entityUrlTemplate"`
	// The URL returned to the CodePipeline console that contains a link to the top-level landing page for the external system, such as the console page for CodeDeploy.
	//
	// This link is shown on the pipeline view page in the CodePipeline console and provides a link to the execution entity of the external action.
	ExecutionUrlTemplate *string `field:"optional" json:"executionUrlTemplate" yaml:"executionUrlTemplate"`
	// The URL returned to the CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action.
	RevisionUrlTemplate *string `field:"optional" json:"revisionUrlTemplate" yaml:"revisionUrlTemplate"`
	// The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
	ThirdPartyConfigurationUrl *string `field:"optional" json:"thirdPartyConfigurationUrl" yaml:"thirdPartyConfigurationUrl"`
}

// Properties for defining a `CfnCustomActionType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomActionTypeProps := &cfnCustomActionTypeProps{
//   	category: jsii.String("category"),
//   	inputArtifactDetails: &artifactDetailsProperty{
//   		maximumCount: jsii.Number(123),
//   		minimumCount: jsii.Number(123),
//   	},
//   	outputArtifactDetails: &artifactDetailsProperty{
//   		maximumCount: jsii.Number(123),
//   		minimumCount: jsii.Number(123),
//   	},
//   	provider: jsii.String("provider"),
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	configurationProperties: []interface{}{
//   		&configurationPropertiesProperty{
//   			key: jsii.Boolean(false),
//   			name: jsii.String("name"),
//   			required: jsii.Boolean(false),
//   			secret: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			queryable: jsii.Boolean(false),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	settings: &settingsProperty{
//   		entityUrlTemplate: jsii.String("entityUrlTemplate"),
//   		executionUrlTemplate: jsii.String("executionUrlTemplate"),
//   		revisionUrlTemplate: jsii.String("revisionUrlTemplate"),
//   		thirdPartyConfigurationUrl: jsii.String("thirdPartyConfigurationUrl"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCustomActionTypeProps struct {
	// The category of the custom action, such as a build action or a test action.
	Category *string `field:"required" json:"category" yaml:"category"`
	// The details of the input artifact for the action, such as its commit ID.
	InputArtifactDetails interface{} `field:"required" json:"inputArtifactDetails" yaml:"inputArtifactDetails"`
	// The details of the output artifact of the action, such as its commit ID.
	OutputArtifactDetails interface{} `field:"required" json:"outputArtifactDetails" yaml:"outputArtifactDetails"`
	// The provider of the service used in the custom action, such as CodeDeploy.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// The version identifier of the custom action.
	Version *string `field:"required" json:"version" yaml:"version"`
	// The configuration properties for the custom action.
	//
	// > You can refer to a name in the configuration properties of the custom action within the URL templates by following the format of {Config:name}, as long as the configuration property is both required and not secret. For more information, see [Create a Custom Action for a Pipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html) .
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// URLs that provide users information about this custom action.
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// The tags for the custom action.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::CodePipeline::Pipeline`.
//
// The `AWS::CodePipeline::Pipeline` resource creates a CodePipeline pipeline that describes how software changes go through a release process. For more information, see [What Is CodePipeline?](https://docs.aws.amazon.com/codepipeline/latest/userguide/welcome.html) in the *AWS CodePipeline User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   cfnPipeline := awscdk.Aws_codepipeline.NewCfnPipeline(this, jsii.String("MyCfnPipeline"), &cfnPipelineProps{
//   	roleArn: jsii.String("roleArn"),
//   	stages: []interface{}{
//   		&stageDeclarationProperty{
//   			actions: []interface{}{
//   				&actionDeclarationProperty{
//   					actionTypeId: &actionTypeIdProperty{
//   						category: jsii.String("category"),
//   						owner: jsii.String("owner"),
//   						provider: jsii.String("provider"),
//   						version: jsii.String("version"),
//   					},
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					configuration: configuration,
//   					inputArtifacts: []interface{}{
//   						&inputArtifactProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					namespace: jsii.String("namespace"),
//   					outputArtifacts: []interface{}{
//   						&outputArtifactProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					region: jsii.String("region"),
//   					roleArn: jsii.String("roleArn"),
//   					runOrder: jsii.Number(123),
//   				},
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			blockers: []interface{}{
//   				&blockerDeclarationProperty{
//   					name: jsii.String("name"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	artifactStore: &artifactStoreProperty{
//   		location: jsii.String("location"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		encryptionKey: &encryptionKeyProperty{
//   			id: jsii.String("id"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	artifactStores: []interface{}{
//   		&artifactStoreMapProperty{
//   			artifactStore: &artifactStoreProperty{
//   				location: jsii.String("location"),
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				encryptionKey: &encryptionKeyProperty{
//   					id: jsii.String("id"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			region: jsii.String("region"),
//   		},
//   	},
//   	disableInboundStageTransitions: []interface{}{
//   		&stageTransitionProperty{
//   			reason: jsii.String("reason"),
//   			stageName: jsii.String("stageName"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	restartExecutionOnUpdate: jsii.Boolean(false),
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
	// The S3 bucket where artifacts for the pipeline are stored.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	ArtifactStore() interface{}
	SetArtifactStore(val interface{})
	// A mapping of `artifactStore` objects and their corresponding AWS Regions.
	//
	// There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	ArtifactStores() interface{}
	SetArtifactStores(val interface{})
	// The version of the pipeline.
	//
	// > A new pipeline is always assigned a version number of 1. This number increments when a pipeline is updated.
	AttrVersion() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Represents the input of a `DisableStageTransition` action.
	DisableInboundStageTransitions() interface{}
	SetDisableInboundStageTransitions(val interface{})
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
	// The name of the pipeline.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Indicates whether to rerun the CodePipeline pipeline after you update it.
	RestartExecutionOnUpdate() interface{}
	SetRestartExecutionOnUpdate(val interface{})
	// The Amazon Resource Name (ARN) for CodePipeline to use to either perform actions with no `actionRoleArn` , or to use to assume roles for actions with an `actionRoleArn` .
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Represents information about a stage and its definition.
	Stages() interface{}
	SetStages(val interface{})
	// Specifies the tags applied to the pipeline.
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

func (j *jsiiProxy_CfnPipeline) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnPipeline(scope constructs.Construct, id *string, props *CfnPipelineProps) CfnPipeline {
	_init_.Initialize()

	j := jsiiProxy_CfnPipeline{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodePipeline::Pipeline`.
func NewCfnPipeline_Override(c CfnPipeline, scope constructs.Construct, id *string, props *CfnPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline",
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
func CfnPipeline_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline",
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
		"aws-cdk-lib.aws_codepipeline.CfnPipeline",
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
		"aws-cdk-lib.aws_codepipeline.CfnPipeline",
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
		"aws-cdk-lib.aws_codepipeline.CfnPipeline",
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

// Represents information about an action declaration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   actionDeclarationProperty := &actionDeclarationProperty{
//   	actionTypeId: &actionTypeIdProperty{
//   		category: jsii.String("category"),
//   		owner: jsii.String("owner"),
//   		provider: jsii.String("provider"),
//   		version: jsii.String("version"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	configuration: configuration,
//   	inputArtifacts: []interface{}{
//   		&inputArtifactProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	namespace: jsii.String("namespace"),
//   	outputArtifacts: []interface{}{
//   		&outputArtifactProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	region: jsii.String("region"),
//   	roleArn: jsii.String("roleArn"),
//   	runOrder: jsii.Number(123),
//   }
//
type CfnPipeline_ActionDeclarationProperty struct {
	// Specifies the action type and the provider of the action.
	ActionTypeId interface{} `field:"required" json:"actionTypeId" yaml:"actionTypeId"`
	// The action declaration's name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The action's configuration.
	//
	// These are key-value pairs that specify input values for an action. For more information, see [Action Structure Requirements in CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements) . For the list of configuration properties for the AWS CloudFormation action type in CodePipeline, see [Configuration Properties Reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-action-reference.html) in the *AWS CloudFormation User Guide* . For template snippets with examples, see [Using Parameter Override Functions with CodePipeline Pipelines](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-parameter-override-functions.html) in the *AWS CloudFormation User Guide* .
	//
	// The values can be represented in either JSON or YAML format. For example, the JSON configuration item format is as follows:
	//
	// *JSON:*
	//
	// `"Configuration" : { Key : Value },`.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The name or ID of the artifact consumed by the action, such as a test or build artifact.
	//
	// > For a CodeBuild action with multiple input artifacts, one of your input sources must be designated the PrimarySource. For more information, see the [CodeBuild action reference page](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeBuild.html) in the *AWS CodePipeline User Guide* .
	InputArtifacts interface{} `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// The variable namespace associated with the action.
	//
	// All variables produced as output by this action fall under this namespace.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name or ID of the result of the action declaration, such as a test or build artifact.
	OutputArtifacts interface{} `field:"optional" json:"outputArtifacts" yaml:"outputArtifacts"`
	// The action declaration's AWS Region, such as us-east-1.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The ARN of the IAM service role that performs the declared action.
	//
	// This is assumed through the roleArn for the pipeline.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The order in which actions are run.
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
}

// Represents information about an action type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionTypeIdProperty := &actionTypeIdProperty{
//   	category: jsii.String("category"),
//   	owner: jsii.String("owner"),
//   	provider: jsii.String("provider"),
//   	version: jsii.String("version"),
//   }
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
	// - `Approval`.
	Category *string `field:"required" json:"category" yaml:"category"`
	// The creator of the action being called.
	//
	// There are three valid values for the `Owner` field in the action category section within your pipeline structure: `AWS` , `ThirdParty` , and `Custom` . For more information, see [Valid Action Types and Providers in CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#actions-valid-providers) .
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The provider of the service being called by the action.
	//
	// Valid providers are determined by the action category. For example, an action in the Deploy category type might have a provider of CodeDeploy, which would be specified as `CodeDeploy` . For more information, see [Valid Action Types and Providers in CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#actions-valid-providers) .
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// A string that describes the action version.
	Version *string `field:"required" json:"version" yaml:"version"`
}

// A mapping of `artifactStore` objects and their corresponding AWS Regions.
//
// There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.
//
// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactStoreMapProperty := &artifactStoreMapProperty{
//   	artifactStore: &artifactStoreProperty{
//   		location: jsii.String("location"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		encryptionKey: &encryptionKeyProperty{
//   			id: jsii.String("id"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	region: jsii.String("region"),
//   }
//
type CfnPipeline_ArtifactStoreMapProperty struct {
	// Represents information about the S3 bucket where artifacts are stored for the pipeline.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	ArtifactStore interface{} `field:"required" json:"artifactStore" yaml:"artifactStore"`
	// The action declaration's AWS Region, such as us-east-1.
	Region *string `field:"required" json:"region" yaml:"region"`
}

// The S3 bucket where artifacts for the pipeline are stored.
//
// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactStoreProperty := &artifactStoreProperty{
//   	location: jsii.String("location"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	encryptionKey: &encryptionKeyProperty{
//   		id: jsii.String("id"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnPipeline_ArtifactStoreProperty struct {
	// The S3 bucket used for storing the artifacts for a pipeline.
	//
	// You can specify the name of an S3 bucket but not a folder in the bucket. A folder to contain the pipeline artifacts is created for you based on the name of the pipeline. You can use any S3 bucket in the same AWS Region as the pipeline to store your pipeline artifacts.
	Location *string `field:"required" json:"location" yaml:"location"`
	// The type of the artifact store, such as S3.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The encryption key used to encrypt the data in the artifact store, such as an AWS Key Management Service ( AWS KMS) key.
	//
	// If this is undefined, the default key for Amazon S3 is used. To see an example artifact store encryption key field, see the example structure here: [AWS::CodePipeline::Pipeline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html) .
	EncryptionKey interface{} `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

// Reserved for future use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockerDeclarationProperty := &blockerDeclarationProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//   }
//
type CfnPipeline_BlockerDeclarationProperty struct {
	// Reserved for future use.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Reserved for future use.
	Type *string `field:"required" json:"type" yaml:"type"`
}

// Represents information about the key used to encrypt data in the artifact store, such as an AWS Key Management Service ( AWS KMS) key.
//
// `EncryptionKey` is a property of the [ArtifactStore](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionKeyProperty := &encryptionKeyProperty{
//   	id: jsii.String("id"),
//   	type: jsii.String("type"),
//   }
//
type CfnPipeline_EncryptionKeyProperty struct {
	// The ID used to identify the key.
	//
	// For an AWS KMS key, you can use the key ID, the key ARN, or the alias ARN.
	//
	// > Aliases are recognized only in the account that created the AWS KMS key. For cross-account actions, you can only use the key ID or key ARN to identify the key.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The type of encryption key, such as an AWS KMS key.
	//
	// When creating or updating a pipeline, the value must be set to 'KMS'.
	Type *string `field:"required" json:"type" yaml:"type"`
}

// Represents information about an artifact to be worked on, such as a test or build artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputArtifactProperty := &inputArtifactProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnPipeline_InputArtifactProperty struct {
	// The name of the artifact to be worked on (for example, "My App").
	//
	// The input artifact of an action must exactly match the output artifact declared in a preceding action, but the input artifact does not have to be the next action in strict sequence from the action that provided the output artifact. Actions in parallel can declare different output artifacts, which are in turn consumed by different following actions.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Represents information about the output of an action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputArtifactProperty := &outputArtifactProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnPipeline_OutputArtifactProperty struct {
	// The name of the output of an artifact, such as "My App".
	//
	// The output artifact name must exactly match the input artifact declared for a downstream action. However, the downstream action's input artifact does not have to be the next action in strict sequence from the action that provided the output artifact. Actions in parallel can declare different output artifacts, which are in turn consumed by different following actions.
	//
	// Output artifact names must be unique within a pipeline.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Represents information about a stage and its definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   stageDeclarationProperty := &stageDeclarationProperty{
//   	actions: []interface{}{
//   		&actionDeclarationProperty{
//   			actionTypeId: &actionTypeIdProperty{
//   				category: jsii.String("category"),
//   				owner: jsii.String("owner"),
//   				provider: jsii.String("provider"),
//   				version: jsii.String("version"),
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			configuration: configuration,
//   			inputArtifacts: []interface{}{
//   				&inputArtifactProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			namespace: jsii.String("namespace"),
//   			outputArtifacts: []interface{}{
//   				&outputArtifactProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			region: jsii.String("region"),
//   			roleArn: jsii.String("roleArn"),
//   			runOrder: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	blockers: []interface{}{
//   		&blockerDeclarationProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnPipeline_StageDeclarationProperty struct {
	// The actions included in a stage.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The name of the stage.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Reserved for future use.
	Blockers interface{} `field:"optional" json:"blockers" yaml:"blockers"`
}

// The name of the pipeline in which you want to disable the flow of artifacts from one stage to another.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageTransitionProperty := &stageTransitionProperty{
//   	reason: jsii.String("reason"),
//   	stageName: jsii.String("stageName"),
//   }
//
type CfnPipeline_StageTransitionProperty struct {
	// The reason given to the user that a stage is disabled, such as waiting for manual approval or manual tests.
	//
	// This message is displayed in the pipeline console UI.
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// The name of the stage where you want to disable the inbound or outbound transition of artifacts.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
}

// Properties for defining a `CfnPipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   cfnPipelineProps := &cfnPipelineProps{
//   	roleArn: jsii.String("roleArn"),
//   	stages: []interface{}{
//   		&stageDeclarationProperty{
//   			actions: []interface{}{
//   				&actionDeclarationProperty{
//   					actionTypeId: &actionTypeIdProperty{
//   						category: jsii.String("category"),
//   						owner: jsii.String("owner"),
//   						provider: jsii.String("provider"),
//   						version: jsii.String("version"),
//   					},
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					configuration: configuration,
//   					inputArtifacts: []interface{}{
//   						&inputArtifactProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					namespace: jsii.String("namespace"),
//   					outputArtifacts: []interface{}{
//   						&outputArtifactProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					region: jsii.String("region"),
//   					roleArn: jsii.String("roleArn"),
//   					runOrder: jsii.Number(123),
//   				},
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			blockers: []interface{}{
//   				&blockerDeclarationProperty{
//   					name: jsii.String("name"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	artifactStore: &artifactStoreProperty{
//   		location: jsii.String("location"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		encryptionKey: &encryptionKeyProperty{
//   			id: jsii.String("id"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	artifactStores: []interface{}{
//   		&artifactStoreMapProperty{
//   			artifactStore: &artifactStoreProperty{
//   				location: jsii.String("location"),
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				encryptionKey: &encryptionKeyProperty{
//   					id: jsii.String("id"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			region: jsii.String("region"),
//   		},
//   	},
//   	disableInboundStageTransitions: []interface{}{
//   		&stageTransitionProperty{
//   			reason: jsii.String("reason"),
//   			stageName: jsii.String("stageName"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	restartExecutionOnUpdate: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipelineProps struct {
	// The Amazon Resource Name (ARN) for CodePipeline to use to either perform actions with no `actionRoleArn` , or to use to assume roles for actions with an `actionRoleArn` .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Represents information about a stage and its definition.
	Stages interface{} `field:"required" json:"stages" yaml:"stages"`
	// The S3 bucket where artifacts for the pipeline are stored.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	ArtifactStore interface{} `field:"optional" json:"artifactStore" yaml:"artifactStore"`
	// A mapping of `artifactStore` objects and their corresponding AWS Regions.
	//
	// There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	ArtifactStores interface{} `field:"optional" json:"artifactStores" yaml:"artifactStores"`
	// Represents the input of a `DisableStageTransition` action.
	DisableInboundStageTransitions interface{} `field:"optional" json:"disableInboundStageTransitions" yaml:"disableInboundStageTransitions"`
	// The name of the pipeline.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates whether to rerun the CodePipeline pipeline after you update it.
	RestartExecutionOnUpdate interface{} `field:"optional" json:"restartExecutionOnUpdate" yaml:"restartExecutionOnUpdate"`
	// Specifies the tags applied to the pipeline.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::CodePipeline::Webhook`.
//
// The `AWS::CodePipeline::Webhook` resource creates and registers your webhook. After the webhook is created and registered, it triggers your pipeline to start every time an external event occurs. For more information, see [Configure Your GitHub Pipelines to Use Webhooks for Change Detection](https://docs.aws.amazon.com/codepipeline/latest/userguide/pipelines-webhooks-migration.html) in the *AWS CodePipeline User Guide* .
//
// We strongly recommend that you use AWS Secrets Manager to store your credentials. If you use Secrets Manager, you must have already configured and stored your secret parameters in Secrets Manager. For more information, see [Using Dynamic References to Specify Template Values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) .
//
// > When passing secret parameters, do not enter the value directly into the template. The value is rendered as plaintext and is therefore readable. For security reasons, do not use plaintext in your AWS CloudFormation template to store your credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebhook := awscdk.Aws_codepipeline.NewCfnWebhook(this, jsii.String("MyCfnWebhook"), &cfnWebhookProps{
//   	authentication: jsii.String("authentication"),
//   	authenticationConfiguration: &webhookAuthConfigurationProperty{
//   		allowedIpRange: jsii.String("allowedIpRange"),
//   		secretToken: jsii.String("secretToken"),
//   	},
//   	filters: []interface{}{
//   		&webhookFilterRuleProperty{
//   			jsonPath: jsii.String("jsonPath"),
//
//   			// the properties below are optional
//   			matchEquals: jsii.String("matchEquals"),
//   		},
//   	},
//   	targetAction: jsii.String("targetAction"),
//   	targetPipeline: jsii.String("targetPipeline"),
//   	targetPipelineVersion: jsii.Number(123),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	registerWithThirdParty: jsii.Boolean(false),
//   })
//
type CfnWebhook interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The webhook URL generated by AWS CodePipeline , such as `https://eu-central-1.webhooks.aws/trigger123456` .
	AttrUrl() *string
	// Supported options are GITHUB_HMAC, IP, and UNAUTHENTICATED.
	//
	// - For information about the authentication scheme implemented by GITHUB_HMAC, see [Securing your webhooks](https://docs.aws.amazon.com/https://developer.github.com/webhooks/securing/) on the GitHub Developer website.
	// - IP rejects webhooks trigger requests unless they originate from an IP address in the IP range whitelisted in the authentication configuration.
	// - UNAUTHENTICATED accepts all webhook trigger requests regardless of origin.
	Authentication() *string
	SetAuthentication(val *string)
	// Properties that configure the authentication applied to incoming webhook trigger requests.
	//
	// The required properties depend on the authentication type. For GITHUB_HMAC, only the `SecretToken` property must be set. For IP, only the `AllowedIPRange` property must be set to a valid CIDR range. For UNAUTHENTICATED, no properties can be set.
	AuthenticationConfiguration() interface{}
	SetAuthenticationConfiguration(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A list of rules applied to the body/payload sent in the POST request to a webhook URL.
	//
	// All defined rules must pass for the request to be accepted and the pipeline started.
	Filters() interface{}
	SetFilters(val interface{})
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
	// The name of the webhook.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Configures a connection between the webhook that was created and the external tool with events to be detected.
	RegisterWithThirdParty() interface{}
	SetRegisterWithThirdParty(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The name of the action in a pipeline you want to connect to the webhook.
	//
	// The action must be from the source (first) stage of the pipeline.
	TargetAction() *string
	SetTargetAction(val *string)
	// The name of the pipeline you want to connect to the webhook.
	TargetPipeline() *string
	SetTargetPipeline(val *string)
	// The version number of the pipeline to be connected to the trigger request.
	//
	// Required: Yes
	//
	// Type: Integer
	//
	// Update requires: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)
	TargetPipelineVersion() *float64
	SetTargetPipelineVersion(val *float64)
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

func (j *jsiiProxy_CfnWebhook) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnWebhook(scope constructs.Construct, id *string, props *CfnWebhookProps) CfnWebhook {
	_init_.Initialize()

	j := jsiiProxy_CfnWebhook{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.CfnWebhook",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodePipeline::Webhook`.
func NewCfnWebhook_Override(c CfnWebhook, scope constructs.Construct, id *string, props *CfnWebhookProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.CfnWebhook",
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
func CfnWebhook_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.CfnWebhook",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnWebhook_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.CfnWebhook",
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
func CfnWebhook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.CfnWebhook",
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
		"aws-cdk-lib.aws_codepipeline.CfnWebhook",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWebhook) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnWebhook) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWebhook) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnWebhook) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnWebhook) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnWebhook) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnWebhook) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnWebhook) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnWebhook) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

func (c *jsiiProxy_CfnWebhook) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The authentication applied to incoming webhook trigger requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webhookAuthConfigurationProperty := &webhookAuthConfigurationProperty{
//   	allowedIpRange: jsii.String("allowedIpRange"),
//   	secretToken: jsii.String("secretToken"),
//   }
//
type CfnWebhook_WebhookAuthConfigurationProperty struct {
	// The property used to configure acceptance of webhooks in an IP address range.
	//
	// For IP, only the `AllowedIPRange` property must be set. This property must be set to a valid CIDR range.
	AllowedIpRange *string `field:"optional" json:"allowedIpRange" yaml:"allowedIpRange"`
	// The property used to configure GitHub authentication.
	//
	// For GITHUB_HMAC, only the `SecretToken` property must be set.
	SecretToken *string `field:"optional" json:"secretToken" yaml:"secretToken"`
}

// The event criteria that specify when a webhook notification is sent to your URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webhookFilterRuleProperty := &webhookFilterRuleProperty{
//   	jsonPath: jsii.String("jsonPath"),
//
//   	// the properties below are optional
//   	matchEquals: jsii.String("matchEquals"),
//   }
//
type CfnWebhook_WebhookFilterRuleProperty struct {
	// A JsonPath expression that is applied to the body/payload of the webhook.
	//
	// The value selected by the JsonPath expression must match the value specified in the `MatchEquals` field. Otherwise, the request is ignored. For more information, see [Java JsonPath implementation](https://docs.aws.amazon.com/https://github.com/json-path/JsonPath) in GitHub.
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
	// The value selected by the `JsonPath` expression must match what is supplied in the `MatchEquals` field.
	//
	// Otherwise, the request is ignored. Properties from the target action configuration can be included as placeholders in this value by surrounding the action configuration key with curly brackets. For example, if the value supplied here is "refs/heads/{Branch}" and the target action has an action configuration property called "Branch" with a value of "main", the `MatchEquals` value is evaluated as "refs/heads/main". For a list of action configuration properties for built-in action types, see [Pipeline Structure Reference Action Requirements](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements) .
	MatchEquals *string `field:"optional" json:"matchEquals" yaml:"matchEquals"`
}

// Properties for defining a `CfnWebhook`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebhookProps := &cfnWebhookProps{
//   	authentication: jsii.String("authentication"),
//   	authenticationConfiguration: &webhookAuthConfigurationProperty{
//   		allowedIpRange: jsii.String("allowedIpRange"),
//   		secretToken: jsii.String("secretToken"),
//   	},
//   	filters: []interface{}{
//   		&webhookFilterRuleProperty{
//   			jsonPath: jsii.String("jsonPath"),
//
//   			// the properties below are optional
//   			matchEquals: jsii.String("matchEquals"),
//   		},
//   	},
//   	targetAction: jsii.String("targetAction"),
//   	targetPipeline: jsii.String("targetPipeline"),
//   	targetPipelineVersion: jsii.Number(123),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	registerWithThirdParty: jsii.Boolean(false),
//   }
//
type CfnWebhookProps struct {
	// Supported options are GITHUB_HMAC, IP, and UNAUTHENTICATED.
	//
	// - For information about the authentication scheme implemented by GITHUB_HMAC, see [Securing your webhooks](https://docs.aws.amazon.com/https://developer.github.com/webhooks/securing/) on the GitHub Developer website.
	// - IP rejects webhooks trigger requests unless they originate from an IP address in the IP range whitelisted in the authentication configuration.
	// - UNAUTHENTICATED accepts all webhook trigger requests regardless of origin.
	Authentication *string `field:"required" json:"authentication" yaml:"authentication"`
	// Properties that configure the authentication applied to incoming webhook trigger requests.
	//
	// The required properties depend on the authentication type. For GITHUB_HMAC, only the `SecretToken` property must be set. For IP, only the `AllowedIPRange` property must be set to a valid CIDR range. For UNAUTHENTICATED, no properties can be set.
	AuthenticationConfiguration interface{} `field:"required" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// A list of rules applied to the body/payload sent in the POST request to a webhook URL.
	//
	// All defined rules must pass for the request to be accepted and the pipeline started.
	Filters interface{} `field:"required" json:"filters" yaml:"filters"`
	// The name of the action in a pipeline you want to connect to the webhook.
	//
	// The action must be from the source (first) stage of the pipeline.
	TargetAction *string `field:"required" json:"targetAction" yaml:"targetAction"`
	// The name of the pipeline you want to connect to the webhook.
	TargetPipeline *string `field:"required" json:"targetPipeline" yaml:"targetPipeline"`
	// The version number of the pipeline to be connected to the trigger request.
	//
	// Required: Yes
	//
	// Type: Integer
	//
	// Update requires: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)
	TargetPipelineVersion *float64 `field:"required" json:"targetPipelineVersion" yaml:"targetPipelineVersion"`
	// The name of the webhook.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Configures a connection between the webhook that was created and the external tool with events to be detected.
	RegisterWithThirdParty interface{} `field:"optional" json:"registerWithThirdParty" yaml:"registerWithThirdParty"`
}

// Common properties shared by all Actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   commonActionProps := &commonActionProps{
//   	actionName: jsii.String("actionName"),
//
//   	// the properties below are optional
//   	runOrder: jsii.Number(123),
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   }
//
type CommonActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
}

// Common properties shared by all Actions whose {@link ActionProperties.owner} field is 'AWS' (or unset, as 'AWS' is the default).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   commonAwsActionProps := &commonAwsActionProps{
//   	actionName: jsii.String("actionName"),
//
//   	// the properties below are optional
//   	role: role,
//   	runOrder: jsii.Number(123),
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   }
//
type CommonAwsActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

// An interface representing resources generated in order to support the cross-region capabilities of CodePipeline.
//
// You get instances of this interface from the {@link Pipeline#crossRegionSupport} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var stack stack
//
//   crossRegionSupport := &crossRegionSupport{
//   	replicationBucket: bucket,
//   	stack: stack,
//   }
//
type CrossRegionSupport struct {
	// The replication Bucket used by CodePipeline to operate in this region.
	//
	// Belongs to {@link stack}.
	ReplicationBucket awss3.IBucket `field:"required" json:"replicationBucket" yaml:"replicationBucket"`
	// The Stack that has been created to house the replication Bucket required for this  region.
	Stack awscdk.Stack `field:"required" json:"stack" yaml:"stack"`
}

// The creation attributes used for defining a configuration property of a custom Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionProperty := &customActionProperty{
//   	name: jsii.String("name"),
//   	required: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	key: jsii.Boolean(false),
//   	queryable: jsii.Boolean(false),
//   	secret: jsii.Boolean(false),
//   	type: jsii.String("type"),
//   }
//
type CustomActionProperty struct {
	// The name of the property.
	//
	// You use this name in the `configuration` attribute when defining your custom Action class.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether this property is required.
	Required *bool `field:"required" json:"required" yaml:"required"`
	// The description of the property.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether this property is a key.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-key
	//
	Key *bool `field:"optional" json:"key" yaml:"key"`
	// Whether this property is queryable.
	//
	// Note that only a single property of a custom Action can be queryable.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-queryable
	//
	Queryable *bool `field:"optional" json:"queryable" yaml:"queryable"`
	// Whether this property is secret, like a password, or access key.
	Secret *bool `field:"optional" json:"secret" yaml:"secret"`
	// The type of the property, like 'String', 'Number', or 'Boolean'.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// The resource representing registering a custom Action with CodePipeline.
//
// For the Action to be usable, it has to be registered for every region and every account it's used in.
// In addition to this class, you should most likely also provide your clients a class
// representing your custom Action, extending the Action class,
// and taking the `actionProperties` as properly typed, construction properties.
//
// Example:
//   // Make a custom CodePipeline Action
//   // Make a custom CodePipeline Action
//   codepipeline.NewCustomActionRegistration(this, jsii.String("GenericGitSourceProviderResource"), &customActionRegistrationProps{
//   	category: codepipeline.actionCategory_SOURCE,
//   	artifactBounds: &actionArtifactBounds{
//   		minInputs: jsii.Number(0),
//   		maxInputs: jsii.Number(0),
//   		minOutputs: jsii.Number(1),
//   		maxOutputs: jsii.Number(1),
//   	},
//   	provider: jsii.String("GenericGitSource"),
//   	version: jsii.String("1"),
//   	entityUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
//   	executionUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
//   	actionProperties: []customActionProperty{
//   		&customActionProperty{
//   			name: jsii.String("Branch"),
//   			required: jsii.Boolean(true),
//   			key: jsii.Boolean(false),
//   			secret: jsii.Boolean(false),
//   			queryable: jsii.Boolean(false),
//   			description: jsii.String("Git branch to pull"),
//   			type: jsii.String("String"),
//   		},
//   		&customActionProperty{
//   			name: jsii.String("GitUrl"),
//   			required: jsii.Boolean(true),
//   			key: jsii.Boolean(false),
//   			secret: jsii.Boolean(false),
//   			queryable: jsii.Boolean(false),
//   			description: jsii.String("SSH git clone URL"),
//   			type: jsii.String("String"),
//   		},
//   	},
//   })
//
type CustomActionRegistration interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CustomActionRegistration
type jsiiProxy_CustomActionRegistration struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CustomActionRegistration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewCustomActionRegistration(scope constructs.Construct, id *string, props *CustomActionRegistrationProps) CustomActionRegistration {
	_init_.Initialize()

	j := jsiiProxy_CustomActionRegistration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.CustomActionRegistration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCustomActionRegistration_Override(c CustomActionRegistration, scope constructs.Construct, id *string, props *CustomActionRegistrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.CustomActionRegistration",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CustomActionRegistration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.CustomActionRegistration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

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

// Properties of registering a custom Action.
//
// Example:
//   // Make a custom CodePipeline Action
//   // Make a custom CodePipeline Action
//   codepipeline.NewCustomActionRegistration(this, jsii.String("GenericGitSourceProviderResource"), &customActionRegistrationProps{
//   	category: codepipeline.actionCategory_SOURCE,
//   	artifactBounds: &actionArtifactBounds{
//   		minInputs: jsii.Number(0),
//   		maxInputs: jsii.Number(0),
//   		minOutputs: jsii.Number(1),
//   		maxOutputs: jsii.Number(1),
//   	},
//   	provider: jsii.String("GenericGitSource"),
//   	version: jsii.String("1"),
//   	entityUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
//   	executionUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
//   	actionProperties: []customActionProperty{
//   		&customActionProperty{
//   			name: jsii.String("Branch"),
//   			required: jsii.Boolean(true),
//   			key: jsii.Boolean(false),
//   			secret: jsii.Boolean(false),
//   			queryable: jsii.Boolean(false),
//   			description: jsii.String("Git branch to pull"),
//   			type: jsii.String("String"),
//   		},
//   		&customActionProperty{
//   			name: jsii.String("GitUrl"),
//   			required: jsii.Boolean(true),
//   			key: jsii.Boolean(false),
//   			secret: jsii.Boolean(false),
//   			queryable: jsii.Boolean(false),
//   			description: jsii.String("SSH git clone URL"),
//   			type: jsii.String("String"),
//   		},
//   	},
//   })
//
type CustomActionRegistrationProps struct {
	// The artifact bounds of the Action.
	ArtifactBounds *ActionArtifactBounds `field:"required" json:"artifactBounds" yaml:"artifactBounds"`
	// The category of the Action.
	Category ActionCategory `field:"required" json:"category" yaml:"category"`
	// The provider of the Action.
	//
	// For example, `'MyCustomActionProvider'`.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// The properties used for customizing the instance of your Action.
	ActionProperties *[]*CustomActionProperty `field:"optional" json:"actionProperties" yaml:"actionProperties"`
	// The URL shown for the entire Action in the Pipeline UI.
	EntityUrl *string `field:"optional" json:"entityUrl" yaml:"entityUrl"`
	// The URL shown for a particular execution of an Action in the Pipeline UI.
	ExecutionUrl *string `field:"optional" json:"executionUrl" yaml:"executionUrl"`
	// The version of your Action.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// The CodePipeline variables that are global, not bound to a specific action.
//
// This class defines a bunch of static fields that represent the different variables.
// These can be used can be used in any action configuration.
//
// Example:
//   // OtherAction is some action type that produces variables, like EcrSourceAction
//   // OtherAction is some action type that produces variables, like EcrSourceAction
//   NewOtherAction(&otherActionProps{
//   	// ...
//   	config: codepipeline.globalVariables_ExecutionId(),
//   	actionName: jsii.String("otherAction"),
//   })
//
type GlobalVariables interface {
}

// The jsii proxy struct for GlobalVariables
type jsiiProxy_GlobalVariables struct {
	_ byte // padding
}

func NewGlobalVariables() GlobalVariables {
	_init_.Initialize()

	j := jsiiProxy_GlobalVariables{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.GlobalVariables",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewGlobalVariables_Override(g GlobalVariables) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.GlobalVariables",
		nil, // no parameters
		g,
	)
}

func GlobalVariables_ExecutionId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codepipeline.GlobalVariables",
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
type IAction interface {
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *ActionProperties
}

// The jsii proxy for IAction
type jsiiProxy_IAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IAction) Bind(scope constructs.Construct, stage IStage, options *ActionBindOptions) *ActionConfig {
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
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Action execution" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	NotifyOnAnyActionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Manual approval" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	NotifyOnAnyManualApprovalStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Stage execution" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	NotifyOnAnyStageStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Pipeline execution" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	NotifyOnExecutionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an event rule triggered by this CodePipeline.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Define an event rule triggered by the "CodePipeline Pipeline Execution State Change" event emitted from this pipeline.
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The ARN of the Pipeline.
	PipelineArn() *string
	// The name of the Pipeline.
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

func (j *jsiiProxy_IPipeline) Node() constructs.Node {
	var returns constructs.Node
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
type IStage interface {
	AddAction(action IAction)
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// The actions belonging to this stage.
	Actions() *[]IAction
	Pipeline() IPipeline
	// The physical, human-readable name of this Pipeline Stage.
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
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // create a pipeline
//   import codecommit "github.com/aws-samples/dummy/awscdkawscodecommit"
//
//   // add a source action to the stage
//   var repo repository
//   var sourceArtifact artifact
//
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("Pipeline"))
//
//   // add a stage
//   sourceStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   })
//   sourceStage.addAction(codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	output: sourceArtifact,
//   	repository: repo,
//   }))
//
type Pipeline interface {
	awscdk.Resource
	IPipeline
	// Bucket used to store output artifacts.
	ArtifactBucket() awss3.IBucket
	// Returns all of the {@link CrossRegionSupportStack}s that were generated automatically when dealing with Actions that reside in a different region than the Pipeline itself.
	CrossRegionSupport() *map[string]*CrossRegionSupport
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// ARN of this pipeline.
	PipelineArn() *string
	// The name of the pipeline.
	PipelineName() *string
	// The version of the pipeline.
	PipelineVersion() *string
	// The IAM role AWS CodePipeline will use to perform actions or assume roles for actions with a more specific IAM role.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Get the number of Stages in this Pipeline.
	StageCount() *float64
	// Returns the stages that comprise the pipeline.
	//
	// **Note**: the returned array is a defensive copy,
	// so adding elements to it has no effect.
	// Instead, use the {@link addStage} method if you want to add more stages
	// to the pipeline.
	Stages() *[]IStage
	// Creates a new Stage, and adds it to this Pipeline.
	//
	// Returns: the newly created Stage.
	AddStage(props *StageOptions) IStage
	// Adds a statement to the pipeline role.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Returns a source configuration for notification rule.
	BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Defines a CodeStar notification rule triggered when the pipeline events emitted by you specified, it very similar to `onEvent` API.
	//
	// You can also use the methods `notifyOnExecutionStateChange`, `notifyOnAnyStageStateChange`,
	// `notifyOnAnyActionStateChange` and `notifyOnAnyManualApprovalStateChange`
	// to define rules for these specific event emitted.
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Action execution" events emitted from this pipeline.
	NotifyOnAnyActionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Manual approval" events emitted from this pipeline.
	NotifyOnAnyManualApprovalStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Stage execution" events emitted from this pipeline.
	NotifyOnAnyStageStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Pipeline execution" events emitted from this pipeline.
	NotifyOnExecutionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines an event rule triggered by this CodePipeline.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an event rule triggered by the "CodePipeline Pipeline Execution State Change" event emitted from this pipeline.
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Access one of the pipeline's stages by stage name.
	Stage(stageName *string) IStage
	// Returns a string representation of this construct.
	ToString() *string
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

func (j *jsiiProxy_Pipeline) Node() constructs.Node {
	var returns constructs.Node
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


func NewPipeline(scope constructs.Construct, id *string, props *PipelineProps) Pipeline {
	_init_.Initialize()

	j := jsiiProxy_Pipeline{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Pipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPipeline_Override(p Pipeline, scope constructs.Construct, id *string, props *PipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Pipeline",
		[]interface{}{scope, id, props},
		p,
	)
}

// Import a pipeline into this app.
func Pipeline_FromPipelineArn(scope constructs.Construct, id *string, pipelineArn *string) IPipeline {
	_init_.Initialize()

	var returns IPipeline

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.Pipeline",
		"fromPipelineArn",
		[]interface{}{scope, id, pipelineArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Pipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.Pipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Pipeline_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.Pipeline",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

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

func (p *jsiiProxy_Pipeline) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		p,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (p *jsiiProxy_Pipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

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

// The list of event types for AWS Codepipeline Pipeline.
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
//
type PipelineNotificationEvents string

const (
	// Trigger notification when pipeline execution failed.
	PipelineNotificationEvents_PIPELINE_EXECUTION_FAILED PipelineNotificationEvents = "PIPELINE_EXECUTION_FAILED"
	// Trigger notification when pipeline execution canceled.
	PipelineNotificationEvents_PIPELINE_EXECUTION_CANCELED PipelineNotificationEvents = "PIPELINE_EXECUTION_CANCELED"
	// Trigger notification when pipeline execution started.
	PipelineNotificationEvents_PIPELINE_EXECUTION_STARTED PipelineNotificationEvents = "PIPELINE_EXECUTION_STARTED"
	// Trigger notification when pipeline execution resumed.
	PipelineNotificationEvents_PIPELINE_EXECUTION_RESUMED PipelineNotificationEvents = "PIPELINE_EXECUTION_RESUMED"
	// Trigger notification when pipeline execution succeeded.
	PipelineNotificationEvents_PIPELINE_EXECUTION_SUCCEEDED PipelineNotificationEvents = "PIPELINE_EXECUTION_SUCCEEDED"
	// Trigger notification when pipeline execution superseded.
	PipelineNotificationEvents_PIPELINE_EXECUTION_SUPERSEDED PipelineNotificationEvents = "PIPELINE_EXECUTION_SUPERSEDED"
	// Trigger notification when pipeline stage execution started.
	PipelineNotificationEvents_STAGE_EXECUTION_STARTED PipelineNotificationEvents = "STAGE_EXECUTION_STARTED"
	// Trigger notification when pipeline stage execution succeeded.
	PipelineNotificationEvents_STAGE_EXECUTION_SUCCEEDED PipelineNotificationEvents = "STAGE_EXECUTION_SUCCEEDED"
	// Trigger notification when pipeline stage execution resumed.
	PipelineNotificationEvents_STAGE_EXECUTION_RESUMED PipelineNotificationEvents = "STAGE_EXECUTION_RESUMED"
	// Trigger notification when pipeline stage execution canceled.
	PipelineNotificationEvents_STAGE_EXECUTION_CANCELED PipelineNotificationEvents = "STAGE_EXECUTION_CANCELED"
	// Trigger notification when pipeline stage execution failed.
	PipelineNotificationEvents_STAGE_EXECUTION_FAILED PipelineNotificationEvents = "STAGE_EXECUTION_FAILED"
	// Trigger notification when pipeline action execution succeeded.
	PipelineNotificationEvents_ACTION_EXECUTION_SUCCEEDED PipelineNotificationEvents = "ACTION_EXECUTION_SUCCEEDED"
	// Trigger notification when pipeline action execution failed.
	PipelineNotificationEvents_ACTION_EXECUTION_FAILED PipelineNotificationEvents = "ACTION_EXECUTION_FAILED"
	// Trigger notification when pipeline action execution canceled.
	PipelineNotificationEvents_ACTION_EXECUTION_CANCELED PipelineNotificationEvents = "ACTION_EXECUTION_CANCELED"
	// Trigger notification when pipeline action execution started.
	PipelineNotificationEvents_ACTION_EXECUTION_STARTED PipelineNotificationEvents = "ACTION_EXECUTION_STARTED"
	// Trigger notification when pipeline manual approval failed.
	PipelineNotificationEvents_MANUAL_APPROVAL_FAILED PipelineNotificationEvents = "MANUAL_APPROVAL_FAILED"
	// Trigger notification when pipeline manual approval needed.
	PipelineNotificationEvents_MANUAL_APPROVAL_NEEDED PipelineNotificationEvents = "MANUAL_APPROVAL_NEEDED"
	// Trigger notification when pipeline manual approval succeeded.
	PipelineNotificationEvents_MANUAL_APPROVAL_SUCCEEDED PipelineNotificationEvents = "MANUAL_APPROVAL_SUCCEEDED"
)

// Additional options to pass to the notification rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineNotifyOnOptions := &pipelineNotifyOnOptions{
//   	events: []pipelineNotificationEvents{
//   		awscdk.Aws_codepipeline.*pipelineNotificationEvents_PIPELINE_EXECUTION_FAILED,
//   	},
//
//   	// the properties below are optional
//   	detailType: awscdk.Aws_codestarnotifications.detailType_BASIC,
//   	enabled: jsii.Boolean(false),
//   	notificationRuleName: jsii.String("notificationRuleName"),
//   }
//
type PipelineNotifyOnOptions struct {
	// The level of detail to include in the notifications for this resource.
	//
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	DetailType awscodestarnotifications.DetailType `field:"optional" json:"detailType" yaml:"detailType"`
	// The status of the notification rule.
	//
	// If the enabled is set to DISABLED, notifications aren't sent for the notification rule.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account.
	NotificationRuleName *string `field:"optional" json:"notificationRuleName" yaml:"notificationRuleName"`
	// A list of event types associated with this notification rule for CodePipeline Pipeline.
	//
	// For a complete list of event types and IDs, see Notification concepts in the Developer Tools Console User Guide.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#concepts-api
	//
	Events *[]PipelineNotificationEvents `field:"required" json:"events" yaml:"events"`
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var project pipelineProject
//
//   repository := codecommit.NewRepository(this, jsii.String("MyRepository"), &repositoryProps{
//   	repositoryName: jsii.String("MyRepository"),
//   })
//   project := codebuild.NewPipelineProject(this, jsii.String("MyProject"))
//
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("CodeCommit"),
//   	repository: repository,
//   	output: sourceOutput,
//   })
//   buildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("CodeBuild"),
//   	project: project,
//   	input: sourceOutput,
//   	outputs: []artifact{
//   		codepipeline.NewArtifact(),
//   	},
//   	 // optional
//   	executeBatchBuild: jsii.Boolean(true),
//   	 // optional, defaults to false
//   	combineBatchBuildArtifacts: jsii.Boolean(true),
//   })
//
//   codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		&stageProps{
//   			stageName: jsii.String("Source"),
//   			actions: []iAction{
//   				sourceAction,
//   			},
//   		},
//   		&stageProps{
//   			stageName: jsii.String("Build"),
//   			actions: []*iAction{
//   				buildAction,
//   			},
//   		},
//   	},
//   })
//
type PipelineProps struct {
	// The S3 bucket used by this Pipeline to store artifacts.
	ArtifactBucket awss3.IBucket `field:"optional" json:"artifactBucket" yaml:"artifactBucket"`
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
	CrossAccountKeys *bool `field:"optional" json:"crossAccountKeys" yaml:"crossAccountKeys"`
	// A map of region to S3 bucket name used for cross-region CodePipeline.
	//
	// For every Action that you specify targeting a different region than the Pipeline itself,
	// if you don't provide an explicit Bucket for that region using this property,
	// the construct will automatically create a Stack containing an S3 Bucket in that region.
	CrossRegionReplicationBuckets *map[string]awss3.IBucket `field:"optional" json:"crossRegionReplicationBuckets" yaml:"crossRegionReplicationBuckets"`
	// Enable KMS key rotation for the generated KMS keys.
	//
	// By default KMS key rotation is disabled, but will add an additional $1/month
	// for each year the key exists when enabled.
	EnableKeyRotation *bool `field:"optional" json:"enableKeyRotation" yaml:"enableKeyRotation"`
	// Name of the pipeline.
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// Indicates whether to rerun the AWS CodePipeline pipeline after you update it.
	RestartExecutionOnUpdate *bool `field:"optional" json:"restartExecutionOnUpdate" yaml:"restartExecutionOnUpdate"`
	// Reuse the same cross region support stack for all pipelines in the App.
	ReuseCrossRegionSupportStacks *bool `field:"optional" json:"reuseCrossRegionSupportStacks" yaml:"reuseCrossRegionSupportStacks"`
	// The IAM role to be assumed by this Pipeline.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The list of Stages, in order, to create this Pipeline with.
	//
	// You can always add more Stages later by calling {@link Pipeline#addStage}.
	Stages *[]*StageProps `field:"optional" json:"stages" yaml:"stages"`
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import stepfunctions "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   startState := stepfunctions.NewPass(this, jsii.String("StartState"))
//   simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &stateMachineProps{
//   	definition: startState,
//   })
//   stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&stepFunctionsInvokeActionProps{
//   	actionName: jsii.String("Invoke"),
//   	stateMachine: simpleStateMachine,
//   	stateMachineInput: codepipeline_actions.stateMachineInput.literal(map[string]*bool{
//   		"IsHelloWorldExample": jsii.Boolean(true),
//   	}),
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("StepFunctions"),
//   	actions: []iAction{
//   		stepFunctionAction,
//   	},
//   })
//
type StageOptions struct {
	// The physical, human-readable name to assign to this Pipeline Stage.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The list of Actions to create this Stage with.
	//
	// You can always add more Actions later by calling {@link IStage#addAction}.
	Actions *[]IAction `field:"optional" json:"actions" yaml:"actions"`
	// The reason for disabling transition to this stage.
	//
	// Only applicable
	// if `transitionToEnabled` is set to `false`.
	TransitionDisabledReason *string `field:"optional" json:"transitionDisabledReason" yaml:"transitionDisabledReason"`
	// Whether to enable transition to this stage.
	TransitionToEnabled *bool `field:"optional" json:"transitionToEnabled" yaml:"transitionToEnabled"`
	Placement *StagePlacement `field:"optional" json:"placement" yaml:"placement"`
}

// Allows you to control where to place a new Stage when it's added to the Pipeline.
//
// Note that you can provide only one of the below properties -
// specifying more than one will result in a validation error.
//
// Example:
//   // Insert a new Stage at an arbitrary point
//   var pipeline pipeline
//   var anotherStage iStage
//   var yetAnotherStage iStage
//
//
//   someStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("SomeStage"),
//   	placement: &stagePlacement{
//   		// note: you can only specify one of the below properties
//   		rightBefore: anotherStage,
//   		justAfter: yetAnotherStage,
//   	},
//   })
//
// See: #justAfter.
//
type StagePlacement struct {
	// Inserts the new Stage as a child of the given Stage (changing its current child Stage, if it had one).
	JustAfter IStage `field:"optional" json:"justAfter" yaml:"justAfter"`
	// Inserts the new Stage as a parent of the given Stage (changing its current parent Stage, if it had one).
	RightBefore IStage `field:"optional" json:"rightBefore" yaml:"rightBefore"`
}

// Construction properties of a Pipeline Stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var action action
//
//   stageProps := &stageProps{
//   	stageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	actions: []iAction{
//   		action,
//   	},
//   	transitionDisabledReason: jsii.String("transitionDisabledReason"),
//   	transitionToEnabled: jsii.Boolean(false),
//   }
//
type StageProps struct {
	// The physical, human-readable name to assign to this Pipeline Stage.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The list of Actions to create this Stage with.
	//
	// You can always add more Actions later by calling {@link IStage#addAction}.
	Actions *[]IAction `field:"optional" json:"actions" yaml:"actions"`
	// The reason for disabling transition to this stage.
	//
	// Only applicable
	// if `transitionToEnabled` is set to `false`.
	TransitionDisabledReason *string `field:"optional" json:"transitionDisabledReason" yaml:"transitionDisabledReason"`
	// Whether to enable transition to this stage.
	TransitionToEnabled *bool `field:"optional" json:"transitionToEnabled" yaml:"transitionToEnabled"`
}

