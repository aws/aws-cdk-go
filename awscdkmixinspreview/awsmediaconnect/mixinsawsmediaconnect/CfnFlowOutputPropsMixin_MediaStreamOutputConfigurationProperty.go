package mixinsawsmediaconnect


// The media stream that is associated with the output, and the parameters for that association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaStreamOutputConfigurationProperty := &MediaStreamOutputConfigurationProperty{
//   	DestinationConfigurations: []interface{}{
//   		&DestinationConfigurationProperty{
//   			DestinationIp: jsii.String("destinationIp"),
//   			DestinationPort: jsii.Number(123),
//   			Interface: &InterfaceProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	EncodingName: jsii.String("encodingName"),
//   	EncodingParameters: &EncodingParametersProperty{
//   		CompressionFactor: jsii.Number(123),
//   		EncoderProfile: jsii.String("encoderProfile"),
//   	},
//   	MediaStreamName: jsii.String("mediaStreamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-mediastreamoutputconfiguration.html
//
type CfnFlowOutputPropsMixin_MediaStreamOutputConfigurationProperty struct {
	// The transport parameters that are associated with each outbound media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-mediastreamoutputconfiguration.html#cfn-mediaconnect-flowoutput-mediastreamoutputconfiguration-destinationconfigurations
	//
	DestinationConfigurations interface{} `field:"optional" json:"destinationConfigurations" yaml:"destinationConfigurations"`
	// The format that was used to encode the data.
	//
	// For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video, 2110 streams, set the encoding name to raw. For video, JPEG XS streams, set the encoding name to jxsv.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-mediastreamoutputconfiguration.html#cfn-mediaconnect-flowoutput-mediastreamoutputconfiguration-encodingname
	//
	EncodingName *string `field:"optional" json:"encodingName" yaml:"encodingName"`
	// A collection of parameters that determine how MediaConnect will convert the content.
	//
	// These fields only apply to outputs on flows that have a CDI source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-mediastreamoutputconfiguration.html#cfn-mediaconnect-flowoutput-mediastreamoutputconfiguration-encodingparameters
	//
	EncodingParameters interface{} `field:"optional" json:"encodingParameters" yaml:"encodingParameters"`
	// The name of the media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-mediastreamoutputconfiguration.html#cfn-mediaconnect-flowoutput-mediastreamoutputconfiguration-mediastreamname
	//
	MediaStreamName *string `field:"optional" json:"mediaStreamName" yaml:"mediaStreamName"`
}

