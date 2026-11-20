package awscognito


// Properties for defining a `CfnUserPoolClientSessionConfigurationAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolClientSessionConfigurationAttachmentProps := &CfnUserPoolClientSessionConfigurationAttachmentProps{
//   	ClientId: jsii.String("clientId"),
//   	SsoSession: &SSOSessionProperty{
//   		MaxValidityMinutes: jsii.Number(123),
//   	},
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclientsessionconfigurationattachment.html
//
type CfnUserPoolClientSessionConfigurationAttachmentProps struct {
	// The ID of the app client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclientsessionconfigurationattachment.html#cfn-cognito-userpoolclientsessionconfigurationattachment-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Managed login single sign-on session settings for the app client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclientsessionconfigurationattachment.html#cfn-cognito-userpoolclientsessionconfigurationattachment-ssosession
	//
	SsoSession interface{} `field:"required" json:"ssoSession" yaml:"ssoSession"`
	// The ID of the user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclientsessionconfigurationattachment.html#cfn-cognito-userpoolclientsessionconfigurationattachment-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

