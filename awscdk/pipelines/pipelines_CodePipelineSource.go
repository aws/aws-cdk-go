package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipelineactions"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Factory for CodePipeline source steps.
//
// This class contains a number of factory methods for the different types
// of sources that CodePipeline supports.
//
// Example:
//   // Access the CommitId of a GitHub source in the synth
//   source := pipelines.codePipelineSource.gitHub(jsii.String("owner/repo"), jsii.String("main"))
//
//   pipeline := pipelines.NewCodePipeline(*scope, jsii.String("MyPipeline"), &codePipelineProps{
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: source,
//   		commands: []*string{
//   		},
//   		env: map[string]*string{
//   			"COMMIT_ID": source.sourceAttribute(jsii.String("CommitId")),
//   		},
//   	}),
//   })
//
// Experimental.
type CodePipelineSource interface {
	Step
	ICodePipelineActionFactory
	// Return the steps this step depends on, based on the FileSets it requires.
	// Experimental.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	// Experimental.
	DependencyFileSets() *[]FileSet
	// Identifier for this step.
	// Experimental.
	Id() *string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	// Experimental.
	IsSource() *bool
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	// Experimental.
	PrimaryOutput() FileSet
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	// Experimental.
	AddDependencyFileSet(fs FileSet)
	// Add a dependency on another step.
	// Experimental.
	AddStepDependency(step Step)
	// Configure the given FileSet as the primary output of this step.
	// Experimental.
	ConfigurePrimaryOutput(fs FileSet)
	// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
	//
	// Should be called in the constructor of subclasses based on what the user
	// passes in as construction properties. The format of the structure passed in
	// here does not have to correspond exactly to what gets rendered into the
	// engine, it just needs to contain the same data.
	// Experimental.
	DiscoverReferencedOutputs(structure interface{})
	// Experimental.
	GetAction(output awscodepipeline.Artifact, actionName *string, runOrder *float64, variablesNamespace *string) awscodepipelineactions.Action
	// Create the desired Action and add it to the pipeline.
	// Experimental.
	ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult
	// Return an attribute of the current source revision.
	//
	// These values can be passed into the environment variables of pipeline steps,
	// so your steps can access information about the source revision.
	//
	// Pipeline synth step has some source attributes predefined in the environment.
	// If these suffice, you don't need to use this method for the synth step.
	//
	// Example:
	//   // Access the CommitId of a GitHub source in the synth
	//   source := pipelines.codePipelineSource.gitHub(jsii.String("owner/repo"), jsii.String("main"))
	//
	//   pipeline := pipelines.NewCodePipeline(*scope, jsii.String("MyPipeline"), &codePipelineProps{
	//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
	//   		input: source,
	//   		commands: []*string{
	//   		},
	//   		env: map[string]*string{
	//   			"COMMIT_ID": source.sourceAttribute(jsii.String("CommitId")),
	//   		},
	//   	}),
	//   })
	//
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-variables.html#reference-variables-list
	//
	// Experimental.
	SourceAttribute(name *string) *string
	// Return a string representation of this Step.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodePipelineSource
type jsiiProxy_CodePipelineSource struct {
	jsiiProxy_Step
	jsiiProxy_ICodePipelineActionFactory
}

func (j *jsiiProxy_CodePipelineSource) Dependencies() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineSource) DependencyFileSets() *[]FileSet {
	var returns *[]FileSet
	_jsii_.Get(
		j,
		"dependencyFileSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineSource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineSource) IsSource() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineSource) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}


// Experimental.
func NewCodePipelineSource_Override(c CodePipelineSource, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.CodePipelineSource",
		[]interface{}{id},
		c,
	)
}

// Returns a CodeCommit source.
//
// If you need access to symlinks or the repository history, be sure to set
// `codeBuildCloneOutput`.
//
// Example:
//   var repository iRepository
//
//   pipelines.codePipelineSource.codeCommit(repository, jsii.String("main"))
//
// Experimental.
func CodePipelineSource_CodeCommit(repository awscodecommit.IRepository, branch *string, props *CodeCommitSourceOptions) CodePipelineSource {
	_init_.Initialize()

	if err := validateCodePipelineSource_CodeCommitParameters(repository, branch, props); err != nil {
		panic(err)
	}
	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"codeCommit",
		[]interface{}{repository, branch, props},
		&returns,
	)

	return returns
}

