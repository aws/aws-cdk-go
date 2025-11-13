package interfacesawsqldb


// A reference to a Stream resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamReference := &StreamReference{
//   	LedgerName: jsii.String("ledgerName"),
//   	StreamArn: jsii.String("streamArn"),
//   	StreamId: jsii.String("streamId"),
//   }
//
type StreamReference struct {
	// The LedgerName of the Stream resource.
	LedgerName *string `field:"required" json:"ledgerName" yaml:"ledgerName"`
	// The ARN of the Stream resource.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
	// The Id of the Stream resource.
	StreamId *string `field:"required" json:"streamId" yaml:"streamId"`
}

