package mixinsawscognito


// An image file from a managed login branding style in a user pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   assetTypeProperty := &AssetTypeProperty{
//   	Bytes: jsii.String("bytes"),
//   	Category: jsii.String("category"),
//   	ColorMode: jsii.String("colorMode"),
//   	Extension: jsii.String("extension"),
//   	ResourceId: jsii.String("resourceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html
//
type CfnManagedLoginBrandingPropsMixin_AssetTypeProperty struct {
	// The image file, in Base64-encoded binary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html#cfn-cognito-managedloginbranding-assettype-bytes
	//
	Bytes *string `field:"optional" json:"bytes" yaml:"bytes"`
	// The category that the image corresponds to in your managed login configuration.
	//
	// Managed login has asset categories for different types of logos, backgrounds, and icons.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html#cfn-cognito-managedloginbranding-assettype-category
	//
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The display-mode target of the asset: light, dark, or browser-adaptive.
	//
	// For example, Amazon Cognito displays a dark-mode image only when the browser or application is in dark mode, but displays a browser-adaptive file in all contexts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html#cfn-cognito-managedloginbranding-assettype-colormode
	//
	ColorMode *string `field:"optional" json:"colorMode" yaml:"colorMode"`
	// The file type of the image file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html#cfn-cognito-managedloginbranding-assettype-extension
	//
	Extension *string `field:"optional" json:"extension" yaml:"extension"`
	// The ID of the asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html#cfn-cognito-managedloginbranding-assettype-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

