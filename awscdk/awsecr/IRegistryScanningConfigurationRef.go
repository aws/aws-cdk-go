package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RegistryScanningConfiguration.
// Experimental.
type IRegistryScanningConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRegistryScanningConfigurationRef
type jsiiProxy_IRegistryScanningConfigurationRef struct {
	internal.Type__constructsIConstruct
}

