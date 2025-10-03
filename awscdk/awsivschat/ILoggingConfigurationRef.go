package awsivschat

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivschat/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoggingConfiguration.
// Experimental.
type ILoggingConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILoggingConfigurationRef
type jsiiProxy_ILoggingConfigurationRef struct {
	internal.Type__constructsIConstruct
}

