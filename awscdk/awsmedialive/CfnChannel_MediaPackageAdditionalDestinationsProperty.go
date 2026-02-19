package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaPackageAdditionalDestinationsProperty := &MediaPackageAdditionalDestinationsProperty{
//   	Destination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackageadditionaldestinations.html
//
type CfnChannel_MediaPackageAdditionalDestinationsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediapackageadditionaldestinations.html#cfn-medialive-channel-mediapackageadditionaldestinations-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
}

