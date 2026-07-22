package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnHub`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHubProps := &CfnHubProps{
//   	HubDescription: jsii.String("hubDescription"),
//   	HubName: jsii.String("hubName"),
//
//   	// the properties below are optional
//   	HubDisplayName: jsii.String("hubDisplayName"),
//   	HubSearchKeywords: []*string{
//   		jsii.String("hubSearchKeywords"),
//   	},
//   	S3StorageConfig: &S3StorageConfigProperty{
//   		S3OutputPath: jsii.String("s3OutputPath"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-hub.html
//
type CfnHubProps struct {
	// A description of the hub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-hub.html#cfn-sagemaker-hub-hubdescription
	//
	HubDescription *string `field:"required" json:"hubDescription" yaml:"hubDescription"`
	// The name of the hub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-hub.html#cfn-sagemaker-hub-hubname
	//
	HubName *string `field:"required" json:"hubName" yaml:"hubName"`
	// The display name of the hub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-hub.html#cfn-sagemaker-hub-hubdisplayname
	//
	HubDisplayName *string `field:"optional" json:"hubDisplayName" yaml:"hubDisplayName"`
	// The searchable keywords for the hub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-hub.html#cfn-sagemaker-hub-hubsearchkeywords
	//
	HubSearchKeywords *[]*string `field:"optional" json:"hubSearchKeywords" yaml:"hubSearchKeywords"`
	// The Amazon S3 storage configuration for the hub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-hub.html#cfn-sagemaker-hub-s3storageconfig
	//
	S3StorageConfig interface{} `field:"optional" json:"s3StorageConfig" yaml:"s3StorageConfig"`
	// Tags to associate with the hub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-hub.html#cfn-sagemaker-hub-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

