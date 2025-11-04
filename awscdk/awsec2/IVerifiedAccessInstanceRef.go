package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VerifiedAccessInstance.
// Experimental.
type IVerifiedAccessInstanceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VerifiedAccessInstance resource.
	// Experimental.
	VerifiedAccessInstanceRef() *VerifiedAccessInstanceReference
}

// The jsii proxy for IVerifiedAccessInstanceRef
type jsiiProxy_IVerifiedAccessInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IVerifiedAccessInstanceRef) VerifiedAccessInstanceRef() *VerifiedAccessInstanceReference {
	var returns *VerifiedAccessInstanceReference
	_jsii_.Get(
		j,
		"verifiedAccessInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVerifiedAccessInstanceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVerifiedAccessInstanceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

