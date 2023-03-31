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
//   cfnDatasetProps := &cfnDatasetProps{
//   	datasetGroupArn: jsii.String("datasetGroupArn"),
//   	datasetType: jsii.String("datasetType"),
//   	name: jsii.String("name"),
//   	schemaArn: jsii.String("schemaArn"),
//
//   	// the properties below are optional
//   	datasetImportJob: &datasetImportJobProperty{
//   		datasetArn: jsii.String("datasetArn"),
//   		datasetImportJobArn: jsii.String("datasetImportJobArn"),
//   		dataSource: dataSource,
//   		jobName: jsii.String("jobName"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   }
//
type CfnDatasetProps struct {
	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string `field:"required" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// One of the following values:.
	//
	// - Interactions
	// - Items
	// - Users.
	DatasetType *string `field:"required" json:"datasetType" yaml:"datasetType"`
	// The name of the dataset.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the associated schema.
	SchemaArn *string `field:"required" json:"schemaArn" yaml:"schemaArn"`
	// Describes a job that imports training data from a data source (Amazon S3 bucket) to an Amazon Personalize dataset.
	DatasetImportJob interface{} `field:"optional" json:"datasetImportJob" yaml:"datasetImportJob"`
}

