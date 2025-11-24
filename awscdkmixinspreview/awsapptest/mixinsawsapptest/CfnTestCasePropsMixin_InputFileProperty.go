package mixinsawsapptest


// Specifies the input file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputFileProperty := &InputFileProperty{
//   	FileMetadata: &FileMetadataProperty{
//   		DatabaseCdc: &DatabaseCDCProperty{
//   			SourceMetadata: &SourceDatabaseMetadataProperty{
//   				CaptureTool: jsii.String("captureTool"),
//   				Type: jsii.String("type"),
//   			},
//   			TargetMetadata: &TargetDatabaseMetadataProperty{
//   				CaptureTool: jsii.String("captureTool"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		DataSets: []interface{}{
//   			&DataSetProperty{
//   				Ccsid: jsii.String("ccsid"),
//   				Format: jsii.String("format"),
//   				Length: jsii.Number(123),
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	SourceLocation: jsii.String("sourceLocation"),
//   	TargetLocation: jsii.String("targetLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-inputfile.html
//
type CfnTestCasePropsMixin_InputFileProperty struct {
	// The file metadata of the input file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-inputfile.html#cfn-apptest-testcase-inputfile-filemetadata
	//
	FileMetadata interface{} `field:"optional" json:"fileMetadata" yaml:"fileMetadata"`
	// The source location of the input file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-inputfile.html#cfn-apptest-testcase-inputfile-sourcelocation
	//
	SourceLocation *string `field:"optional" json:"sourceLocation" yaml:"sourceLocation"`
	// The target location of the input file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-inputfile.html#cfn-apptest-testcase-inputfile-targetlocation
	//
	TargetLocation *string `field:"optional" json:"targetLocation" yaml:"targetLocation"`
}

