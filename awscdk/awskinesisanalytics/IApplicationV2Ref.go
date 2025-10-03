package awskinesisanalytics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Application.
// Experimental.
type IApplicationV2Ref interface {
	constructs.IConstruct
}

// The jsii proxy for IApplicationV2Ref
type jsiiProxy_IApplicationV2Ref struct {
	internal.Type__constructsIConstruct
}

