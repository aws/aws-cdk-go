package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An attribute available from a third party identity provider.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   cognito.NewUserPoolIdentityProviderAmazon(this, jsii.String("Amazon"), &UserPoolIdentityProviderAmazonProps{
//   	ClientId: jsii.String("amzn-client-id"),
//   	ClientSecret: jsii.String("amzn-client-secret"),
//   	UserPool: userpool,
//   	AttributeMapping: &AttributeMapping{
//   		Email: cognito.ProviderAttribute_AMAZON_EMAIL(),
//   		Website: cognito.ProviderAttribute_Other(jsii.String("url")),
//   		 // use other() when an attribute is not pre-defined in the CDK
//   		Custom: map[string]providerAttribute{
//   			// custom user pool attributes go here
//   			"uniqueId": cognito.*providerAttribute_AMAZON_USER_ID(),
//   		},
//   	},
//   })
//
type ProviderAttribute interface {
	// The attribute value string as recognized by the provider.
	AttributeName() *string
}

// The jsii proxy struct for ProviderAttribute
type jsiiProxy_ProviderAttribute struct {
	_ byte // padding
}

func (j *jsiiProxy_ProviderAttribute) AttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeName",
		&returns,
	)
	return returns
}


// Use this to specify an attribute from the identity provider that is not pre-defined in the CDK.
func ProviderAttribute_Other(attributeName *string) ProviderAttribute {
	_init_.Initialize()

	if err := validateProviderAttribute_OtherParameters(attributeName); err != nil {
		panic(err)
	}
	var returns ProviderAttribute

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"other",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func ProviderAttribute_AMAZON_EMAIL() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"AMAZON_EMAIL",
		&returns,
	)
	return returns
}

func ProviderAttribute_AMAZON_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"AMAZON_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_AMAZON_POSTAL_CODE() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"AMAZON_POSTAL_CODE",
		&returns,
	)
	return returns
}

func ProviderAttribute_AMAZON_USER_ID() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"AMAZON_USER_ID",
		&returns,
	)
	return returns
}

func ProviderAttribute_APPLE_EMAIL() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"APPLE_EMAIL",
		&returns,
	)
	return returns
}

func ProviderAttribute_APPLE_FIRST_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"APPLE_FIRST_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_APPLE_LAST_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"APPLE_LAST_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_APPLE_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"APPLE_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_BIRTHDAY() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"FACEBOOK_BIRTHDAY",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_EMAIL() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"FACEBOOK_EMAIL",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_FIRST_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"FACEBOOK_FIRST_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_GENDER() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"FACEBOOK_GENDER",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_ID() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"FACEBOOK_ID",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_LAST_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"FACEBOOK_LAST_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_LOCALE() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"FACEBOOK_LOCALE",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_MIDDLE_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"FACEBOOK_MIDDLE_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_FACEBOOK_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"FACEBOOK_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_BIRTHDAYS() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"GOOGLE_BIRTHDAYS",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_EMAIL() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"GOOGLE_EMAIL",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_FAMILY_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"GOOGLE_FAMILY_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_GENDER() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"GOOGLE_GENDER",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_GIVEN_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"GOOGLE_GIVEN_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_NAME() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"GOOGLE_NAME",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_NAMES() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"GOOGLE_NAMES",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_PHONE_NUMBERS() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"GOOGLE_PHONE_NUMBERS",
		&returns,
	)
	return returns
}

func ProviderAttribute_GOOGLE_PICTURE() ProviderAttribute {
	_init_.Initialize()
	var returns ProviderAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.ProviderAttribute",
		"GOOGLE_PICTURE",
		&returns,
	)
	return returns
}

