package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// IAM actions for AgentCore API key credential providers (Token Vault).
// See: https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonbedrockagentcore.html
//
type ApiKeyCredentialProviderIdentityPerms interface {
}

// The jsii proxy struct for ApiKeyCredentialProviderIdentityPerms
type jsiiProxy_ApiKeyCredentialProviderIdentityPerms struct {
	_ byte // padding
}

func ApiKeyCredentialProviderIdentityPerms_ADMIN_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialProviderIdentityPerms",
		"ADMIN_PERMS",
		&returns,
	)
	return returns
}

func ApiKeyCredentialProviderIdentityPerms_FULL_ACCESS_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialProviderIdentityPerms",
		"FULL_ACCESS_PERMS",
		&returns,
	)
	return returns
}

func ApiKeyCredentialProviderIdentityPerms_LIST_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialProviderIdentityPerms",
		"LIST_PERMS",
		&returns,
	)
	return returns
}

func ApiKeyCredentialProviderIdentityPerms_READ_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialProviderIdentityPerms",
		"READ_PERMS",
		&returns,
	)
	return returns
}

func ApiKeyCredentialProviderIdentityPerms_USE_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialProviderIdentityPerms",
		"USE_PERMS",
		&returns,
	)
	return returns
}

