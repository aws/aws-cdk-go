package awscognito


// Properties for CfnUserPoolSessionConfigurationAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnUserPoolSessionConfigurationAttachmentMixinProps := &CfnUserPoolSessionConfigurationAttachmentMixinProps{
//   	TokenSession: &TokenSessionProperty{
//   		MaxPerUser: jsii.Number(123),
//   	},
//   	UserPoolId: jsii.String("userPoolId"),
//   	UserSessionTracking: jsii.String("userSessionTracking"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolsessionconfigurationattachment.html
//
type CfnUserPoolSessionConfigurationAttachmentMixinProps struct {
	// Token session settings that apply to every app client in the user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolsessionconfigurationattachment.html#cfn-cognito-userpoolsessionconfigurationattachment-tokensession
	//
	TokenSession interface{} `field:"optional" json:"tokenSession" yaml:"tokenSession"`
	// The ID of the user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolsessionconfigurationattachment.html#cfn-cognito-userpoolsessionconfigurationattachment-userpoolid
	//
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
	// Whether server-side user session tracking is enabled for the user pool.
	//
	// Requires the ESSENTIALS or PLUS feature plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolsessionconfigurationattachment.html#cfn-cognito-userpoolsessionconfigurationattachment-usersessiontracking
	//
	UserSessionTracking *string `field:"optional" json:"userSessionTracking" yaml:"userSessionTracking"`
}

