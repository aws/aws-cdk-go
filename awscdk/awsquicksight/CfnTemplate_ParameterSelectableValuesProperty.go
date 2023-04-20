package awsquicksight


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
type CfnTemplate_ParameterSelectableValuesProperty struct {
	// `CfnTemplate.ParameterSelectableValuesProperty.LinkToDataSetColumn`.
	LinkToDataSetColumn interface{} `field:"optional" json:"linkToDataSetColumn" yaml:"linkToDataSetColumn"`
	// `CfnTemplate.ParameterSelectableValuesProperty.Values`.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

