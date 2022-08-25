package awspinpoint


// Specifies demographic-based criteria, such as device platform, for the segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   demographicProperty := &demographicProperty{
//   	appVersion: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	channel: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	deviceType: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	make: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	model: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	platform: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   }
//
type CfnSegment_DemographicProperty struct {
	// The app version criteria for the segment.
	AppVersion interface{} `field:"optional" json:"appVersion" yaml:"appVersion"`
	// The channel criteria for the segment.
	Channel interface{} `field:"optional" json:"channel" yaml:"channel"`
	// The device type criteria for the segment.
	DeviceType interface{} `field:"optional" json:"deviceType" yaml:"deviceType"`
	// The device make criteria for the segment.
	Make interface{} `field:"optional" json:"make" yaml:"make"`
	// The device model criteria for the segment.
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// The device platform criteria for the segment.
	Platform interface{} `field:"optional" json:"platform" yaml:"platform"`
}

