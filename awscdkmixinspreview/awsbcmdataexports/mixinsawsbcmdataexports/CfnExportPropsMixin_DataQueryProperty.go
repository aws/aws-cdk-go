package mixinsawsbcmdataexports


// The SQL query of column selections and row filters from the data table you want.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataQueryProperty := &DataQueryProperty{
//   	QueryStatement: jsii.String("queryStatement"),
//   	TableConfigurations: map[string]interface{}{
//   		"tableConfigurationsKey": map[string]*string{
//   			"tableConfigurationsKey": jsii.String("tableConfigurations"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-dataquery.html
//
type CfnExportPropsMixin_DataQueryProperty struct {
	// The query statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-dataquery.html#cfn-bcmdataexports-export-dataquery-querystatement
	//
	QueryStatement *string `field:"optional" json:"queryStatement" yaml:"queryStatement"`
	// The table configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-dataquery.html#cfn-bcmdataexports-export-dataquery-tableconfigurations
	//
	TableConfigurations interface{} `field:"optional" json:"tableConfigurations" yaml:"tableConfigurations"`
}

