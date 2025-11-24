package mixinsawslookoutmetrics


// Contains information about the column used to track time in a source data file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timestampColumnProperty := &TimestampColumnProperty{
//   	ColumnFormat: jsii.String("columnFormat"),
//   	ColumnName: jsii.String("columnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-timestampcolumn.html
//
type CfnAnomalyDetectorPropsMixin_TimestampColumnProperty struct {
	// The format of the timestamp column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-timestampcolumn.html#cfn-lookoutmetrics-anomalydetector-timestampcolumn-columnformat
	//
	ColumnFormat *string `field:"optional" json:"columnFormat" yaml:"columnFormat"`
	// The name of the timestamp column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-timestampcolumn.html#cfn-lookoutmetrics-anomalydetector-timestampcolumn-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
}

