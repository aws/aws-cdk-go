package awsglobalaccelerator

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCrossAccountAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCrossAccountAttachmentProps := &CfnCrossAccountAttachmentProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Principals: []*string{
//   		jsii.String("principals"),
//   	},
//   	Resources: []interface{}{
//   		&ResourceProperty{
//   			EndpointId: jsii.String("endpointId"),
//
//   			// the properties below are optional
//   			Region: jsii.String("region"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-crossaccountattachment.html
//
type CfnCrossAccountAttachmentProps struct {
	// The Friendly identifier of the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-crossaccountattachment.html#cfn-globalaccelerator-crossaccountattachment-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Principals to share the resources with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-crossaccountattachment.html#cfn-globalaccelerator-crossaccountattachment-principals
	//
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// Resources shared using the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-crossaccountattachment.html#cfn-globalaccelerator-crossaccountattachment-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-crossaccountattachment.html#cfn-globalaccelerator-crossaccountattachment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

