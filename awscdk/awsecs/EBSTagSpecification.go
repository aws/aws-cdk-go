package awsecs


// Tag Specification for EBS volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eBSTagSpecification := &EBSTagSpecification{
//   	PropagateTags: awscdk.Aws_ecs.EbsPropagatedTagSource_SERVICE,
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type EBSTagSpecification struct {
	// Specifies whether to propagate the tags from the task definition or the service to the task.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION
	// Default: - undefined.
	//
	PropagateTags EbsPropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The tags to apply to the volume.
	// Default: - No tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

