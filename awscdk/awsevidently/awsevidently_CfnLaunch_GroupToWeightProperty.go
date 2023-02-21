package awsevidently


// A structure containing the percentage of launch traffic to allocate to one launch group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupToWeightProperty := &groupToWeightProperty{
//   	groupName: jsii.String("groupName"),
//   	splitWeight: jsii.Number(123),
//   }
//
type CfnLaunch_GroupToWeightProperty struct {
	// The name of the launch group.
	//
	// It can include up to 127 characters.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The portion of launch traffic to allocate to this launch group.
	//
	// This is represented in thousandths of a percent. For example, specify 20,000 to allocate 20% of the launch audience to this launch group.
	SplitWeight *float64 `field:"required" json:"splitWeight" yaml:"splitWeight"`
}

