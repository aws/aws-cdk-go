package awspipes


// The parameters for using a Timestream for LiveAnalytics table as a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetTimestreamParametersProperty := &PipeTargetTimestreamParametersProperty{
//   	DimensionMappings: []interface{}{
//   		&DimensionMappingProperty{
//   			DimensionName: jsii.String("dimensionName"),
//   			DimensionValue: jsii.String("dimensionValue"),
//   			DimensionValueType: jsii.String("dimensionValueType"),
//   		},
//   	},
//   	TimeValue: jsii.String("timeValue"),
//   	VersionValue: jsii.String("versionValue"),
//
//   	// the properties below are optional
//   	EpochTimeUnit: jsii.String("epochTimeUnit"),
//   	MultiMeasureMappings: []interface{}{
//   		&MultiMeasureMappingProperty{
//   			MultiMeasureAttributeMappings: []interface{}{
//   				&MultiMeasureAttributeMappingProperty{
//   					MeasureValue: jsii.String("measureValue"),
//   					MeasureValueType: jsii.String("measureValueType"),
//   					MultiMeasureAttributeName: jsii.String("multiMeasureAttributeName"),
//   				},
//   			},
//   			MultiMeasureName: jsii.String("multiMeasureName"),
//   		},
//   	},
//   	SingleMeasureMappings: []interface{}{
//   		&SingleMeasureMappingProperty{
//   			MeasureName: jsii.String("measureName"),
//   			MeasureValue: jsii.String("measureValue"),
//   			MeasureValueType: jsii.String("measureValueType"),
//   		},
//   	},
//   	TimeFieldType: jsii.String("timeFieldType"),
//   	TimestampFormat: jsii.String("timestampFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html
//
type CfnPipe_PipeTargetTimestreamParametersProperty struct {
	// Map source data to dimensions in the target Timestream for LiveAnalytics table.
	//
	// For more information, see [Amazon Timestream for LiveAnalytics concepts](https://docs.aws.amazon.com/timestream/latest/developerguide/concepts.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-dimensionmappings
	//
	DimensionMappings interface{} `field:"required" json:"dimensionMappings" yaml:"dimensionMappings"`
	// Dynamic path to the source data field that represents the time value for your data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-timevalue
	//
	TimeValue *string `field:"required" json:"timeValue" yaml:"timeValue"`
	// 64 bit version value or source data field that represents the version value for your data.
	//
	// Write requests with a higher version number will update the existing measure values of the record and version. In cases where the measure value is the same, the version will still be updated.
	//
	// Default value is 1.
	//
	// Timestream for LiveAnalytics does not support updating partial measure values in a record.
	//
	// Write requests for duplicate data with a higher version number will update the existing measure value and version. In cases where the measure value is the same, `Version` will still be updated. Default value is `1` .
	//
	// > `Version` must be `1` or greater, or you will receive a `ValidationException` error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-versionvalue
	//
	VersionValue *string `field:"required" json:"versionValue" yaml:"versionValue"`
	// The granularity of the time units used. Default is `MILLISECONDS` .
	//
	// Required if `TimeFieldType` is specified as `EPOCH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-epochtimeunit
	//
	EpochTimeUnit *string `field:"optional" json:"epochTimeUnit" yaml:"epochTimeUnit"`
	// Maps multiple measures from the source event to the same record in the specified Timestream for LiveAnalytics table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-multimeasuremappings
	//
	MultiMeasureMappings interface{} `field:"optional" json:"multiMeasureMappings" yaml:"multiMeasureMappings"`
	// Mappings of single source data fields to individual records in the specified Timestream for LiveAnalytics table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-singlemeasuremappings
	//
	SingleMeasureMappings interface{} `field:"optional" json:"singleMeasureMappings" yaml:"singleMeasureMappings"`
	// The type of time value used.
	//
	// The default is `EPOCH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-timefieldtype
	//
	TimeFieldType *string `field:"optional" json:"timeFieldType" yaml:"timeFieldType"`
	// How to format the timestamps. For example, `yyyy-MM-dd'T'HH:mm:ss'Z'` .
	//
	// Required if `TimeFieldType` is specified as `TIMESTAMP_FORMAT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-timestampformat
	//
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
}

