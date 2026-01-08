package previewawsworkspaceswebmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsworkspacesweb/previewawsworkspaceswebmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// This resource specifies user settings that can be associated with a web portal.
//
// Once associated with a web portal, user settings control how users can transfer data between a streaming session and the their local devices.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserSettingsPropsMixin := awscdkmixinspreview.Mixins.NewCfnUserSettingsPropsMixin(&CfnUserSettingsMixinProps{
//   	AdditionalEncryptionContext: map[string]*string{
//   		"additionalEncryptionContextKey": jsii.String("additionalEncryptionContext"),
//   	},
//   	BrandingConfiguration: &BrandingConfigurationProperty{
//   		ColorTheme: jsii.String("colorTheme"),
//   		Favicon: jsii.String("favicon"),
//   		FaviconMetadata: &ImageMetadataProperty{
//   			FileExtension: jsii.String("fileExtension"),
//   			LastUploadTimestamp: jsii.String("lastUploadTimestamp"),
//   			MimeType: jsii.String("mimeType"),
//   		},
//   		LocalizedStrings: map[string]interface{}{
//   			"localizedStringsKey": &LocalizedBrandingStringsProperty{
//   				"browserTabTitle": jsii.String("browserTabTitle"),
//   				"contactButtonText": jsii.String("contactButtonText"),
//   				"contactLink": jsii.String("contactLink"),
//   				"loadingText": jsii.String("loadingText"),
//   				"loginButtonText": jsii.String("loginButtonText"),
//   				"loginDescription": jsii.String("loginDescription"),
//   				"loginTitle": jsii.String("loginTitle"),
//   				"welcomeText": jsii.String("welcomeText"),
//   			},
//   		},
//   		Logo: jsii.String("logo"),
//   		LogoMetadata: &ImageMetadataProperty{
//   			FileExtension: jsii.String("fileExtension"),
//   			LastUploadTimestamp: jsii.String("lastUploadTimestamp"),
//   			MimeType: jsii.String("mimeType"),
//   		},
//   		TermsOfService: jsii.String("termsOfService"),
//   		Wallpaper: jsii.String("wallpaper"),
//   		WallpaperMetadata: &ImageMetadataProperty{
//   			FileExtension: jsii.String("fileExtension"),
//   			LastUploadTimestamp: jsii.String("lastUploadTimestamp"),
//   			MimeType: jsii.String("mimeType"),
//   		},
//   	},
//   	CookieSynchronizationConfiguration: &CookieSynchronizationConfigurationProperty{
//   		Allowlist: []interface{}{
//   			&CookieSpecificationProperty{
//   				Domain: jsii.String("domain"),
//   				Name: jsii.String("name"),
//   				Path: jsii.String("path"),
//   			},
//   		},
//   		Blocklist: []interface{}{
//   			&CookieSpecificationProperty{
//   				Domain: jsii.String("domain"),
//   				Name: jsii.String("name"),
//   				Path: jsii.String("path"),
//   			},
//   		},
//   	},
//   	CopyAllowed: jsii.String("copyAllowed"),
//   	CustomerManagedKey: jsii.String("customerManagedKey"),
//   	DeepLinkAllowed: jsii.String("deepLinkAllowed"),
//   	DisconnectTimeoutInMinutes: jsii.Number(123),
//   	DownloadAllowed: jsii.String("downloadAllowed"),
//   	IdleDisconnectTimeoutInMinutes: jsii.Number(123),
//   	PasteAllowed: jsii.String("pasteAllowed"),
//   	PrintAllowed: jsii.String("printAllowed"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ToolbarConfiguration: &ToolbarConfigurationProperty{
//   		HiddenToolbarItems: []*string{
//   			jsii.String("hiddenToolbarItems"),
//   		},
//   		MaxDisplayResolution: jsii.String("maxDisplayResolution"),
//   		ToolbarType: jsii.String("toolbarType"),
//   		VisualMode: jsii.String("visualMode"),
//   	},
//   	UploadAllowed: jsii.String("uploadAllowed"),
//   	WebAuthnAllowed: jsii.String("webAuthnAllowed"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html
//
type CfnUserSettingsPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnUserSettingsMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserSettingsPropsMixin
type jsiiProxy_CfnUserSettingsPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnUserSettingsPropsMixin) Props() *CfnUserSettingsMixinProps {
	var returns *CfnUserSettingsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettingsPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WorkSpacesWeb::UserSettings`.
func NewCfnUserSettingsPropsMixin(props *CfnUserSettingsMixinProps, options *mixins.CfnPropertyMixinOptions) CfnUserSettingsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserSettingsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserSettingsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnUserSettingsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WorkSpacesWeb::UserSettings`.
func NewCfnUserSettingsPropsMixin_Override(c CfnUserSettingsPropsMixin, props *CfnUserSettingsMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnUserSettingsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnUserSettingsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserSettingsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnUserSettingsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserSettingsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_workspacesweb.mixins.CfnUserSettingsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserSettingsPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnUserSettingsPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

