package awspersonalize


// Describes the data source that contains the data to upload to a dataset.
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
type CfnDataset_DataSourceProperty struct {
	// The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.
	//
	// For example:
	//
	// `s3://bucket-name/folder-name/`.
	DataLocation *string `field:"optional" json:"dataLocation" yaml:"dataLocation"`
}

