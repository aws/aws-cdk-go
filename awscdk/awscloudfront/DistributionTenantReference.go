package awscloudfront


// A reference to a DistributionTenant resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   distributionTenantReference := &DistributionTenantReference{
//   	DistributionTenantArn: jsii.String("distributionTenantArn"),
//   	DistributionTenantId: jsii.String("distributionTenantId"),
//   }
//
type DistributionTenantReference struct {
	// The ARN of the DistributionTenant resource.
	DistributionTenantArn *string `field:"required" json:"distributionTenantArn" yaml:"distributionTenantArn"`
	// The Id of the DistributionTenant resource.
	DistributionTenantId *string `field:"required" json:"distributionTenantId" yaml:"distributionTenantId"`
}

