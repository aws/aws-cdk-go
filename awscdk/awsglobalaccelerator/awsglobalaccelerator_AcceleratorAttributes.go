package awsglobalaccelerator


// Attributes required to import an existing accelerator to the stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorAttributes := &acceleratorAttributes{
//   	acceleratorArn: jsii.String("acceleratorArn"),
//   	dnsName: jsii.String("dnsName"),
//   }
//
type AcceleratorAttributes struct {
	// The ARN of the accelerator.
	AcceleratorArn *string `field:"required" json:"acceleratorArn" yaml:"acceleratorArn"`
	// The DNS name of the accelerator.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
}

