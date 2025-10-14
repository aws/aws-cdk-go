package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AcceptedPortfolioShare.
// Experimental.
type IAcceptedPortfolioShareRef interface {
	constructs.IConstruct
	// A reference to a AcceptedPortfolioShare resource.
	// Experimental.
	AcceptedPortfolioShareRef() *AcceptedPortfolioShareReference
}

// The jsii proxy for IAcceptedPortfolioShareRef
type jsiiProxy_IAcceptedPortfolioShareRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAcceptedPortfolioShareRef) AcceptedPortfolioShareRef() *AcceptedPortfolioShareReference {
	var returns *AcceptedPortfolioShareReference
	_jsii_.Get(
		j,
		"acceptedPortfolioShareRef",
		&returns,
	)
	return returns
}

