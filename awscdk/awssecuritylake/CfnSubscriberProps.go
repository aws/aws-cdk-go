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
	// The Amazon S3 or AWS Lake Formation access type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-accesstypes
	//
	AccessTypes *[]*string `field:"required" json:"accessTypes" yaml:"accessTypes"`
	// The ARN for the data lake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-datalakearn
	//
	DataLakeArn *string `field:"required" json:"dataLakeArn" yaml:"dataLakeArn"`
	// The supported AWS services from which logs and events are collected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-sources
	//
	Sources interface{} `field:"required" json:"sources" yaml:"sources"`
	// The AWS identity used to access your data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-subscriberidentity
	//
	SubscriberIdentity interface{} `field:"required" json:"subscriberIdentity" yaml:"subscriberIdentity"`
	// The name of your Security Lake subscriber account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-subscribername
	//
	SubscriberName *string `field:"required" json:"subscriberName" yaml:"subscriberName"`
	// The description for your subscriber account in Security Lake.
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

