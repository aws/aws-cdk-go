package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Built-in OAuth2 vendors supported by `AWS::BedrockAgentCore::OAuth2CredentialProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuth2CredentialProviderVendor := awscdk.Aws_bedrockagentcore.OAuth2CredentialProviderVendor_ATLASSIAN()
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html
//
type OAuth2CredentialProviderVendor interface {
	// The vendor string value.
	Value() *string
	// Returns the string value.
	ToString() *string
}

// The jsii proxy struct for OAuth2CredentialProviderVendor
type jsiiProxy_OAuth2CredentialProviderVendor struct {
	_ byte // padding
}

func (j *jsiiProxy_OAuth2CredentialProviderVendor) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom OAuth2 vendor not yet defined in this class.
func OAuth2CredentialProviderVendor_Of(value *string) OAuth2CredentialProviderVendor {
	_init_.Initialize()

	if err := validateOAuth2CredentialProviderVendor_OfParameters(value); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProviderVendor

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func OAuth2CredentialProviderVendor_ATLASSIAN() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"ATLASSIAN",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_AUTH0() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"AUTH0",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_COGNITO() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"COGNITO",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_CUSTOM() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"CUSTOM",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_CYBER_ARK() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"CYBER_ARK",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_DROPBOX() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"DROPBOX",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_FACEBOOK() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"FACEBOOK",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_FUSION_AUTH() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"FUSION_AUTH",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_GITHUB() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"GITHUB",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_GOOGLE() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"GOOGLE",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_HUBSPOT() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"HUBSPOT",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_LINKEDIN() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"LINKEDIN",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_MICROSOFT() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"MICROSOFT",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_NOTION() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"NOTION",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_OKTA() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"OKTA",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_ONE_LOGIN() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"ONE_LOGIN",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_PING_ONE() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"PING_ONE",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_REDDIT() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"REDDIT",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_SALESFORCE() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"SALESFORCE",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_SLACK() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"SLACK",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_SPOTIFY() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"SPOTIFY",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_TWITCH() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"TWITCH",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_X() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"X",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_YANDEX() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"YANDEX",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderVendor_ZOOM() OAuth2CredentialProviderVendor {
	_init_.Initialize()
	var returns OAuth2CredentialProviderVendor
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderVendor",
		"ZOOM",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OAuth2CredentialProviderVendor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

