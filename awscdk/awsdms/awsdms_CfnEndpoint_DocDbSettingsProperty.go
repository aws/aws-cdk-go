package awsdms


// Provides information that defines a DocumentDB endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Using extra connections attributes with Amazon DocumentDB as a source](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DocumentDB.html#CHAP_Source.DocumentDB.ECAs) and [Using Amazon DocumentDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DocumentDB.html) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   docDbSettingsProperty := &docDbSettingsProperty{
//   	docsToInvestigate: jsii.Number(123),
//   	extractDocId: jsii.Boolean(false),
//   	nestingLevel: jsii.String("nestingLevel"),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   }
//
type CfnEndpoint_DocDbSettingsProperty struct {
	// Indicates the number of documents to preview to determine the document organization.
	//
	// Use this setting when `NestingLevel` is set to `"one"` .
	//
	// Must be a positive value greater than `0` . Default value is `1000` .
	DocsToInvestigate *float64 `field:"optional" json:"docsToInvestigate" yaml:"docsToInvestigate"`
	// Specifies the document ID. Use this setting when `NestingLevel` is set to `"none"` .
	//
	// Default value is `"false"` .
	ExtractDocId interface{} `field:"optional" json:"extractDocId" yaml:"extractDocId"`
	// Specifies either document or table mode.
	//
	// Default value is `"none"` . Specify `"none"` to use document mode. Specify `"one"` to use table mode.
	NestingLevel *string `field:"optional" json:"nestingLevel" yaml:"nestingLevel"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the DocumentDB endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the DocumentDB endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

