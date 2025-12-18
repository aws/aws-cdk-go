package previewawssecurityhubmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssecurityhub/previewawssecurityhubmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Grants permission to create a connectorV2 based on input parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectorV2PropsMixin := awscdkmixinspreview.Mixins.NewCfnConnectorV2PropsMixin(&CfnConnectorV2MixinProps{
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Name: jsii.String("name"),
//   	Provider: &ProviderProperty{
//   		JiraCloud: &JiraCloudProperty{
//   			AuthStatus: jsii.String("authStatus"),
//   			AuthUrl: jsii.String("authUrl"),
//   			CloudId: jsii.String("cloudId"),
//   			Domain: jsii.String("domain"),
//   			ProjectKey: jsii.String("projectKey"),
//   		},
//   		ServiceNow: &ServiceNowProperty{
//   			AuthStatus: jsii.String("authStatus"),
//   			InstanceName: jsii.String("instanceName"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-connectorv2.html
//
type CfnConnectorV2PropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConnectorV2MixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConnectorV2PropsMixin
type jsiiProxy_CfnConnectorV2PropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConnectorV2PropsMixin) Props() *CfnConnectorV2MixinProps {
	var returns *CfnConnectorV2MixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorV2PropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityHub::ConnectorV2`.
func NewCfnConnectorV2PropsMixin(props *CfnConnectorV2MixinProps, options *mixins.CfnPropertyMixinOptions) CfnConnectorV2PropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConnectorV2PropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnectorV2PropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnConnectorV2PropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityHub::ConnectorV2`.
func NewCfnConnectorV2PropsMixin_Override(c CfnConnectorV2PropsMixin, props *CfnConnectorV2MixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnConnectorV2PropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConnectorV2PropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectorV2PropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnConnectorV2PropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectorV2PropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnConnectorV2PropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectorV2PropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnConnectorV2PropsMixin) Supports(construct constructs.IConstruct) *bool {
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

