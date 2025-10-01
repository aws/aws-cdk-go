package awsmsk


// A reference to a VpcConnection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectionReference := &VpcConnectionReference{
//   	VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   }
//
type VpcConnectionReference struct {
	// The Arn of the VpcConnection resource.
	VpcConnectionArn *string `field:"required" json:"vpcConnectionArn" yaml:"vpcConnectionArn"`
}

