package awsmedialive


// The input specification for this channel.
//
// It specifies the key characteristics of CDI inputs for this channel, when those characteristics are different from other inputs.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cdiInputSpecificationProperty := &cdiInputSpecificationProperty{
//   	resolution: jsii.String("resolution"),
//   }
//
type CfnChannel_CdiInputSpecificationProperty struct {
	// Maximum CDI input resolution.
	Resolution *string `field:"optional" json:"resolution" yaml:"resolution"`
}

