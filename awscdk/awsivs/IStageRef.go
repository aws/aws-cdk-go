package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Stage.
// Experimental.
type IStageRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStageRef
type jsiiProxy_IStageRef struct {
	internal.Type__constructsIConstruct
}

