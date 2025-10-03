package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SiteToSiteVpnAttachment.
// Experimental.
type ISiteToSiteVpnAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISiteToSiteVpnAttachmentRef
type jsiiProxy_ISiteToSiteVpnAttachmentRef struct {
	internal.Type__constructsIConstruct
}

