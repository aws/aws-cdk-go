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
	// The user pool ID for the user pool where the user will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.
	//
	// You create custom workflows by assigning AWS Lambda functions to user pool triggers. When you use the AdminCreateUser API action, Amazon Cognito invokes the function that is assigned to the *pre sign-up* trigger. When Amazon Cognito invokes this function, it passes a JSON payload, which the function receives as input. This payload contains a `clientMetadata` attribute, which provides the data that you assigned to the ClientMetadata parameter in your AdminCreateUser request. In your function code in AWS Lambda , you can process the `clientMetadata` value to enhance your workflow for your specific needs.
	//
	// For more information, see [Customizing user pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html) in the *Amazon Cognito Developer Guide* .
	//
	// > When you use the ClientMetadata parameter, remember that Amazon Cognito won't do the following:
	// >
	// > - Store the ClientMetadata value. This data is available only to AWS Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration doesn't include triggers, the ClientMetadata parameter serves no purpose.
	// > - Validate the ClientMetadata value.
	// > - Encrypt the ClientMetadata value. Don't use Amazon Cognito to provide sensitive information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-clientmetadata
	//
	ClientMetadata interface{} `field:"optional" json:"clientMetadata" yaml:"clientMetadata"`
	// Specify `"EMAIL"` if email will be used to send the welcome message.
	//
	// Specify `"SMS"` if the phone number will be used. The default value is `"SMS"` . You can specify more than one value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-desireddeliverymediums
	//
	DesiredDeliveryMediums *[]*string `field:"optional" json:"desiredDeliveryMediums" yaml:"desiredDeliveryMediums"`
	// This parameter is used only if the `phone_number_verified` or `email_verified` attribute is set to `True` .
	//
	// Otherwise, it is ignored.
	//
	// If this parameter is set to `True` and the phone number or email address specified in the UserAttributes parameter already exists as an alias with a different user, the API call will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias.
	//
	// If this parameter is set to `False` , the API throws an `AliasExistsException` error if the alias already exists. The default value is `False` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-forcealiascreation
	//
	ForceAliasCreation interface{} `field:"optional" json:"forceAliasCreation" yaml:"forceAliasCreation"`
	// Set to `RESEND` to resend the invitation message to a user that already exists and reset the expiration limit on the user's account.
	//
	// Set to `SUPPRESS` to suppress sending the message. You can specify only one value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-messageaction
	//
	MessageAction *string `field:"optional" json:"messageAction" yaml:"messageAction"`
	// An array of name-value pairs that contain user attributes and attribute values.
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
	// Your Lambda function can analyze this additional data and act on it. Your function might perform external API operations like logging user attributes and validation data to Amazon CloudWatch Logs. Validation data might also affect the response that your function returns to Amazon Cognito, like automatically confirming the user if they sign up from within your network.
	//
	// For more information about the pre sign-up Lambda trigger, see [Pre sign-up Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-sign-up.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-validationdata
	//
	ValidationData interface{} `field:"optional" json:"validationData" yaml:"validationData"`
}

