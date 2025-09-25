package awslightsail


// A reference to a Disk resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   diskReference := &DiskReference{
//   	DiskArn: jsii.String("diskArn"),
//   	DiskName: jsii.String("diskName"),
//   }
//
type DiskReference struct {
	// The ARN of the Disk resource.
	DiskArn *string `field:"required" json:"diskArn" yaml:"diskArn"`
	// The DiskName of the Disk resource.
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
}

