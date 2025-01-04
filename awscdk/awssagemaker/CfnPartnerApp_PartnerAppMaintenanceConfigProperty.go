package awssagemaker


// Maintenance configuration settings for the SageMaker Partner AI App.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partnerAppMaintenanceConfigProperty := &PartnerAppMaintenanceConfigProperty{
//   	MaintenanceWindowStart: jsii.String("maintenanceWindowStart"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-partnerapp-partnerappmaintenanceconfig.html
//
type CfnPartnerApp_PartnerAppMaintenanceConfigProperty struct {
	// The day and time of the week in Coordinated Universal Time (UTC) 24-hour standard time that weekly maintenance updates are scheduled.
	//
	// This value must take the following format: `3-letter-day:24-h-hour:minute` . For example: `TUE:03:30` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-partnerapp-partnerappmaintenanceconfig.html#cfn-sagemaker-partnerapp-partnerappmaintenanceconfig-maintenancewindowstart
	//
	MaintenanceWindowStart *string `field:"required" json:"maintenanceWindowStart" yaml:"maintenanceWindowStart"`
}

