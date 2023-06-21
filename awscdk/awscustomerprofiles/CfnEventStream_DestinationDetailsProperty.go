package awscustomerprofiles


// Details regarding the Kinesis stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationDetailsProperty := &DestinationDetailsProperty{
//   	Status: jsii.String("status"),
//   	Uri: jsii.String("uri"),
//   }
//
type CfnEventStream_DestinationDetailsProperty struct {
	// The status of enabling the Kinesis stream as a destination for export.
	Status *string `field:"required" json:"status" yaml:"status"`
	// The StreamARN of the destination to deliver profile events to.
	//
	// For example, arn:aws:kinesis:region:account-id:stream/stream-name.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

