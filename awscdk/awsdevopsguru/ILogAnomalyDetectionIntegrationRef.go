package awsdevopsguru

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevopsguru/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogAnomalyDetectionIntegration.
// Experimental.
type ILogAnomalyDetectionIntegrationRef interface {
	constructs.IConstruct
	// A reference to a LogAnomalyDetectionIntegration resource.
	// Experimental.
	LogAnomalyDetectionIntegrationRef() *LogAnomalyDetectionIntegrationReference
}

// The jsii proxy for ILogAnomalyDetectionIntegrationRef
type jsiiProxy_ILogAnomalyDetectionIntegrationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILogAnomalyDetectionIntegrationRef) LogAnomalyDetectionIntegrationRef() *LogAnomalyDetectionIntegrationReference {
	var returns *LogAnomalyDetectionIntegrationReference
	_jsii_.Get(
		j,
		"logAnomalyDetectionIntegrationRef",
		&returns,
	)
	return returns
}

