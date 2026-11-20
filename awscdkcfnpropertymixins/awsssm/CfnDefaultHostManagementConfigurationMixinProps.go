package awsssm


// Properties for CfnDefaultHostManagementConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDefaultHostManagementConfigurationMixinProps := &CfnDefaultHostManagementConfigurationMixinProps{
//   	CustomerManagedRoleArn: jsii.String("customerManagedRoleArn"),
//   	PermissionModel: jsii.String("permissionModel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-defaulthostmanagementconfiguration.html
//
type CfnDefaultHostManagementConfigurationMixinProps struct {
	// The ARN of the customer-provided IAM role to register for DHMC.
	//
	// Required when PermissionModel is SELF_MANAGED and must not be set when PermissionModel is SERVICE_MANAGED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-defaulthostmanagementconfiguration.html#cfn-ssm-defaulthostmanagementconfiguration-customermanagedrolearn
	//
	CustomerManagedRoleArn *string `field:"optional" json:"customerManagedRoleArn" yaml:"customerManagedRoleArn"`
	// The permission model used to enable DHMC.
	//
	// SERVICE_MANAGED lets SSM create and manage the DHMC IAM role on your behalf. SELF_MANAGED registers an IAM role that you provide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-defaulthostmanagementconfiguration.html#cfn-ssm-defaulthostmanagementconfiguration-permissionmodel
	//
	PermissionModel *string `field:"optional" json:"permissionModel" yaml:"permissionModel"`
}

