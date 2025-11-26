package previewawsapptestmixins


// Specifies a file metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fileMetadataProperty := &FileMetadataProperty{
//   	DatabaseCdc: &DatabaseCDCProperty{
//   		SourceMetadata: &SourceDatabaseMetadataProperty{
//   			CaptureTool: jsii.String("captureTool"),
//   			Type: jsii.String("type"),
//   		},
//   		TargetMetadata: &TargetDatabaseMetadataProperty{
//   			CaptureTool: jsii.String("captureTool"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	DataSets: []interface{}{
//   		&DataSetProperty{
//   			Ccsid: jsii.String("ccsid"),
//   			Format: jsii.String("format"),
//   			Length: jsii.Number(123),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-filemetadata.html
//
type CfnTestCasePropsMixin_FileMetadataProperty struct {
	// The database CDC of the file metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-filemetadata.html#cfn-apptest-testcase-filemetadata-databasecdc
	//
	DatabaseCdc interface{} `field:"optional" json:"databaseCdc" yaml:"databaseCdc"`
	// The data sets of the file metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-filemetadata.html#cfn-apptest-testcase-filemetadata-datasets
	//
	DataSets interface{} `field:"optional" json:"dataSets" yaml:"dataSets"`
}

