package awsathena


// A reference to a WorkGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workGroupReference := &WorkGroupReference{
//   	WorkGroupName: jsii.String("workGroupName"),
//   }
//
type WorkGroupReference struct {
	// The Name of the WorkGroup resource.
	WorkGroupName *string `field:"required" json:"workGroupName" yaml:"workGroupName"`
}

