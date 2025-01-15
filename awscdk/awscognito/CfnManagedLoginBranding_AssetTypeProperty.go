package awscognito


// An image file from a managed login branding style in a user pool.
//
// This data type is a request parameter of `API_CreateManagedLoginBranding` and `API_UpdateManagedLoginBranding` , and a response parameter of `API_DescribeManagedLoginBranding` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetTypeProperty := &AssetTypeProperty{
//   	Category: jsii.String("category"),
//   	ColorMode: jsii.String("colorMode"),
//   	Extension: jsii.String("extension"),
//
//   	// the properties below are optional
//   	Bytes: jsii.String("bytes"),
//   	ResourceId: jsii.String("resourceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html
//
type CfnManagedLoginBranding_AssetTypeProperty struct {
	// The category that the image corresponds to in your managed login configuration.
	//
	// Managed login has asset categories for different types of logos, backgrounds, and icons.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html#cfn-cognito-managedloginbranding-assettype-category
	//
	Category *string `field:"required" json:"category" yaml:"category"`
	// The display-mode target of the asset: light, dark, or browser-adaptive.
	//
	// For example, Amazon Cognito displays a dark-mode image only when the browser or application is in dark mode, but displays a browser-adaptive file in all contexts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html#cfn-cognito-managedloginbranding-assettype-colormode
	//
	ColorMode *string `field:"required" json:"colorMode" yaml:"colorMode"`
	// The file type of the image file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html#cfn-cognito-managedloginbranding-assettype-extension
	//
	Extension *string `field:"required" json:"extension" yaml:"extension"`
	// The image file, in Base64-encoded binary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html#cfn-cognito-managedloginbranding-assettype-bytes
	//
	Bytes *string `field:"optional" json:"bytes" yaml:"bytes"`
	// The ID of the asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-managedloginbranding-assettype.html#cfn-cognito-managedloginbranding-assettype-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

