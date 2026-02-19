package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UsageProfile.
// Experimental.
type IUsageProfileRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a UsageProfile resource.
	// Experimental.
	UsageProfileRef() *UsageProfileReference
}

// The jsii proxy for IUsageProfileRef
type jsiiProxy_IUsageProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IUsageProfileRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IUsageProfileRef) UsageProfileRef() *UsageProfileReference {
	var returns *UsageProfileReference
	_jsii_.Get(
		j,
		"usageProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUsageProfileRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUsageProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

