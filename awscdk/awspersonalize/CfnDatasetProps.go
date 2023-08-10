package awspersonalize


// Properties for defining a `CfnDataset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataSource interface{}
//
//   cfnDatasetProps := &CfnDatasetProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	DatasetType: jsii.String("datasetType"),
//   	Name: jsii.String("name"),
//   	SchemaArn: jsii.String("schemaArn"),
//
//   	// the properties below are optional
//   	DatasetImportJob: &DatasetImportJobProperty{
//   		DatasetArn: jsii.String("datasetArn"),
//   		DatasetImportJobArn: jsii.String("datasetImportJobArn"),
//   		DataSource: dataSource,
//   		JobName: jsii.String("jobName"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-dataset.html
//
type CfnDatasetProps struct {
	// The Amazon Resource Name (ARN) of the dataset group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-dataset.html#cfn-personalize-dataset-datasetgrouparn
	//
	DatasetGroupArn *string `field:"required" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// One of the following values:.
	//
	// - Interactions
	// - Items
	// - Users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-dataset.html#cfn-personalize-dataset-datasettype
	//
	DatasetType *string `field:"required" json:"datasetType" yaml:"datasetType"`
	// The name of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-dataset.html#cfn-personalize-dataset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the associated schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-dataset.html#cfn-personalize-dataset-schemaarn
	//
	SchemaArn *string `field:"required" json:"schemaArn" yaml:"schemaArn"`
	// Describes a job that imports training data from a data source (Amazon S3 bucket) to an Amazon Personalize dataset.
	//
	// If you specify a dataset import job as part of a dataset, all dataset import job fields are required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-dataset.html#cfn-personalize-dataset-datasetimportjob
	//
	DatasetImportJob interface{} `field:"optional" json:"datasetImportJob" yaml:"datasetImportJob"`
}

