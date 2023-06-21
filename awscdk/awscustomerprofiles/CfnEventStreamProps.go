package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEventStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventStreamProps := &CfnEventStreamProps{
//   	DomainName: jsii.String("domainName"),
//   	EventStreamName: jsii.String("eventStreamName"),
//   	Uri: jsii.String("uri"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEventStreamProps struct {
	// The unique name of the domain.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The name of the event stream.
	EventStreamName *string `field:"required" json:"eventStreamName" yaml:"eventStreamName"`
	// The StreamARN of the destination to deliver profile events to.
	//
	// For example, arn:aws:kinesis:region:account-id:stream/stream-name.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

