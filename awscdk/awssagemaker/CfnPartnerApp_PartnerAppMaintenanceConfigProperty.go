package awssagemaker


// A collection of settings that specify the maintenance schedule for the PartnerApp.
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
	// The maintenance window start day and time for the PartnerApp.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-partnerapp-partnerappmaintenanceconfig.html#cfn-sagemaker-partnerapp-partnerappmaintenanceconfig-maintenancewindowstart
	//
	MaintenanceWindowStart *string `field:"required" json:"maintenanceWindowStart" yaml:"maintenanceWindowStart"`
}

