package awsapptest


// Specifies a target database metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetDatabaseMetadataProperty := &TargetDatabaseMetadataProperty{
//   	CaptureTool: jsii.String("captureTool"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-targetdatabasemetadata.html
//
type CfnTestCase_TargetDatabaseMetadataProperty struct {
	// The capture tool of the target database metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-targetdatabasemetadata.html#cfn-apptest-testcase-targetdatabasemetadata-capturetool
	//
	CaptureTool *string `field:"required" json:"captureTool" yaml:"captureTool"`
	// The type of the target database metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-targetdatabasemetadata.html#cfn-apptest-testcase-targetdatabasemetadata-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

