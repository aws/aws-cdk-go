package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComputationModel.
// Experimental.
type IComputationModelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IComputationModelRef
type jsiiProxy_IComputationModelRef struct {
	internal.Type__constructsIConstruct
}

