package awscognito


// Properties for defining a `CfnManagedLoginBranding`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var settings interface{}
//
//   cfnManagedLoginBrandingProps := &CfnManagedLoginBrandingProps{
//   	UserPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	Assets: []interface{}{
//   		&AssetTypeProperty{
//   			Category: jsii.String("category"),
//   			ColorMode: jsii.String("colorMode"),
//   			Extension: jsii.String("extension"),
//
//   			// the properties below are optional
//   			Bytes: jsii.String("bytes"),
//   			ResourceId: jsii.String("resourceId"),
//   		},
//   	},
//   	ClientId: jsii.String("clientId"),
//   	ReturnMergedResources: jsii.Boolean(false),
//   	Settings: settings,
//   	UseCognitoProvidedValues: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html
//
type CfnManagedLoginBrandingProps struct {
	// The user pool where the branding style is assigned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// An array of image files that you want to apply to roles like backgrounds, logos, and icons.
	//
	// Each object must also indicate whether it is for dark mode, light mode, or browser-adaptive mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-assets
	//
	Assets interface{} `field:"optional" json:"assets" yaml:"assets"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-returnmergedresources
	//
	ReturnMergedResources interface{} `field:"optional" json:"returnMergedResources" yaml:"returnMergedResources"`
	// A JSON file, encoded as a `Document` type, with the the settings that you want to apply to your style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-settings
	//
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// When true, applies the default branding style options.
	//
	// This option reverts to a "blank" style that you can modify later in the branding designer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html#cfn-cognito-managedloginbranding-usecognitoprovidedvalues
	//
	UseCognitoProvidedValues interface{} `field:"optional" json:"useCognitoProvidedValues" yaml:"useCognitoProvidedValues"`
}

