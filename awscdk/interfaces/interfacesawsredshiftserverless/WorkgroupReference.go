package interfacesawsredshiftserverless


// A reference to a Workgroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workgroupReference := &WorkgroupReference{
//   	WorkgroupName: jsii.String("workgroupName"),
//   }
//
type WorkgroupReference struct {
	// The WorkgroupName of the Workgroup resource.
	WorkgroupName *string `field:"required" json:"workgroupName" yaml:"workgroupName"`
}

