package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SiteToSiteVpnAttachment.
// Experimental.
type ISiteToSiteVpnAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SiteToSiteVpnAttachment resource.
	// Experimental.
	SiteToSiteVpnAttachmentRef() *SiteToSiteVpnAttachmentReference
}

// The jsii proxy for ISiteToSiteVpnAttachmentRef
type jsiiProxy_ISiteToSiteVpnAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISiteToSiteVpnAttachmentRef) SiteToSiteVpnAttachmentRef() *SiteToSiteVpnAttachmentReference {
	var returns *SiteToSiteVpnAttachmentReference
	_jsii_.Get(
		j,
		"siteToSiteVpnAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISiteToSiteVpnAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISiteToSiteVpnAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

