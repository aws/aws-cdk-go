package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HostedConfigurationVersion.
// Experimental.
type IHostedConfigurationVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IHostedConfigurationVersionRef
type jsiiProxy_IHostedConfigurationVersionRef struct {
	internal.Type__constructsIConstruct
}

