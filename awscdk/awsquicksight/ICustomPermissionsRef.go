package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomPermissions.
// Experimental.
type ICustomPermissionsRef interface {
	constructs.IConstruct
	// A reference to a CustomPermissions resource.
	// Experimental.
	CustomPermissionsRef() *CustomPermissionsReference
}

// The jsii proxy for ICustomPermissionsRef
type jsiiProxy_ICustomPermissionsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICustomPermissionsRef) CustomPermissionsRef() *CustomPermissionsReference {
	var returns *CustomPermissionsReference
	_jsii_.Get(
		j,
		"customPermissionsRef",
		&returns,
	)
	return returns
}

