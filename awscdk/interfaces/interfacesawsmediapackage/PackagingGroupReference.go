package interfacesawsmediapackage


// A reference to a PackagingGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   packagingGroupReference := &PackagingGroupReference{
//   	PackagingGroupArn: jsii.String("packagingGroupArn"),
//   	PackagingGroupId: jsii.String("packagingGroupId"),
//   }
//
type PackagingGroupReference struct {
	// The ARN of the PackagingGroup resource.
	PackagingGroupArn *string `field:"required" json:"packagingGroupArn" yaml:"packagingGroupArn"`
	// The Id of the PackagingGroup resource.
	PackagingGroupId *string `field:"required" json:"packagingGroupId" yaml:"packagingGroupId"`
}

