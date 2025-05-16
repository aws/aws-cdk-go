package awspinpoint


// Specifies demographic-based criteria, such as device platform, for the segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   demographicProperty := &DemographicProperty{
//   	AppVersion: &SetDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Channel: &SetDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	DeviceType: &SetDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Make: &SetDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Model: &SetDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Platform: &SetDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-demographic.html
//
type CfnSegment_DemographicProperty struct {
	// The app version criteria for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-demographic.html#cfn-pinpoint-segment-demographic-appversion
	//
	AppVersion interface{} `field:"optional" json:"appVersion" yaml:"appVersion"`
	// The channel criteria for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-demographic.html#cfn-pinpoint-segment-demographic-channel
	//
	Channel interface{} `field:"optional" json:"channel" yaml:"channel"`
	// The device type criteria for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-demographic.html#cfn-pinpoint-segment-demographic-devicetype
	//
	DeviceType interface{} `field:"optional" json:"deviceType" yaml:"deviceType"`
	// The device make criteria for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-demographic.html#cfn-pinpoint-segment-demographic-make
	//
	Make interface{} `field:"optional" json:"make" yaml:"make"`
	// The device model criteria for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-demographic.html#cfn-pinpoint-segment-demographic-model
	//
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// The device platform criteria for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-demographic.html#cfn-pinpoint-segment-demographic-platform
	//
	Platform interface{} `field:"optional" json:"platform" yaml:"platform"`
}

