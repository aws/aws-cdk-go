package awsworkspacesthinclient

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
//   	DesktopArn: jsii.String("desktopArn"),
//
//   	// the properties below are optional
//   	DesiredSoftwareSetId: jsii.String("desiredSoftwareSetId"),
//   	DesktopEndpoint: jsii.String("desktopEndpoint"),
//   	DeviceCreationTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	MaintenanceWindow: &MaintenanceWindowProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		ApplyTimeOf: jsii.String("applyTimeOf"),
//   		DaysOfTheWeek: []*string{
//   			jsii.String("daysOfTheWeek"),
//   		},
//   		EndTimeHour: jsii.Number(123),
//   		EndTimeMinute: jsii.Number(123),
//   		StartTimeHour: jsii.Number(123),
//   		StartTimeMinute: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	SoftwareSetUpdateMode: jsii.String("softwareSetUpdateMode"),
//   	SoftwareSetUpdateSchedule: jsii.String("softwareSetUpdateSchedule"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html
//
type CfnEnvironmentProps struct {
	// The Amazon Resource Name (ARN) of the desktop to stream from Amazon WorkSpaces, WorkSpaces Web, or AppStream 2.0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-desktoparn
	//
	DesktopArn *string `field:"required" json:"desktopArn" yaml:"desktopArn"`
	// The ID of the software set to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-desiredsoftwaresetid
	//
	DesiredSoftwareSetId *string `field:"optional" json:"desiredSoftwareSetId" yaml:"desiredSoftwareSetId"`
	// The URL for the identity provider login (only for environments that use AppStream 2.0).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-desktopendpoint
	//
	DesktopEndpoint *string `field:"optional" json:"desktopEndpoint" yaml:"desktopEndpoint"`
	// The tag keys and optional values for the newly created devices for this environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-devicecreationtags
	//
	DeviceCreationTags interface{} `field:"optional" json:"deviceCreationTags" yaml:"deviceCreationTags"`
	// The Amazon Resource Name (ARN) of the AWS Key Management Service key used to encrypt the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// A specification for a time window to apply software updates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-maintenancewindow
	//
	MaintenanceWindow interface{} `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The name of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An option to define which software updates to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-softwaresetupdatemode
	//
	SoftwareSetUpdateMode *string `field:"optional" json:"softwareSetUpdateMode" yaml:"softwareSetUpdateMode"`
	// An option to define if software updates should be applied within a maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-softwaresetupdateschedule
	//
	SoftwareSetUpdateSchedule *string `field:"optional" json:"softwareSetUpdateSchedule" yaml:"softwareSetUpdateSchedule"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

