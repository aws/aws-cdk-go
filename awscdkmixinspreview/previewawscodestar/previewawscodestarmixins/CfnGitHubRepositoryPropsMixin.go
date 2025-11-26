package previewawscodestarmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscodestar/previewawscodestarmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CodeStar::GitHubRepository` resource creates a GitHub repository where users can store source code for use with AWS workflows.
//
// You must provide a location for the source code ZIP file in the CloudFormation template, so the code can be uploaded to the created repository. You must have created a personal access token in GitHub to provide in the CloudFormation template. AWS uses this token to connect to GitHub on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGitHubRepositoryPropsMixin := awscdkmixinspreview.Mixins.NewCfnGitHubRepositoryPropsMixin(&CfnGitHubRepositoryMixinProps{
//   	Code: &CodeProperty{
//   		S3: &S3Property{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   			ObjectVersion: jsii.String("objectVersion"),
//   		},
//   	},
//   	ConnectionArn: jsii.String("connectionArn"),
//   	EnableIssues: jsii.Boolean(false),
//   	IsPrivate: jsii.Boolean(false),
//   	RepositoryAccessToken: jsii.String("repositoryAccessToken"),
//   	RepositoryDescription: jsii.String("repositoryDescription"),
//   	RepositoryName: jsii.String("repositoryName"),
//   	RepositoryOwner: jsii.String("repositoryOwner"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html
//
type CfnGitHubRepositoryPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGitHubRepositoryMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGitHubRepositoryPropsMixin
type jsiiProxy_CfnGitHubRepositoryPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGitHubRepositoryPropsMixin) Props() *CfnGitHubRepositoryMixinProps {
	var returns *CfnGitHubRepositoryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepositoryPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodeStar::GitHubRepository`.
func NewCfnGitHubRepositoryPropsMixin(props *CfnGitHubRepositoryMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGitHubRepositoryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGitHubRepositoryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGitHubRepositoryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codestar.mixins.CfnGitHubRepositoryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodeStar::GitHubRepository`.
func NewCfnGitHubRepositoryPropsMixin_Override(c CfnGitHubRepositoryPropsMixin, props *CfnGitHubRepositoryMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codestar.mixins.CfnGitHubRepositoryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGitHubRepositoryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGitHubRepositoryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codestar.mixins.CfnGitHubRepositoryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGitHubRepositoryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codestar.mixins.CfnGitHubRepositoryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGitHubRepositoryPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGitHubRepositoryPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

