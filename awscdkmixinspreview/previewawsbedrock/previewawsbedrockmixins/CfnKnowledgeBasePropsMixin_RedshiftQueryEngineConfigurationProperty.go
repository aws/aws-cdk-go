package previewawsbedrockmixins


// Contains configurations for an Amazon Redshift query engine.
//
// Specify the type of query engine in `type` and include the corresponding field. For more information, see [Build a knowledge base by connecting to a structured data source](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-build-structured.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftQueryEngineConfigurationProperty := &RedshiftQueryEngineConfigurationProperty{
//   	ProvisionedConfiguration: &RedshiftProvisionedConfigurationProperty{
//   		AuthConfiguration: &RedshiftProvisionedAuthConfigurationProperty{
//   			DatabaseUser: jsii.String("databaseUser"),
//   			Type: jsii.String("type"),
//   			UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   		},
//   		ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	},
//   	ServerlessConfiguration: &RedshiftServerlessConfigurationProperty{
//   		AuthConfiguration: &RedshiftServerlessAuthConfigurationProperty{
//   			Type: jsii.String("type"),
//   			UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   		},
//   		WorkgroupArn: jsii.String("workgroupArn"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineconfiguration.html
//
type CfnKnowledgeBasePropsMixin_RedshiftQueryEngineConfigurationProperty struct {
	// Specifies configurations for a provisioned Amazon Redshift query engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryengineconfiguration-provisionedconfiguration
	//
	ProvisionedConfiguration interface{} `field:"optional" json:"provisionedConfiguration" yaml:"provisionedConfiguration"`
	// Specifies configurations for a serverless Amazon Redshift query engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryengineconfiguration-serverlessconfiguration
	//
	ServerlessConfiguration interface{} `field:"optional" json:"serverlessConfiguration" yaml:"serverlessConfiguration"`
	// The type of query engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftqueryengineconfiguration.html#cfn-bedrock-knowledgebase-redshiftqueryengineconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

