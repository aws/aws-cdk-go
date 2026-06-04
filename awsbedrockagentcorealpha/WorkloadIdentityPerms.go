package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// IAM actions for AgentCore workload identities.
// See: https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonbedrockagentcore.html
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type WorkloadIdentityPerms interface {
}

// The jsii proxy struct for WorkloadIdentityPerms
type jsiiProxy_WorkloadIdentityPerms struct {
	_ byte // padding
}

func WorkloadIdentityPerms_ADMIN_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.WorkloadIdentityPerms",
		"ADMIN_PERMS",
		&returns,
	)
	return returns
}

func WorkloadIdentityPerms_FULL_ACCESS_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.WorkloadIdentityPerms",
		"FULL_ACCESS_PERMS",
		&returns,
	)
	return returns
}

func WorkloadIdentityPerms_LIST_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.WorkloadIdentityPerms",
		"LIST_PERMS",
		&returns,
	)
	return returns
}

func WorkloadIdentityPerms_READ_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.WorkloadIdentityPerms",
		"READ_PERMS",
		&returns,
	)
	return returns
}

func WorkloadIdentityPerms_USE_PERMS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.WorkloadIdentityPerms",
		"USE_PERMS",
		&returns,
	)
	return returns
}

