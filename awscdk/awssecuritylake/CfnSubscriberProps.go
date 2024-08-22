package awssecuritylake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSubscriber`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriberProps := &CfnSubscriberProps{
//   	AccessTypes: []*string{
//   		jsii.String("accessTypes"),
//   	},
//   	DataLakeArn: jsii.String("dataLakeArn"),
//   	Sources: []interface{}{
//   		&SourceProperty{
//   			AwsLogSource: &AwsLogSourceProperty{
//   				SourceName: jsii.String("sourceName"),
//   				SourceVersion: jsii.String("sourceVersion"),
//   			},
//   			CustomLogSource: &CustomLogSourceProperty{
//   				SourceName: jsii.String("sourceName"),
//   				SourceVersion: jsii.String("sourceVersion"),
//   			},
//   		},
//   	},
//   	SubscriberIdentity: &SubscriberIdentityProperty{
//   		ExternalId: jsii.String("externalId"),
//   		Principal: jsii.String("principal"),
//   	},
//   	SubscriberName: jsii.String("subscriberName"),
//
//   	// the properties below are optional
//   	SubscriberDescription: jsii.String("subscriberDescription"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html
//
type CfnSubscriberProps struct {
	// You can choose to notify subscribers of new objects with an Amazon Simple Queue Service (Amazon SQS) queue or through messaging to an HTTPS endpoint provided by the subscriber.
	//
	// Subscribers can consume data by directly querying AWS Lake Formation tables in your Amazon S3 bucket through services like Amazon Athena. This subscription type is defined as `LAKEFORMATION` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-accesstypes
	//
	AccessTypes *[]*string `field:"required" json:"accessTypes" yaml:"accessTypes"`
	// The Amazon Resource Name (ARN) used to create the data lake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-datalakearn
	//
	DataLakeArn *string `field:"required" json:"dataLakeArn" yaml:"dataLakeArn"`
	// Amazon Security Lake supports log and event collection for natively supported AWS-services .
	//
	// For more information, see the [Amazon Security Lake User Guide](https://docs.aws.amazon.com//security-lake/latest/userguide/source-management.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-sources
	//
	Sources interface{} `field:"required" json:"sources" yaml:"sources"`
	// The AWS identity used to access your data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-subscriberidentity
	//
	SubscriberIdentity interface{} `field:"required" json:"subscriberIdentity" yaml:"subscriberIdentity"`
	// The name of your Amazon Security Lake subscriber account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-subscribername
	//
	SubscriberName *string `field:"required" json:"subscriberName" yaml:"subscriberName"`
	// The subscriber descriptions for a subscriber account.
	//
	// The description for a subscriber includes `subscriberName` , `accountID` , `externalID` , and `subscriberId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-subscriberdescription
	//
	SubscriberDescription *string `field:"optional" json:"subscriberDescription" yaml:"subscriberDescription"`
	// An array of objects, one for each tag to associate with the subscriber.
	//
	// For each tag, you must specify both a tag key and a tag value. A tag value cannot be null, but it can be an empty string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

