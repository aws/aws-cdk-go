package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Site.
// Experimental.
type ISiteRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Site resource.
	// Experimental.
	SiteRef() *SiteReference
}

// The jsii proxy for ISiteRef
type jsiiProxy_ISiteRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISiteRef) SiteRef() *SiteReference {
	var returns *SiteReference
	_jsii_.Get(
		j,
		"siteRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISiteRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISiteRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

