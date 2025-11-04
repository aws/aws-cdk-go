package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TLSInspectionConfiguration.
// Experimental.
type ITLSInspectionConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TLSInspectionConfiguration resource.
	// Experimental.
	TlsInspectionConfigurationRef() *TLSInspectionConfigurationReference
}

// The jsii proxy for ITLSInspectionConfigurationRef
type jsiiProxy_ITLSInspectionConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ITLSInspectionConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

