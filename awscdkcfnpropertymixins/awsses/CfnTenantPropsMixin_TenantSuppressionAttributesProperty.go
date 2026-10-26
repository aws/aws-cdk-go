package awsses


// The suppression attributes for the tenant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
type CfnTenantPropsMixin_TenantSuppressionAttributesProperty struct {
	// The reasons that cause messages to be added to the suppression list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-tenant-tenantsuppressionattributes.html#cfn-ses-tenant-tenantsuppressionattributes-suppressedreasons
	//
	SuppressedReasons *[]*string `field:"optional" json:"suppressedReasons" yaml:"suppressedReasons"`
	// The scope of the suppression list for the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-tenant-tenantsuppressionattributes.html#cfn-ses-tenant-tenantsuppressionattributes-suppressionscope
	//
	SuppressionScope *string `field:"optional" json:"suppressionScope" yaml:"suppressionScope"`
}

