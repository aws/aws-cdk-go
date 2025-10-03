package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnycastIpList.
// Experimental.
type IAnycastIpListRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAnycastIpListRef
type jsiiProxy_IAnycastIpListRef struct {
	internal.Type__constructsIConstruct
}

