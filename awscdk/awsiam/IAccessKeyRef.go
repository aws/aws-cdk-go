package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessKey.
// Experimental.
type IAccessKeyRef interface {
	constructs.IConstruct
	// A reference to a AccessKey resource.
	// Experimental.
	AccessKeyRef() *AccessKeyReference
}

// The jsii proxy for IAccessKeyRef
type jsiiProxy_IAccessKeyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccessKeyRef) AccessKeyRef() *AccessKeyReference {
	var returns *AccessKeyReference
	_jsii_.Get(
		j,
		"accessKeyRef",
		&returns,
	)
	return returns
}

