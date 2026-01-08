package previewawsworkspaceswebmixins


// The branding configuration that customizes the appearance of the web portal for end users.
//
// This includes a custom logo, favicon, wallpaper, localized strings, color theme, and an optional terms of service.
//
// > The `LogoMetadata` , `FaviconMetadata` , and `WallpaperMetadata` properties are read-only and cannot be specified in your template. They are automatically populated by the service after you upload images and can be retrieved using the `Fn::GetAtt` intrinsic function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   brandingConfigurationProperty := &BrandingConfigurationProperty{
//   	ColorTheme: jsii.String("colorTheme"),
//   	Favicon: jsii.String("favicon"),
//   	FaviconMetadata: &ImageMetadataProperty{
//   		FileExtension: jsii.String("fileExtension"),
//   		LastUploadTimestamp: jsii.String("lastUploadTimestamp"),
//   		MimeType: jsii.String("mimeType"),
//   	},
//   	LocalizedStrings: map[string]interface{}{
//   		"localizedStringsKey": &LocalizedBrandingStringsProperty{
//   			"browserTabTitle": jsii.String("browserTabTitle"),
//   			"contactButtonText": jsii.String("contactButtonText"),
//   			"contactLink": jsii.String("contactLink"),
//   			"loadingText": jsii.String("loadingText"),
//   			"loginButtonText": jsii.String("loginButtonText"),
//   			"loginDescription": jsii.String("loginDescription"),
//   			"loginTitle": jsii.String("loginTitle"),
//   			"welcomeText": jsii.String("welcomeText"),
//   		},
//   	},
//   	Logo: jsii.String("logo"),
//   	LogoMetadata: &ImageMetadataProperty{
//   		FileExtension: jsii.String("fileExtension"),
//   		LastUploadTimestamp: jsii.String("lastUploadTimestamp"),
//   		MimeType: jsii.String("mimeType"),
//   	},
//   	TermsOfService: jsii.String("termsOfService"),
//   	Wallpaper: jsii.String("wallpaper"),
//   	WallpaperMetadata: &ImageMetadataProperty{
//   		FileExtension: jsii.String("fileExtension"),
//   		LastUploadTimestamp: jsii.String("lastUploadTimestamp"),
//   		MimeType: jsii.String("mimeType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html
//
type CfnUserSettingsPropsMixin_BrandingConfigurationProperty struct {
	// The color theme for components on the web portal.
	//
	// Choose `Light` if you upload a dark wallpaper, or `Dark` for a light wallpaper.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-colortheme
	//
	ColorTheme *string `field:"optional" json:"colorTheme" yaml:"colorTheme"`
	// The favicon image for the portal.
	//
	// Provide either a binary image file or an S3 URI pointing to the image file. Maximum 100 KB in JPEG, PNG, or ICO format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-favicon
	//
	Favicon *string `field:"optional" json:"favicon" yaml:"favicon"`
	// Read-only.
	//
	// Metadata for the favicon image file, including the MIME type, file extension, and upload timestamp. This property is automatically populated by the service and cannot be specified in your template. It can be retrieved using the `Fn::GetAtt` intrinsic function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-faviconmetadata
	//
	FaviconMetadata interface{} `field:"optional" json:"faviconMetadata" yaml:"faviconMetadata"`
	// A map of localized text strings for different languages, allowing the portal to display content in the user's preferred language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-localizedstrings
	//
	LocalizedStrings interface{} `field:"optional" json:"localizedStrings" yaml:"localizedStrings"`
	// The logo image for the portal.
	//
	// Provide either a binary image file or an S3 URI pointing to the image file. Maximum 100 KB in JPEG, PNG, or ICO format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-logo
	//
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Read-only.
	//
	// Metadata for the logo image file, including the MIME type, file extension, and upload timestamp. This property is automatically populated by the service and cannot be specified in your template. It can be retrieved using the `Fn::GetAtt` intrinsic function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-logometadata
	//
	LogoMetadata interface{} `field:"optional" json:"logoMetadata" yaml:"logoMetadata"`
	// The terms of service text in Markdown format that users must accept before accessing the portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-termsofservice
	//
	TermsOfService *string `field:"optional" json:"termsOfService" yaml:"termsOfService"`
	// The wallpaper image for the portal.
	//
	// Provide either a binary image file or an S3 URI pointing to the image file. Maximum 5 MB in JPEG or PNG format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-wallpaper
	//
	Wallpaper *string `field:"optional" json:"wallpaper" yaml:"wallpaper"`
	// Read-only.
	//
	// Metadata for the wallpaper image file, including the MIME type, file extension, and upload timestamp. This property is automatically populated by the service and cannot be specified in your template. It can be retrieved using the `Fn::GetAtt` intrinsic function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-wallpapermetadata
	//
	WallpaperMetadata interface{} `field:"optional" json:"wallpaperMetadata" yaml:"wallpaperMetadata"`
}

