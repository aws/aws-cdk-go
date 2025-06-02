package awsbedrock


// The endpoint information to connect to your Salesforce data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceSourceConfigurationProperty := &SalesforceSourceConfigurationProperty{
//   	AuthType: jsii.String("authType"),
//   	CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   	HostUrl: jsii.String("hostUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcesourceconfiguration.html
//
type CfnDataSource_SalesforceSourceConfigurationProperty struct {
	// The supported authentication type to authenticate and connect to your Salesforce instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcesourceconfiguration.html#cfn-bedrock-datasource-salesforcesourceconfiguration-authtype
	//
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your Salesforce instance URL.
	//
	// For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see [Salesforce connection configuration](https://docs.aws.amazon.com/bedrock/latest/userguide/salesforce-data-source-connector.html#configuration-salesforce-connector) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcesourceconfiguration.html#cfn-bedrock-datasource-salesforcesourceconfiguration-credentialssecretarn
	//
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The Salesforce host URL or instance URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcesourceconfiguration.html#cfn-bedrock-datasource-salesforcesourceconfiguration-hosturl
	//
	HostUrl *string `field:"required" json:"hostUrl" yaml:"hostUrl"`
}

