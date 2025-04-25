package awsquicksight


// The data path options for the pivot table field options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableDataPathOptionProperty := &PivotTableDataPathOptionProperty{
//   	DataPathList: []interface{}{
//   		&DataPathValueProperty{
//   			DataPathType: &DataPathTypeProperty{
//   				PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//   			FieldValue: jsii.String("fieldValue"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Width: jsii.String("width"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottabledatapathoption.html
//
type CfnTemplate_PivotTableDataPathOptionProperty struct {
	// The list of data path values for the data path options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottabledatapathoption.html#cfn-quicksight-template-pivottabledatapathoption-datapathlist
	//
	DataPathList interface{} `field:"required" json:"dataPathList" yaml:"dataPathList"`
	// The width of the data path option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottabledatapathoption.html#cfn-quicksight-template-pivottabledatapathoption-width
	//
	Width *string `field:"optional" json:"width" yaml:"width"`
}

