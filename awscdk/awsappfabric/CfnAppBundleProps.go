package awsappfabric

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAppBundle`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppBundleProps := &CfnAppBundleProps{
//   	CustomerManagedKeyIdentifier: jsii.String("customerManagedKeyIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-appbundle.html
//
type CfnAppBundleProps struct {
	// The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS) key to use to encrypt the application data.
	//
	// If this is not specified, an AWS owned key is used for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-appbundle.html#cfn-appfabric-appbundle-customermanagedkeyidentifier
	//
	CustomerManagedKeyIdentifier *string `field:"optional" json:"customerManagedKeyIdentifier" yaml:"customerManagedKeyIdentifier"`
	// A map of the key-value pairs of the tag or tags to assign to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-appbundle.html#cfn-appfabric-appbundle-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

