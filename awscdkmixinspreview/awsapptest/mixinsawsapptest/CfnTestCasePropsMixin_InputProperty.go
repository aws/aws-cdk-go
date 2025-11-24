package mixinsawsapptest


// Specifies the input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputProperty := &InputProperty{
//   	File: &InputFileProperty{
//   		FileMetadata: &FileMetadataProperty{
//   			DatabaseCdc: &DatabaseCDCProperty{
//   				SourceMetadata: &SourceDatabaseMetadataProperty{
//   					CaptureTool: jsii.String("captureTool"),
//   					Type: jsii.String("type"),
//   				},
//   				TargetMetadata: &TargetDatabaseMetadataProperty{
//   					CaptureTool: jsii.String("captureTool"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			DataSets: []interface{}{
//   				&DataSetProperty{
//   					Ccsid: jsii.String("ccsid"),
//   					Format: jsii.String("format"),
//   					Length: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		SourceLocation: jsii.String("sourceLocation"),
//   		TargetLocation: jsii.String("targetLocation"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-input.html
//
type CfnTestCasePropsMixin_InputProperty struct {
	// The file in the input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-input.html#cfn-apptest-testcase-input-file
	//
	File interface{} `field:"optional" json:"file" yaml:"file"`
}

