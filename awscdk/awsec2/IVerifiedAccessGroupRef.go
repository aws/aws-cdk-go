package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VerifiedAccessGroup.
// Experimental.
type IVerifiedAccessGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VerifiedAccessGroup resource.
	// Experimental.
	VerifiedAccessGroupRef() *VerifiedAccessGroupReference
}

// The jsii proxy for IVerifiedAccessGroupRef
type jsiiProxy_IVerifiedAccessGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IVerifiedAccessGroupRef) VerifiedAccessGroupRef() *VerifiedAccessGroupReference {
	var returns *VerifiedAccessGroupReference
	_jsii_.Get(
		j,
		"verifiedAccessGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVerifiedAccessGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVerifiedAccessGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

