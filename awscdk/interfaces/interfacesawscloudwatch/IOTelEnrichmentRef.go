package interfacesawscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OTelEnrichment.
// Experimental.
type IOTelEnrichmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OTelEnrichment resource.
	// Experimental.
	OTelEnrichmentRef() *OTelEnrichmentReference
}

// The jsii proxy for IOTelEnrichmentRef
type jsiiProxy_IOTelEnrichmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IOTelEnrichmentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOTelEnrichmentRef) OTelEnrichmentRef() *OTelEnrichmentReference {
	var returns *OTelEnrichmentReference
	_jsii_.Get(
		j,
		"oTelEnrichmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOTelEnrichmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOTelEnrichmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

