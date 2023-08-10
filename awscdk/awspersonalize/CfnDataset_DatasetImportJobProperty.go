package awspersonalize


// Describes a job that imports training data from a data source (Amazon S3 bucket) to an Amazon Personalize dataset.
//
// A dataset import job can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// If you specify a dataset import job as part of a dataset, all dataset import job fields are required.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataSource interface{}
//
//   datasetImportJobProperty := &DatasetImportJobProperty{
//   	DatasetArn: jsii.String("datasetArn"),
//   	DatasetImportJobArn: jsii.String("datasetImportJobArn"),
//   	DataSource: dataSource,
//   	JobName: jsii.String("jobName"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-dataset-datasetimportjob.html
//
type CfnDataset_DatasetImportJobProperty struct {
	// The Amazon Resource Name (ARN) of the dataset that receives the imported data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-dataset-datasetimportjob.html#cfn-personalize-dataset-datasetimportjob-datasetarn
	//
	DatasetArn *string `field:"optional" json:"datasetArn" yaml:"datasetArn"`
	// The ARN of the dataset import job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-dataset-datasetimportjob.html#cfn-personalize-dataset-datasetimportjob-datasetimportjobarn
	//
	DatasetImportJobArn *string `field:"optional" json:"datasetImportJobArn" yaml:"datasetImportJobArn"`
	// The Amazon S3 bucket that contains the training data to import.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-dataset-datasetimportjob.html#cfn-personalize-dataset-datasetimportjob-datasource
	//
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The name of the import job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-dataset-datasetimportjob.html#cfn-personalize-dataset-datasetimportjob-jobname
	//
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// The ARN of the IAM role that has permissions to read from the Amazon S3 data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-dataset-datasetimportjob.html#cfn-personalize-dataset-datasetimportjob-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

