package awsssmquicksetup


// The definition of a Quick Setup configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationDefinitionProperty := &ConfigurationDefinitionProperty{
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Id: jsii.String("id"),
//   	LocalDeploymentAdministrationRoleArn: jsii.String("localDeploymentAdministrationRoleArn"),
//   	LocalDeploymentExecutionRoleName: jsii.String("localDeploymentExecutionRoleName"),
//   	TypeVersion: jsii.String("typeVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-configurationdefinition.html
//
type CfnConfigurationManager_ConfigurationDefinitionProperty struct {
	// The parameters for the configuration definition type.
	//
	// Parameters for configuration definitions vary based the configuration type. The following lists outline the parameters for each configuration type.
	//
	// - **AWS Config Recording (Type: AWS QuickSetupType-CFGRecording)** - - `RecordAllResources`
	//
	// - Description: (Optional) A boolean value that determines whether all supported resources are recorded. The default value is " `true` ".
	// - `ResourceTypesToRecord`
	//
	// - Description: (Optional) A comma separated list of resource types you want to record.
	// - `RecordGlobalResourceTypes`
	//
	// - Description: (Optional) A boolean value that determines whether global resources are recorded with all resource configurations. The default value is " `false` ".
	// - `GlobalResourceTypesRegion`
	//
	// - Description: (Optional) Determines the AWS Region where global resources are recorded.
	// - `UseCustomBucket`
	//
	// - Description: (Optional) A boolean value that determines whether a custom Amazon S3 bucket is used for delivery. The default value is " `false` ".
	// - `DeliveryBucketName`
	//
	// - Description: (Optional) The name of the Amazon S3 bucket you want AWS Config to deliver configuration snapshots and configuration history files to.
	// - `DeliveryBucketPrefix`
	//
	// - Description: (Optional) The key prefix you want to use in the custom Amazon S3 bucket.
	// - `NotificationOptions`
	//
	// - Description: (Optional) Determines the notification configuration for the recorder. The valid values are `NoStreaming` , `UseExistingTopic` , and `CreateTopic` . The default value is `NoStreaming` .
	// - `CustomDeliveryTopicAccountId`
	//
	// - Description: (Optional) The ID of the AWS account where the Amazon SNS topic you want to use for notifications resides. You must specify a value for this parameter if you use the `UseExistingTopic` notification option.
	// - `CustomDeliveryTopicName`
	//
	// - Description: (Optional) The name of the Amazon SNS topic you want to use for notifications. You must specify a value for this parameter if you use the `UseExistingTopic` notification option.
	// - `RemediationSchedule`
	//
	// - Description: (Optional) A rate expression that defines the schedule for drift remediation. The valid values are `rate(30 days)` , `rate(7 days)` , `rate(1 days)` , and `none` . The default value is " `none` ".
	// - `TargetAccounts`
	//
	// - Description: (Optional) The ID of the AWS account initiating the configuration deployment. You only need to provide a value for this parameter if you want to deploy the configuration locally. A value must be provided for either `TargetAccounts` or `TargetOrganizationalUnits` .
	// - `TargetOrganizationalUnits`
	//
	// - Description: (Optional) The ID of the root of your Organization. This configuration type doesn't currently support choosing specific OUs. The configuration will be deployed to all the OUs in the Organization.
	// - `TargetRegions`
	//
	// - Description: (Required) A comma separated list of AWS Regions you want to deploy the configuration to.
	// - **Change Manager (Type: AWS QuickSetupType-SSMChangeMgr)** - - `DelegatedAccountId`
	//
	// - Description: (Required) The ID of the delegated administrator account.
	// - `JobFunction`
	//
	// - Description: (Required) The name for the Change Manager job function.
	// - `PermissionType`
	//
	// - Description: (Optional) Specifies whether you want to use default administrator permissions for the job function role, or provide a custom IAM policy. The valid values are `CustomPermissions` and `AdminPermissions` . The default value for the parameter is `CustomerPermissions` .
	// - `CustomPermissions`
	//
	// - Description: (Optional) A JSON string containing the IAM policy you want your job function to use. You must provide a value for this parameter if you specify `CustomPermissions` for the `PermissionType` parameter.
	// - `TargetOrganizationalUnits`
	//
	// - Description: (Required) A comma separated list of organizational units (OUs) you want to deploy the configuration to.
	// - `TargetRegions`
	//
	// - Description: (Required) A comma separated list of AWS Regions you want to deploy the configuration to.
	// - **Conformance Packs (Type: AWS QuickSetupType-CFGCPacks)** - - `DelegatedAccountId`
	//
	// - Description: (Optional) The ID of the delegated administrator account. This parameter is required for Organization deployments.
	// - `RemediationSchedule`
	//
	// - Description: (Optional) A rate expression that defines the schedule for drift remediation. The valid values are `rate(30 days)` , `rate(14 days)` , `rate(2 days)` , and `none` . The default value is " `none` ".
	// - `CPackNames`
	//
	// - Description: (Required) A comma separated list of AWS Config conformance packs.
	// - `TargetAccounts`
	//
	// - Description: (Optional) The ID of the AWS account initiating the configuration deployment. You only need to provide a value for this parameter if you want to deploy the configuration locally. A value must be provided for either `TargetAccounts` or `TargetOrganizationalUnits` .
	// - `TargetOrganizationalUnits`
	//
	// - Description: (Optional) The ID of the root of your Organization. This configuration type doesn't currently support choosing specific OUs. The configuration will be deployed to all the OUs in the Organization.
	// - `TargetRegions`
	//
	// - Description: (Required) A comma separated list of AWS Regions you want to deploy the configuration to.
	// - **Default Host Management Configuration (Type: AWS QuickSetupType-DHMC)** - - `UpdateSSMAgent`
	//
	// - Description: (Optional) A boolean value that determines whether the SSM Agent is updated on the target instances every 2 weeks. The default value is " `true` ".
	// - `TargetOrganizationalUnits`
	//
	// - Description: (Required) A comma separated list of organizational units (OUs) you want to deploy the configuration to.
	// - `TargetRegions`
	//
	// - Description: (Required) A comma separated list of AWS Regions you want to deploy the configuration to.
	// - **DevOps Guru (Type: AWS QuickSetupType-DevOpsGuru)** - - `AnalyseAllResources`
	//
	// - Description: (Optional) A boolean value that determines whether DevOps Guru analyzes all AWS CloudFormation stacks in the account. The default value is " `false` ".
	// - `EnableSnsNotifications`
	//
	// - Description: (Optional) A boolean value that determines whether DevOps Guru sends notifications when an insight is created. The default value is " `true` ".
	// - `EnableSsmOpsItems`
	//
	// - Description: (Optional) A boolean value that determines whether DevOps Guru creates an OpsCenter OpsItem when an insight is created. The default value is " `true` ".
	// - `EnableDriftRemediation`
	//
	// - Description: (Optional) A boolean value that determines whether a drift remediation schedule is used. The default value is " `false` ".
	// - `RemediationSchedule`
	//
	// - Description: (Optional) A rate expression that defines the schedule for drift remediation. The valid values are `rate(30 days)` , `rate(14 days)` , `rate(1 days)` , and `none` . The default value is " `none` ".
	// - `TargetAccounts`
	//
	// - Description: (Optional) The ID of the AWS account initiating the configuration deployment. You only need to provide a value for this parameter if you want to deploy the configuration locally. A value must be provided for either `TargetAccounts` or `TargetOrganizationalUnits` .
	// - `TargetOrganizationalUnits`
	//
	// - Description: (Optional) A comma separated list of organizational units (OUs) you want to deploy the configuration to.
	// - `TargetRegions`
	//
	// - Description: (Required) A comma separated list of AWS Regions you want to deploy the configuration to.
	// - **Distributor (Type: AWS QuickSetupType-Distributor)** - - `PackagesToInstall`
	//
	// - Description: (Required) A comma separated list of packages you want to install on the target instances. The valid values are `AWSEFSTools` , `AWSCWAgent` , and `AWSEC2LaunchAgent` .
	// - `RemediationSchedule`
	//
	// - Description: (Optional) A rate expression that defines the schedule for drift remediation. The valid values are `rate(30 days)` , `rate(14 days)` , `rate(2 days)` , and `none` . The default value is " `rate(30 days)` ".
	// - `IsPolicyAttachAllowed`
	//
	// - Description: (Optional) A boolean value that determines whether Quick Setup attaches policies to instances profiles already associated with the target instances. The default value is " `false` ".
	// - `TargetType`
	//
	// - Description: (Optional) Determines how instances are targeted for local account deployments. Don't specify a value for this parameter if you're deploying to OUs. The valid values are `*` , `InstanceIds` , `ResourceGroups` , and `Tags` . Use `*` to target all instances in the account.
	// - `TargetInstances`
	//
	// - Description: (Optional) A comma separated list of instance IDs. You must provide a value for this parameter if you specify `InstanceIds` for the `TargetType` parameter.
	// - `TargetTagKey`
	//
	// - Description: (Required) The tag key assigned to the instances you want to target. You must provide a value for this parameter if you specify `Tags` for the `TargetType` parameter.
	// - `TargetTagValue`
	//
	// - Description: (Required) The value of the tag key assigned to the instances you want to target. You must provide a value for this parameter if you specify `Tags` for the `TargetType` parameter.
	// - `ResourceGroupName`
	//
	// - Description: (Required) The name of the resource group associated with the instances you want to target. You must provide a value for this parameter if you specify `ResourceGroups` for the `TargetType` parameter.
	// - `TargetAccounts`
	//
	// - Description: (Optional) The ID of the AWS account initiating the configuration deployment. You only need to provide a value for this parameter if you want to deploy the configuration locally. A value must be provided for either `TargetAccounts` or `TargetOrganizationalUnits` .
	// - `TargetOrganizationalUnits`
	//
	// - Description: (Optional) A comma separated list of organizational units (OUs) you want to deploy the configuration to.
	// - `TargetRegions`
	//
	// - Description: (Required) A comma separated list of AWS Regions you want to deploy the configuration to.
	// - **Host Management (Type: AWS QuickSetupType-SSMHostMgmt)** - - `UpdateSSMAgent`
	//
	// - Description: (Optional) A boolean value that determines whether the SSM Agent is updated on the target instances every 2 weeks. The default value is " `true` ".
	// - `UpdateEc2LaunchAgent`
	//
	// - Description: (Optional) A boolean value that determines whether the EC2 Launch agent is updated on the target instances every month. The default value is " `false` ".
	// - `CollectInventory`
	//
	// - Description: (Optional) A boolean value that determines whether instance metadata is collected on the target instances every 30 minutes. The default value is " `true` ".
	// - `ScanInstances`
	//
	// - Description: (Optional) A boolean value that determines whether the target instances are scanned daily for available patches. The default value is " `true` ".
	// - `InstallCloudWatchAgent`
	//
	// - Description: (Optional) A boolean value that determines whether the Amazon CloudWatch agent is installed on the target instances. The default value is " `false` ".
	// - `UpdateCloudWatchAgent`
	//
	// - Description: (Optional) A boolean value that determines whether the Amazon CloudWatch agent is updated on the target instances every month. The default value is " `false` ".
	// - `IsPolicyAttachAllowed`
	//
	// - Description: (Optional) A boolean value that determines whether Quick Setup attaches policies to instances profiles already associated with the target instances. The default value is " `false` ".
	// - `TargetType`
	//
	// - Description: (Optional) Determines how instances are targeted for local account deployments. Don't specify a value for this parameter if you're deploying to OUs. The valid values are `*` , `InstanceIds` , `ResourceGroups` , and `Tags` . Use `*` to target all instances in the account.
	// - `TargetInstances`
	//
	// - Description: (Optional) A comma separated list of instance IDs. You must provide a value for this parameter if you specify `InstanceIds` for the `TargetType` parameter.
	// - `TargetTagKey`
	//
	// - Description: (Optional) The tag key assigned to the instances you want to target. You must provide a value for this parameter if you specify `Tags` for the `TargetType` parameter.
	// - `TargetTagValue`
	//
	// - Description: (Optional) The value of the tag key assigned to the instances you want to target. You must provide a value for this parameter if you specify `Tags` for the `TargetType` parameter.
	// - `ResourceGroupName`
	//
	// - Description: (Optional) The name of the resource group associated with the instances you want to target. You must provide a value for this parameter if you specify `ResourceGroups` for the `TargetType` parameter.
	// - `TargetAccounts`
	//
	// - Description: (Optional) The ID of the AWS account initiating the configuration deployment. You only need to provide a value for this parameter if you want to deploy the configuration locally. A value must be provided for either `TargetAccounts` or `TargetOrganizationalUnits` .
	// - `TargetOrganizationalUnits`
	//
	// - Description: (Optional) A comma separated list of organizational units (OUs) you want to deploy the configuration to.
	// - `TargetRegions`
	//
	// - Description: (Required) A comma separated list of AWS Regions you want to deploy the configuration to.
	// - **OpsCenter (Type: AWS QuickSetupType-SSMOpsCenter)** - - `DelegatedAccountId`
	//
	// - Description: (Required) The ID of the delegated administrator account.
	// - `TargetOrganizationalUnits`
	//
	// - Description: (Required) A comma separated list of organizational units (OUs) you want to deploy the configuration to.
	// - `TargetRegions`
	//
	// - Description: (Required) A comma separated list of AWS Regions you want to deploy the configuration to.
	// - **Patch Policy (Type: AWS QuickSetupType-PatchPolicy)** - - `PatchPolicyName`
	//
	// - Description: (Required) A name for the patch policy. The value you provide is applied to target Amazon EC2 instances as a tag.
	// - `SelectedPatchBaselines`
	//
	// - Description: (Required) An array of JSON objects containing the information for the patch baselines to include in your patch policy.
	// - `PatchBaselineUseDefault`
	//
	// - Description: (Optional) A value that determines whether the selected patch baselines are all AWS provided. Supported values are `default` and `custom` .
	// - `PatchBaselineRegion`
	//
	// - Description: (Required) The AWS Region where the patch baseline exist.
	// - `ConfigurationOptionsPatchOperation`
	//
	// - Description: (Optional) Determines whether target instances scan for available patches, or scan and install available patches. The valid values are `Scan` and `ScanAndInstall` . The default value for the parameter is `Scan` .
	// - `ConfigurationOptionsScanValue`
	//
	// - Description: (Optional) A cron expression that is used as the schedule for when instances scan for available patches.
	// - `ConfigurationOptionsInstallValue`
	//
	// - Description: (Optional) A cron expression that is used as the schedule for when instances install available patches.
	// - `ConfigurationOptionsScanNextInterval`
	//
	// - Description: (Optional) A boolean value that determines whether instances should scan for available patches at the next cron interval. The default value is " `false` ".
	// - `ConfigurationOptionsInstallNextInterval`
	//
	// - Description: (Optional) A boolean value that determines whether instances should scan for available patches at the next cron interval. The default value is " `false` ".
	// - `RebootOption`
	//
	// - Description: (Optional) Determines whether instances are rebooted after patches are installed. Valid values are `RebootIfNeeded` and `NoReboot` .
	// - `IsPolicyAttachAllowed`
	//
	// - Description: (Optional) A boolean value that determines whether Quick Setup attaches policies to instances profiles already associated with the target instances. The default value is " `false` ".
	// - `OutputLogEnableS3`
	//
	// - Description: (Optional) A boolean value that determines whether command output logs are sent to Amazon S3.
	// - `OutputS3Location`
	//
	// - Description: (Optional) Information about the Amazon S3 bucket where you want to store the output details of the request.
	//
	// - `OutputBucketRegion`
	//
	// - Description: (Optional) The AWS Region where the Amazon S3 bucket you want to deliver command output to is located.
	// - `OutputS3BucketName`
	//
	// - Description: (Optional) The name of the Amazon S3 bucket you want to deliver command output to.
	// - `OutputS3KeyPrefix`
	//
	// - Description: (Optional) The key prefix you want to use in the custom Amazon S3 bucket.
	// - `TargetType`
	//
	// - Description: (Optional) Determines how instances are targeted for local account deployments. Don't specify a value for this parameter if you're deploying to OUs. The valid values are `*` , `InstanceIds` , `ResourceGroups` , and `Tags` . Use `*` to target all instances in the account.
	// - `TargetInstances`
	//
	// - Description: (Optional) A comma separated list of instance IDs. You must provide a value for this parameter if you specify `InstanceIds` for the `TargetType` parameter.
	// - `TargetTagKey`
	//
	// - Description: (Required) The tag key assigned to the instances you want to target. You must provide a value for this parameter if you specify `Tags` for the `TargetType` parameter.
	// - `TargetTagValue`
	//
	// - Description: (Required) The value of the tag key assigned to the instances you want to target. You must provide a value for this parameter if you specify `Tags` for the `TargetType` parameter.
	// - `ResourceGroupName`
	//
	// - Description: (Required) The name of the resource group associated with the instances you want to target. You must provide a value for this parameter if you specify `ResourceGroups` for the `TargetType` parameter.
	// - `TargetAccounts`
	//
	// - Description: (Optional) The ID of the AWS account initiating the configuration deployment. You only need to provide a value for this parameter if you want to deploy the configuration locally. A value must be provided for either `TargetAccounts` or `TargetOrganizationalUnits` .
	// - `TargetOrganizationalUnits`
	//
	// - Description: (Optional) A comma separated list of organizational units (OUs) you want to deploy the configuration to.
	// - `TargetRegions`
	//
	// - Description: (Required) A comma separated list of AWS Regions you want to deploy the configuration to.
	// - **Resource Explorer (Type: AWS QuickSetupType-ResourceExplorer)** - - `SelectedAggregatorRegion`
	//
	// - Description: (Required) The AWS Region where you want to create the aggregator index.
	// - `ReplaceExistingAggregator`
	//
	// - Description: (Required) A boolean value that determines whether to demote an existing aggregator if it is in a Region that differs from the value you specify for the `SelectedAggregatorRegion` .
	// - `TargetOrganizationalUnits`
	//
	// - Description: (Required) A comma separated list of organizational units (OUs) you want to deploy the configuration to.
	// - `TargetRegions`
	//
	// - Description: (Required) A comma separated list of AWS Regions you want to deploy the configuration to.
	// - **Resource Scheduler (Type: AWS QuickSetupType-Scheduler)** - - `TargetTagKey`
	//
	// - Description: (Required) The tag key assigned to the instances you want to target.
	// - `TargetTagValue`
	//
	// - Description: (Required) The value of the tag key assigned to the instances you want to target.
	// - `ICalendarString`
	//
	// - Description: (Required) An iCalendar formatted string containing the schedule you want Change Manager to use.
	// - `TargetAccounts`
	//
	// - Description: (Optional) The ID of the AWS account initiating the configuration deployment. You only need to provide a value for this parameter if you want to deploy the configuration locally. A value must be provided for either `TargetAccounts` or `TargetOrganizationalUnits` .
	// - `TargetOrganizationalUnits`
	//
	// - Description: (Optional) A comma separated list of organizational units (OUs) you want to deploy the configuration to.
	// - `TargetRegions`
	//
	// - Description: (Required) A comma separated list of AWS Regions you want to deploy the configuration to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-configurationdefinition.html#cfn-ssmquicksetup-configurationmanager-configurationdefinition-parameters
	//
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// The type of the Quick Setup configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-configurationdefinition.html#cfn-ssmquicksetup-configurationmanager-configurationdefinition-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The ID of the configuration definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-configurationdefinition.html#cfn-ssmquicksetup-configurationmanager-configurationdefinition-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ARN of the IAM role used to administrate local configuration deployments.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only for organizational deployments of types other than `AWSQuickSetupType-PatchPolicy` . A value must be provided when you are running an organizational deployment for a patch policy or running any type of deployment for a single account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-configurationdefinition.html#cfn-ssmquicksetup-configurationmanager-configurationdefinition-localdeploymentadministrationrolearn
	//
	LocalDeploymentAdministrationRoleArn *string `field:"optional" json:"localDeploymentAdministrationRoleArn" yaml:"localDeploymentAdministrationRoleArn"`
	// The name of the IAM role used to deploy local configurations.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only for organizational deployments of types other than `AWSQuickSetupType-PatchPolicy` . A value must be provided when you are running an organizational deployment for a patch policy or running any type of deployment for a single account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-configurationdefinition.html#cfn-ssmquicksetup-configurationmanager-configurationdefinition-localdeploymentexecutionrolename
	//
	LocalDeploymentExecutionRoleName *string `field:"optional" json:"localDeploymentExecutionRoleName" yaml:"localDeploymentExecutionRoleName"`
	// The version of the Quick Setup type used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-configurationdefinition.html#cfn-ssmquicksetup-configurationmanager-configurationdefinition-typeversion
	//
	TypeVersion *string `field:"optional" json:"typeVersion" yaml:"typeVersion"`
}

