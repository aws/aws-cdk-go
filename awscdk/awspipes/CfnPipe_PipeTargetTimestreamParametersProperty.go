package awspipes


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-dimensionmappings
	//
	DimensionMappings interface{} `field:"required" json:"dimensionMappings" yaml:"dimensionMappings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-timevalue
	//
	TimeValue *string `field:"required" json:"timeValue" yaml:"timeValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-versionvalue
	//
	VersionValue *string `field:"required" json:"versionValue" yaml:"versionValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-epochtimeunit
	//
	EpochTimeUnit *string `field:"optional" json:"epochTimeUnit" yaml:"epochTimeUnit"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-multimeasuremappings
	//
	MultiMeasureMappings interface{} `field:"optional" json:"multiMeasureMappings" yaml:"multiMeasureMappings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-singlemeasuremappings
	//
	SingleMeasureMappings interface{} `field:"optional" json:"singleMeasureMappings" yaml:"singleMeasureMappings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-timefieldtype
	//
	TimeFieldType *string `field:"optional" json:"timeFieldType" yaml:"timeFieldType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargettimestreamparameters.html#cfn-pipes-pipe-pipetargettimestreamparameters-timestampformat
	//
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
}

