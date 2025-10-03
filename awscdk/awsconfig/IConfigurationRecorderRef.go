package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationRecorder.
// Experimental.
type IConfigurationRecorderRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfigurationRecorderRef
type jsiiProxy_IConfigurationRecorderRef struct {
	internal.Type__constructsIConstruct
}

