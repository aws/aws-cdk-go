package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrustStore.
// Experimental.
type ITrustStoreRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITrustStoreRef
type jsiiProxy_ITrustStoreRef struct {
	internal.Type__constructsIConstruct
}

