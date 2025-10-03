package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClientCertificate.
// Experimental.
type IClientCertificateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IClientCertificateRef
type jsiiProxy_IClientCertificateRef struct {
	internal.Type__constructsIConstruct
}

