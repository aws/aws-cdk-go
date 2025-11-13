package interfacesawscloudfront


// A reference to a VpcOrigin resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcOriginReference := &VpcOriginReference{
//   	VpcOriginArn: jsii.String("vpcOriginArn"),
//   	VpcOriginId: jsii.String("vpcOriginId"),
//   }
//
type VpcOriginReference struct {
	// The ARN of the VpcOrigin resource.
	VpcOriginArn *string `field:"required" json:"vpcOriginArn" yaml:"vpcOriginArn"`
	// The Id of the VpcOrigin resource.
	VpcOriginId *string `field:"required" json:"vpcOriginId" yaml:"vpcOriginId"`
}

