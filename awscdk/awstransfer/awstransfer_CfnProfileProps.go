package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProfileProps := &cfnProfileProps{
//   	as2Id: jsii.String("as2Id"),
//   	profileType: jsii.String("profileType"),
//
//   	// the properties below are optional
//   	certificateIds: []*string{
//   		jsii.String("certificateIds"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProfileProps struct {
	// `AWS::Transfer::Profile.As2Id`.
	As2Id *string `field:"required" json:"as2Id" yaml:"as2Id"`
	// `AWS::Transfer::Profile.ProfileType`.
	ProfileType *string `field:"required" json:"profileType" yaml:"profileType"`
	// `AWS::Transfer::Profile.CertificateIds`.
	CertificateIds *[]*string `field:"optional" json:"certificateIds" yaml:"certificateIds"`
	// `AWS::Transfer::Profile.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

