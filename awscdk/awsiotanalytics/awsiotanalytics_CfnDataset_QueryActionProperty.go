package awsiotanalytics


// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryActionProperty := &queryActionProperty{
//   	sqlQuery: jsii.String("sqlQuery"),
//
//   	// the properties below are optional
//   	filters: []interface{}{
//   		&filterProperty{
//   			deltaTime: &deltaTimeProperty{
//   				offsetSeconds: jsii.Number(123),
//   				timeExpression: jsii.String("timeExpression"),
//   			},
//   		},
//   	},
//   }
//
type CfnDataset_QueryActionProperty struct {
	// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
	SqlQuery *string `field:"required" json:"sqlQuery" yaml:"sqlQuery"`
	// Pre-filters applied to message data.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

