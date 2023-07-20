package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainProps := &CfnDomainProps{
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	DeadLetterQueueUrl: jsii.String("deadLetterQueueUrl"),
//   	DefaultEncryptionKey: jsii.String("defaultEncryptionKey"),
//   	DefaultExpirationDays: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html
//
type CfnDomainProps struct {
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The URL of the SQS dead letter queue, which is used for reporting errors associated with ingesting data from third party applications.
	//
	// You must set up a policy on the DeadLetterQueue for the SendMessage operation to enable Amazon Connect Customer Profiles to send messages to the DeadLetterQueue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-deadletterqueueurl
	//
	DeadLetterQueueUrl *string `field:"optional" json:"deadLetterQueueUrl" yaml:"deadLetterQueueUrl"`
	// The default encryption key, which is an AWS managed key, is used when no specific type of encryption key is specified.
	//
	// It is used to encrypt all data before it is placed in permanent or semi-permanent storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-defaultencryptionkey
	//
	DefaultEncryptionKey *string `field:"optional" json:"defaultEncryptionKey" yaml:"defaultEncryptionKey"`
	// The default number of days until the data within the domain expires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-defaultexpirationdays
	//
	DefaultExpirationDays *float64 `field:"optional" json:"defaultExpirationDays" yaml:"defaultExpirationDays"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-domain.html#cfn-customerprofiles-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

