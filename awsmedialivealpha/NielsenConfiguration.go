package awsmedialivealpha


// Nielsen watermark configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var nielsenPcmToId3TaggingState NielsenPcmToId3TaggingState
//
//   nielsenConfiguration := &NielsenConfiguration{
//   	DistributorId: jsii.String("distributorId"),
//   	NielsenPcmToId3Tagging: nielsenPcmToId3TaggingState,
//   }
//
// Experimental.
type NielsenConfiguration struct {
	// The Distributor ID assigned to your organization by Nielsen.
	// Default: - no distributor ID.
	//
	// Experimental.
	DistributorId *string `field:"optional" json:"distributorId" yaml:"distributorId"`
	// Whether to enable Nielsen PCM to ID3 tagging.
	// Default: - service default.
	//
	// Experimental.
	NielsenPcmToId3Tagging NielsenPcmToId3TaggingState `field:"optional" json:"nielsenPcmToId3Tagging" yaml:"nielsenPcmToId3Tagging"`
}

