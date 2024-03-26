package awskinesisanalytics


// For a SQL-based Kinesis Data Analytics application, provides additional mapping information when JSON is the record format on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jSONMappingParametersProperty := &JSONMappingParametersProperty{
//   	RecordRowPath: jsii.String("recordRowPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-jsonmappingparameters.html
//
type CfnApplicationReferenceDataSourceV2_JSONMappingParametersProperty struct {
	// The path to the top-level parent that contains the records.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-jsonmappingparameters.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-jsonmappingparameters-recordrowpath
	//
	RecordRowPath *string `field:"required" json:"recordRowPath" yaml:"recordRowPath"`
}

