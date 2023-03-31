package awsdms


// Provides information that defines a MongoDB endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Endpoint configuration settings when using MongoDB as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html#CHAP_Source.MongoDB.Configuration) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mongoDbSettingsProperty := &mongoDbSettingsProperty{
//   	authMechanism: jsii.String("authMechanism"),
//   	authSource: jsii.String("authSource"),
//   	authType: jsii.String("authType"),
//   	databaseName: jsii.String("databaseName"),
//   	docsToInvestigate: jsii.String("docsToInvestigate"),
//   	extractDocId: jsii.String("extractDocId"),
//   	nestingLevel: jsii.String("nestingLevel"),
//   	password: jsii.String("password"),
//   	port: jsii.Number(123),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	serverName: jsii.String("serverName"),
//   	username: jsii.String("username"),
//   }
//
type CfnEndpoint_MongoDbSettingsProperty struct {
	// The authentication mechanism you use to access the MongoDB source endpoint.
	//
	// For the default value, in MongoDB version 2.x, `"default"` is `"mongodb_cr"` . For MongoDB version 3.x or later, `"default"` is `"scram_sha_1"` . This setting isn't used when `AuthType` is set to `"no"` .
	AuthMechanism *string `field:"optional" json:"authMechanism" yaml:"authMechanism"`
	// The MongoDB database name. This setting isn't used when `AuthType` is set to `"no"` .
	//
	// The default is `"admin"` .
	AuthSource *string `field:"optional" json:"authSource" yaml:"authSource"`
	// The authentication type you use to access the MongoDB source endpoint.
	//
	// When set to `"no"` , user name and password parameters are not used and can be empty.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The database name on the MongoDB source endpoint.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Indicates the number of documents to preview to determine the document organization.
	//
	// Use this setting when `NestingLevel` is set to `"one"` .
	//
	// Must be a positive value greater than `0` . Default value is `1000` .
	DocsToInvestigate *string `field:"optional" json:"docsToInvestigate" yaml:"docsToInvestigate"`
	// Specifies the document ID. Use this setting when `NestingLevel` is set to `"none"` .
	//
	// Default value is `"false"` .
	ExtractDocId *string `field:"optional" json:"extractDocId" yaml:"extractDocId"`
	// Specifies either document or table mode.
	//
	// Default value is `"none"` . Specify `"none"` to use document mode. Specify `"one"` to use table mode.
	NestingLevel *string `field:"optional" json:"nestingLevel" yaml:"nestingLevel"`
	// The password for the user account you use to access the MongoDB source endpoint.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port value for the MongoDB source endpoint.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the MongoDB endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the MongoDB endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// The name of the server on the MongoDB source endpoint.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The user name you use to access the MongoDB source endpoint.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

