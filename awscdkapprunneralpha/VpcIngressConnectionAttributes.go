package awscdkapprunneralpha


// Attributes for the App Runner VPC Ingress Connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   vpcIngressConnectionAttributes := &VpcIngressConnectionAttributes{
//   	DomainName: jsii.String("domainName"),
//   	Status: jsii.String("status"),
//   	VpcIngressConnectionArn: jsii.String("vpcIngressConnectionArn"),
//   	VpcIngressConnectionName: jsii.String("vpcIngressConnectionName"),
//   }
//
// Experimental.
type VpcIngressConnectionAttributes struct {
	// The domain name associated with the VPC Ingress Connection resource.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The current status of the VPC Ingress Connection.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// The Amazon Resource Name (ARN) of the VPC Ingress Connection.
	// Experimental.
	VpcIngressConnectionArn *string `field:"required" json:"vpcIngressConnectionArn" yaml:"vpcIngressConnectionArn"`
	// The name of the VPC Ingress Connection.
	// Experimental.
	VpcIngressConnectionName *string `field:"required" json:"vpcIngressConnectionName" yaml:"vpcIngressConnectionName"`
}

