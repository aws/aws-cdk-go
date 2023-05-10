package awsglobalaccelerator


// Attributes required to import an existing accelerator to the stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorAttributes := &AcceleratorAttributes{
//   	AcceleratorArn: jsii.String("acceleratorArn"),
//   	DnsName: jsii.String("dnsName"),
//   }
//
// Experimental.
type AcceleratorAttributes struct {
	// The ARN of the accelerator.
	// Experimental.
	AcceleratorArn *string `field:"required" json:"acceleratorArn" yaml:"acceleratorArn"`
	// The DNS name of the accelerator.
	// Experimental.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
}

