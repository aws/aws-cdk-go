package awsworkspaces


// The default values that are used to create WorkSpaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultWorkspaceCreationPropertiesProperty := &DefaultWorkspaceCreationPropertiesProperty{
//   	CustomSecurityGroupId: jsii.String("customSecurityGroupId"),
//   	DefaultOu: jsii.String("defaultOu"),
//   	EnableInternetAccess: jsii.Boolean(false),
//   	EnableMaintenanceMode: jsii.Boolean(false),
//   	InstanceIamRoleArn: jsii.String("instanceIamRoleArn"),
//   	UserEnabledAsLocalAdministrator: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-defaultworkspacecreationproperties.html
//
type CfnDirectory_DefaultWorkspaceCreationPropertiesProperty struct {
	// The identifier of the default security group to apply to WorkSpaces when they are created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-defaultworkspacecreationproperties.html#cfn-workspaces-directory-defaultworkspacecreationproperties-customsecuritygroupid
	//
	CustomSecurityGroupId *string `field:"optional" json:"customSecurityGroupId" yaml:"customSecurityGroupId"`
	// The organizational unit (OU) in the directory for the WorkSpace machine accounts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-defaultworkspacecreationproperties.html#cfn-workspaces-directory-defaultworkspacecreationproperties-defaultou
	//
	DefaultOu *string `field:"optional" json:"defaultOu" yaml:"defaultOu"`
	// Specifies whether to automatically assign an Elastic public IP address to WorkSpaces in this directory by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-defaultworkspacecreationproperties.html#cfn-workspaces-directory-defaultworkspacecreationproperties-enableinternetaccess
	//
	EnableInternetAccess interface{} `field:"optional" json:"enableInternetAccess" yaml:"enableInternetAccess"`
	// Specifies whether maintenance mode is enabled for WorkSpaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-defaultworkspacecreationproperties.html#cfn-workspaces-directory-defaultworkspacecreationproperties-enablemaintenancemode
	//
	EnableMaintenanceMode interface{} `field:"optional" json:"enableMaintenanceMode" yaml:"enableMaintenanceMode"`
	// Indicates the IAM role ARN of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-defaultworkspacecreationproperties.html#cfn-workspaces-directory-defaultworkspacecreationproperties-instanceiamrolearn
	//
	InstanceIamRoleArn *string `field:"optional" json:"instanceIamRoleArn" yaml:"instanceIamRoleArn"`
	// Specifies whether WorkSpace users are local administrators on their WorkSpaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-defaultworkspacecreationproperties.html#cfn-workspaces-directory-defaultworkspacecreationproperties-userenabledaslocaladministrator
	//
	UserEnabledAsLocalAdministrator interface{} `field:"optional" json:"userEnabledAsLocalAdministrator" yaml:"userEnabledAsLocalAdministrator"`
}

