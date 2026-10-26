package awsmedialive


// A MediaConnect Router destination associated with a MediaLive Input.
//
// RouterOutputArn is server-assigned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   routerDestinationSettingsProperty := &RouterDestinationSettingsProperty{
//   	AvailabilityZoneName: jsii.String("availabilityZoneName"),
//   	RouterOutputArn: jsii.String("routerOutputArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routerdestinationsettings.html
//
type CfnInputPropsMixin_RouterDestinationSettingsProperty struct {
	// The Availability Zone (AZ) name this destination is created in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routerdestinationsettings.html#cfn-medialive-input-routerdestinationsettings-availabilityzonename
	//
	AvailabilityZoneName *string `field:"optional" json:"availabilityZoneName" yaml:"availabilityZoneName"`
	// The ARN of the MediaConnect Router output currently connected to this input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routerdestinationsettings.html#cfn-medialive-input-routerdestinationsettings-routeroutputarn
	//
	RouterOutputArn *string `field:"optional" json:"routerOutputArn" yaml:"routerOutputArn"`
}

