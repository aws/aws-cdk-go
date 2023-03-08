package awsappstream


// The credentials for the service account used by the streaming instance to connect to the directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceAccountCredentialsProperty := &serviceAccountCredentialsProperty{
//   	accountName: jsii.String("accountName"),
//   	accountPassword: jsii.String("accountPassword"),
//   }
//
type CfnDirectoryConfig_ServiceAccountCredentialsProperty struct {
	// The user name of the account.
	//
	// This account must have the following privileges: create computer objects, join computers to the domain, and change/reset the password on descendant computer objects for the organizational units specified.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// The password for the account.
	AccountPassword *string `field:"required" json:"accountPassword" yaml:"accountPassword"`
}

