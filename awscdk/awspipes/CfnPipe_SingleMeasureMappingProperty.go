package awspipes


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-singlemeasuremapping.html#cfn-pipes-pipe-singlemeasuremapping-measurename
	//
	MeasureName *string `field:"required" json:"measureName" yaml:"measureName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-singlemeasuremapping.html#cfn-pipes-pipe-singlemeasuremapping-measurevalue
	//
	MeasureValue *string `field:"required" json:"measureValue" yaml:"measureValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-singlemeasuremapping.html#cfn-pipes-pipe-singlemeasuremapping-measurevaluetype
	//
	MeasureValueType *string `field:"required" json:"measureValueType" yaml:"measureValueType"`
}

