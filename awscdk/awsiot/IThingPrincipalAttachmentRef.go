package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThingPrincipalAttachment.
// Experimental.
type IThingPrincipalAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IThingPrincipalAttachmentRef
type jsiiProxy_IThingPrincipalAttachmentRef struct {
	internal.Type__constructsIConstruct
}

