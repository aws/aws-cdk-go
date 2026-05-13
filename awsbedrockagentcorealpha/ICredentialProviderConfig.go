package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Abstract interface for gateway credential provider configuration.
// Experimental.
type ICredentialProviderConfig interface {
	// Grant the gateway's execution role the permissions needed for outbound authentication.
	// Experimental.
	GrantNeededPermissionsToRole(gateway IGateway) awsiam.Grant
	// The credential provider type.
	// Experimental.
	CredentialProviderType() CredentialProviderType
}

// The jsii proxy for ICredentialProviderConfig
type jsiiProxy_ICredentialProviderConfig struct {
	_ byte // padding
}

func (i *jsiiProxy_ICredentialProviderConfig) GrantNeededPermissionsToRole(gateway IGateway) awsiam.Grant {
	if err := i.validateGrantNeededPermissionsToRoleParameters(gateway); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantNeededPermissionsToRole",
		[]interface{}{gateway},
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

