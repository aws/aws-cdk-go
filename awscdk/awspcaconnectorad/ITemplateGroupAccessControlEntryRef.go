package awspcaconnectorad

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspcaconnectorad/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TemplateGroupAccessControlEntry.
// Experimental.
type ITemplateGroupAccessControlEntryRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITemplateGroupAccessControlEntryRef
type jsiiProxy_ITemplateGroupAccessControlEntryRef struct {
	internal.Type__constructsIConstruct
}

