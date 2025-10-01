package awsiotwireless


// A reference to a FuotaTask resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fuotaTaskReference := &FuotaTaskReference{
//   	FuotaTaskArn: jsii.String("fuotaTaskArn"),
//   	FuotaTaskId: jsii.String("fuotaTaskId"),
//   }
//
type FuotaTaskReference struct {
	// The ARN of the FuotaTask resource.
	FuotaTaskArn *string `field:"required" json:"fuotaTaskArn" yaml:"fuotaTaskArn"`
	// The Id of the FuotaTask resource.
	FuotaTaskId *string `field:"required" json:"fuotaTaskId" yaml:"fuotaTaskId"`
}

