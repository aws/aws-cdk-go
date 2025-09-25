package awssecuritylake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecuritylake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AwsLogSource.
// Experimental.
type IAwsLogSourceRef interface {
	constructs.IConstruct
	// A reference to a AwsLogSource resource.
	// Experimental.
	AwsLogSourceRef() *AwsLogSourceReference
}

// The jsii proxy for IAwsLogSourceRef
type jsiiProxy_IAwsLogSourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAwsLogSourceRef) AwsLogSourceRef() *AwsLogSourceReference {
	var returns *AwsLogSourceReference
	_jsii_.Get(
		j,
		"awsLogSourceRef",
		&returns,
	)
	return returns
}

