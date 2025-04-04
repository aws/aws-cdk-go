package awsapptest


// Specifies the input file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnTestCase_InputFileProperty struct {
	// The file metadata of the input file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-inputfile.html#cfn-apptest-testcase-inputfile-filemetadata
	//
	FileMetadata interface{} `field:"required" json:"fileMetadata" yaml:"fileMetadata"`
	// The source location of the input file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-inputfile.html#cfn-apptest-testcase-inputfile-sourcelocation
	//
	SourceLocation *string `field:"required" json:"sourceLocation" yaml:"sourceLocation"`
	// The target location of the input file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-inputfile.html#cfn-apptest-testcase-inputfile-targetlocation
	//
	TargetLocation *string `field:"required" json:"targetLocation" yaml:"targetLocation"`
}

