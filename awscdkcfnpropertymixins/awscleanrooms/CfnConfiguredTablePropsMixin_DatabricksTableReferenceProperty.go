package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
type CfnConfiguredTablePropsMixin_DatabricksTableReferenceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databrickstablereference.html#cfn-cleanrooms-configuredtable-databrickstablereference-connectiontype
	//
	ConnectionType interface{} `field:"optional" json:"connectionType" yaml:"connectionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databrickstablereference.html#cfn-cleanrooms-configuredtable-databrickstablereference-schemaname
	//
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databrickstablereference.html#cfn-cleanrooms-configuredtable-databrickstablereference-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databrickstablereference.html#cfn-cleanrooms-configuredtable-databrickstablereference-workspaceurl
	//
	WorkspaceUrl *string `field:"optional" json:"workspaceUrl" yaml:"workspaceUrl"`
}

