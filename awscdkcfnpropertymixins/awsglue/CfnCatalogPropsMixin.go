package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a catalog in the Glue Data Catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnCatalogPropsMixin := awscdkcfnpropertymixins.Aws_glue.NewCfnCatalogPropsMixin(&CfnCatalogMixinProps{
//   	AllowFullTableExternalDataAccess: jsii.String("allowFullTableExternalDataAccess"),
//   	CatalogProperties: &CatalogPropertiesProperty{
//   		CustomProperties: map[string]*string{
//   			"customPropertiesKey": jsii.String("customProperties"),
//   		},
//   		DataLakeAccessProperties: &DataLakeAccessPropertiesProperty{
//   			AllowFullTableExternalDataAccess: jsii.String("allowFullTableExternalDataAccess"),
//   			CatalogType: jsii.String("catalogType"),
//   			DataLakeAccess: jsii.Boolean(false),
//   			DataTransferRole: jsii.String("dataTransferRole"),
//   			KmsKey: jsii.String("kmsKey"),
//   			ManagedWorkgroupName: jsii.String("managedWorkgroupName"),
//   			ManagedWorkgroupStatus: jsii.String("managedWorkgroupStatus"),
//   			RedshiftDatabaseName: jsii.String("redshiftDatabaseName"),
//   		},
//   	},
//   	CreateDatabaseDefaultPermissions: []interface{}{
//   		&PrincipalPermissionsProperty{
//   			Permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   			Principal: &DataLakePrincipalProperty{
//   				DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   			},
//   		},
//   	},
//   	CreateTableDefaultPermissions: []interface{}{
//   		&PrincipalPermissionsProperty{
//   			Permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   			Principal: &DataLakePrincipalProperty{
//   				DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	FederatedCatalog: &FederatedCatalogProperty{
//   		ConnectionName: jsii.String("connectionName"),
//   		Identifier: jsii.String("identifier"),
//   	},
//   	Name: jsii.String("name"),
//   	OverwriteChildResourcePermissionsWithDefault: jsii.String("overwriteChildResourcePermissionsWithDefault"),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetRedshiftCatalog: &TargetRedshiftCatalogProperty{
//   		CatalogArn: jsii.String("catalogArn"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html
//
type CfnCatalogPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCatalogMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCatalogPropsMixin
type jsiiProxy_CfnCatalogPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCatalogPropsMixin) Props() *CfnCatalogMixinProps {
	var returns *CfnCatalogMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalogPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::Catalog`.
func NewCfnCatalogPropsMixin(props *CfnCatalogMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCatalogPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCatalogPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCatalogPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::Catalog`.
func NewCfnCatalogPropsMixin_Override(c CfnCatalogPropsMixin, props *CfnCatalogMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCatalogPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCatalogPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCatalogPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnCatalogPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCatalogPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCatalogPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

