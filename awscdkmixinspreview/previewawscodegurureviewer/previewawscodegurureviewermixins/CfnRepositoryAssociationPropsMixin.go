package previewawscodegurureviewermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscodegurureviewer/previewawscodegurureviewermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// This resource configures how Amazon CodeGuru Reviewer retrieves the source code to be reviewed.
//
// You can use an AWS CloudFormation template to create an association with the following repository types:
//
// - AWS CodeCommit - For more information, see [Create an AWS CodeCommit repository association](https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/create-codecommit-association.html) in the *Amazon CodeGuru Reviewer User Guide* .
// - Bitbucket - For more information, see [Create a Bitbucket repository association](https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/create-bitbucket-association.html) in the *Amazon CodeGuru Reviewer User Guide* .
// - GitHub Enterprise Server - For more information, see [Create a GitHub Enterprise Server repository association](https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/create-github-enterprise-association.html) in the *Amazon CodeGuru Reviewer User Guide* .
// - S3Bucket - For more information, see [Create code reviews with GitHub Actions](https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/working-with-cicd.html) in the *Amazon CodeGuru Reviewer User Guide* .
//
// > You cannot use a CloudFormation template to create an association with a GitHub repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRepositoryAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnRepositoryAssociationPropsMixin(&CfnRepositoryAssociationMixinProps{
//   	BucketName: jsii.String("bucketName"),
//   	ConnectionArn: jsii.String("connectionArn"),
//   	Name: jsii.String("name"),
//   	Owner: jsii.String("owner"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codegurureviewer-repositoryassociation.html
//
type CfnRepositoryAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRepositoryAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRepositoryAssociationPropsMixin
type jsiiProxy_CfnRepositoryAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRepositoryAssociationPropsMixin) Props() *CfnRepositoryAssociationMixinProps {
	var returns *CfnRepositoryAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepositoryAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodeGuruReviewer::RepositoryAssociation`.
func NewCfnRepositoryAssociationPropsMixin(props *CfnRepositoryAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRepositoryAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRepositoryAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRepositoryAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codegurureviewer.mixins.CfnRepositoryAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodeGuruReviewer::RepositoryAssociation`.
func NewCfnRepositoryAssociationPropsMixin_Override(c CfnRepositoryAssociationPropsMixin, props *CfnRepositoryAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codegurureviewer.mixins.CfnRepositoryAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRepositoryAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRepositoryAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codegurureviewer.mixins.CfnRepositoryAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRepositoryAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codegurureviewer.mixins.CfnRepositoryAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRepositoryAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRepositoryAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

