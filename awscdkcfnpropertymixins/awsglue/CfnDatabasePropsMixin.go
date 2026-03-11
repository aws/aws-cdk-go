package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Glue::Database` resource specifies a logical grouping of tables in AWS Glue .
//
// For more information, see [Defining a Database in Your Data Catalog](https://docs.aws.amazon.com/glue/latest/dg/define-database.html) and [Database Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-databases.html#aws-glue-api-catalog-databases-Database) in the *AWS Glue Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var parameters interface{}
//
//   cfnDatabasePropsMixin := awscdkcfnpropertymixins.Aws_glue.NewCfnDatabasePropsMixin(&CfnDatabaseMixinProps{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseInput: &DatabaseInputProperty{
//   		CreateTableDefaultPermissions: []interface{}{
//   			&PrincipalPrivilegesProperty{
//   				Permissions: []*string{
//   					jsii.String("permissions"),
//   				},
//   				Principal: &DataLakePrincipalProperty{
//   					DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   				},
//   			},
//   		},
//   		Description: jsii.String("description"),
//   		FederatedDatabase: &FederatedDatabaseProperty{
//   			ConnectionName: jsii.String("connectionName"),
//   			Identifier: jsii.String("identifier"),
//   		},
//   		LocationUri: jsii.String("locationUri"),
//   		Name: jsii.String("name"),
//   		Parameters: parameters,
//   		TargetDatabase: &DatabaseIdentifierProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Region: jsii.String("region"),
//   		},
//   	},
//   	DatabaseName: jsii.String("databaseName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-database.html
//
type CfnDatabasePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDatabaseMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDatabasePropsMixin
type jsiiProxy_CfnDatabasePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDatabasePropsMixin) Props() *CfnDatabaseMixinProps {
	var returns *CfnDatabaseMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabasePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::Database`.
func NewCfnDatabasePropsMixin(props *CfnDatabaseMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDatabasePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDatabasePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDatabasePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDatabasePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::Database`.
func NewCfnDatabasePropsMixin_Override(c CfnDatabasePropsMixin, props *CfnDatabaseMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDatabasePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDatabasePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDatabasePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDatabasePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatabasePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnDatabasePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDatabasePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDatabasePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

