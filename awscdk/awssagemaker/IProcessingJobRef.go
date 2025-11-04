package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProcessingJob.
// Experimental.
type IProcessingJobRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ProcessingJob resource.
	// Experimental.
	ProcessingJobRef() *ProcessingJobReference
}

// The jsii proxy for IProcessingJobRef
type jsiiProxy_IProcessingJobRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IProcessingJobRef) ProcessingJobRef() *ProcessingJobReference {
	var returns *ProcessingJobReference
	_jsii_.Get(
		j,
		"processingJobRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProcessingJobRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProcessingJobRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

