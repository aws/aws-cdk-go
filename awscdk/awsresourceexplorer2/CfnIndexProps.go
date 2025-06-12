package awsresourceexplorer2


// Properties for defining a `CfnIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIndexProps := &CfnIndexProps{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-index.html
//
type CfnIndexProps struct {
	// Specifies the type of the index in this Region.
	//
	// For information about the aggregator index and how it differs from a local index, see [Turning on cross-Region search by creating an aggregator index](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html) in the *AWS Resource Explorer User Guide.* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-index.html#cfn-resourceexplorer2-index-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The specified tags are attached to only the index created in this AWS Region .
	//
	// The tags don't attach to any of the resources listed in the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-index.html#cfn-resourceexplorer2-index-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

