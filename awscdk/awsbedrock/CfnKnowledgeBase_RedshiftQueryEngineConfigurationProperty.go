package awsbedrock


// Contains configurations for an Amazon Redshift query engine.
//
// Specify the type of query engine in `type` and include the corresponding field. For more information, see [Build a knowledge base by connecting to a structured data source](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-build-structured.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftQueryEngineConfigurationProperty := &RedshiftQueryEngineConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ProvisionedConfiguration: &RedshiftProvisionedConfigurationProperty{
//   		AuthConfiguration: &RedshiftProvisionedAuthConfigurationProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			DatabaseUser: jsii.String("databaseUser"),
//   			UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   		},
//   		ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	},
//   	ServerlessConfiguration: &RedshiftServerlessConfigurationProperty{
//   		AuthConfiguration: &RedshiftServerlessAuthConfigurationProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   		},
//   		WorkgroupArn: jsii.String("workgroupArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineconfiguration.html
//
type CfnKnowledgeBase_RedshiftQueryEngineConfigurationProperty struct {
	// The type of query engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryengineconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies configurations for a provisioned Amazon Redshift query engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryengineconfiguration-provisionedconfiguration
	//
	ProvisionedConfiguration interface{} `field:"optional" json:"provisionedConfiguration" yaml:"provisionedConfiguration"`
	// Specifies configurations for a serverless Amazon Redshift query engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryengineconfiguration-serverlessconfiguration
	//
	ServerlessConfiguration interface{} `field:"optional" json:"serverlessConfiguration" yaml:"serverlessConfiguration"`
}

