package awscertificatemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscertificatemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::CertificateManager::AcmeDomainValidation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAcmeDomainValidationPropsMixin := awscdkcfnpropertymixins.Aws_certificatemanager.NewCfnAcmeDomainValidationPropsMixin(&CfnAcmeDomainValidationMixinProps{
//   	AcmeEndpointArn: jsii.String("acmeEndpointArn"),
//   	DomainName: jsii.String("domainName"),
//   	PrevalidationOptions: &PrevalidationOptionsProperty{
//   		DnsPrevalidation: &DnsPrevalidationOptionsProperty{
//   			DomainScope: &DomainScopeProperty{
//   				ExactDomain: jsii.String("exactDomain"),
//   				Subdomains: jsii.String("subdomains"),
//   				Wildcards: jsii.String("wildcards"),
//   			},
//   			HostedZoneId: jsii.String("hostedZoneId"),
//   		},
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmedomainvalidation.html
//
type CfnAcmeDomainValidationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAcmeDomainValidationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAcmeDomainValidationPropsMixin
type jsiiProxy_CfnAcmeDomainValidationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAcmeDomainValidationPropsMixin) Props() *CfnAcmeDomainValidationMixinProps {
	var returns *CfnAcmeDomainValidationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcmeDomainValidationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CertificateManager::AcmeDomainValidation`.
func NewCfnAcmeDomainValidationPropsMixin(props *CfnAcmeDomainValidationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAcmeDomainValidationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAcmeDomainValidationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAcmeDomainValidationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeDomainValidationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CertificateManager::AcmeDomainValidation`.
func NewCfnAcmeDomainValidationPropsMixin_Override(c CfnAcmeDomainValidationPropsMixin, props *CfnAcmeDomainValidationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeDomainValidationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAcmeDomainValidationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAcmeDomainValidationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeDomainValidationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAcmeDomainValidationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeDomainValidationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAcmeDomainValidationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAcmeDomainValidationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

