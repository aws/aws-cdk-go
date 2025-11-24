package mixinsawskinesisanalyticsv2


// For a SQL-based Kinesis Data Analytics application, describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
//
// Also used to describe the format of the reference data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recordColumnProperty := &RecordColumnProperty{
//   	Mapping: jsii.String("mapping"),
//   	Name: jsii.String("name"),
//   	SqlType: jsii.String("sqlType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-recordcolumn.html
//
type CfnApplicationReferenceDataSourcePropsMixin_RecordColumnProperty struct {
	// A reference to the data element in the streaming input or the reference data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-recordcolumn.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-recordcolumn-mapping
	//
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
	// The name of the column that is created in the in-application input stream or reference table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-recordcolumn.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-recordcolumn-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of column created in the in-application input stream or reference table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-recordcolumn.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-recordcolumn-sqltype
	//
	SqlType *string `field:"optional" json:"sqlType" yaml:"sqlType"`
}

