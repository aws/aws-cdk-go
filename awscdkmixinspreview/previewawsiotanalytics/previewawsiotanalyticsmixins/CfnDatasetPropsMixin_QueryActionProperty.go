package previewawsiotanalyticsmixins


// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   queryActionProperty := &QueryActionProperty{
//   	Filters: []interface{}{
//   		&FilterProperty{
//   			DeltaTime: &DeltaTimeProperty{
//   				OffsetSeconds: jsii.Number(123),
//   				TimeExpression: jsii.String("timeExpression"),
//   			},
//   		},
//   	},
//   	SqlQuery: jsii.String("sqlQuery"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html
//
type CfnDatasetPropsMixin_QueryActionProperty struct {
	// Pre-filters applied to message data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html#cfn-iotanalytics-dataset-queryaction-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html#cfn-iotanalytics-dataset-queryaction-sqlquery
	//
	SqlQuery *string `field:"optional" json:"sqlQuery" yaml:"sqlQuery"`
}

