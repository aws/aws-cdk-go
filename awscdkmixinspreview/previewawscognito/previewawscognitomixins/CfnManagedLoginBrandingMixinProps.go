package previewawscognitomixins


// Properties for CfnManagedLoginBrandingPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var settings interface{}
//
//   cfnManagedLoginBrandingMixinProps := &CfnManagedLoginBrandingMixinProps{
//   	Assets: []interface{}{
//   		&AssetTypeProperty{
//   			Bytes: jsii.String("bytes"),
//   			Category: jsii.String("category"),
//   			ColorMode: jsii.String("colorMode"),
//   			Extension: jsii.String("extension"),
//   			ResourceId: jsii.String("resourceId"),
//   		},
//   	},
//   	ClientId: jsii.String("clientId"),
//   	ReturnMergedResources: jsii.Boolean(false),
//   	Settings: settings,
//   	UseCognitoProvidedValues: jsii.Boolean(false),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html
//
type CfnManagedLoginBrandingMixinProps struct {
	// An array of image files that you want to apply to roles like backgrounds, logos, and icons.
	//
	// Each object must also indicate whether it is for dark mode, light mode, or browser-adaptive mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-assets
	//
	Assets interface{} `field:"optional" json:"assets" yaml:"assets"`
	// The app client that you want to assign the branding style to.
	//
	// Each style is linked to an app client until you delete it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// When `true` , returns values for branding options that are unchanged from Amazon Cognito defaults.
	//
	// When `false` or when you omit this parameter, returns only values that you customized in your branding style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-returnmergedresources
	//
	ReturnMergedResources interface{} `field:"optional" json:"returnMergedResources" yaml:"returnMergedResources"`
	// A JSON file, encoded as a `Document` type, with the the settings that you want to apply to your style.
	//
	// The following components are not currently implemented and reserved for future use:
	//
	// - `signUp`
	// - `instructions`
	// - `sessionTimerDisplay`
	// - `languageSelector` (for localization, see [Managed login localization)](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html#managed-login-localization)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-settings
	//
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// When true, applies the default branding style options.
	//
	// This option reverts to default style options that are managed by Amazon Cognito. You can modify them later in the branding editor.
	//
	// When you specify `true` for this option, you must also omit values for `Settings` and `Assets` in the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-usecognitoprovidedvalues
	//
	UseCognitoProvidedValues interface{} `field:"optional" json:"useCognitoProvidedValues" yaml:"useCognitoProvidedValues"`
	// The user pool where the branding style is assigned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-userpoolid
	//
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

