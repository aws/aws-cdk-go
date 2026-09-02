package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnWebhookPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnWebhookMixinProps := &CfnWebhookMixinProps{
//   	AppId: jsii.String("appId"),
//   	BranchName: jsii.String("branchName"),
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-webhook.html
//
type CfnWebhookMixinProps struct {
	// The unique ID for an Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-webhook.html#cfn-amplify-webhook-appid
	//
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// The name for a branch that is part of an Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-webhook.html#cfn-amplify-webhook-branchname
	//
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// The description for a webhook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-webhook.html#cfn-amplify-webhook-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags for the webhook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-webhook.html#cfn-amplify-webhook-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

