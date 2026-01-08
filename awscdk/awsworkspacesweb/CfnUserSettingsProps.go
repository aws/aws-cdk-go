package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnUserSettings`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserSettingsProps := &CfnUserSettingsProps{
//   	CopyAllowed: jsii.String("copyAllowed"),
//   	DownloadAllowed: jsii.String("downloadAllowed"),
//   	PasteAllowed: jsii.String("pasteAllowed"),
//   	PrintAllowed: jsii.String("printAllowed"),
//   	UploadAllowed: jsii.String("uploadAllowed"),
//
//   	// the properties below are optional
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
//   				"welcomeText": jsii.String("welcomeText"),
//
//   				// the properties below are optional
//   				"contactButtonText": jsii.String("contactButtonText"),
//   				"contactLink": jsii.String("contactLink"),
//   				"loadingText": jsii.String("loadingText"),
//   				"loginButtonText": jsii.String("loginButtonText"),
//   				"loginDescription": jsii.String("loginDescription"),
//   				"loginTitle": jsii.String("loginTitle"),
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
//
//   				// the properties below are optional
//   				Name: jsii.String("name"),
//   				Path: jsii.String("path"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Blocklist: []interface{}{
//   			&CookieSpecificationProperty{
//   				Domain: jsii.String("domain"),
//
//   				// the properties below are optional
//   				Name: jsii.String("name"),
//   				Path: jsii.String("path"),
//   			},
//   		},
//   	},
//   	CustomerManagedKey: jsii.String("customerManagedKey"),
//   	DeepLinkAllowed: jsii.String("deepLinkAllowed"),
//   	DisconnectTimeoutInMinutes: jsii.Number(123),
//   	IdleDisconnectTimeoutInMinutes: jsii.Number(123),
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
//   	WebAuthnAllowed: jsii.String("webAuthnAllowed"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html
//
type CfnUserSettingsProps struct {
	// Specifies whether the user can copy text from the streaming session to the local device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-copyallowed
	//
	CopyAllowed *string `field:"required" json:"copyAllowed" yaml:"copyAllowed"`
	// Specifies whether the user can download files from the streaming session to the local device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-downloadallowed
	//
	DownloadAllowed *string `field:"required" json:"downloadAllowed" yaml:"downloadAllowed"`
	// Specifies whether the user can paste text from the local device to the streaming session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-pasteallowed
	//
	PasteAllowed *string `field:"required" json:"pasteAllowed" yaml:"pasteAllowed"`
	// Specifies whether the user can print to the local device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-printallowed
	//
	PrintAllowed *string `field:"required" json:"printAllowed" yaml:"printAllowed"`
	// Specifies whether the user can upload files from the local device to the streaming session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-uploadallowed
	//
	UploadAllowed *string `field:"required" json:"uploadAllowed" yaml:"uploadAllowed"`
	// The additional encryption context of the user settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-additionalencryptioncontext
	//
	AdditionalEncryptionContext interface{} `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// The branding configuration that customizes the appearance of the web portal for end users.
	//
	// This includes a custom logo, favicon, wallpaper, localized strings, color theme, and an optional terms of service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-brandingconfiguration
	//
	BrandingConfiguration interface{} `field:"optional" json:"brandingConfiguration" yaml:"brandingConfiguration"`
	// The configuration that specifies which cookies should be synchronized from the end user's local browser to the remote browser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-cookiesynchronizationconfiguration
	//
	CookieSynchronizationConfiguration interface{} `field:"optional" json:"cookieSynchronizationConfiguration" yaml:"cookieSynchronizationConfiguration"`
	// The customer managed key used to encrypt sensitive information in the user settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-customermanagedkey
	//
	CustomerManagedKey *string `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// Specifies whether the user can use deep links that open automatically when connecting to a session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-deeplinkallowed
	//
	DeepLinkAllowed *string `field:"optional" json:"deepLinkAllowed" yaml:"deepLinkAllowed"`
	// The amount of time that a streaming session remains active after users disconnect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-disconnecttimeoutinminutes
	//
	DisconnectTimeoutInMinutes *float64 `field:"optional" json:"disconnectTimeoutInMinutes" yaml:"disconnectTimeoutInMinutes"`
	// The amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the disconnect timeout interval begins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-idledisconnecttimeoutinminutes
	//
	IdleDisconnectTimeoutInMinutes *float64 `field:"optional" json:"idleDisconnectTimeoutInMinutes" yaml:"idleDisconnectTimeoutInMinutes"`
	// The tags to add to the user settings resource.
	//
	// A tag is a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration of the toolbar.
	//
	// This allows administrators to select the toolbar type and visual mode, set maximum display resolution for sessions, and choose which items are visible to end users during their sessions. If administrators do not modify these settings, end users retain control over their toolbar preferences.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-toolbarconfiguration
	//
	ToolbarConfiguration interface{} `field:"optional" json:"toolbarConfiguration" yaml:"toolbarConfiguration"`
	// Specifies whether the user can use WebAuthn redirection for passwordless login to websites within the streaming session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html#cfn-workspacesweb-usersettings-webauthnallowed
	//
	WebAuthnAllowed *string `field:"optional" json:"webAuthnAllowed" yaml:"webAuthnAllowed"`
}

