package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamSAMPTProperty := &streamSAMPTProperty{
//   	streamName: jsii.String("streamName"),
//   }
//
type CfnFunction_StreamSAMPTProperty struct {
	// `CfnFunction.StreamSAMPTProperty.StreamName`.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
}

