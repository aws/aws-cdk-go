package previewawsorganizationsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnResourcePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var content interface{}
//
//   cfnResourcePolicyMixinProps := &CfnResourcePolicyMixinProps{
//   	Content: content,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-resourcepolicy.html
//
type CfnResourcePolicyMixinProps struct {
	// The policy text of the organization resource policy.
	//
	// You can specify the resource policy content as a JSON object or a JSON string.
	//
	// > When you specify the resource policy content as a JSON string, you can't perform drift detection on the CloudFormation stack. For this reason, we recommend specifying the resource policy content as a JSON object instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-resourcepolicy.html#cfn-organizations-resourcepolicy-content
	//
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// A list of tags that you want to attach to the newly created resource policy.
	//
	// For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to `null` . For more information about tagging, see [Tagging AWS Organizations resources](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html) in the *AWS Organizations User Guide* .
	//
	// > If any one of the tags is not valid or if you exceed the allowed number of tags for the resource policy, then the entire request fails and the resource policy is not created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-resourcepolicy.html#cfn-organizations-resourcepolicy-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

