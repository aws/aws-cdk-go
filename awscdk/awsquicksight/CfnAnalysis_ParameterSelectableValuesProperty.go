package awsquicksight


// A list of selectable values that are used in a control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterSelectableValuesProperty := &ParameterSelectableValuesProperty{
//   	LinkToDataSetColumn: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterselectablevalues.html
//
type CfnAnalysis_ParameterSelectableValuesProperty struct {
	// The column identifier that fetches values from the data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterselectablevalues.html#cfn-quicksight-analysis-parameterselectablevalues-linktodatasetcolumn
	//
	LinkToDataSetColumn interface{} `field:"optional" json:"linkToDataSetColumn" yaml:"linkToDataSetColumn"`
	// The values that are used in `ParameterSelectableValues` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-parameterselectablevalues.html#cfn-quicksight-analysis-parameterselectablevalues-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

