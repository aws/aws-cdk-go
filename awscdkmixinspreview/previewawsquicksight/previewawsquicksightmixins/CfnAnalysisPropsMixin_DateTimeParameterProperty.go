package previewawsquicksightmixins


// A date-time parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dateTimeParameterProperty := &DateTimeParameterProperty{
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeparameter.html
//
type CfnAnalysisPropsMixin_DateTimeParameterProperty struct {
	// A display name for the date-time parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeparameter.html#cfn-quicksight-analysis-datetimeparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The values for the date-time parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeparameter.html#cfn-quicksight-analysis-datetimeparameter-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

