package awsbedrock


// Configurations for Redshift query engine Redshift backed storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftQueryEngineRedshiftStorageConfigurationProperty := &RedshiftQueryEngineRedshiftStorageConfigurationProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineredshiftstorageconfiguration.html
//
type CfnKnowledgeBase_RedshiftQueryEngineRedshiftStorageConfigurationProperty struct {
	// Redshift database name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineredshiftstorageconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryengineredshiftstorageconfiguration-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
}

