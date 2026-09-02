package interfacesawsinvoicing

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsinvoicing/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProcurementPortalPreference.
// Experimental.
type IProcurementPortalPreferenceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ProcurementPortalPreference resource.
	// Experimental.
	ProcurementPortalPreferenceRef() *ProcurementPortalPreferenceReference
}

// The jsii proxy for IProcurementPortalPreferenceRef
type jsiiProxy_IProcurementPortalPreferenceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IProcurementPortalPreferenceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IProcurementPortalPreferenceRef) ProcurementPortalPreferenceRef() *ProcurementPortalPreferenceReference {
	var returns *ProcurementPortalPreferenceReference
	_jsii_.Get(
		j,
		"procurementPortalPreferenceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProcurementPortalPreferenceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProcurementPortalPreferenceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

