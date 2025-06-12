package awspersonalize


// Describes the data source that contains the data to upload to a dataset, or the list of records to delete from Amazon Personalize.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceProperty := &DataSourceProperty{
//   	DataLocation: jsii.String("dataLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-dataset-datasource.html
//
type CfnDataset_DataSourceProperty struct {
	// For dataset import jobs, the path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.
	//
	// For data deletion jobs, the path to the Amazon S3 bucket that stores the list of records to delete.
	//
	// For example:
	//
	// `s3://bucket-name/folder-name/fileName.csv`
	//
	// If your CSV files are in a folder in your Amazon S3 bucket and you want your import job or data deletion job to consider multiple files, you can specify the path to the folder. With a data deletion job, Amazon Personalize uses all files in the folder and any sub folder. Use the following syntax with a `/` after the folder name:
	//
	// `s3://bucket-name/folder-name/`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-dataset-datasource.html#cfn-personalize-dataset-datasource-datalocation
	//
	DataLocation *string `field:"optional" json:"dataLocation" yaml:"dataLocation"`
}

