package awspcaconnectorad

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspcaconnectorad/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServicePrincipalName.
// Experimental.
type IServicePrincipalNameRef interface {
	constructs.IConstruct
	// A reference to a ServicePrincipalName resource.
	// Experimental.
	ServicePrincipalNameRef() *ServicePrincipalNameReference
}

// The jsii proxy for IServicePrincipalNameRef
type jsiiProxy_IServicePrincipalNameRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServicePrincipalNameRef) ServicePrincipalNameRef() *ServicePrincipalNameReference {
	var returns *ServicePrincipalNameReference
	_jsii_.Get(
		j,
		"servicePrincipalNameRef",
		&returns,
	)
	return returns
}

