package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
)

// Deployment of a single Stack.
//
// You don't need to instantiate this class -- it will
// be automatically instantiated as necessary when you
// add a `Stage` to a pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cloudFormationStackArtifact cloudFormationStackArtifact
//
//   stackDeployment := awscdk.Pipelines.stackDeployment.fromArtifact(cloudFormationStackArtifact)
//
type StackDeployment interface {
	// Template path on disk to CloudAssembly.
	AbsoluteTemplatePath() *string
	// Account where the stack should be deployed.
	Account() *string
	// Assets referenced by this stack.
	Assets() *[]*StackAsset
	// Role to assume before deploying this stack.
	AssumeRoleArn() *string
	// Steps that take place after stack is prepared but before stack deploys.
	//
	// Your pipeline engine may not disable `prepareStep`.
	ChangeSet() *[]Step
	// Construct path for this stack.
	ConstructPath() *string
	// Execution role to pass to CloudFormation.
	ExecutionRoleArn() *string
	// Steps to execute after stack deploys.
	Post() *[]Step
	// Steps that take place before stack is prepared.
	//
	// If your pipeline engine disables 'prepareStep', then this will happen before stack deploys.
	Pre() *[]Step
	// Region where the stack should be deployed.
	Region() *string
	// Artifact ID for this stack.
	StackArtifactId() *string
	// Other stacks this stack depends on.
	StackDependencies() *[]StackDeployment
	// Name for this stack.
	StackName() *string
	// Tags to apply to the stack.
	Tags() *map[string]*string
	// The asset that represents the CloudFormation template for this stack.
	TemplateAsset() *StackAsset
	// The S3 URL which points to the template asset location in the publishing bucket.
	//
	// This is `undefined` if the stack template is not published. Use the
	// `DefaultStackSynthesizer` to ensure it is.
	//
	// Example value: `https://bucket.s3.amazonaws.com/object/key`
	TemplateUrl() *string
	// Add a dependency on another stack.
	AddStackDependency(stackDeployment StackDeployment)
	// Adds steps to each phase of the stack.
	AddStackSteps(pre *[]Step, changeSet *[]Step, post *[]Step)
}

// The jsii proxy struct for StackDeployment
type jsiiProxy_StackDeployment struct {
	_ byte // padding
}

func (j *jsiiProxy_StackDeployment) AbsoluteTemplatePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absoluteTemplatePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) Assets() *[]*StackAsset {
	var returns *[]*StackAsset
	_jsii_.Get(
		j,
		"assets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) AssumeRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) ChangeSet() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"changeSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) ConstructPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) Post() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"post",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) Pre() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"pre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) StackArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackArtifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) StackDependencies() *[]StackDeployment {
	var returns *[]StackDeployment
	_jsii_.Get(
		j,
		"stackDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) TemplateAsset() *StackAsset {
	var returns *StackAsset
	_jsii_.Get(
		j,
		"templateAsset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) TemplateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrl",
		&returns,
	)
	return returns
}


// Build a `StackDeployment` from a Stack Artifact in a Cloud Assembly.
func StackDeployment_FromArtifact(stackArtifact cxapi.CloudFormationStackArtifact) StackDeployment {
	_init_.Initialize()

	var returns StackDeployment

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.StackDeployment",
		"fromArtifact",
		[]interface{}{stackArtifact},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackDeployment) AddStackDependency(stackDeployment StackDeployment) {
	_jsii_.InvokeVoid(
		s,
		"addStackDependency",
		[]interface{}{stackDeployment},
	)
}

func (s *jsiiProxy_StackDeployment) AddStackSteps(pre *[]Step, changeSet *[]Step, post *[]Step) {
	_jsii_.InvokeVoid(
		s,
		"addStackSteps",
		[]interface{}{pre, changeSet, post},
	)
}

