package awssystemsmanagersap


// The credentials of your SAP application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialProperty := &CredentialProperty{
//   	CredentialType: jsii.String("credentialType"),
//   	DatabaseName: jsii.String("databaseName"),
//   	SecretId: jsii.String("secretId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-credential.html
//
type CfnApplication_CredentialProperty struct {
	// The type of the application credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-credential.html#cfn-systemsmanagersap-application-credential-credentialtype
	//
	CredentialType *string `field:"optional" json:"credentialType" yaml:"credentialType"`
	// The name of the SAP HANA database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-credential.html#cfn-systemsmanagersap-application-credential-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The secret ID created in AWS Secrets Manager to store the credentials of the SAP application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-credential.html#cfn-systemsmanagersap-application-credential-secretid
	//
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
}

