package mixinsawsworkspacesweb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsworkspacesweb/mixinsawsworkspacesweb/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The data protection settings resource that can be associated with a web portal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataProtectionSettingsPropsMixin := awscdkmixinspreview.Mixins.NewCfnDataProtectionSettingsPropsMixin(&CfnDataProtectionSettingsMixinProps{
//   	AdditionalEncryptionContext: map[string]*string{
//   		"additionalEncryptionContextKey": jsii.String("additionalEncryptionContext"),
//   	},
//   	CustomerManagedKey: jsii.String("customerManagedKey"),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	InlineRedactionConfiguration: &InlineRedactionConfigurationProperty{
//   		GlobalConfidenceLevel: jsii.Number(123),
//   		GlobalEnforcedUrls: []*string{
//   			jsii.String("globalEnforcedUrls"),
//   		},
//   		GlobalExemptUrls: []*string{
//   			jsii.String("globalExemptUrls"),
//   		},
//   		InlineRedactionPatterns: []interface{}{
//   			&InlineRedactionPatternProperty{
//   				BuiltInPatternId: jsii.String("builtInPatternId"),
//   				ConfidenceLevel: jsii.Number(123),
//   				CustomPattern: &CustomPatternProperty{
//   					KeywordRegex: jsii.String("keywordRegex"),
//   					PatternDescription: jsii.String("patternDescription"),
//   					PatternName: jsii.String("patternName"),
//   					PatternRegex: jsii.String("patternRegex"),
//   				},
//   				EnforcedUrls: []*string{
//   					jsii.String("enforcedUrls"),
//   				},
//   				ExemptUrls: []*string{
//   					jsii.String("exemptUrls"),
//   				},
//   				RedactionPlaceHolder: &RedactionPlaceHolderProperty{
//   					RedactionPlaceHolderText: jsii.String("redactionPlaceHolderText"),
//   					RedactionPlaceHolderType: jsii.String("redactionPlaceHolderType"),
//   				},
//   			},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-dataprotectionsettings.html
//
type CfnDataProtectionSettingsPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataProtectionSettingsMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataProtectionSettingsPropsMixin
type jsiiProxy_CfnDataProtectionSettingsPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataProtectionSettingsPropsMixin) Props() *CfnDataProtectionSettingsMixinProps {
	var returns *CfnDataProtectionSettingsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataProtectionSettingsPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WorkSpacesWeb::DataProtectionSettings`.
func NewCfnDataProtectionSettingsPropsMixin(props *CfnDataProtectionSettingsMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataProtectionSettingsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataProtectionSettingsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataProtectionSettingsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnDataProtectionSettingsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WorkSpacesWeb::DataProtectionSettings`.
func NewCfnDataProtectionSettingsPropsMixin_Override(c CfnDataProtectionSettingsPropsMixin, props *CfnDataProtectionSettingsMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnDataProtectionSettingsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataProtectionSettingsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataProtectionSettingsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnDataProtectionSettingsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataProtectionSettingsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnDataProtectionSettingsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataProtectionSettingsPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDataProtectionSettingsPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

