package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databricksTableReferenceProperty := &DatabricksTableReferenceProperty{
//   	ConnectionType: &DatabricksConnectionTypeProperty{
//   		DeltaSharing: &DatabricksDeltaSharingConnectionProperty{
//   			DeltaShareEndpoint: jsii.String("deltaShareEndpoint"),
//   			ShareName: jsii.String("shareName"),
//   		},
//   	},
//   	SchemaName: jsii.String("schemaName"),
//   	TableName: jsii.String("tableName"),
//   	WorkspaceUrl: jsii.String("workspaceUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databrickstablereference.html
//
type CfnConfiguredTable_DatabricksTableReferenceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databrickstablereference.html#cfn-cleanrooms-configuredtable-databrickstablereference-connectiontype
	//
	ConnectionType interface{} `field:"required" json:"connectionType" yaml:"connectionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databrickstablereference.html#cfn-cleanrooms-configuredtable-databrickstablereference-schemaname
	//
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databrickstablereference.html#cfn-cleanrooms-configuredtable-databrickstablereference-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databrickstablereference.html#cfn-cleanrooms-configuredtable-databrickstablereference-workspaceurl
	//
	WorkspaceUrl *string `field:"required" json:"workspaceUrl" yaml:"workspaceUrl"`
}

