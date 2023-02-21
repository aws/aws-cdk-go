package awsappflow


// The connector-specific profile credentials required by Infor Nexus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inforNexusConnectorProfileCredentialsProperty := &InforNexusConnectorProfileCredentialsProperty{
//   	AccessKeyId: jsii.String("accessKeyId"),
//   	Datakey: jsii.String("datakey"),
//   	SecretAccessKey: jsii.String("secretAccessKey"),
//   	UserId: jsii.String("userId"),
//   }
//
type CfnConnectorProfile_InforNexusConnectorProfileCredentialsProperty struct {
	// The Access Key portion of the credentials.
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// The encryption keys used to encrypt data.
	Datakey *string `field:"required" json:"datakey" yaml:"datakey"`
	// The secret key used to sign requests.
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
	// The identifier for the user.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

