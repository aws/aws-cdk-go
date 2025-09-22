package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainNameV2.
// Experimental.
type IDomainNameV2Ref interface {
	constructs.IConstruct
	// A reference to a DomainNameV2 resource.
	// Experimental.
	DomainNameV2Ref() *DomainNameV2Reference
}

// The jsii proxy for IDomainNameV2Ref
type jsiiProxy_IDomainNameV2Ref struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDomainNameV2Ref) DomainNameV2Ref() *DomainNameV2Reference {
	var returns *DomainNameV2Reference
	_jsii_.Get(
		j,
		"domainNameV2Ref",
		&returns,
	)
	return returns
}

