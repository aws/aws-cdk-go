package awsworkspaces


// Specifies the configurations of the Microsoft Entra.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   microsoftEntraConfigProperty := &MicrosoftEntraConfigProperty{
//   	ApplicationConfigSecretArn: jsii.String("applicationConfigSecretArn"),
//   	TenantId: jsii.String("tenantId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-microsoftentraconfig.html
//
type CfnDirectoryPropsMixin_MicrosoftEntraConfigProperty struct {
	// The Amazon Resource Name (ARN) of the application config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-microsoftentraconfig.html#cfn-workspaces-directory-microsoftentraconfig-applicationconfigsecretarn
	//
	ApplicationConfigSecretArn *string `field:"optional" json:"applicationConfigSecretArn" yaml:"applicationConfigSecretArn"`
	// The identifier of the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-microsoftentraconfig.html#cfn-workspaces-directory-microsoftentraconfig-tenantid
	//
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

