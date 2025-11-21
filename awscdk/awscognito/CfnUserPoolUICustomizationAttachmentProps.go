package awscognito


// Properties for defining a `CfnUserPoolUICustomizationAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolUICustomizationAttachmentProps := &CfnUserPoolUICustomizationAttachmentProps{
//   	ClientId: jsii.String("clientId"),
//   	UserPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	Css: jsii.String("css"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html
//
type CfnUserPoolUICustomizationAttachmentProps struct {
	// The app client ID for your UI customization.
	//
	// When this value isn't present, the customization applies to all user pool app clients that don't have client-level settings..
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html#cfn-cognito-userpooluicustomizationattachment-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The ID of the user pool where you want to apply branding to the classic hosted UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html#cfn-cognito-userpooluicustomizationattachment-userpoolid
	//
	UserPoolId interface{} `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// A plaintext CSS file that contains the custom fields that you want to apply to your user pool or app client.
	//
	// To download a template, go to the Amazon Cognito console. Navigate to your user pool *App clients* tab, select *Login pages* , edit *Hosted UI (classic) style* , and select the link to `CSS template.css` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html#cfn-cognito-userpooluicustomizationattachment-css
	//
	Css *string `field:"optional" json:"css" yaml:"css"`
}

