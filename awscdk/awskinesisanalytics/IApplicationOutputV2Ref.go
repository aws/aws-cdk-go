package awskinesisanalytics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationOutput.
// Experimental.
type IApplicationOutputV2Ref interface {
	constructs.IConstruct
}

// The jsii proxy for IApplicationOutputV2Ref
type jsiiProxy_IApplicationOutputV2Ref struct {
	internal.Type__constructsIConstruct
}

