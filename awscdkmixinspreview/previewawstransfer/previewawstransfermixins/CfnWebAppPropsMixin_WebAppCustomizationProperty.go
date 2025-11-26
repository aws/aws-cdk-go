package previewawstransfermixins


// A structure that contains the customization fields for the web app.
//
// You can provide a title, logo, and icon to customize the appearance of your web app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   webAppCustomizationProperty := &WebAppCustomizationProperty{
//   	FaviconFile: jsii.String("faviconFile"),
//   	LogoFile: jsii.String("logoFile"),
//   	Title: jsii.String("title"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-webappcustomization.html
//
type CfnWebAppPropsMixin_WebAppCustomizationProperty struct {
	// Returns an icon file data string (in base64 encoding).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-webappcustomization.html#cfn-transfer-webapp-webappcustomization-faviconfile
	//
	FaviconFile *string `field:"optional" json:"faviconFile" yaml:"faviconFile"`
	// Returns a logo file data string (in base64 encoding).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-webappcustomization.html#cfn-transfer-webapp-webappcustomization-logofile
	//
	LogoFile *string `field:"optional" json:"logoFile" yaml:"logoFile"`
	// Returns the page title that you defined for your web app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-webappcustomization.html#cfn-transfer-webapp-webappcustomization-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
}

