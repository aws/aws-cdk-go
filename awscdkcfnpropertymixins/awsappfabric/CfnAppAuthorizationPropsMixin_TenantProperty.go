package awsappfabric


// Contains information about an application tenant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tenantProperty := &TenantProperty{
//   	TenantDisplayName: jsii.String("tenantDisplayName"),
//   	TenantIdentifier: jsii.String("tenantIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-appauthorization-tenant.html
//
type CfnAppAuthorizationPropsMixin_TenantProperty struct {
	// The display name of the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-appauthorization-tenant.html#cfn-appfabric-appauthorization-tenant-tenantdisplayname
	//
	TenantDisplayName *string `field:"optional" json:"tenantDisplayName" yaml:"tenantDisplayName"`
	// The ID of the application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-appauthorization-tenant.html#cfn-appfabric-appauthorization-tenant-tenantidentifier
	//
	TenantIdentifier *string `field:"optional" json:"tenantIdentifier" yaml:"tenantIdentifier"`
}

