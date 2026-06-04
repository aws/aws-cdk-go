package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore"
	"github.com/aws/constructs-go/constructs/v10"
)

// An API key credential provider registered in AgentCore Token Vault.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type IApiKeyCredentialProvider interface {
	interfacesawsbedrockagentcore.IApiKeyCredentialProviderRef
	awsiam.IGrantable
	awscdk.IResource
	// ARNs for use with gateway targets (`GatewayCredentialProvider.fromApiKeyIdentity` or `fromApiKeyIdentityArn`).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	BindForGatewayApiKeyTarget() *GatewayApiKeyIdentityBinding
	// Grants IAM actions to the IAM principal.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant control plane permissions to manage this provider.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GrantAdmin(grantee awsiam.IGrantable) awsiam.Grant
	// Grant read, admin, and credential retrieval permissions.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant
	// Grant `GetApiKeyCredentialProvider` and `ListApiKeyCredentialProviders`, scoped to this provider and parent resources required by the Bedrock AgentCore authorization model.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permission to retrieve API key material for outbound calls (`GetResourceApiKey`).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GrantUse(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the Secrets Manager secret that stores the API key after the resource is created.
	//
	// May be undefined for resources imported without this attribute.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ApiKeySecretArn() *string
	// Timestamp when the credential provider was created.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CreatedTime() *string
	// The ARN of this credential provider.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderArn() *string
	// Timestamp when the credential provider was last updated.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	LastUpdatedTime() *string
}

// The jsii proxy for IApiKeyCredentialProvider
type jsiiProxy_IApiKeyCredentialProvider struct {
	internal.Type__interfacesawsbedrockagentcoreIApiKeyCredentialProviderRef
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IApiKeyCredentialProvider) BindForGatewayApiKeyTarget() *GatewayApiKeyIdentityBinding {
	var returns *GatewayApiKeyIdentityBinding

	_jsii_.Invoke(
		i,
		"bindForGatewayApiKeyTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiKeyCredentialProvider) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiKeyCredentialProvider) GrantAdmin(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantAdminParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantAdmin",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiKeyCredentialProvider) GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantFullAccessParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantFullAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiKeyCredentialProvider) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiKeyCredentialProvider) GrantUse(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantUseParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantUse",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiKeyCredentialProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IApiKeyCredentialProvider) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IApiKeyCredentialProvider) ApiKeySecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKeyCredentialProvider) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKeyCredentialProvider) CredentialProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKeyCredentialProvider) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKeyCredentialProvider) ApiKeyCredentialProviderRef() *interfacesawsbedrockagentcore.ApiKeyCredentialProviderReference {
	var returns *interfacesawsbedrockagentcore.ApiKeyCredentialProviderReference
	_jsii_.Get(
		j,
		"apiKeyCredentialProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKeyCredentialProvider) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKeyCredentialProvider) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKeyCredentialProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKeyCredentialProvider) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

