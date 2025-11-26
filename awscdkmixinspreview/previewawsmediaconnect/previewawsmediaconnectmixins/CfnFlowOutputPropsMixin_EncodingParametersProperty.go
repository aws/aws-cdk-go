package previewawsmediaconnectmixins


// A collection of parameters that determine how MediaConnect will convert the content.
//
// These fields only apply to outputs on flows that have a CDI source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encodingParametersProperty := &EncodingParametersProperty{
//   	CompressionFactor: jsii.Number(123),
//   	EncoderProfile: jsii.String("encoderProfile"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-encodingparameters.html
//
type CfnFlowOutputPropsMixin_EncodingParametersProperty struct {
	// A value that is used to calculate compression for an output.
	//
	// The bitrate of the output is calculated as follows: Output bitrate = (1 / compressionFactor) * (source bitrate) This property only applies to outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol. Valid values are floating point numbers in the range of 3.0 to 10.0, inclusive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-encodingparameters.html#cfn-mediaconnect-flowoutput-encodingparameters-compressionfactor
	//
	CompressionFactor *float64 `field:"optional" json:"compressionFactor" yaml:"compressionFactor"`
	// A setting on the encoder that drives compression settings.
	//
	// This property only applies to video media streams associated with outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-encodingparameters.html#cfn-mediaconnect-flowoutput-encodingparameters-encoderprofile
	//
	EncoderProfile *string `field:"optional" json:"encoderProfile" yaml:"encoderProfile"`
}

