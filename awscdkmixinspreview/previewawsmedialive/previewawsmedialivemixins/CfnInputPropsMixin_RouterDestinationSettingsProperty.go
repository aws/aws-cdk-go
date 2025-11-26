package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   routerDestinationSettingsProperty := &RouterDestinationSettingsProperty{
//   	AvailabilityZoneName: jsii.String("availabilityZoneName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routerdestinationsettings.html
//
type CfnInputPropsMixin_RouterDestinationSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routerdestinationsettings.html#cfn-medialive-input-routerdestinationsettings-availabilityzonename
	//
	AvailabilityZoneName *string `field:"optional" json:"availabilityZoneName" yaml:"availabilityZoneName"`
}

