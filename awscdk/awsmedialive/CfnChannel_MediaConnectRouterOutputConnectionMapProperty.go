package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaConnectRouterOutputConnectionMapProperty := &MediaConnectRouterOutputConnectionMapProperty{
//   	Pipeline0: jsii.String("pipeline0"),
//   	Pipeline1: jsii.String("pipeline1"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediaconnectrouteroutputconnectionmap.html
//
type CfnChannel_MediaConnectRouterOutputConnectionMapProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediaconnectrouteroutputconnectionmap.html#cfn-medialive-channel-mediaconnectrouteroutputconnectionmap-pipeline0
	//
	Pipeline0 *string `field:"optional" json:"pipeline0" yaml:"pipeline0"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediaconnectrouteroutputconnectionmap.html#cfn-medialive-channel-mediaconnectrouteroutputconnectionmap-pipeline1
	//
	Pipeline1 *string `field:"optional" json:"pipeline1" yaml:"pipeline1"`
}

