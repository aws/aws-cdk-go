package awsaiops

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsaiops/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InvestigationGroup.
// Experimental.
type IInvestigationGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInvestigationGroupRef
type jsiiProxy_IInvestigationGroupRef struct {
	internal.Type__constructsIConstruct
}

