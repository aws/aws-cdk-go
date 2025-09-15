package awsec2


// A reference to a PlacementGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementGroupReference := &PlacementGroupReference{
//   	GroupName: jsii.String("groupName"),
//   }
//
type PlacementGroupReference struct {
	// The GroupName of the PlacementGroup resource.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
}

