package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeliveryStream.
// Experimental.
type IDeliveryStreamRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeliveryStreamRef
type jsiiProxy_IDeliveryStreamRef struct {
	internal.Type__constructsIConstruct
}

