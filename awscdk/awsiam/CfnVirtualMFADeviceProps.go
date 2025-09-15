package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVirtualMFADevice`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVirtualMFADeviceProps := &CfnVirtualMFADeviceProps{
//   	Users: []*string{
//   		jsii.String("users"),
//   	},
//
//   	// the properties below are optional
//   	Path: jsii.String("path"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualMfaDeviceName: jsii.String("virtualMfaDeviceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-virtualmfadevice.html
//
type CfnVirtualMFADeviceProps struct {
	// The IAM user associated with this virtual MFA device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-virtualmfadevice.html#cfn-iam-virtualmfadevice-users
	//
	Users *[]*string `field:"required" json:"users" yaml:"users"`
	// The path for the virtual MFA device.
	//
	// For more information about paths, see [IAM identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide* .
	//
	// This parameter is optional. If it is not included, it defaults to a slash (/).
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! ( `\u0021` ) through the DEL character ( `\u007F` ), including most punctuation characters, digits, and upper and lowercased letters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-virtualmfadevice.html#cfn-iam-virtualmfadevice-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// A list of tags that you want to attach to the new IAM virtual MFA device.
	//
	// Each tag consists of a key name and an associated value. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide* .
	//
	// > If any one of the tags is invalid or if you exceed the allowed maximum number of tags, then the entire request fails and the resource is not created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-virtualmfadevice.html#cfn-iam-virtualmfadevice-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the virtual MFA device, which must be unique.
	//
	// Use with path to uniquely identify a virtual MFA device.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-virtualmfadevice.html#cfn-iam-virtualmfadevice-virtualmfadevicename
	//
	VirtualMfaDeviceName *string `field:"optional" json:"virtualMfaDeviceName" yaml:"virtualMfaDeviceName"`
}

