package awspersonalize


// Describes a job that imports training data from a data source (Amazon S3 bucket) to an Amazon Personalize dataset.
//
// For more information, see [CreateDatasetImportJob](https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetImportJob.html) .
//
// A dataset import job can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataSource interface{}
//
//   datasetImportJobProperty := &datasetImportJobProperty{
//   	datasetArn: jsii.String("datasetArn"),
//   	datasetImportJobArn: jsii.String("datasetImportJobArn"),
//   	dataSource: dataSource,
//   	jobName: jsii.String("jobName"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnDataset_DatasetImportJobProperty struct {
	// The Amazon Resource Name (ARN) of the dataset that receives the imported data.
	DatasetArn *string `field:"optional" json:"datasetArn" yaml:"datasetArn"`
	// The ARN of the dataset import job.
	DatasetImportJobArn *string `field:"optional" json:"datasetImportJobArn" yaml:"datasetImportJobArn"`
	// The Amazon S3 bucket that contains the training data to import.
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The name of the import job.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// The ARN of the IAM role that has permissions to read from the Amazon S3 data source.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

