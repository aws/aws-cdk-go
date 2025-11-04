package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PortfolioProductAssociation.
// Experimental.
type IPortfolioProductAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PortfolioProductAssociation resource.
	// Experimental.
	PortfolioProductAssociationRef() *PortfolioProductAssociationReference
}

// The jsii proxy for IPortfolioProductAssociationRef
type jsiiProxy_IPortfolioProductAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPortfolioProductAssociationRef) PortfolioProductAssociationRef() *PortfolioProductAssociationReference {
	var returns *PortfolioProductAssociationReference
	_jsii_.Get(
		j,
		"portfolioProductAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortfolioProductAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortfolioProductAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

