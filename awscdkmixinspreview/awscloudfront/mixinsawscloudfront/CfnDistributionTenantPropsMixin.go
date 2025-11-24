package mixinsawscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscloudfront/mixinsawscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The distribution tenant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDistributionTenantPropsMixin := awscdkmixinspreview.Mixins.NewCfnDistributionTenantPropsMixin(&CfnDistributionTenantMixinProps{
//   	ConnectionGroupId: jsii.String("connectionGroupId"),
//   	Customizations: &CustomizationsProperty{
//   		Certificate: &CertificateProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		GeoRestrictions: &GeoRestrictionCustomizationProperty{
//   			Locations: []*string{
//   				jsii.String("locations"),
//   			},
//   			RestrictionType: jsii.String("restrictionType"),
//   		},
//   		WebAcl: &WebAclCustomizationProperty{
//   			Action: jsii.String("action"),
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   	DistributionId: jsii.String("distributionId"),
//   	Domains: []*string{
//   		jsii.String("domains"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   	ManagedCertificateRequest: &ManagedCertificateRequestProperty{
//   		CertificateTransparencyLoggingPreference: jsii.String("certificateTransparencyLoggingPreference"),
//   		PrimaryDomainName: jsii.String("primaryDomainName"),
//   		ValidationTokenHost: jsii.String("validationTokenHost"),
//   	},
//   	Name: jsii.String("name"),
//   	Parameters: []interface{}{
//   		&ParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html
//
type CfnDistributionTenantPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDistributionTenantMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDistributionTenantPropsMixin
type jsiiProxy_CfnDistributionTenantPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDistributionTenantPropsMixin) Props() *CfnDistributionTenantMixinProps {
	var returns *CfnDistributionTenantMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionTenantPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFront::DistributionTenant`.
func NewCfnDistributionTenantPropsMixin(props *CfnDistributionTenantMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDistributionTenantPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDistributionTenantPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDistributionTenantPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFront::DistributionTenant`.
func NewCfnDistributionTenantPropsMixin_Override(c CfnDistributionTenantPropsMixin, props *CfnDistributionTenantMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDistributionTenantPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDistributionTenantPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDistributionTenantPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDistributionTenantPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDistributionTenantPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

