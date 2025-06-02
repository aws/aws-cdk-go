package awsmediapackage

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPackagingGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackagingGroupProps := &CfnPackagingGroupProps{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	Authorization: &AuthorizationProperty{
//   		CdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   		SecretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   	EgressAccessLogs: &LogConfigurationProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html
//
type CfnPackagingGroupProps struct {
	// Unique identifier that you assign to the packaging group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html#cfn-mediapackage-packaginggroup-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// Parameters for CDN authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html#cfn-mediapackage-packaginggroup-authorization
	//
	Authorization interface{} `field:"optional" json:"authorization" yaml:"authorization"`
	// The configuration parameters for egress access logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html#cfn-mediapackage-packaginggroup-egressaccesslogs
	//
	EgressAccessLogs interface{} `field:"optional" json:"egressAccessLogs" yaml:"egressAccessLogs"`
	// The tags to assign to the packaging group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html#cfn-mediapackage-packaginggroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

