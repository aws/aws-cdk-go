package awscodepipeline


// Git push filter for trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitPushFilter := &GitPushFilter{
//   	TagsExcludes: []*string{
//   		jsii.String("tagsExcludes"),
//   	},
//   	TagsIncludes: []*string{
//   		jsii.String("tagsIncludes"),
//   	},
//   }
//
type GitPushFilter struct {
	// The list of patterns of Git tags that, when pushed, are to be excluded from starting the pipeline.
	//
	// You can filter with glob patterns. The `tagsExcludes` takes priority
	// over the `tagsIncludes`.
	//
	// Maximum length of this array is 8.
	// Default: - no tags.
	//
	TagsExcludes *[]*string `field:"optional" json:"tagsExcludes" yaml:"tagsExcludes"`
	// The list of patterns of Git tags that, when pushed, are to be included as criteria that starts the pipeline.
	//
	// You can filter with glob patterns. The `tagsExcludes` takes priority
	// over the `tagsIncludes`.
	//
	// Maximum length of this array is 8.
	// Default: - no tags.
	//
	TagsIncludes *[]*string `field:"optional" json:"tagsIncludes" yaml:"tagsIncludes"`
}

