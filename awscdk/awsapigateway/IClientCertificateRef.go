package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClientCertificate.
// Experimental.
type IClientCertificateRef interface {
	constructs.IConstruct
	// A reference to a ClientCertificate resource.
	// Experimental.
	ClientCertificateRef() *ClientCertificateReference
}

// The jsii proxy for IClientCertificateRef
type jsiiProxy_IClientCertificateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IClientCertificateRef) ClientCertificateRef() *ClientCertificateReference {
	var returns *ClientCertificateReference
	_jsii_.Get(
		j,
		"clientCertificateRef",
		&returns,
	)
	return returns
}

