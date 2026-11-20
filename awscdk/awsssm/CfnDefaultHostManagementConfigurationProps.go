package awsssm


// Properties for defining a `CfnDefaultHostManagementConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDefaultHostManagementConfigurationProps := &CfnDefaultHostManagementConfigurationProps{
//   	PermissionModel: jsii.String("permissionModel"),
//
//   	// the properties below are optional
//   	CustomerManagedRoleArn: jsii.String("customerManagedRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-defaulthostmanagementconfiguration.html
//
type CfnDefaultHostManagementConfigurationProps struct {
	// The permission model used to enable DHMC.
	//
	// SERVICE_MANAGED lets SSM create and manage the DHMC IAM role on your behalf. SELF_MANAGED registers an IAM role that you provide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-defaulthostmanagementconfiguration.html#cfn-ssm-defaulthostmanagementconfiguration-permissionmodel
	//
	PermissionModel *string `field:"required" json:"permissionModel" yaml:"permissionModel"`
	// The ARN of the customer-provided IAM role to register for DHMC.
	//
	// Required when PermissionModel is SELF_MANAGED and must not be set when PermissionModel is SERVICE_MANAGED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-defaulthostmanagementconfiguration.html#cfn-ssm-defaulthostmanagementconfiguration-customermanagedrolearn
	//
	CustomerManagedRoleArn *string `field:"optional" json:"customerManagedRoleArn" yaml:"customerManagedRoleArn"`
}

