package previewawssecurityhubmixins


// Properties for CfnHubV2PropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHubV2MixinProps := &CfnHubV2MixinProps{
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hubv2.html
//
type CfnHubV2MixinProps struct {
	// The tags to add to the hub V2 resource when you enable Security Hub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hubv2.html#cfn-securityhub-hubv2-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

