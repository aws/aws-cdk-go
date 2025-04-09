package awspipes


// Maps a single source data field to a single record in the specified Timestream for LiveAnalytics table.
//
// For more information, see [Amazon Timestream for LiveAnalytics concepts](https://docs.aws.amazon.com/timestream/latest/developerguide/concepts.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleMeasureMappingProperty := &SingleMeasureMappingProperty{
//   	MeasureName: jsii.String("measureName"),
//   	MeasureValue: jsii.String("measureValue"),
//   	MeasureValueType: jsii.String("measureValueType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-singlemeasuremapping.html
//
type CfnPipe_SingleMeasureMappingProperty struct {
	// Target measure name for the measurement attribute in the Timestream table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-singlemeasuremapping.html#cfn-pipes-pipe-singlemeasuremapping-measurename
	//
	MeasureName *string `field:"required" json:"measureName" yaml:"measureName"`
	// Dynamic path of the source field to map to the measure in the record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-singlemeasuremapping.html#cfn-pipes-pipe-singlemeasuremapping-measurevalue
	//
	MeasureValue *string `field:"required" json:"measureValue" yaml:"measureValue"`
	// Data type of the source field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-singlemeasuremapping.html#cfn-pipes-pipe-singlemeasuremapping-measurevaluetype
	//
	MeasureValueType *string `field:"required" json:"measureValueType" yaml:"measureValueType"`
}

