package interfacesawsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PortfolioShare.
// Experimental.
type IPortfolioShareRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PortfolioShare resource.
	// Experimental.
	PortfolioShareRef() *PortfolioShareReference
}

// The jsii proxy for IPortfolioShareRef
type jsiiProxy_IPortfolioShareRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPortfolioShareRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IPortfolioShareRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortfolioShareRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

