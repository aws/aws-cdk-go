package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AcceptedPortfolioShare.
// Experimental.
type IAcceptedPortfolioShareRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AcceptedPortfolioShare resource.
	// Experimental.
	AcceptedPortfolioShareRef() *AcceptedPortfolioShareReference
}

// The jsii proxy for IAcceptedPortfolioShareRef
type jsiiProxy_IAcceptedPortfolioShareRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IAcceptedPortfolioShareRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAcceptedPortfolioShareRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

