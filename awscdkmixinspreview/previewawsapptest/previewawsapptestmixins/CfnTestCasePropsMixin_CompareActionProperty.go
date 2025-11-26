package previewawsapptestmixins


// Compares the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   compareActionProperty := &CompareActionProperty{
//   	Input: &InputProperty{
//   		File: &InputFileProperty{
//   			FileMetadata: &FileMetadataProperty{
//   				DatabaseCdc: &DatabaseCDCProperty{
//   					SourceMetadata: &SourceDatabaseMetadataProperty{
//   						CaptureTool: jsii.String("captureTool"),
//   						Type: jsii.String("type"),
//   					},
//   					TargetMetadata: &TargetDatabaseMetadataProperty{
//   						CaptureTool: jsii.String("captureTool"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				DataSets: []interface{}{
//   					&DataSetProperty{
//   						Ccsid: jsii.String("ccsid"),
//   						Format: jsii.String("format"),
//   						Length: jsii.Number(123),
//   						Name: jsii.String("name"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			SourceLocation: jsii.String("sourceLocation"),
//   			TargetLocation: jsii.String("targetLocation"),
//   		},
//   	},
//   	Output: &OutputProperty{
//   		File: &OutputFileProperty{
//   			FileLocation: jsii.String("fileLocation"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-compareaction.html
//
type CfnTestCasePropsMixin_CompareActionProperty struct {
	// The input of the compare action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-compareaction.html#cfn-apptest-testcase-compareaction-input
	//
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// The output of the compare action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-compareaction.html#cfn-apptest-testcase-compareaction-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
}

