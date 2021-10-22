// The CDK Construct Library for AWS::CodeStar
package awscdkcodestaralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcodestaralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdkcodestaralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The GitHubRepository resource.
// Experimental.
type GitHubRepository interface {
	awscdk.Resource
	IGitHubRepository
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	Owner() *string
	PhysicalName() *string
	Repo() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for GitHubRepository
type jsiiProxy_GitHubRepository struct {
	internal.Type__awscdkResource
	jsiiProxy_IGitHubRepository
}

func (j *jsiiProxy_GitHubRepository) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Repo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitHubRepository(scope constructs.Construct, id *string, props *GitHubRepositoryProps) GitHubRepository {
	_init_.Initialize()

	j := jsiiProxy_GitHubRepository{}

	_jsii_.Create(
		"@aws-cdk/aws-codestar-alpha.GitHubRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubRepository_Override(g GitHubRepository, scope constructs.Construct, id *string, props *GitHubRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-codestar-alpha.GitHubRepository",
		[]interface{}{scope, id, props},
		g,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GitHubRepository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-codestar-alpha.GitHubRepository",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func GitHubRepository_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-codestar-alpha.GitHubRepository",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
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
func (g *jsiiProxy_GitHubRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (g *jsiiProxy_GitHubRepository) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
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
func (g *jsiiProxy_GitHubRepository) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		g,
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
func (g *jsiiProxy_GitHubRepository) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (g *jsiiProxy_GitHubRepository) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties of {@link GitHubRepository}.
// Experimental.
type GitHubRepositoryProps struct {
	// The GitHub user's personal access token for the GitHub repository.
	// Experimental.
	AccessToken awscdk.SecretValue `json:"accessToken"`
	// The name of the Amazon S3 bucket that contains the ZIP file with the content to be committed to the new repository.
	// Experimental.
	ContentsBucket awss3.IBucket `json:"contentsBucket"`
	// The S3 object key or file name for the ZIP file.
	// Experimental.
	ContentsKey *string `json:"contentsKey"`
	// The GitHub user name for the owner of the GitHub repository to be created.
	//
	// If this
	// repository should be owned by a GitHub organization, provide its name
	// Experimental.
	Owner *string `json:"owner"`
	// The name of the repository you want to create in GitHub with AWS CloudFormation stack creation.
	// Experimental.
	RepositoryName *string `json:"repositoryName"`
	// The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket.
	// Experimental.
	ContentsS3Version *string `json:"contentsS3Version"`
	// A comment or description about the new repository.
	//
	// This description is displayed in GitHub after the repository
	// is created.
	// Experimental.
	Description *string `json:"description"`
	// Indicates whether to enable issues for the GitHub repository.
	//
	// You can use GitHub issues to track information
	// and bugs for your repository.
	// Experimental.
	EnableIssues *bool `json:"enableIssues"`
	// Indicates whether the GitHub repository is a private repository.
	//
	// If so, you choose who can see and commit to
	// this repository.
	// Experimental.
	Visibility RepositoryVisibility `json:"visibility"`
}

// GitHubRepository resource interface.
// Experimental.
type IGitHubRepository interface {
	awscdk.IResource
	// the repository owner.
	// Experimental.
	Owner() *string
	// the repository name.
	// Experimental.
	Repo() *string
}

// The jsii proxy for IGitHubRepository
type jsiiProxy_IGitHubRepository struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IGitHubRepository) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGitHubRepository) Repo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repo",
		&returns,
	)
	return returns
}

// Visibility of the GitHubRepository.
// Experimental.
type RepositoryVisibility string

const (
	RepositoryVisibility_PRIVATE RepositoryVisibility = "PRIVATE"
	RepositoryVisibility_PUBLIC RepositoryVisibility = "PUBLIC"
)

