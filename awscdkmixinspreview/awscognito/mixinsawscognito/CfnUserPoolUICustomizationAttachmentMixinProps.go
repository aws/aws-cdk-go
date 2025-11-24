package mixinsawscognito


// Properties for CfnUserPoolUICustomizationAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserPoolUICustomizationAttachmentMixinProps := &CfnUserPoolUICustomizationAttachmentMixinProps{
//   	ClientId: jsii.String("clientId"),
//   	Css: jsii.String("css"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html
//
type CfnUserPoolUICustomizationAttachmentMixinProps struct {
	// The app client ID for your UI customization.
	//
	// When this value isn't present, the customization applies to all user pool app clients that don't have client-level settings..
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html#cfn-cognito-userpooluicustomizationattachment-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// A plaintext CSS file that contains the custom fields that you want to apply to your user pool or app client.
	//
	// To download a template, go to the Amazon Cognito console. Navigate to your user pool *App clients* tab, select *Login pages* , edit *Hosted UI (classic) style* , and select the link to `CSS template.css` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html#cfn-cognito-userpooluicustomizationattachment-css
	//
	Css *string `field:"optional" json:"css" yaml:"css"`
	// The ID of the user pool where you want to apply branding to the classic hosted UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html#cfn-cognito-userpooluicustomizationattachment-userpoolid
	//
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

