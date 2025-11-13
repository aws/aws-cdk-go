package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UsagePlanKey.
// Experimental.
type IUsagePlanKeyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a UsagePlanKey resource.
	// Experimental.
	UsagePlanKeyRef() *UsagePlanKeyReference
}

// The jsii proxy for IUsagePlanKeyRef
type jsiiProxy_IUsagePlanKeyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IUsagePlanKeyRef) UsagePlanKeyRef() *UsagePlanKeyReference {
	var returns *UsagePlanKeyReference
	_jsii_.Get(
		j,
		"usagePlanKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUsagePlanKeyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUsagePlanKeyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

