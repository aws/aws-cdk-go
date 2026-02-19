package interfacesawsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TLSInspectionConfiguration.
// Experimental.
type ITLSInspectionConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TLSInspectionConfiguration resource.
	// Experimental.
	TlsInspectionConfigurationRef() *TLSInspectionConfigurationReference
}

// The jsii proxy for ITLSInspectionConfigurationRef
type jsiiProxy_ITLSInspectionConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITLSInspectionConfigurationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITLSInspectionConfigurationRef) TlsInspectionConfigurationRef() *TLSInspectionConfigurationReference {
	var returns *TLSInspectionConfigurationReference
	_jsii_.Get(
		j,
		"tlsInspectionConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITLSInspectionConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITLSInspectionConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

