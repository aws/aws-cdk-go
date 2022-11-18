package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAgreement`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAgreementProps := &cfnAgreementProps{
//   	accessRole: jsii.String("accessRole"),
//   	baseDirectory: jsii.String("baseDirectory"),
//   	localProfileId: jsii.String("localProfileId"),
//   	partnerProfileId: jsii.String("partnerProfileId"),
//   	serverId: jsii.String("serverId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	status: jsii.String("status"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAgreementProps struct {
	// `AWS::Transfer::Agreement.AccessRole`.
	AccessRole *string `field:"required" json:"accessRole" yaml:"accessRole"`
	// `AWS::Transfer::Agreement.BaseDirectory`.
	BaseDirectory *string `field:"required" json:"baseDirectory" yaml:"baseDirectory"`
	// `AWS::Transfer::Agreement.LocalProfileId`.
	LocalProfileId *string `field:"required" json:"localProfileId" yaml:"localProfileId"`
	// `AWS::Transfer::Agreement.PartnerProfileId`.
	PartnerProfileId *string `field:"required" json:"partnerProfileId" yaml:"partnerProfileId"`
	// `AWS::Transfer::Agreement.ServerId`.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// `AWS::Transfer::Agreement.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Transfer::Agreement.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// `AWS::Transfer::Agreement.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

