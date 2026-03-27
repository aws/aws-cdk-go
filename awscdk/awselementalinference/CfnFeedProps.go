package awselementalinference


// Properties for defining a `CfnFeed`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cropping interface{}
//
//   cfnFeedProps := &CfnFeedProps{
//   	Name: jsii.String("name"),
//   	Outputs: []interface{}{
//   		&GetOutputProperty{
//   			Name: jsii.String("name"),
//   			OutputConfig: &OutputConfigProperty{
//   				Clipping: &ClippingConfigProperty{
//   					CallbackMetadata: jsii.String("callbackMetadata"),
//   				},
//   				Cropping: cropping,
//   			},
//   			Status: jsii.String("status"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-feed.html
//
type CfnFeedProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-feed.html#cfn-elementalinference-feed-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-feed.html#cfn-elementalinference-feed-outputs
	//
	Outputs interface{} `field:"required" json:"outputs" yaml:"outputs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-feed.html#cfn-elementalinference-feed-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

