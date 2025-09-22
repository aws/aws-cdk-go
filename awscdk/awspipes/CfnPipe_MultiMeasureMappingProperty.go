package awspipes


// Maps multiple measures from the source event to the same Timestream for LiveAnalytics record.
//
// For more information, see [Amazon Timestream for LiveAnalytics concepts](https://docs.aws.amazon.com/timestream/latest/developerguide/concepts.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiMeasureMappingProperty := &MultiMeasureMappingProperty{
//   	MultiMeasureAttributeMappings: []interface{}{
//   		&MultiMeasureAttributeMappingProperty{
//   			MeasureValue: jsii.String("measureValue"),
//   			MeasureValueType: jsii.String("measureValueType"),
//   			MultiMeasureAttributeName: jsii.String("multiMeasureAttributeName"),
//   		},
//   	},
//   	MultiMeasureName: jsii.String("multiMeasureName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-multimeasuremapping.html
//
type CfnPipe_MultiMeasureMappingProperty struct {
	// Mappings that represent multiple source event fields mapped to measures in the same Timestream for LiveAnalytics record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-multimeasuremapping.html#cfn-pipes-pipe-multimeasuremapping-multimeasureattributemappings
	//
	MultiMeasureAttributeMappings interface{} `field:"required" json:"multiMeasureAttributeMappings" yaml:"multiMeasureAttributeMappings"`
	// The name of the multiple measurements per record (multi-measure).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-multimeasuremapping.html#cfn-pipes-pipe-multimeasuremapping-multimeasurename
	//
	MultiMeasureName *string `field:"required" json:"multiMeasureName" yaml:"multiMeasureName"`
}

