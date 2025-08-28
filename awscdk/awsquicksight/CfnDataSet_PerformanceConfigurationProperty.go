package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   performanceConfigurationProperty := &PerformanceConfigurationProperty{
//   	UniqueKeys: []interface{}{
//   		&UniqueKeyProperty{
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-performanceconfiguration.html
//
type CfnDataSet_PerformanceConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-performanceconfiguration.html#cfn-quicksight-dataset-performanceconfiguration-uniquekeys
	//
	UniqueKeys interface{} `field:"optional" json:"uniqueKeys" yaml:"uniqueKeys"`
}

