package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RemediationConfiguration.
// Experimental.
type IRemediationConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRemediationConfigurationRef
type jsiiProxy_IRemediationConfigurationRef struct {
	internal.Type__constructsIConstruct
}

