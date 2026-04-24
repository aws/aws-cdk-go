package awssynthetics


// Properties for defining a CloudWatch Synthetics Group.
//
// Example:
//   // First, declare your canaries
//   var canary1 ICanary
//   var canary2 ICanary
//
//
//   group := synthetics.NewGroup(this, jsii.String("MyCanaryGroup"), &GroupProps{
//   	GroupName: jsii.String("production-canaries"),
//   	Canaries: []ICanary{
//   		canary1,
//   		canary2,
//   	},
//   })
//
type GroupProps struct {
	// List of canaries to associate with this group.
	//
	// Each group can contain as many as 10 canaries.
	// Default: - No canaries are associated with the group initially.
	//
	Canaries *[]ICanary `field:"optional" json:"canaries" yaml:"canaries"`
	// A name for the group. Must contain only lowercase alphanumeric characters, hyphens, or underscores, and be at most 64 characters.
	//
	// The names for all groups in your account, across all Regions, must be unique.
	// Default: - A unique name will be generated from the construct ID.
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
}

