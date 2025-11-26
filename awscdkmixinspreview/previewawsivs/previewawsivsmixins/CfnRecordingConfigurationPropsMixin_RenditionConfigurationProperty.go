package previewawsivsmixins


// The RenditionConfiguration property type describes which renditions should be recorded for a stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   renditionConfigurationProperty := &RenditionConfigurationProperty{
//   	Renditions: []*string{
//   		jsii.String("renditions"),
//   	},
//   	RenditionSelection: jsii.String("renditionSelection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-renditionconfiguration.html
//
type CfnRecordingConfigurationPropsMixin_RenditionConfigurationProperty struct {
	// A list of which renditions are recorded for a stream, if `renditionSelection` is `CUSTOM` ;
	//
	// otherwise, this field is irrelevant. The selected renditions are recorded if they are available during the stream. If a selected rendition is unavailable, the best available rendition is recorded. For details on the resolution dimensions of each rendition, see [Auto-Record to Amazon S3](https://docs.aws.amazon.com//ivs/latest/LowLatencyUserGuide/record-to-s3.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-renditionconfiguration.html#cfn-ivs-recordingconfiguration-renditionconfiguration-renditions
	//
	Renditions *[]*string `field:"optional" json:"renditions" yaml:"renditions"`
	// The set of renditions are recorded for a stream.
	//
	// For `BASIC` channels, the `CUSTOM` value has no effect. If `CUSTOM` is specified, a set of renditions can be specified in the `renditions` field. Default: `ALL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-renditionconfiguration.html#cfn-ivs-recordingconfiguration-renditionconfiguration-renditionselection
	//
	// Default: - "ALL".
	//
	RenditionSelection *string `field:"optional" json:"renditionSelection" yaml:"renditionSelection"`
}

