package awslakeformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::LakeFormation::DataLakeSettings` resource is an AWS Lake Formation resource type that manages the data lake settings for your account.
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
//   cfnDataLakeSettingsPropsMixin := awscdkcfnpropertymixins.Aws_lakeformation.NewCfnDataLakeSettingsPropsMixin(&CfnDataLakeSettingsMixinProps{
//   	Admins: []interface{}{
//   		&DataLakePrincipalProperty{
//   			DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   		},
//   	},
//   	AllowExternalDataFiltering: jsii.Boolean(false),
//   	AllowFullTableExternalDataAccess: jsii.Boolean(false),
//   	AuthorizedSessionTagValueList: []*string{
//   		jsii.String("authorizedSessionTagValueList"),
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
//   	ExternalDataFilteringAllowList: []interface{}{
//   		&DataLakePrincipalProperty{
//   			DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   		},
//   	},
//   	MutationType: jsii.String("mutationType"),
//   	Parameters: parameters,
//   	ReadOnlyAdmins: []interface{}{
//   		&DataLakePrincipalProperty{
//   			DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   		},
//   	},
//   	TrustedResourceOwners: []*string{
//   		jsii.String("trustedResourceOwners"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html
//
type CfnDataLakeSettingsPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDataLakeSettingsMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataLakeSettingsPropsMixin
type jsiiProxy_CfnDataLakeSettingsPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDataLakeSettingsPropsMixin) Props() *CfnDataLakeSettingsMixinProps {
	var returns *CfnDataLakeSettingsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettingsPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LakeFormation::DataLakeSettings`.
func NewCfnDataLakeSettingsPropsMixin(props *CfnDataLakeSettingsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDataLakeSettingsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataLakeSettingsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataLakeSettingsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataLakeSettingsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LakeFormation::DataLakeSettings`.
func NewCfnDataLakeSettingsPropsMixin_Override(c CfnDataLakeSettingsPropsMixin, props *CfnDataLakeSettingsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataLakeSettingsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDataLakeSettingsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataLakeSettingsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataLakeSettingsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataLakeSettingsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_lakeformation.CfnDataLakeSettingsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataLakeSettingsPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDataLakeSettingsPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

