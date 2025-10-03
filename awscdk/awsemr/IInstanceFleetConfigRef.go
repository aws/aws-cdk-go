package awsemr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceFleetConfig.
// Experimental.
type IInstanceFleetConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInstanceFleetConfigRef
type jsiiProxy_IInstanceFleetConfigRef struct {
	internal.Type__constructsIConstruct
}

