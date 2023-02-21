package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents a Cognito UserPool.
type IUserPool interface {
	awscdk.IResource
	// Add a new app client to this user pool.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-client-apps.html
	//
	AddClient(id *string, options *UserPoolClientOptions) UserPoolClient
	// Associate a domain to this user pool.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-assign-domain.html
	//
	AddDomain(id *string, options *UserPoolDomainOptions) UserPoolDomain
	// Add a new resource server to this user pool.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-resource-servers.html
	//
	AddResourceServer(id *string, options *UserPoolResourceServerOptions) UserPoolResourceServer
	// Adds an IAM policy statement associated with this user pool to an IAM principal's policy.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Register an identity provider with this user pool.
	RegisterIdentityProvider(provider IUserPoolIdentityProvider)
	// Get all identity providers registered with this user pool.
	IdentityProviders() *[]IUserPoolIdentityProvider
	// The ARN of this user pool resource.
	UserPoolArn() *string
	// The physical ID of this user pool resource.
	UserPoolId() *string
}

// The jsii proxy for IUserPool
type jsiiProxy_IUserPool struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IUserPool) AddClient(id *string, options *UserPoolClientOptions) UserPoolClient {
	if err := i.validateAddClientParameters(id, options); err != nil {
		panic(err)
	}
	var returns UserPoolClient

	_jsii_.Invoke(
		i,
		"addClient",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IUserPool) AddDomain(id *string, options *UserPoolDomainOptions) UserPoolDomain {
	if err := i.validateAddDomainParameters(id, options); err != nil {
		panic(err)
	}
	var returns UserPoolDomain

	_jsii_.Invoke(
		i,
		"addDomain",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IUserPool) AddResourceServer(id *string, options *UserPoolResourceServerOptions) UserPoolResourceServer {
	if err := i.validateAddResourceServerParameters(id, options); err != nil {
		panic(err)
	}
	var returns UserPoolResourceServer

	_jsii_.Invoke(
		i,
		"addResourceServer",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IUserPool) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IUserPool) RegisterIdentityProvider(provider IUserPoolIdentityProvider) {
	if err := i.validateRegisterIdentityProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerIdentityProvider",
		[]interface{}{provider},
	)
}

func (j *jsiiProxy_IUserPool) IdentityProviders() *[]IUserPoolIdentityProvider {
	var returns *[]IUserPoolIdentityProvider
	_jsii_.Get(
		j,
		"identityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPool) UserPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPool) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

