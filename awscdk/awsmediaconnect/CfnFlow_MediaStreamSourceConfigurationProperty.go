package awsmediaconnect


// The media stream that is associated with the source, and the parameters for that association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaStreamSourceConfigurationProperty := &MediaStreamSourceConfigurationProperty{
//   	EncodingName: jsii.String("encodingName"),
//   	MediaStreamName: jsii.String("mediaStreamName"),
//
//   	// the properties below are optional
//   	InputConfigurations: []interface{}{
//   		&InputConfigurationProperty{
//   			InputPort: jsii.Number(123),
//   			Interface: &InterfaceProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastreamsourceconfiguration.html
//
type CfnFlow_MediaStreamSourceConfigurationProperty struct {
	// The format that was used to encode the data.
	//
	// For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video, 2110 streams, set the encoding name to raw. For video, JPEG XS streams, set the encoding name to jxsv.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastreamsourceconfiguration.html#cfn-mediaconnect-flow-mediastreamsourceconfiguration-encodingname
	//
	EncodingName *string `field:"required" json:"encodingName" yaml:"encodingName"`
	// A name that helps you distinguish one media stream from another.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastreamsourceconfiguration.html#cfn-mediaconnect-flow-mediastreamsourceconfiguration-mediastreamname
	//
	MediaStreamName *string `field:"required" json:"mediaStreamName" yaml:"mediaStreamName"`
	// The media streams that you want to associate with the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastreamsourceconfiguration.html#cfn-mediaconnect-flow-mediastreamsourceconfiguration-inputconfigurations
	//
	InputConfigurations interface{} `field:"optional" json:"inputConfigurations" yaml:"inputConfigurations"`
}

