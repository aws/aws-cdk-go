package awsmsk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Configuration.
// Experimental.
type IConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfigurationRef
type jsiiProxy_IConfigurationRef struct {
	internal.Type__constructsIConstruct
}

