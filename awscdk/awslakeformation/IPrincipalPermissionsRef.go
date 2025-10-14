package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrincipalPermissions.
// Experimental.
type IPrincipalPermissionsRef interface {
	constructs.IConstruct
	// A reference to a PrincipalPermissions resource.
	// Experimental.
	PrincipalPermissionsRef() *PrincipalPermissionsReference
}

// The jsii proxy for IPrincipalPermissionsRef
type jsiiProxy_IPrincipalPermissionsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPrincipalPermissionsRef) PrincipalPermissionsRef() *PrincipalPermissionsReference {
	var returns *PrincipalPermissionsReference
	_jsii_.Get(
		j,
		"principalPermissionsRef",
		&returns,
	)
	return returns
}

