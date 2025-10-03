package awsredshiftserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshiftserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workgroup.
// Experimental.
type IWorkgroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWorkgroupRef
type jsiiProxy_IWorkgroupRef struct {
	internal.Type__constructsIConstruct
}

