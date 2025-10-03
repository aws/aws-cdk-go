package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Portfolio.
// Experimental.
type IPortfolioRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPortfolioRef
type jsiiProxy_IPortfolioRef struct {
	internal.Type__constructsIConstruct
}

