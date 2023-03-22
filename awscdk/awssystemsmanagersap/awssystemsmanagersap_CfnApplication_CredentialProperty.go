package awssystemsmanagersap


// The credentials of your SAP application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialProperty := &credentialProperty{
//   	credentialType: jsii.String("credentialType"),
//   	databaseName: jsii.String("databaseName"),
//   	secretId: jsii.String("secretId"),
//   }
//
type CfnApplication_CredentialProperty struct {
	// The type of the application credentials.
	CredentialType *string `field:"optional" json:"credentialType" yaml:"credentialType"`
	// The name of the SAP HANA database.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The secret ID created in AWS Secrets Manager to store the credentials of the SAP application.
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
}

