package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   streamProperty := &StreamProperty{
//   	FileId: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-stream.html
//
type CfnOTAUpdatePropsMixin_StreamProperty struct {
	// The ID of a file associated with a stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-stream.html#cfn-iot-otaupdate-stream-fileid
	//
	FileId *float64 `field:"optional" json:"fileId" yaml:"fileId"`
	// The stream ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-stream.html#cfn-iot-otaupdate-stream-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

