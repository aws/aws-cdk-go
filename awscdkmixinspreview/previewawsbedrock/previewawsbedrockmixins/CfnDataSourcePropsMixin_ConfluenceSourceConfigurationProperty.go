package previewawsbedrockmixins


// The endpoint information to connect to your Confluence data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   confluenceSourceConfigurationProperty := &ConfluenceSourceConfigurationProperty{
//   	AuthType: jsii.String("authType"),
//   	CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   	HostType: jsii.String("hostType"),
//   	HostUrl: jsii.String("hostUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-confluencesourceconfiguration.html
//
type CfnDataSourcePropsMixin_ConfluenceSourceConfigurationProperty struct {
	// The supported authentication type to authenticate and connect to your Confluence instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-confluencesourceconfiguration.html#cfn-bedrock-datasource-confluencesourceconfiguration-authtype
	//
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your Confluence instance URL.
	//
	// For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see [Confluence connection configuration](https://docs.aws.amazon.com/bedrock/latest/userguide/confluence-data-source-connector.html#configuration-confluence-connector) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-confluencesourceconfiguration.html#cfn-bedrock-datasource-confluencesourceconfiguration-credentialssecretarn
	//
	CredentialsSecretArn *string `field:"optional" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The supported host type, whether online/cloud or server/on-premises.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-confluencesourceconfiguration.html#cfn-bedrock-datasource-confluencesourceconfiguration-hosttype
	//
	HostType *string `field:"optional" json:"hostType" yaml:"hostType"`
	// The Confluence host URL or instance URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-confluencesourceconfiguration.html#cfn-bedrock-datasource-confluencesourceconfiguration-hosturl
	//
	HostUrl *string `field:"optional" json:"hostUrl" yaml:"hostUrl"`
}

