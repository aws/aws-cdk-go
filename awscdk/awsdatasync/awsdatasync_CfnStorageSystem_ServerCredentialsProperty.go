package awsdatasync


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
	// `CfnStorageSystem.ServerCredentialsProperty.Password`.
	Password *string `field:"required" json:"password" yaml:"password"`
	// `CfnStorageSystem.ServerCredentialsProperty.Username`.
	Username *string `field:"required" json:"username" yaml:"username"`
}

