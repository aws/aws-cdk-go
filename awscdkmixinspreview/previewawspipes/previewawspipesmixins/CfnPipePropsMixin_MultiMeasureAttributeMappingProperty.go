package previewawspipesmixins


// A mapping of a source event data field to a measure in a Timestream for LiveAnalytics record.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiMeasureAttributeMappingProperty := &MultiMeasureAttributeMappingProperty{
//   	MeasureValue: jsii.String("measureValue"),
//   	MeasureValueType: jsii.String("measureValueType"),
//   	MultiMeasureAttributeName: jsii.String("multiMeasureAttributeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-multimeasureattributemapping.html
//
type CfnPipePropsMixin_MultiMeasureAttributeMappingProperty struct {
	// Dynamic path to the measurement attribute in the source event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-multimeasureattributemapping.html#cfn-pipes-pipe-multimeasureattributemapping-measurevalue
	//
	MeasureValue *string `field:"optional" json:"measureValue" yaml:"measureValue"`
	// Data type of the measurement attribute in the source event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-multimeasureattributemapping.html#cfn-pipes-pipe-multimeasureattributemapping-measurevaluetype
	//
	MeasureValueType *string `field:"optional" json:"measureValueType" yaml:"measureValueType"`
	// Target measure name to be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-multimeasureattributemapping.html#cfn-pipes-pipe-multimeasureattributemapping-multimeasureattributename
	//
	MultiMeasureAttributeName *string `field:"optional" json:"multiMeasureAttributeName" yaml:"multiMeasureAttributeName"`
}

