package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Portfolio.
// Experimental.
type IPortfolioRef interface {
	constructs.IConstruct
	// A reference to a Portfolio resource.
	// Experimental.
	PortfolioRef() *PortfolioReference
}

// The jsii proxy for IPortfolioRef
type jsiiProxy_IPortfolioRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPortfolioRef) PortfolioRef() *PortfolioReference {
	var returns *PortfolioReference
	_jsii_.Get(
		j,
		"portfolioRef",
		&returns,
	)
	return returns
}

