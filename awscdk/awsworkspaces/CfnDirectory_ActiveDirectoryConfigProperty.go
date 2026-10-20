package awsworkspaces


// Information about the Active Directory config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activeDirectoryConfigProperty := &ActiveDirectoryConfigProperty{
//   	DomainName: jsii.String("domainName"),
//   	ServiceAccountSecretArn: jsii.String("serviceAccountSecretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-activedirectoryconfig.html
//
type CfnDirectory_ActiveDirectoryConfigProperty struct {
	// The name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-activedirectoryconfig.html#cfn-workspaces-directory-activedirectoryconfig-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Indicates the secret ARN on the service account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-activedirectoryconfig.html#cfn-workspaces-directory-activedirectoryconfig-serviceaccountsecretarn
	//
	ServiceAccountSecretArn *string `field:"required" json:"serviceAccountSecretArn" yaml:"serviceAccountSecretArn"`
}

