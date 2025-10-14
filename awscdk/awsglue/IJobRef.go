package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Job.
// Experimental.
type IJobRef interface {
	constructs.IConstruct
	// A reference to a Job resource.
	// Experimental.
	JobRef() *JobReference
}

// The jsii proxy for IJobRef
type jsiiProxy_IJobRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IJobRef) JobRef() *JobReference {
	var returns *JobReference
	_jsii_.Get(
		j,
		"jobRef",
		&returns,
	)
	return returns
}

