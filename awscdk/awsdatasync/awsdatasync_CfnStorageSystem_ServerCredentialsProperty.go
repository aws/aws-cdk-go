package awsdatasync


// The credentials that provide DataSync Discovery read access to your on-premises storage system's management interface.
//
// DataSync Discovery stores these credentials in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) . For more information, see [Accessing your on-premises storage system](https://docs.aws.amazon.com/datasync/latest/userguide/discovery-configure-storage.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverCredentialsProperty := &ServerCredentialsProperty{
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//   }
//
type CfnStorageSystem_ServerCredentialsProperty struct {
	// Specifies the password for your storage system's management interface.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies the user name for your storage system's management interface.
	Username *string `field:"required" json:"username" yaml:"username"`
}

