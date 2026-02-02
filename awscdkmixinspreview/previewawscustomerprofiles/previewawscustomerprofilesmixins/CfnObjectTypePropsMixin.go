package previewawscustomerprofilesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscustomerprofiles/previewawscustomerprofilesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an Amazon Connect Customer Profiles Object Type Mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnObjectTypePropsMixin := awscdkmixinspreview.Mixins.NewCfnObjectTypePropsMixin(&CfnObjectTypeMixinProps{
//   	AllowProfileCreation: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DomainName: jsii.String("domainName"),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	ExpirationDays: jsii.Number(123),
//   	Fields: []interface{}{
//   		&FieldMapProperty{
//   			Name: jsii.String("name"),
//   			ObjectTypeField: &ObjectTypeFieldProperty{
//   				ContentType: jsii.String("contentType"),
//   				Source: jsii.String("source"),
//   				Target: jsii.String("target"),
//   			},
//   		},
//   	},
//   	Keys: []interface{}{
//   		&KeyMapProperty{
//   			Name: jsii.String("name"),
//   			ObjectTypeKeyList: []interface{}{
//   				&ObjectTypeKeyProperty{
//   					FieldNames: []*string{
//   						jsii.String("fieldNames"),
//   					},
//   					StandardIdentifiers: []*string{
//   						jsii.String("standardIdentifiers"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	MaxProfileObjectCount: jsii.Number(123),
//   	ObjectTypeName: jsii.String("objectTypeName"),
//   	SourceLastUpdatedTimestampFormat: jsii.String("sourceLastUpdatedTimestampFormat"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateId: jsii.String("templateId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-objecttype.html
//
type CfnObjectTypePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnObjectTypeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnObjectTypePropsMixin
type jsiiProxy_CfnObjectTypePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnObjectTypePropsMixin) Props() *CfnObjectTypeMixinProps {
	var returns *CfnObjectTypeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectTypePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CustomerProfiles::ObjectType`.
func NewCfnObjectTypePropsMixin(props *CfnObjectTypeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnObjectTypePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnObjectTypePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnObjectTypePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnObjectTypePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CustomerProfiles::ObjectType`.
func NewCfnObjectTypePropsMixin_Override(c CfnObjectTypePropsMixin, props *CfnObjectTypeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnObjectTypePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnObjectTypePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnObjectTypePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnObjectTypePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnObjectTypePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnObjectTypePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnObjectTypePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnObjectTypePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

