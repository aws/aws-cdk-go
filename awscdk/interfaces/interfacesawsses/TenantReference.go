package interfacesawsses


// A reference to a Tenant resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tenantReference := &TenantReference{
//   	TenantArn: jsii.String("tenantArn"),
//   	TenantName: jsii.String("tenantName"),
//   }
//
type TenantReference struct {
	// The ARN of the Tenant resource.
	TenantArn *string `field:"required" json:"tenantArn" yaml:"tenantArn"`
	// The TenantName of the Tenant resource.
	TenantName *string `field:"required" json:"tenantName" yaml:"tenantName"`
}

