package previewawskendramixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawskendra/previewawskendramixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon Kendra index.
//
// Once the index is active you can add documents to your index using the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation or using one of the supported data sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIndexPropsMixin := awscdkmixinspreview.Mixins.NewCfnIndexPropsMixin(&CfnIndexMixinProps{
//   	CapacityUnits: &CapacityUnitsConfigurationProperty{
//   		QueryCapacityUnits: jsii.Number(123),
//   		StorageCapacityUnits: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	DocumentMetadataConfigurations: []interface{}{
//   		&DocumentMetadataConfigurationProperty{
//   			Name: jsii.String("name"),
//   			Relevance: &RelevanceProperty{
//   				Duration: jsii.String("duration"),
//   				Freshness: jsii.Boolean(false),
//   				Importance: jsii.Number(123),
//   				RankOrder: jsii.String("rankOrder"),
//   				ValueImportanceItems: []interface{}{
//   					&ValueImportanceItemProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Search: &SearchProperty{
//   				Displayable: jsii.Boolean(false),
//   				Facetable: jsii.Boolean(false),
//   				Searchable: jsii.Boolean(false),
//   				Sortable: jsii.Boolean(false),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Edition: jsii.String("edition"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	ServerSideEncryptionConfiguration: &ServerSideEncryptionConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserContextPolicy: jsii.String("userContextPolicy"),
//   	UserTokenConfigurations: []interface{}{
//   		&UserTokenConfigurationProperty{
//   			JsonTokenTypeConfiguration: &JsonTokenTypeConfigurationProperty{
//   				GroupAttributeField: jsii.String("groupAttributeField"),
//   				UserNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   			JwtTokenTypeConfiguration: &JwtTokenTypeConfigurationProperty{
//   				ClaimRegex: jsii.String("claimRegex"),
//   				GroupAttributeField: jsii.String("groupAttributeField"),
//   				Issuer: jsii.String("issuer"),
//   				KeyLocation: jsii.String("keyLocation"),
//   				SecretManagerArn: jsii.String("secretManagerArn"),
//   				Url: jsii.String("url"),
//   				UserNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html
//
type CfnIndexPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIndexMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIndexPropsMixin
type jsiiProxy_CfnIndexPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIndexPropsMixin) Props() *CfnIndexMixinProps {
	var returns *CfnIndexMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndexPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Kendra::Index`.
func NewCfnIndexPropsMixin(props *CfnIndexMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIndexPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIndexPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIndexPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kendra.mixins.CfnIndexPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Kendra::Index`.
func NewCfnIndexPropsMixin_Override(c CfnIndexPropsMixin, props *CfnIndexMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kendra.mixins.CfnIndexPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIndexPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIndexPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kendra.mixins.CfnIndexPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIndexPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_kendra.mixins.CfnIndexPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIndexPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIndexPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

