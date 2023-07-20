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
	// The client ID for the client app.
	//
	// You can specify the UI customization settings for a single client (with a specific clientId) or for all clients (by setting the clientId to `ALL` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html#cfn-cognito-userpooluicustomizationattachment-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The user pool ID for the user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html#cfn-cognito-userpooluicustomizationattachment-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// The CSS values in the UI customization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluicustomizationattachment.html#cfn-cognito-userpooluicustomizationattachment-css
	//
	Css *string `field:"optional" json:"css" yaml:"css"`
}

