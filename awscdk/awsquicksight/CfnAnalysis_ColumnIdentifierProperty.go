package awsquicksight


// A column of a data set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnIdentifierProperty := &ColumnIdentifierProperty{
//   	ColumnName: jsii.String("columnName"),
//
//   	// the properties below are optional
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	TopicIdentifier: jsii.String("topicIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnidentifier.html
//
type CfnAnalysis_ColumnIdentifierProperty struct {
	// The name of the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnidentifier.html#cfn-quicksight-analysis-columnidentifier-columnname
	//
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The data set that the column belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnidentifier.html#cfn-quicksight-analysis-columnidentifier-datasetidentifier
	//
	// Default: - "".
	//
	DataSetIdentifier *string `field:"optional" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnidentifier.html#cfn-quicksight-analysis-columnidentifier-topicidentifier
	//
	TopicIdentifier *string `field:"optional" json:"topicIdentifier" yaml:"topicIdentifier"`
}

