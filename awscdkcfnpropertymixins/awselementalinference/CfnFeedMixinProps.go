package awselementalinference


// Properties for CfnFeedPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var cropping interface{}
//
//   cfnFeedMixinProps := &CfnFeedMixinProps{
//   	Name: jsii.String("name"),
//   	Outputs: []interface{}{
//   		&GetOutputProperty{
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   			OutputConfig: &OutputConfigProperty{
//   				Clipping: &ClippingConfigProperty{
//   					CallbackMetadata: jsii.String("callbackMetadata"),
//   				},
//   				Cropping: cropping,
//   			},
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-feed.html
//
type CfnFeedMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-feed.html#cfn-elementalinference-feed-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-feed.html#cfn-elementalinference-feed-outputs
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-feed.html#cfn-elementalinference-feed-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

