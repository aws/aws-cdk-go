package mixinsawsappstream


// The credentials for the service account used by the streaming instance to connect to the directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceAccountCredentialsProperty := &ServiceAccountCredentialsProperty{
//   	AccountName: jsii.String("accountName"),
//   	AccountPassword: jsii.String("accountPassword"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html
//
type CfnDirectoryConfigPropsMixin_ServiceAccountCredentialsProperty struct {
	// The user name of the account.
	//
	// This account must have the following privileges: create computer objects, join computers to the domain, and change/reset the password on descendant computer objects for the organizational units specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html#cfn-appstream-directoryconfig-serviceaccountcredentials-accountname
	//
	AccountName *string `field:"optional" json:"accountName" yaml:"accountName"`
	// The password for the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html#cfn-appstream-directoryconfig-serviceaccountcredentials-accountpassword
	//
	AccountPassword *string `field:"optional" json:"accountPassword" yaml:"accountPassword"`
}

