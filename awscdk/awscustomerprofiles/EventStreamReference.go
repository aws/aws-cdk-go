package awscustomerprofiles


// A reference to a EventStream resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventStreamReference := &EventStreamReference{
//   	DomainName: jsii.String("domainName"),
//   	EventStreamArn: jsii.String("eventStreamArn"),
//   	EventStreamName: jsii.String("eventStreamName"),
//   }
//
type EventStreamReference struct {
	// The DomainName of the EventStream resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The ARN of the EventStream resource.
	EventStreamArn *string `field:"required" json:"eventStreamArn" yaml:"eventStreamArn"`
	// The EventStreamName of the EventStream resource.
	EventStreamName *string `field:"required" json:"eventStreamName" yaml:"eventStreamName"`
}

