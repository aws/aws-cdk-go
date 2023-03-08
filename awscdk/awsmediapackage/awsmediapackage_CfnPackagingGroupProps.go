package awsmediapackage

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPackagingGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackagingGroupProps := &cfnPackagingGroupProps{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	authorization: &authorizationProperty{
//   		cdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   		secretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   	egressAccessLogs: &logConfigurationProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPackagingGroupProps struct {
	// Unique identifier that you assign to the packaging group.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Parameters for CDN authorization.
	Authorization interface{} `field:"optional" json:"authorization" yaml:"authorization"`
	// The configuration parameters for egress access logging.
	EgressAccessLogs interface{} `field:"optional" json:"egressAccessLogs" yaml:"egressAccessLogs"`
	// The tags to assign to the packaging group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

