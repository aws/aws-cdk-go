package awscognito


// The settings for updates to user attributes.
//
// These settings include the property `AttributesRequireVerificationBeforeUpdate` ,
// a user-pool setting that tells Amazon Cognito how to handle changes to the value of your users' email address and phone number attributes. For
// more information, see [Verifying updates to email addresses and phone numbers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html#user-pool-settings-verifications-verify-attribute-updates) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userAttributeUpdateSettingsProperty := &UserAttributeUpdateSettingsProperty{
//   	AttributesRequireVerificationBeforeUpdate: []*string{
//   		jsii.String("attributesRequireVerificationBeforeUpdate"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userattributeupdatesettings.html
//
type CfnUserPool_UserAttributeUpdateSettingsProperty struct {
	// Requires that your user verifies their email address, phone number, or both before Amazon Cognito updates the value of that attribute.
	//
	// When you update a user attribute that has this option activated, Amazon Cognito sends a verification message to the new phone number or email address. Amazon Cognito doesn’t change the value of the attribute until your user responds to the verification message and confirms the new value.
	//
	// You can verify an updated email address or phone number with a `API_VerifyUserAttribute` API request. You can also call the `API_AdminUpdateUserAttributes` API and set `email_verified` or `phone_number_verified` to true.
	//
	// When `AttributesRequireVerificationBeforeUpdate` is false, your user pool doesn't require that your users verify attribute changes before Amazon Cognito updates them. In a user pool where `AttributesRequireVerificationBeforeUpdate` is false, API operations that change attribute values can immediately update a user’s `email` or `phone_number` attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userattributeupdatesettings.html#cfn-cognito-userpool-userattributeupdatesettings-attributesrequireverificationbeforeupdate
	//
	AttributesRequireVerificationBeforeUpdate *[]*string `field:"required" json:"attributesRequireVerificationBeforeUpdate" yaml:"attributesRequireVerificationBeforeUpdate"`
}

