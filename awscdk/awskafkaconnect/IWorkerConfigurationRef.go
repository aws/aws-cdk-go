package awskafkaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskafkaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkerConfiguration.
// Experimental.
type IWorkerConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWorkerConfigurationRef
type jsiiProxy_IWorkerConfigurationRef struct {
	internal.Type__constructsIConstruct
}

