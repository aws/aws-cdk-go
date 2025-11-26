package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SageMaker::NotebookInstance` resource creates an Amazon SageMaker notebook instance.
//
// A notebook instance is a machine learning (ML) compute instance running on a Jupyter notebook. For more information, see [Use Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNotebookInstancePropsMixin := awscdkmixinspreview.Mixins.NewCfnNotebookInstancePropsMixin(&CfnNotebookInstanceMixinProps{
//   	AcceleratorTypes: []*string{
//   		jsii.String("acceleratorTypes"),
//   	},
//   	AdditionalCodeRepositories: []*string{
//   		jsii.String("additionalCodeRepositories"),
//   	},
//   	DefaultCodeRepository: jsii.String("defaultCodeRepository"),
//   	DirectInternetAccess: jsii.String("directInternetAccess"),
//   	InstanceMetadataServiceConfiguration: &InstanceMetadataServiceConfigurationProperty{
//   		MinimumInstanceMetadataServiceVersion: jsii.String("minimumInstanceMetadataServiceVersion"),
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LifecycleConfigName: jsii.String("lifecycleConfigName"),
//   	NotebookInstanceName: jsii.String("notebookInstanceName"),
//   	PlatformIdentifier: jsii.String("platformIdentifier"),
//   	RoleArn: jsii.String("roleArn"),
//   	RootAccess: jsii.String("rootAccess"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VolumeSizeInGb: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html
//
type CfnNotebookInstancePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNotebookInstanceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNotebookInstancePropsMixin
type jsiiProxy_CfnNotebookInstancePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnNotebookInstancePropsMixin) Props() *CfnNotebookInstanceMixinProps {
	var returns *CfnNotebookInstanceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstancePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::NotebookInstance`.
func NewCfnNotebookInstancePropsMixin(props *CfnNotebookInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNotebookInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNotebookInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNotebookInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnNotebookInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::NotebookInstance`.
func NewCfnNotebookInstancePropsMixin_Override(c CfnNotebookInstancePropsMixin, props *CfnNotebookInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnNotebookInstancePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNotebookInstancePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNotebookInstancePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnNotebookInstancePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNotebookInstancePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnNotebookInstancePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNotebookInstancePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnNotebookInstancePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

