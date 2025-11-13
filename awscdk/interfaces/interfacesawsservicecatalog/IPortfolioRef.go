package interfacesawsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Portfolio.
// Experimental.
type IPortfolioRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Portfolio resource.
	// Experimental.
	PortfolioRef() *PortfolioReference
}

// The jsii proxy for IPortfolioRef
type jsiiProxy_IPortfolioRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IPortfolioRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortfolioRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

