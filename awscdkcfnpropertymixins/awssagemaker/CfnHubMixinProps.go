package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnHubPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnHubMixinProps := &CfnHubMixinProps{
//   	HubDescription: jsii.String("hubDescription"),
//   	HubDisplayName: jsii.String("hubDisplayName"),
//   	HubName: jsii.String("hubName"),
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
type CfnHubMixinProps struct {
	// A description of the hub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-hub.html#cfn-sagemaker-hub-hubdescription
	//
	HubDescription *string `field:"optional" json:"hubDescription" yaml:"hubDescription"`
	// The display name of the hub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-hub.html#cfn-sagemaker-hub-hubdisplayname
	//
	HubDisplayName *string `field:"optional" json:"hubDisplayName" yaml:"hubDisplayName"`
	// The name of the hub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-hub.html#cfn-sagemaker-hub-hubname
	//
	HubName *string `field:"optional" json:"hubName" yaml:"hubName"`
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

