package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TLSInspectionConfiguration.
// Experimental.
type ITLSInspectionConfigurationRef interface {
	constructs.IConstruct
	// A reference to a TLSInspectionConfiguration resource.
	// Experimental.
	TlsInspectionConfigurationRef() *TLSInspectionConfigurationReference
}

// The jsii proxy for ITLSInspectionConfigurationRef
type jsiiProxy_ITLSInspectionConfigurationRef struct {
	internal.Type__constructsIConstruct
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

