package awsapptest


// Defines the Change Data Capture (CDC) of the database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseCDCProperty := &DatabaseCDCProperty{
//   	SourceMetadata: &SourceDatabaseMetadataProperty{
//   		CaptureTool: jsii.String("captureTool"),
//   		Type: jsii.String("type"),
//   	},
//   	TargetMetadata: &TargetDatabaseMetadataProperty{
//   		CaptureTool: jsii.String("captureTool"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-databasecdc.html
//
type CfnTestCase_DatabaseCDCProperty struct {
	// The source metadata of the database CDC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-databasecdc.html#cfn-apptest-testcase-databasecdc-sourcemetadata
	//
	SourceMetadata interface{} `field:"required" json:"sourceMetadata" yaml:"sourceMetadata"`
	// The target metadata of the database CDC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-databasecdc.html#cfn-apptest-testcase-databasecdc-targetmetadata
	//
	TargetMetadata interface{} `field:"required" json:"targetMetadata" yaml:"targetMetadata"`
}

