package previewawsdeadlinemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdeadline/previewawsdeadlinemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a storage profile that specifies the operating system, file type, and file location of resources used on a farm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStorageProfilePropsMixin := awscdkmixinspreview.Mixins.NewCfnStorageProfilePropsMixin(&CfnStorageProfileMixinProps{
//   	DisplayName: jsii.String("displayName"),
//   	FarmId: jsii.String("farmId"),
//   	FileSystemLocations: []interface{}{
//   		&FileSystemLocationProperty{
//   			Name: jsii.String("name"),
//   			Path: jsii.String("path"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	OsFamily: jsii.String("osFamily"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-storageprofile.html
//
type CfnStorageProfilePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStorageProfileMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStorageProfilePropsMixin
type jsiiProxy_CfnStorageProfilePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStorageProfilePropsMixin) Props() *CfnStorageProfileMixinProps {
	var returns *CfnStorageProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageProfilePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Deadline::StorageProfile`.
func NewCfnStorageProfilePropsMixin(props *CfnStorageProfileMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStorageProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStorageProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStorageProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnStorageProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Deadline::StorageProfile`.
func NewCfnStorageProfilePropsMixin_Override(c CfnStorageProfilePropsMixin, props *CfnStorageProfileMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnStorageProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStorageProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStorageProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnStorageProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStorageProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnStorageProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStorageProfilePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnStorageProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

