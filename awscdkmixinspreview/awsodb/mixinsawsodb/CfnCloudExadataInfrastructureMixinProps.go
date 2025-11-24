package mixinsawsodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCloudExadataInfrastructurePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCloudExadataInfrastructureMixinProps := &CfnCloudExadataInfrastructureMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   	ComputeCount: jsii.Number(123),
//   	CustomerContactsToSendToOci: []interface{}{
//   		&CustomerContactProperty{
//   			Email: jsii.String("email"),
//   		},
//   	},
//   	DatabaseServerType: jsii.String("databaseServerType"),
//   	DisplayName: jsii.String("displayName"),
//   	MaintenanceWindow: &MaintenanceWindowProperty{
//   		CustomActionTimeoutInMins: jsii.Number(123),
//   		DaysOfWeek: []*string{
//   			jsii.String("daysOfWeek"),
//   		},
//   		HoursOfDay: []interface{}{
//   			jsii.Number(123),
//   		},
//   		IsCustomActionTimeoutEnabled: jsii.Boolean(false),
//   		LeadTimeInWeeks: jsii.Number(123),
//   		Months: []*string{
//   			jsii.String("months"),
//   		},
//   		PatchingMode: jsii.String("patchingMode"),
//   		Preference: jsii.String("preference"),
//   		WeeksOfMonth: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   	Shape: jsii.String("shape"),
//   	StorageCount: jsii.Number(123),
//   	StorageServerType: jsii.String("storageServerType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html
//
type CfnCloudExadataInfrastructureMixinProps struct {
	// The name of the Availability Zone (AZ) where the Exadata infrastructure is located.
	//
	// Required when creating an Exadata infrastructure. Specify either AvailabilityZone or AvailabilityZoneId to define the location of the infrastructure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html#cfn-odb-cloudexadatainfrastructure-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AZ ID of the AZ where the Exadata infrastructure is located.
	//
	// Required when creating an Exadata infrastructure. Specify either AvailabilityZone or AvailabilityZoneId to define the location of the infrastructure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html#cfn-odb-cloudexadatainfrastructure-availabilityzoneid
	//
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The number of database servers for the Exadata infrastructure.
	//
	// Required when creating an Exadata infrastructure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html#cfn-odb-cloudexadatainfrastructure-computecount
	//
	ComputeCount *float64 `field:"optional" json:"computeCount" yaml:"computeCount"`
	// The email addresses of contacts to receive notification from Oracle about maintenance updates for the Exadata infrastructure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html#cfn-odb-cloudexadatainfrastructure-customercontactstosendtooci
	//
	CustomerContactsToSendToOci interface{} `field:"optional" json:"customerContactsToSendToOci" yaml:"customerContactsToSendToOci"`
	// The database server model type of the Exadata infrastructure.
	//
	// For the list of valid model names, use the `ListDbSystemShapes` operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html#cfn-odb-cloudexadatainfrastructure-databaseservertype
	//
	DatabaseServerType *string `field:"optional" json:"databaseServerType" yaml:"databaseServerType"`
	// The user-friendly name for the Exadata infrastructure.
	//
	// Required when creating an Exadata infrastructure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html#cfn-odb-cloudexadatainfrastructure-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The scheduling details for the maintenance window.
	//
	// Patching and system updates take place during the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html#cfn-odb-cloudexadatainfrastructure-maintenancewindow
	//
	MaintenanceWindow interface{} `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The model name of the Exadata infrastructure.
	//
	// Required when creating an Exadata infrastructure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html#cfn-odb-cloudexadatainfrastructure-shape
	//
	Shape *string `field:"optional" json:"shape" yaml:"shape"`
	// The number of storage servers that are activated for the Exadata infrastructure.
	//
	// Required when creating an Exadata infrastructure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html#cfn-odb-cloudexadatainfrastructure-storagecount
	//
	StorageCount *float64 `field:"optional" json:"storageCount" yaml:"storageCount"`
	// The storage server model type of the Exadata infrastructure.
	//
	// For the list of valid model names, use the `ListDbSystemShapes` operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html#cfn-odb-cloudexadatainfrastructure-storageservertype
	//
	StorageServerType *string `field:"optional" json:"storageServerType" yaml:"storageServerType"`
	// Tags to assign to the Exadata Infrastructure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html#cfn-odb-cloudexadatainfrastructure-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

