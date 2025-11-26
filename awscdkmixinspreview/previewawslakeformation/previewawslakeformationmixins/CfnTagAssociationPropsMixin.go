package previewawslakeformationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslakeformation/previewawslakeformationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::LakeFormation::TagAssociation` resource represents an assignment of an LF-tag to a Data Catalog resource (database, table, or column).
//
// During a stack operation, CloudFormation calls AWS Lake Formation `AddLFTagsToResource` API to create a `TagAssociation` resource and calls the `RemoveLFTagsToResource` API to delete it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   cfnTagAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnTagAssociationPropsMixin(&CfnTagAssociationMixinProps{
//   	LfTags: []interface{}{
//   		&LFTagPairProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			TagKey: jsii.String("tagKey"),
//   			TagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   	},
//   	Resource: &ResourceProperty{
//   		Catalog: catalog,
//   		Database: &DatabaseResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			Name: jsii.String("name"),
//   		},
//   		Table: &TableResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//   			TableWildcard: tableWildcard,
//   		},
//   		TableWithColumns: &TableWithColumnsResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tagassociation.html
//
type CfnTagAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTagAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTagAssociationPropsMixin
type jsiiProxy_CfnTagAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTagAssociationPropsMixin) Props() *CfnTagAssociationMixinProps {
	var returns *CfnTagAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LakeFormation::TagAssociation`.
func NewCfnTagAssociationPropsMixin(props *CfnTagAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTagAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTagAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTagAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LakeFormation::TagAssociation`.
func NewCfnTagAssociationPropsMixin_Override(c CfnTagAssociationPropsMixin, props *CfnTagAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTagAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTagAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTagAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnTagAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTagAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTagAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

