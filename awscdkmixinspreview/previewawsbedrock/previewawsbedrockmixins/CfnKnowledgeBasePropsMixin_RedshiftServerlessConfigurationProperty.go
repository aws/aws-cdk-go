package previewawsbedrockmixins


// Contains configurations for authentication to Amazon Redshift Serverless.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftServerlessConfigurationProperty := &RedshiftServerlessConfigurationProperty{
//   	AuthConfiguration: &RedshiftServerlessAuthConfigurationProperty{
//   		Type: jsii.String("type"),
//   		UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   	},
//   	WorkgroupArn: jsii.String("workgroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftserverlessconfiguration.html
//
type CfnKnowledgeBasePropsMixin_RedshiftServerlessConfigurationProperty struct {
	// Specifies configurations for authentication to an Amazon Redshift provisioned data warehouse.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftserverlessconfiguration.html#cfn-bedrock-knowledgebase-redshiftserverlessconfiguration-authconfiguration
	//
	AuthConfiguration interface{} `field:"optional" json:"authConfiguration" yaml:"authConfiguration"`
	// The ARN of the Amazon Redshift workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftserverlessconfiguration.html#cfn-bedrock-knowledgebase-redshiftserverlessconfiguration-workgrouparn
	//
	WorkgroupArn *string `field:"optional" json:"workgroupArn" yaml:"workgroupArn"`
}

