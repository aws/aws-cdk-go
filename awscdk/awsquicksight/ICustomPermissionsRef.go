package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomPermissions.
// Experimental.
type ICustomPermissionsRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CustomPermissions resource.
	// Experimental.
	CustomPermissionsRef() *CustomPermissionsReference
}

// The jsii proxy for ICustomPermissionsRef
type jsiiProxy_ICustomPermissionsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ICustomPermissionsRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomPermissionsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