// Returns a CodeStar connection source.
//
// A CodeStar connection allows AWS CodePipeline to
// access external resources, such as repositories in GitHub, GitHub Enterprise or
// BitBucket.
//
// To use this method, you first need to create a CodeStar connection
// using the AWS console. In the process, you may have to sign in to the external provider
// -- GitHub, for example -- to authorize AWS to read and modify your repository.
// Once you have done this, copy the connection ARN and use it to create the source.
//
// Example:
//
// ```ts
// pipelines.CodePipelineSource.connection('owner/repo', 'main', {
//    connectionArn: 'arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41', // Created using the AWS console
// });
// ```
//
// If you need access to symlinks or the repository history, be sure to set
// `codeBuildCloneOutput`.
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/welcome-connections.html
//
// Experimental.
func CodePipelineSource_Connection(repoString *string, branch *string, props *ConnectionSourceOptions) CodePipelineSource {
	_init_.Initialize()

	if err := validateCodePipelineSource_ConnectionParameters(repoString, branch, props); err != nil {
		panic(err)
	}
	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"connection",
		[]interface{}{repoString, branch, props},
		&returns,
	)

	return returns
}

// Returns an ECR source.
//
// Example:
//   var repository iRepository
//
//   pipelines.codePipelineSource.ecr(repository, &eCRSourceOptions{
//   	imageTag: jsii.String("latest"),
//   })
//
// Experimental.
func CodePipelineSource_Ecr(repository awsecr.IRepository, props *ECRSourceOptions) CodePipelineSource {
	_init_.Initialize()

	if err := validateCodePipelineSource_EcrParameters(repository, props); err != nil {
		panic(err)
	}
	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"ecr",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Returns a GitHub source, using OAuth tokens to authenticate with GitHub and a separate webhook to detect changes.
//
// This is no longer
// the recommended method. Please consider using `connection()`
// instead.
//
// Pass in the owner and repository in a single string, like this:
//
// ```ts
// pipelines.CodePipelineSource.gitHub('owner/repo', 'main');
// ```
//
// Authentication will be done by a secret called `github-token` in AWS
// Secrets Manager (unless specified otherwise).
//
// The token should have these permissions:
//
// * **repo** - to read the repository
// * **admin:repo_hook** - if you plan to use webhooks (true by default)
//
// If you need access to symlinks or the repository history, use a source of type
// `connection` instead.
// Experimental.
func CodePipelineSource_GitHub(repoString *string, branch *string, props *GitHubSourceOptions) CodePipelineSource {
	_init_.Initialize()

	if err := validateCodePipelineSource_GitHubParameters(repoString, branch, props); err != nil {
		panic(err)
	}
	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"gitHub",
		[]interface{}{repoString, branch, props},
		&returns,
	)

	return returns
}

// Returns an S3 source.
//
// Example:
//   var bucket bucket
//
//   pipelines.codePipelineSource.s3(bucket, jsii.String("path/to/file.zip"))
//
// Experimental.
func CodePipelineSource_S3(bucket awss3.IBucket, objectKey *string, props *S3SourceOptions) CodePipelineSource {
	_init_.Initialize()

	if err := validateCodePipelineSource_S3Parameters(bucket, objectKey, props); err != nil {
		panic(err)
	}
	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"s3",
		[]interface{}{bucket, objectKey, props},
		&returns,
	)

	return returns
}

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
// Experimental.
func CodePipelineSource_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	if err := validateCodePipelineSource_SequenceParameters(steps); err != nil {
		panic(err)
	}
	var returns *[]Step

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodePipelineSource) AddDependencyFileSet(fs FileSet) {
	if err := c.validateAddDependencyFileSetParameters(fs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

func (c *jsiiProxy_CodePipelineSource) AddStepDependency(step Step) {
	if err := c.validateAddStepDependencyParameters(step); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addStepDependency",
		[]interface{}{step},
	)
}

func (c *jsiiProxy_CodePipelineSource) ConfigurePrimaryOutput(fs FileSet) {
	if err := c.validateConfigurePrimaryOutputParameters(fs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

func (c *jsiiProxy_CodePipelineSource) DiscoverReferencedOutputs(structure interface{}) {
	if err := c.validateDiscoverReferencedOutputsParameters(structure); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

func (c *jsiiProxy_CodePipelineSource) GetAction(output awscodepipeline.Artifact, actionName *string, runOrder *float64, variablesNamespace *string) awscodepipelineactions.Action {
	if err := c.validateGetActionParameters(output, actionName, runOrder); err != nil {
		panic(err)
	}
	var returns awscodepipelineactions.Action

	_jsii_.Invoke(
		c,
		"getAction",
		[]interface{}{output, actionName, runOrder, variablesNamespace},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodePipelineSource) ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult {
	if err := c.validateProduceActionParameters(stage, options); err != nil {
		panic(err)
	}
	var returns *CodePipelineActionFactoryResult

	_jsii_.Invoke(
		c,
		"produceAction",
		[]interface{}{stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodePipelineSource) SourceAttribute(name *string) *string {
	if err := c.validateSourceAttributeParameters(name); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"sourceAttribute",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodePipelineSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

