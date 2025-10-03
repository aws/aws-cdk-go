package awsgroundstation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgroundstation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Config.
// Experimental.
type IConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfigRef
type jsiiProxy_IConfigRef struct {
	internal.Type__constructsIConstruct
}

