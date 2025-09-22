package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SiteToSiteVpnAttachment.
// Experimental.
type ISiteToSiteVpnAttachmentRef interface {
	constructs.IConstruct
	// A reference to a SiteToSiteVpnAttachment resource.
	// Experimental.
	SiteToSiteVpnAttachmentRef() *SiteToSiteVpnAttachmentReference
}

// The jsii proxy for ISiteToSiteVpnAttachmentRef
type jsiiProxy_ISiteToSiteVpnAttachmentRef struct {
	internal.Type__constructsIConstruct
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

