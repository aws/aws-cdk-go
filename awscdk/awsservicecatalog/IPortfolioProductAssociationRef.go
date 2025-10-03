package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PortfolioProductAssociation.
// Experimental.
type IPortfolioProductAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPortfolioProductAssociationRef
type jsiiProxy_IPortfolioProductAssociationRef struct {
	internal.Type__constructsIConstruct
}

