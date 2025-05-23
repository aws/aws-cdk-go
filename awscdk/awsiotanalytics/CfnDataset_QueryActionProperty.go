package awsiotanalytics


// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryActionProperty := &QueryActionProperty{
//   	SqlQuery: jsii.String("sqlQuery"),
//
//   	// the properties below are optional
//   	Filters: []interface{}{
//   		&FilterProperty{
//   			DeltaTime: &DeltaTimeProperty{
//   				OffsetSeconds: jsii.Number(123),
//   				TimeExpression: jsii.String("timeExpression"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html
//
type CfnDataset_QueryActionProperty struct {
	// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html#cfn-iotanalytics-dataset-queryaction-sqlquery
	//
	SqlQuery *string `field:"required" json:"sqlQuery" yaml:"sqlQuery"`
	// Pre-filters applied to message data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html#cfn-iotanalytics-dataset-queryaction-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

