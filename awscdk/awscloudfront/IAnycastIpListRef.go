package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnycastIpList.
// Experimental.
type IAnycastIpListRef interface {
	constructs.IConstruct
	// A reference to a AnycastIpList resource.
	// Experimental.
	AnycastIpListRef() *AnycastIpListReference
}

// The jsii proxy for IAnycastIpListRef
type jsiiProxy_IAnycastIpListRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAnycastIpListRef) AnycastIpListRef() *AnycastIpListReference {
	var returns *AnycastIpListReference
	_jsii_.Get(
		j,
		"anycastIpListRef",
		&returns,
	)
	return returns
}

