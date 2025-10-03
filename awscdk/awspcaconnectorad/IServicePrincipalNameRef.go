package awspcaconnectorad

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspcaconnectorad/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServicePrincipalName.
// Experimental.
type IServicePrincipalNameRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServicePrincipalNameRef
type jsiiProxy_IServicePrincipalNameRef struct {
	internal.Type__constructsIConstruct
}

