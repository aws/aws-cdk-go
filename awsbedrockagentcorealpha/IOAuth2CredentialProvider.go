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

// An OAuth2 credential provider registered in AgentCore Token Vault.
// Experimental.
type IOAuth2CredentialProvider interface {
	awsiam.IGrantable
	interfacesawsbedrockagentcore.IOAuth2CredentialProviderRef
	awscdk.IResource
	// ARNs and OAuth scopes for gateway targets (`GatewayCredentialProvider.fromOauthIdentity` or `fromOauthIdentityArn`).
	// Experimental.
	BindForGatewayOAuthTarget(scopes *[]*string, customParameters *map[string]*string) *GatewayOAuth2IdentityBinding
	// Grants IAM actions to the IAM principal.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant control plane permissions to manage this provider.
	// Experimental.
	GrantAdmin(grantee awsiam.IGrantable) awsiam.Grant
	// Grant read, admin, and token retrieval permissions.
	// Experimental.
	GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant
	// Grant `GetOauth2CredentialProvider` and `ListOauth2CredentialProviders`, scoped to this provider and parent resources required by the Bedrock AgentCore authorization model.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permission to retrieve OAuth tokens (`GetResourceOauth2Token`, `CompleteResourceTokenAuth`).
	// Experimental.
	GrantUse(grantee awsiam.IGrantable) awsiam.Grant
	// Callback URL for the OAuth2 authorization flow.
	// Experimental.
	CallbackUrl() *string
	// The ARN of the Secrets Manager secret for the OAuth2 client credentials.
	//
	// May be undefined for resources imported without this attribute.
	// Experimental.
	ClientSecretArn() *string
	// Timestamp when the credential provider was created.
	// Experimental.
	CreatedTime() *string
	// The ARN of this credential provider.
	// Experimental.
	CredentialProviderArn() *string
	// OAuth2 vendor string passed to CloudFormation.
	// Experimental.
	CredentialProviderVendor() *string
	// Timestamp when the credential provider was last updated.
	// Experimental.
	LastUpdatedTime() *string
}

// The jsii proxy for IOAuth2CredentialProvider
type jsiiProxy_IOAuth2CredentialProvider struct {
	internal.Type__awsiamIGrantable
	internal.Type__interfacesawsbedrockagentcoreIOAuth2CredentialProviderRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IOAuth2CredentialProvider) BindForGatewayOAuthTarget(scopes *[]*string, customParameters *map[string]*string) *GatewayOAuth2IdentityBinding {
	if err := i.validateBindForGatewayOAuthTargetParameters(scopes); err != nil {
		panic(err)
	}
	var returns *GatewayOAuth2IdentityBinding

	_jsii_.Invoke(
		i,
		"bindForGatewayOAuthTarget",
		[]interface{}{scopes, customParameters},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IOAuth2CredentialProvider) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IOAuth2CredentialProvider) GrantAdmin(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IOAuth2CredentialProvider) GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IOAuth2CredentialProvider) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IOAuth2CredentialProvider) GrantUse(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IOAuth2CredentialProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IOAuth2CredentialProvider) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOAuth2CredentialProvider) CallbackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callbackUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuth2CredentialProvider) ClientSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuth2CredentialProvider) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuth2CredentialProvider) CredentialProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuth2CredentialProvider) CredentialProviderVendor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialProviderVendor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuth2CredentialProvider) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuth2CredentialProvider) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuth2CredentialProvider) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuth2CredentialProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuth2CredentialProvider) OAuth2CredentialProviderRef() *interfacesawsbedrockagentcore.OAuth2CredentialProviderReference {
	var returns *interfacesawsbedrockagentcore.OAuth2CredentialProviderReference
	_jsii_.Get(
		j,
		"oAuth2CredentialProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuth2CredentialProvider) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

