package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentProps := &CfnEnvironmentProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DeletionProtectionCheck: jsii.String("deletionProtectionCheck"),
//   	Description: jsii.String("description"),
//   	Monitors: []interface{}{
//   		&MonitorsProperty{
//   			AlarmArn: jsii.String("alarmArn"),
//   			AlarmRoleArn: jsii.String("alarmRoleArn"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-environment.html
//
type CfnEnvironmentProps struct {
	// The application ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-environment.html#cfn-appconfig-environment-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// A name for the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-environment.html#cfn-appconfig-environment-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A parameter to configure deletion protection.
	//
	// Deletion protection prevents a user from deleting an environment if your application called either [GetLatestConfiguration](https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_appconfigdata_GetLatestConfiguration.html) or [GetConfiguration](https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_GetConfiguration.html) in the environment during the specified interval.
	//
	// This parameter supports the following values:
	//
	// - `BYPASS` : Instructs AWS AppConfig to bypass the deletion protection check and delete a configuration profile even if deletion protection would have otherwise prevented it.
	// - `APPLY` : Instructs the deletion protection check to run, even if deletion protection is disabled at the account level. `APPLY` also forces the deletion protection check to run against resources created in the past hour, which are normally excluded from deletion protection checks.
	// - `ACCOUNT_DEFAULT` : The default setting, which instructs AWS AppConfig to implement the deletion protection value specified in the `UpdateAccountSettings` API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-environment.html#cfn-appconfig-environment-deletionprotectioncheck
	//
	DeletionProtectionCheck *string `field:"optional" json:"deletionProtectionCheck" yaml:"deletionProtectionCheck"`
	// A description of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-environment.html#cfn-appconfig-environment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Amazon CloudWatch alarms to monitor during the deployment process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-environment.html#cfn-appconfig-environment-monitors
	//
	Monitors interface{} `field:"optional" json:"monitors" yaml:"monitors"`
	// Metadata to assign to the environment.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-environment.html#cfn-appconfig-environment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

