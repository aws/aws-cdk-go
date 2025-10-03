package awsdatazone

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentBlueprintConfiguration.
// Experimental.
type IEnvironmentBlueprintConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEnvironmentBlueprintConfigurationRef
type jsiiProxy_IEnvironmentBlueprintConfigurationRef struct {
	internal.Type__constructsIConstruct
}

