package awscdk


// Properties for a tag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   tagProps := &TagProps{
//   	ApplyToLaunchedInstances: jsii.Boolean(false),
//   	ExcludeResourceTypes: []*string{
//   		jsii.String("excludeResourceTypes"),
//   	},
//   	IncludeResourceTypes: []*string{
//   		jsii.String("includeResourceTypes"),
//   	},
//   	Priority: jsii.Number(123),
//   }
//
// Experimental.
type TagProps struct {
	// Whether the tag should be applied to instances in an AutoScalingGroup.
	// Experimental.
	ApplyToLaunchedInstances *bool `field:"optional" json:"applyToLaunchedInstances" yaml:"applyToLaunchedInstances"`
	// An array of Resource Types that will not receive this tag.
	//
	// An empty array will allow this tag to be applied to all resources. A
	// non-empty array will apply this tag only if the Resource type is not in
	// this array.
	// Experimental.
	ExcludeResourceTypes *[]*string `field:"optional" json:"excludeResourceTypes" yaml:"excludeResourceTypes"`
	// An array of Resource Types that will receive this tag.
	//
	// An empty array will match any Resource. A non-empty array will apply this
	// tag only to Resource types that are included in this array.
	// Experimental.
	IncludeResourceTypes *[]*string `field:"optional" json:"includeResourceTypes" yaml:"includeResourceTypes"`
	// Priority of the tag operation.
	//
	// Higher or equal priority tags will take precedence.
	//
	// Setting priority will enable the user to control tags when they need to not
	// follow the default precedence pattern of last applied and closest to the
	// construct in the tree.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

