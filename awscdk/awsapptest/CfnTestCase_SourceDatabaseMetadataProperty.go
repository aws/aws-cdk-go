package awsapptest


// Specifies the source database metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceDatabaseMetadataProperty := &SourceDatabaseMetadataProperty{
//   	CaptureTool: jsii.String("captureTool"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-sourcedatabasemetadata.html
//
type CfnTestCase_SourceDatabaseMetadataProperty struct {
	// The capture tool of the source database metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-sourcedatabasemetadata.html#cfn-apptest-testcase-sourcedatabasemetadata-capturetool
	//
	CaptureTool *string `field:"required" json:"captureTool" yaml:"captureTool"`
	// The type of the source database metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-sourcedatabasemetadata.html#cfn-apptest-testcase-sourcedatabasemetadata-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

