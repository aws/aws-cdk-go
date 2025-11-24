package mixinsawsmediaconnect


// The media stream that is associated with the source, and the parameters for that association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaStreamSourceConfigurationProperty := &MediaStreamSourceConfigurationProperty{
//   	EncodingName: jsii.String("encodingName"),
//   	InputConfigurations: []interface{}{
//   		&InputConfigurationProperty{
//   			InputPort: jsii.Number(123),
//   			Interface: &InterfaceProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	MediaStreamName: jsii.String("mediaStreamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastreamsourceconfiguration.html
//
type CfnFlowPropsMixin_MediaStreamSourceConfigurationProperty struct {
	// The format that was used to encode the data.
	//
	// For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video, 2110 streams, set the encoding name to raw. For video, JPEG XS streams, set the encoding name to jxsv.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastreamsourceconfiguration.html#cfn-mediaconnect-flow-mediastreamsourceconfiguration-encodingname
	//
	EncodingName *string `field:"optional" json:"encodingName" yaml:"encodingName"`
	// The media streams that you want to associate with the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastreamsourceconfiguration.html#cfn-mediaconnect-flow-mediastreamsourceconfiguration-inputconfigurations
	//
	InputConfigurations interface{} `field:"optional" json:"inputConfigurations" yaml:"inputConfigurations"`
	// A name that helps you distinguish one media stream from another.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastreamsourceconfiguration.html#cfn-mediaconnect-flow-mediastreamsourceconfiguration-mediastreamname
	//
	MediaStreamName *string `field:"optional" json:"mediaStreamName" yaml:"mediaStreamName"`
}

