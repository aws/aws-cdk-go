package awsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VodSource.
// Experimental.
type IVodSourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVodSourceRef
type jsiiProxy_IVodSourceRef struct {
	internal.Type__constructsIConstruct
}

