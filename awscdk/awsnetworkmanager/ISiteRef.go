package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Site.
// Experimental.
type ISiteRef interface {
	constructs.IConstruct
	// A reference to a Site resource.
	// Experimental.
	SiteRef() *SiteReference
}

// The jsii proxy for ISiteRef
type jsiiProxy_ISiteRef struct {
	internal.Type__constructsIConstruct
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

