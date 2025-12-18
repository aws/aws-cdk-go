package awsworkspacesweb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   			"welcomeText": jsii.String("welcomeText"),
//
//   			// the properties below are optional
//   			"contactButtonText": jsii.String("contactButtonText"),
//   			"contactLink": jsii.String("contactLink"),
//   			"loadingText": jsii.String("loadingText"),
//   			"loginButtonText": jsii.String("loginButtonText"),
//   			"loginDescription": jsii.String("loginDescription"),
//   			"loginTitle": jsii.String("loginTitle"),
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
type CfnUserSettings_BrandingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-colortheme
	//
	ColorTheme *string `field:"optional" json:"colorTheme" yaml:"colorTheme"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-favicon
	//
	Favicon *string `field:"optional" json:"favicon" yaml:"favicon"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-faviconmetadata
	//
	FaviconMetadata interface{} `field:"optional" json:"faviconMetadata" yaml:"faviconMetadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-localizedstrings
	//
	LocalizedStrings interface{} `field:"optional" json:"localizedStrings" yaml:"localizedStrings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-logo
	//
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-logometadata
	//
	LogoMetadata interface{} `field:"optional" json:"logoMetadata" yaml:"logoMetadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-termsofservice
	//
	TermsOfService *string `field:"optional" json:"termsOfService" yaml:"termsOfService"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-wallpaper
	//
	Wallpaper *string `field:"optional" json:"wallpaper" yaml:"wallpaper"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-brandingconfiguration.html#cfn-workspacesweb-usersettings-brandingconfiguration-wallpapermetadata
	//
	WallpaperMetadata interface{} `field:"optional" json:"wallpaperMetadata" yaml:"wallpaperMetadata"`
}

