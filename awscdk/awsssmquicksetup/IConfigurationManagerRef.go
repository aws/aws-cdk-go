package awsssmquicksetup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmquicksetup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationManager.
// Experimental.
type IConfigurationManagerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfigurationManagerRef
type jsiiProxy_IConfigurationManagerRef struct {
	internal.Type__constructsIConstruct
}

