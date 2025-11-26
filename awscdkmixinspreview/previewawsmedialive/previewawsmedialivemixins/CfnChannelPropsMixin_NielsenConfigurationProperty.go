package previewawsmedialivemixins


// The settings to configure Nielsen watermarks.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   nielsenConfigurationProperty := &NielsenConfigurationProperty{
//   	DistributorId: jsii.String("distributorId"),
//   	NielsenPcmToId3Tagging: jsii.String("nielsenPcmToId3Tagging"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsenconfiguration.html
//
type CfnChannelPropsMixin_NielsenConfigurationProperty struct {
	// Enter the Distributor ID assigned to your organization by Nielsen.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsenconfiguration.html#cfn-medialive-channel-nielsenconfiguration-distributorid
	//
	DistributorId *string `field:"optional" json:"distributorId" yaml:"distributorId"`
	// Enables Nielsen PCM to ID3 tagging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsenconfiguration.html#cfn-medialive-channel-nielsenconfiguration-nielsenpcmtoid3tagging
	//
	NielsenPcmToId3Tagging *string `field:"optional" json:"nielsenPcmToId3Tagging" yaml:"nielsenPcmToId3Tagging"`
}

