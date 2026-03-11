package awsecr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The details of the repository creation template associated with the request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRepositoryCreationTemplatePropsMixin := awscdkcfnpropertymixins.Aws_ecr.NewCfnRepositoryCreationTemplatePropsMixin(&CfnRepositoryCreationTemplateMixinProps{
//   	AppliedFor: []*string{
//   		jsii.String("appliedFor"),
//   	},
//   	CustomRoleArn: jsii.String("customRoleArn"),
//   	Description: jsii.String("description"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	ImageTagMutability: jsii.String("imageTagMutability"),
//   	ImageTagMutabilityExclusionFilters: []interface{}{
//   		&ImageTagMutabilityExclusionFilterProperty{
//   			ImageTagMutabilityExclusionFilterType: jsii.String("imageTagMutabilityExclusionFilterType"),
//   			ImageTagMutabilityExclusionFilterValue: jsii.String("imageTagMutabilityExclusionFilterValue"),
//   		},
//   	},
//   	LifecyclePolicy: jsii.String("lifecyclePolicy"),
//   	Prefix: jsii.String("prefix"),
//   	RepositoryPolicy: jsii.String("repositoryPolicy"),
//   	ResourceTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html
//
type CfnRepositoryCreationTemplatePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRepositoryCreationTemplateMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRepositoryCreationTemplatePropsMixin
type jsiiProxy_CfnRepositoryCreationTemplatePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRepositoryCreationTemplatePropsMixin) Props() *CfnRepositoryCreationTemplateMixinProps {
	var returns *CfnRepositoryCreationTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepositoryCreationTemplatePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ECR::RepositoryCreationTemplate`.
func NewCfnRepositoryCreationTemplatePropsMixin(props *CfnRepositoryCreationTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRepositoryCreationTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRepositoryCreationTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRepositoryCreationTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ecr.CfnRepositoryCreationTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ECR::RepositoryCreationTemplate`.
func NewCfnRepositoryCreationTemplatePropsMixin_Override(c CfnRepositoryCreationTemplatePropsMixin, props *CfnRepositoryCreationTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ecr.CfnRepositoryCreationTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRepositoryCreationTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRepositoryCreationTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ecr.CfnRepositoryCreationTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRepositoryCreationTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ecr.CfnRepositoryCreationTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRepositoryCreationTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRepositoryCreationTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

