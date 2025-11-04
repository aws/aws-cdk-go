package awss3express

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3express/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessPoint.
// Experimental.
type IAccessPointRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AccessPoint resource.
	// Experimental.
	AccessPointRef() *AccessPointReference
}

// The jsii proxy for IAccessPointRef
type jsiiProxy_IAccessPointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAccessPointRef) AccessPointRef() *AccessPointReference {
	var returns *AccessPointReference
	_jsii_.Get(
		j,
		"accessPointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPointRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

