package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWebhook`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebhookProps := &CfnWebhookProps{
//   	BranchName: jsii.String("branchName"),
//
//   	// the properties below are optional
//   	AppId: jsii.String("appId"),
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
type CfnWebhookProps struct {
	// The name for a branch that is part of an Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-webhook.html#cfn-amplify-webhook-branchname
	//
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// The unique ID for an Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-webhook.html#cfn-amplify-webhook-appid
	//
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// The description for a webhook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-webhook.html#cfn-amplify-webhook-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags for the webhook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-webhook.html#cfn-amplify-webhook-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

