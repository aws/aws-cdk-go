package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPAMScope.
// Experimental.
type IIPAMScopeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a IPAMScope resource.
	// Experimental.
	IpamScopeRef() *IPAMScopeReference
}

// The jsii proxy for IIPAMScopeRef
type jsiiProxy_IIPAMScopeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IIPAMScopeRef) IpamScopeRef() *IPAMScopeReference {
	var returns *IPAMScopeReference
	_jsii_.Get(
		j,
		"ipamScopeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIPAMScopeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIPAMScopeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

