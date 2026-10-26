package awslicensemanagerusersubscriptions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslicensemanagerusersubscriptions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::LicenseManagerUserSubscriptions::LicenseServerEndpoint.
//
// A license server endpoint that provides access to Remote Desktop Services Subscriber Access License servers for a registered identity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLicenseServerEndpointPropsMixin := awscdkcfnpropertymixins.Aws_licensemanagerusersubscriptions.NewCfnLicenseServerEndpointPropsMixin(&CfnLicenseServerEndpointMixinProps{
//   	IdentityProviderArn: jsii.String("identityProviderArn"),
//   	LicenseServerSettings: &LicenseServerSettingsProperty{
//   		ServerSettings: &ServerSettingsProperty{
//   			RdsSalSettings: &RdsSalSettingsProperty{
//   				RdsSalCredentialsProvider: &CredentialsProviderProperty{
//   					SecretsManagerCredentialsProvider: &SecretsManagerCredentialsProviderProperty{
//   						SecretId: jsii.String("secretId"),
//   					},
//   				},
//   			},
//   		},
//   		ServerType: jsii.String("serverType"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanagerusersubscriptions-licenseserverendpoint.html
//
type CfnLicenseServerEndpointPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLicenseServerEndpointMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLicenseServerEndpointPropsMixin
type jsiiProxy_CfnLicenseServerEndpointPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLicenseServerEndpointPropsMixin) Props() *CfnLicenseServerEndpointMixinProps {
	var returns *CfnLicenseServerEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicenseServerEndpointPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LicenseManagerUserSubscriptions::LicenseServerEndpoint`.
func NewCfnLicenseServerEndpointPropsMixin(props *CfnLicenseServerEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLicenseServerEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLicenseServerEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLicenseServerEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LicenseManagerUserSubscriptions::LicenseServerEndpoint`.
func NewCfnLicenseServerEndpointPropsMixin_Override(c CfnLicenseServerEndpointPropsMixin, props *CfnLicenseServerEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLicenseServerEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLicenseServerEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLicenseServerEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLicenseServerEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLicenseServerEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

