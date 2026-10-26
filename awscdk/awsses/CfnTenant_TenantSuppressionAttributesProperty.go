package awsses


// The suppression attributes for the tenant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tenantSuppressionAttributesProperty := &TenantSuppressionAttributesProperty{
//   	SuppressedReasons: []*string{
//   		jsii.String("suppressedReasons"),
//   	},
//   	SuppressionScope: jsii.String("suppressionScope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-tenant-tenantsuppressionattributes.html
//
type CfnTenant_TenantSuppressionAttributesProperty struct {
	// The reasons that cause messages to be added to the suppression list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-tenant-tenantsuppressionattributes.html#cfn-ses-tenant-tenantsuppressionattributes-suppressedreasons
	//
	SuppressedReasons *[]*string `field:"required" json:"suppressedReasons" yaml:"suppressedReasons"`
	// The scope of the suppression list for the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-tenant-tenantsuppressionattributes.html#cfn-ses-tenant-tenantsuppressionattributes-suppressionscope
	//
	SuppressionScope *string `field:"required" json:"suppressionScope" yaml:"suppressionScope"`
}

