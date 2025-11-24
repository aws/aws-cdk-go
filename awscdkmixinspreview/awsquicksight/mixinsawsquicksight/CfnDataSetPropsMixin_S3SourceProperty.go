package mixinsawsquicksight


// A physical table type for an S3 data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3SourceProperty := &S3SourceProperty{
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   	InputColumns: []interface{}{
//   		&InputColumnProperty{
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   			SubType: jsii.String("subType"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	UploadSettings: &UploadSettingsProperty{
//   		ContainsHeader: jsii.Boolean(false),
//   		Delimiter: jsii.String("delimiter"),
//   		Format: jsii.String("format"),
//   		StartFromRow: jsii.Number(123),
//   		TextQualifier: jsii.String("textQualifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-s3source.html
//
type CfnDataSetPropsMixin_S3SourceProperty struct {
	// The Amazon Resource Name (ARN) for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-s3source.html#cfn-quicksight-dataset-s3source-datasourcearn
	//
	DataSourceArn *string `field:"optional" json:"dataSourceArn" yaml:"dataSourceArn"`
	// A physical table type for an S3 data source.
	//
	// > For files that aren't JSON, only `STRING` data types are supported in input columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-s3source.html#cfn-quicksight-dataset-s3source-inputcolumns
	//
	InputColumns interface{} `field:"optional" json:"inputColumns" yaml:"inputColumns"`
	// Information about the format for the S3 source file or files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-s3source.html#cfn-quicksight-dataset-s3source-uploadsettings
	//
	UploadSettings interface{} `field:"optional" json:"uploadSettings" yaml:"uploadSettings"`
}

