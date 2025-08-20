package awscognito


// Properties for defining a `CfnUserPoolUser`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolUserProps := &CfnUserPoolUserProps{
//   	UserPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	ClientMetadata: map[string]*string{
//   		"clientMetadataKey": jsii.String("clientMetadata"),
//   	},
//   	DesiredDeliveryMediums: []*string{
//   		jsii.String("desiredDeliveryMediums"),
//   	},
//   	ForceAliasCreation: jsii.Boolean(false),
//   	MessageAction: jsii.String("messageAction"),
//   	UserAttributes: []interface{}{
//   		&AttributeTypeProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Username: jsii.String("username"),
//   	ValidationData: []interface{}{
//   		&AttributeTypeProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html
//
type CfnUserPoolUserProps struct {
	// The ID of the user pool where you want to create a user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.
	//
	// You create custom workflows by assigning AWS Lambda functions to user pool triggers. When you use the AdminCreateUser API action, Amazon Cognito invokes the function that is assigned to the *pre sign-up* trigger. When Amazon Cognito invokes this function, it passes a JSON payload, which the function receives as input. This payload contains a `ClientMetadata` attribute, which provides the data that you assigned to the ClientMetadata parameter in your AdminCreateUser request. In your function code in AWS Lambda , you can process the `clientMetadata` value to enhance your workflow for your specific needs.
	//
	// For more information, see [Using Lambda triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html) in the *Amazon Cognito Developer Guide* .
	//
	// > When you use the `ClientMetadata` parameter, note that Amazon Cognito won't do the following:
	// >
	// > - Store the `ClientMetadata` value. This data is available only to AWS Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration doesn't include triggers, the `ClientMetadata` parameter serves no purpose.
	// > - Validate the `ClientMetadata` value.
	// > - Encrypt the `ClientMetadata` value. Don't send sensitive information in this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-clientmetadata
	//
	ClientMetadata interface{} `field:"optional" json:"clientMetadata" yaml:"clientMetadata"`
	// Specify `EMAIL` if email will be used to send the welcome message.
	//
	// Specify `SMS` if the phone number will be used. The default value is `SMS` . You can specify more than one value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-desireddeliverymediums
	//
	DesiredDeliveryMediums *[]*string `field:"optional" json:"desiredDeliveryMediums" yaml:"desiredDeliveryMediums"`
	// This parameter is used only if the `phone_number_verified` or `email_verified` attribute is set to `True` .
	//
	// Otherwise, it is ignored.
	//
	// If this parameter is set to `True` and the phone number or email address specified in the `UserAttributes` parameter already exists as an alias with a different user, this request migrates the alias from the previous user to the newly-created user. The previous user will no longer be able to log in using that alias.
	//
	// If this parameter is set to `False` , the API throws an `AliasExistsException` error if the alias already exists. The default value is `False` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-forcealiascreation
	//
	ForceAliasCreation interface{} `field:"optional" json:"forceAliasCreation" yaml:"forceAliasCreation"`
	// Set to `RESEND` to resend the invitation message to a user that already exists, and to reset the temporary-password duration with a new temporary password.
	//
	// Set to `SUPPRESS` to suppress sending the message. You can specify only one value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-messageaction
	//
	MessageAction *string `field:"optional" json:"messageAction" yaml:"messageAction"`
	// An array of name-value pairs that contain user attributes and attribute values to be set for the user to be created.
	//
	// You can create a user without specifying any attributes other than `Username` . However, any attributes that you specify as required (when creating a user pool or in the *Attributes* tab of the console) either you should supply (in your call to `AdminCreateUser` ) or the user should supply (when they sign up in response to your welcome message).
	//
	// For custom attributes, you must prepend the `custom:` prefix to the attribute name.
	//
	// To send a message inviting the user to sign up, you must specify the user's email address or phone number. You can do this in your call to AdminCreateUser or in the *Users* tab of the Amazon Cognito console for managing your user pools.
	//
	// You must also provide an email address or phone number when you expect the user to do passwordless sign-in with an email or SMS OTP. These attributes must be provided when passwordless options are the only available, or when you don't submit a `TemporaryPassword` .
	//
	// In your call to `AdminCreateUser` , you can set the `email_verified` attribute to `True` , and you can set the `phone_number_verified` attribute to `True` .
	//
	// - *email* : The email address of the user to whom the message that contains the code and username will be sent. Required if the `email_verified` attribute is set to `True` , or if `"EMAIL"` is specified in the `DesiredDeliveryMediums` parameter.
	// - *phone_number* : The phone number of the user to whom the message that contains the code and username will be sent. Required if the `phone_number_verified` attribute is set to `True` , or if `"SMS"` is specified in the `DesiredDeliveryMediums` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-userattributes
	//
	UserAttributes interface{} `field:"optional" json:"userAttributes" yaml:"userAttributes"`
	// The value that you want to set as the username sign-in attribute.
	//
	// The following conditions apply to the username parameter.
	//
	// - The username can't be a duplicate of another username in the same user pool.
	// - You can't change the value of a username after you create it.
	// - You can only provide a value if usernames are a valid sign-in attribute for your user pool. If your user pool only supports phone numbers or email addresses as sign-in attributes, Amazon Cognito automatically generates a username value. For more information, see [Customizing sign-in attributes](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-aliases) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Temporary user attributes that contribute to the outcomes of your pre sign-up Lambda trigger.
	//
	// This set of key-value pairs are for custom validation of information that you collect from your users but don't need to retain.
	//
	// Your Lambda function can analyze this additional data and act on it. Your function can automatically confirm and verify select users or perform external API operations like logging user attributes and validation data to Amazon CloudWatch Logs.
	//
	// For more information about the pre sign-up Lambda trigger, see [Pre sign-up Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-sign-up.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-validationdata
	//
	ValidationData interface{} `field:"optional" json:"validationData" yaml:"validationData"`
}

