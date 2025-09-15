package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainName.
// Experimental.
type IDomainNameRef interface {
	constructs.IConstruct
	// A reference to a DomainName resource.
	// Experimental.
	DomainNameRef() *DomainNameReference
}

// The jsii proxy for IDomainNameRef
type jsiiProxy_IDomainNameRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDomainNameRef) DomainNameRef() *DomainNameReference {
	var returns *DomainNameReference
	_jsii_.Get(
		j,
		"domainNameRef",
		&returns,
	)
	return returns
}

