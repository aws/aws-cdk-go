package previewawsmediatailormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmediatailor/previewawsmediatailormixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A source location is a container for sources.
//
// For more information about source locations, see [Working with source locations](https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-source-locations.html) in the *MediaTailor User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSourceLocationPropsMixin := awscdkmixinspreview.Mixins.NewCfnSourceLocationPropsMixin(&CfnSourceLocationMixinProps{
//   	AccessConfiguration: &AccessConfigurationProperty{
//   		AccessType: jsii.String("accessType"),
//   		SecretsManagerAccessTokenConfiguration: &SecretsManagerAccessTokenConfigurationProperty{
//   			HeaderName: jsii.String("headerName"),
//   			SecretArn: jsii.String("secretArn"),
//   			SecretStringKey: jsii.String("secretStringKey"),
//   		},
//   	},
//   	DefaultSegmentDeliveryConfiguration: &DefaultSegmentDeliveryConfigurationProperty{
//   		BaseUrl: jsii.String("baseUrl"),
//   	},
//   	HttpConfiguration: &HttpConfigurationProperty{
//   		BaseUrl: jsii.String("baseUrl"),
//   	},
//   	SegmentDeliveryConfigurations: []interface{}{
//   		&SegmentDeliveryConfigurationProperty{
//   			BaseUrl: jsii.String("baseUrl"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html
//
type CfnSourceLocationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSourceLocationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSourceLocationPropsMixin
type jsiiProxy_CfnSourceLocationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSourceLocationPropsMixin) Props() *CfnSourceLocationMixinProps {
	var returns *CfnSourceLocationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceLocationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaTailor::SourceLocation`.
func NewCfnSourceLocationPropsMixin(props *CfnSourceLocationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSourceLocationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSourceLocationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSourceLocationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaTailor::SourceLocation`.
func NewCfnSourceLocationPropsMixin_Override(c CfnSourceLocationPropsMixin, props *CfnSourceLocationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSourceLocationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSourceLocationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSourceLocationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSourceLocationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSourceLocationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

