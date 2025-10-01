package awscloudfront


// A reference to a StreamingDistribution resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamingDistributionReference := &StreamingDistributionReference{
//   	StreamingDistributionId: jsii.String("streamingDistributionId"),
//   }
//
type StreamingDistributionReference struct {
	// The Id of the StreamingDistribution resource.
	StreamingDistributionId *string `field:"required" json:"streamingDistributionId" yaml:"streamingDistributionId"`
}

