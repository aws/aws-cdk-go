package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Abstract interface for gateway credential provider configuration.
// Experimental.
type ICredentialProviderConfig interface {
	// Grant the role the permissions.
	// Experimental.
	GrantNeededPermissionsToRole(role awsiam.IRole) awsiam.Grant
	// The credential provider type.
	// Experimental.
	CredentialProviderType() CredentialProviderType
}

// The jsii proxy for ICredentialProviderConfig
type jsiiProxy_ICredentialProviderConfig struct {
	_ byte // padding
}

func (i *jsiiProxy_ICredentialProviderConfig) GrantNeededPermissionsToRole(role awsiam.IRole) awsiam.Grant {
	if err := i.validateGrantNeededPermissionsToRoleParameters(role); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantNeededPermissionsToRole",
		[]interface{}{role},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICredentialProviderConfig) CredentialProviderType() CredentialProviderType {
	var returns CredentialProviderType
	_jsii_.Get(
		j,
		"credentialProviderType",
		&returns,
	)
	return returns
}

