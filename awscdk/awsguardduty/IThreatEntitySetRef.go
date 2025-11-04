package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThreatEntitySet.
// Experimental.
type IThreatEntitySetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ThreatEntitySet resource.
	// Experimental.
	ThreatEntitySetRef() *ThreatEntitySetReference
}

// The jsii proxy for IThreatEntitySetRef
type jsiiProxy_IThreatEntitySetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IThreatEntitySetRef) ThreatEntitySetRef() *ThreatEntitySetReference {
	var returns *ThreatEntitySetReference
	_jsii_.Get(
		j,
		"threatEntitySetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThreatEntitySetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThreatEntitySetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

