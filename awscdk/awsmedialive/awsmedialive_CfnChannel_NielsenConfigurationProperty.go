package awsmedialive


// The settings to configure Nielsen watermarks.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nielsenConfigurationProperty := &nielsenConfigurationProperty{
//   	distributorId: jsii.String("distributorId"),
//   	nielsenPcmToId3Tagging: jsii.String("nielsenPcmToId3Tagging"),
//   }
//
type CfnChannel_NielsenConfigurationProperty struct {
	// Enter the Distributor ID assigned to your organization by Nielsen.
	DistributorId *string `field:"optional" json:"distributorId" yaml:"distributorId"`
	// Enables Nielsen PCM to ID3 tagging.
	NielsenPcmToId3Tagging *string `field:"optional" json:"nielsenPcmToId3Tagging" yaml:"nielsenPcmToId3Tagging"`
}

