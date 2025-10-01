package awsapprunner


// A reference to a VpcIngressConnection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcIngressConnectionReference := &VpcIngressConnectionReference{
//   	VpcIngressConnectionArn: jsii.String("vpcIngressConnectionArn"),
//   }
//
type VpcIngressConnectionReference struct {
	// The VpcIngressConnectionArn of the VpcIngressConnection resource.
	VpcIngressConnectionArn *string `field:"required" json:"vpcIngressConnectionArn" yaml:"vpcIngressConnectionArn"`
}

