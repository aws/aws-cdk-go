package awsquicksight


// A physical table type for an S3 data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourceProperty := &S3SourceProperty{
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   	InputColumns: []interface{}{
//   		&InputColumnProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	UploadSettings: &UploadSettingsProperty{
//   		ContainsHeader: jsii.Boolean(false),
//   		Delimiter: jsii.String("delimiter"),
//   		Format: jsii.String("format"),
//   		StartFromRow: jsii.Number(123),
//   		TextQualifier: jsii.String("textQualifier"),
//   	},
//   }
//
type CfnDataSet_S3SourceProperty struct {
	// The Amazon Resource Name (ARN) for the data source.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// A physical table type for an S3 data source.
	//
	// > For files that aren't JSON, only `STRING` data types are supported in input columns.
	InputColumns interface{} `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// Information about the format for the S3 source file or files.
	UploadSettings interface{} `field:"optional" json:"uploadSettings" yaml:"uploadSettings"`
}

