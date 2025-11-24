package mixinsawsoam


// Properties for CfnSinkPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policy interface{}
//
//   cfnSinkMixinProps := &CfnSinkMixinProps{
//   	Name: jsii.String("name"),
//   	Policy: policy,
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-sink.html
//
type CfnSinkMixinProps struct {
	// A name for the sink.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-sink.html#cfn-oam-sink-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM policy that grants permissions to source accounts to link to this sink.
	//
	// The policy can grant permission in the following ways:
	//
	// - Include organization IDs or organization paths to permit all accounts in an organization
	// - Include account IDs to permit the specified accounts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-sink.html#cfn-oam-sink-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// An array of key-value pairs to apply to the sink.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-sink.html#cfn-oam-sink-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

