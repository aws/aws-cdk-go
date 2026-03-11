package awsworkspacesweb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// This resource specifies browser settings that can be associated with a web portal.
//
// Once associated with a web portal, browser settings control how the browser will behave once a user starts a streaming session for the web portal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnBrowserSettingsPropsMixin := awscdkcfnpropertymixins.Aws_workspacesweb.NewCfnBrowserSettingsPropsMixin(&CfnBrowserSettingsMixinProps{
//   	AdditionalEncryptionContext: map[string]*string{
//   		"additionalEncryptionContextKey": jsii.String("additionalEncryptionContext"),
//   	},
//   	BrowserPolicy: jsii.String("browserPolicy"),
//   	CustomerManagedKey: jsii.String("customerManagedKey"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WebContentFilteringPolicy: &WebContentFilteringPolicyProperty{
//   		AllowedUrls: []*string{
//   			jsii.String("allowedUrls"),
//   		},
//   		BlockedCategories: []*string{
//   			jsii.String("blockedCategories"),
//   		},
//   		BlockedUrls: []*string{
//   			jsii.String("blockedUrls"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-browsersettings.html
//
type CfnBrowserSettingsPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnBrowserSettingsMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBrowserSettingsPropsMixin
type jsiiProxy_CfnBrowserSettingsPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnBrowserSettingsPropsMixin) Props() *CfnBrowserSettingsMixinProps {
	var returns *CfnBrowserSettingsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBrowserSettingsPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WorkSpacesWeb::BrowserSettings`.
func NewCfnBrowserSettingsPropsMixin(props *CfnBrowserSettingsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnBrowserSettingsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBrowserSettingsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBrowserSettingsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_workspacesweb.CfnBrowserSettingsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WorkSpacesWeb::BrowserSettings`.
func NewCfnBrowserSettingsPropsMixin_Override(c CfnBrowserSettingsPropsMixin, props *CfnBrowserSettingsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_workspacesweb.CfnBrowserSettingsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnBrowserSettingsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBrowserSettingsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_workspacesweb.CfnBrowserSettingsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBrowserSettingsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_workspacesweb.CfnBrowserSettingsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBrowserSettingsPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnBrowserSettingsPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

