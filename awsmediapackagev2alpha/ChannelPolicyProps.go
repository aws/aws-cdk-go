package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for the Channel Policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var channel Channel
//   var policyDocument PolicyDocument
//
//   channelPolicyProps := &ChannelPolicyProps{
//   	Channel: channel,
//
//   	// the properties below are optional
//   	PolicyDocument: policyDocument,
//   }
//
// Experimental.
type ChannelPolicyProps struct {
	// Channel to apply the Channel Policy to.
	// Experimental.
	Channel IChannel `field:"required" json:"channel" yaml:"channel"`
	// Initial policy document to apply.
	// Default: - empty policy document.
	//
	// Experimental.
	PolicyDocument awsiam.PolicyDocument `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

