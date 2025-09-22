package awsmediatailor


// A reference to a VodSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vodSourceReference := &VodSourceReference{
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   	VodSourceArn: jsii.String("vodSourceArn"),
//   	VodSourceName: jsii.String("vodSourceName"),
//   }
//
type VodSourceReference struct {
	// The SourceLocationName of the VodSource resource.
	SourceLocationName *string `field:"required" json:"sourceLocationName" yaml:"sourceLocationName"`
	// The ARN of the VodSource resource.
	VodSourceArn *string `field:"required" json:"vodSourceArn" yaml:"vodSourceArn"`
	// The VodSourceName of the VodSource resource.
	VodSourceName *string `field:"required" json:"vodSourceName" yaml:"vodSourceName"`
}

