package awscognito


// Standard attribute that can be marked as required or mutable.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	standardAttributes: &standardAttributes{
//   		fullname: &standardAttribute{
//   			required: jsii.Boolean(true),
//   			mutable: jsii.Boolean(false),
//   		},
//   		address: &standardAttribute{
//   			required: jsii.Boolean(false),
//   			mutable: jsii.Boolean(true),
//   		},
//   	},
//   	customAttributes: map[string]iCustomAttribute{
//   		"myappid": cognito.NewStringAttribute(&StringAttributeProps{
//   			"minLen": jsii.Number(5),
//   			"maxLen": jsii.Number(15),
//   			"mutable": jsii.Boolean(false),
//   		}),
//   		"callingcode": cognito.NewNumberAttribute(&NumberAttributeProps{
//   			"min": jsii.Number(1),
//   			"max": jsii.Number(3),
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"isEmployee": cognito.NewBooleanAttribute(&CustomAttributeProps{
//   			"mutable": jsii.Boolean(true),
//   		}),
//   		"joinedOn": cognito.NewDateTimeAttribute(),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes
//
type StandardAttribute struct {
	// Specifies whether the value of the attribute can be changed.
	//
	// For any user pool attribute that's mapped to an identity provider attribute, this must be set to `true`.
	// Amazon Cognito updates mapped attributes when users sign in to your application through an identity provider.
	// If an attribute is immutable, Amazon Cognito throws an error when it attempts to update the attribute.
	Mutable *bool `field:"optional" json:"mutable" yaml:"mutable"`
	// Specifies whether the attribute is required upon user registration.
	//
	// If the attribute is required and the user does not provide a value, registration or sign-in will fail.
	Required *bool `field:"optional" json:"required" yaml:"required"`
}

