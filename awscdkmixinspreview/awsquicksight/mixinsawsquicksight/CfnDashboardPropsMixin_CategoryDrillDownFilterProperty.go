package mixinsawsquicksight


// The category drill down filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   categoryDrillDownFilterProperty := &CategoryDrillDownFilterProperty{
//   	CategoryValues: []*string{
//   		jsii.String("categoryValues"),
//   	},
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-categorydrilldownfilter.html
//
type CfnDashboardPropsMixin_CategoryDrillDownFilterProperty struct {
	// A list of the string inputs that are the values of the category drill down filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-categorydrilldownfilter.html#cfn-quicksight-dashboard-categorydrilldownfilter-categoryvalues
	//
	CategoryValues *[]*string `field:"optional" json:"categoryValues" yaml:"categoryValues"`
	// The column that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-categorydrilldownfilter.html#cfn-quicksight-dashboard-categorydrilldownfilter-column
	//
	Column interface{} `field:"optional" json:"column" yaml:"column"`
}

