package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaConnectRouterGroupSettingsProperty := &MediaConnectRouterGroupSettingsProperty{
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediaconnectroutergroupsettings.html
//
type CfnChannel_MediaConnectRouterGroupSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediaconnectroutergroupsettings.html#cfn-medialive-channel-mediaconnectroutergroupsettings-availabilityzones
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
}

