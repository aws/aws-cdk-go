package previewawskinesisanalyticsmixins


// Provides additional mapping information when JSON is the record format on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jSONMappingParametersProperty := &JSONMappingParametersProperty{
//   	RecordRowPath: jsii.String("recordRowPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-jsonmappingparameters.html
//
type CfnApplicationReferenceDataSourcePropsMixin_JSONMappingParametersProperty struct {
	// Path to the top-level parent that contains the records.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-jsonmappingparameters.html#cfn-kinesisanalytics-applicationreferencedatasource-jsonmappingparameters-recordrowpath
	//
	RecordRowPath *string `field:"optional" json:"recordRowPath" yaml:"recordRowPath"`
}

