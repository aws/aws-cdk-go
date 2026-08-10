package awspersonalize


// Properties for defining a `CfnDataDeletionJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataDeletionJobProps := &CfnDataDeletionJobProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	DataSource: &DataSourceProperty{
//   		DataLocation: jsii.String("dataLocation"),
//   	},
//   	JobName: jsii.String("jobName"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datadeletionjob.html
//
type CfnDataDeletionJobProps struct {
	// The Amazon Resource Name (ARN) of the dataset group that has the datasets you want to delete records from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datadeletionjob.html#cfn-personalize-datadeletionjob-datasetgrouparn
	//
	DatasetGroupArn *string `field:"optional" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// The Amazon S3 bucket that contains the list of userIds to delete.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datadeletionjob.html#cfn-personalize-datadeletionjob-datasource
	//
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The name for the data deletion job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datadeletionjob.html#cfn-personalize-datadeletionjob-jobname
	//
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// The Amazon Resource Name (ARN) of the IAM role that has permissions to read from the Amazon S3 data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datadeletionjob.html#cfn-personalize-datadeletionjob-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

