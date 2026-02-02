package previewawslakeformationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslakeformation/previewawslakeformationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::LakeFormation::PrincipalPermissions` resource represents the permissions that a principal has on a Data Catalog resource (such as AWS Glue databases or AWS Glue tables).
//
// When you create a `PrincipalPermissions` resource, the permissions are granted via the AWS Lake Formation `GrantPermissions` API operation. When you delete a `PrincipalPermissions` resource, the permissions on principal-resource pair are revoked via the AWS Lake Formation `RevokePermissions` API operation.
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
//   cfnPrincipalPermissionsPropsMixin := awscdkmixinspreview.Mixins.NewCfnPrincipalPermissionsPropsMixin(&CfnPrincipalPermissionsMixinProps{
//   	Catalog: jsii.String("catalog"),
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	PermissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   	Principal: &DataLakePrincipalProperty{
//   		DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	Resource: &ResourceProperty{
//   		Catalog: catalog,
//   		Database: &DatabaseResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			Name: jsii.String("name"),
//   		},
//   		DataCellsFilter: &DataCellsFilterResourceProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//   			TableCatalogId: jsii.String("tableCatalogId"),
//   			TableName: jsii.String("tableName"),
//   		},
//   		DataLocation: &DataLocationResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   		LfTag: &LFTagKeyResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			TagKey: jsii.String("tagKey"),
//   			TagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   		LfTagPolicy: &LFTagPolicyResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			Expression: []interface{}{
//   				&LFTagProperty{
//   					TagKey: jsii.String("tagKey"),
//   					TagValues: []*string{
//   						jsii.String("tagValues"),
//   					},
//   				},
//   			},
//   			ResourceType: jsii.String("resourceType"),
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
//   			ColumnWildcard: &ColumnWildcardProperty{
//   				ExcludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-principalpermissions.html
//
type CfnPrincipalPermissionsPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPrincipalPermissionsMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPrincipalPermissionsPropsMixin
type jsiiProxy_CfnPrincipalPermissionsPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPrincipalPermissionsPropsMixin) Props() *CfnPrincipalPermissionsMixinProps {
	var returns *CfnPrincipalPermissionsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissionsPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LakeFormation::PrincipalPermissions`.
func NewCfnPrincipalPermissionsPropsMixin(props *CfnPrincipalPermissionsMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPrincipalPermissionsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPrincipalPermissionsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPrincipalPermissionsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LakeFormation::PrincipalPermissions`.
func NewCfnPrincipalPermissionsPropsMixin_Override(c CfnPrincipalPermissionsPropsMixin, props *CfnPrincipalPermissionsMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPrincipalPermissionsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPrincipalPermissionsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPrincipalPermissionsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lakeformation.mixins.CfnPrincipalPermissionsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissionsPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissionsPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

