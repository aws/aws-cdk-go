package awscognito


// Props for constructing a StringAttr.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	StandardAttributes: &StandardAttributes{
//   		Fullname: &StandardAttribute{
//   			Required: jsii.Boolean(true),
//   			Mutable: jsii.Boolean(false),
//   		},
//   		Address: &StandardAttribute{
//   			Required: jsii.Boolean(false),
//   			Mutable: jsii.Boolean(true),
//   		},
//   	},
//   	CustomAttributes: map[string]iCustomAttribute{
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
type StringAttributeProps struct {
	// Maximum length of this attribute.
	// Default: 2048.
	//
	MaxLen *float64 `field:"optional" json:"maxLen" yaml:"maxLen"`
	// Minimum length of this attribute.
	// Default: 0.
	//
	MinLen *float64 `field:"optional" json:"minLen" yaml:"minLen"`
	// Specifies whether the value of the attribute can be changed.
	//
	// For any user pool attribute that's mapped to an identity provider attribute, you must set this parameter to true.
	// Amazon Cognito updates mapped attributes when users sign in to your application through an identity provider.
	// If an attribute is immutable, Amazon Cognito throws an error when it attempts to update the attribute.
	// Default: false.
	//
	Mutable *bool `field:"optional" json:"mutable" yaml:"mutable"`
}

