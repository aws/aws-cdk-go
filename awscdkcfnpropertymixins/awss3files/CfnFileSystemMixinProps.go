package awss3files

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFileSystemPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnFileSystemMixinProps := &CfnFileSystemMixinProps{
//   	AcceptBucketWarning: jsii.Boolean(false),
//   	Bucket: jsii.String("bucket"),
//   	ClientToken: jsii.String("clientToken"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Prefix: jsii.String("prefix"),
//   	RoleArn: jsii.String("roleArn"),
//   	SynchronizationConfiguration: &SynchronizationConfigurationProperty{
//   		ExpirationDataRules: []interface{}{
//   			&ExpirationDataRuleProperty{
//   				DaysAfterLastAccess: jsii.Number(123),
//   			},
//   		},
//   		ImportDataRules: []interface{}{
//   			&ImportDataRuleProperty{
//   				Prefix: jsii.String("prefix"),
//   				SizeLessThan: jsii.Number(123),
//   				Trigger: jsii.String("trigger"),
//   			},
//   		},
//   		LatestVersionNumber: jsii.Number(123),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystem.html
//
type CfnFileSystemMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystem.html#cfn-s3files-filesystem-acceptbucketwarning
	//
	AcceptBucketWarning interface{} `field:"optional" json:"acceptBucketWarning" yaml:"acceptBucketWarning"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystem.html#cfn-s3files-filesystem-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystem.html#cfn-s3files-filesystem-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystem.html#cfn-s3files-filesystem-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystem.html#cfn-s3files-filesystem-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystem.html#cfn-s3files-filesystem-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystem.html#cfn-s3files-filesystem-synchronizationconfiguration
	//
	SynchronizationConfiguration interface{} `field:"optional" json:"synchronizationConfiguration" yaml:"synchronizationConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-filesystem.html#cfn-s3files-filesystem-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

