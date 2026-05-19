package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// IAM actions for AgentCore OAuth2 credential providers (Token Vault).
// See: https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonbedrockagentcore.html
//
type OAuth2CredentialProviderIdentityPerms interface {
}

// The jsii proxy struct for OAuth2CredentialProviderIdentityPerms
type jsiiProxy_OAuth2CredentialProviderIdentityPerms struct {
	_ byte // padding
}

func OAuth2CredentialProviderIdentityPerms_ADMIN_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderIdentityPerms",
		"ADMIN_PERMS",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderIdentityPerms_FULL_ACCESS_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderIdentityPerms",
		"FULL_ACCESS_PERMS",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderIdentityPerms_LIST_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderIdentityPerms",
		"LIST_PERMS",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderIdentityPerms_READ_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderIdentityPerms",
		"READ_PERMS",
		&returns,
	)
	return returns
}

func OAuth2CredentialProviderIdentityPerms_USE_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.OAuth2CredentialProviderIdentityPerms",
		"USE_PERMS",
		&returns,
	)
	return returns
}

