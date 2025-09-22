package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PortfolioShare.
// Experimental.
type IPortfolioShareRef interface {
	constructs.IConstruct
	// A reference to a PortfolioShare resource.
	// Experimental.
	PortfolioShareRef() *PortfolioShareReference
}

// The jsii proxy for IPortfolioShareRef
type jsiiProxy_IPortfolioShareRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPortfolioShareRef) PortfolioShareRef() *PortfolioShareReference {
	var returns *PortfolioShareReference
	_jsii_.Get(
		j,
		"portfolioShareRef",
		&returns,
	)
	return returns
}

