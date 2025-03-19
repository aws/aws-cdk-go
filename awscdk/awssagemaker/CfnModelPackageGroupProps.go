package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnModelPackageGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var modelPackageGroupPolicy interface{}
//
//   cfnModelPackageGroupProps := &CfnModelPackageGroupProps{
//   	ModelPackageGroupName: jsii.String("modelPackageGroupName"),
//
//   	// the properties below are optional
//   	ModelPackageGroupDescription: jsii.String("modelPackageGroupDescription"),
//   	ModelPackageGroupPolicy: modelPackageGroupPolicy,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html
//
type CfnModelPackageGroupProps struct {
	// The name of the model group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-modelpackagegroupname
	//
	ModelPackageGroupName *string `field:"required" json:"modelPackageGroupName" yaml:"modelPackageGroupName"`
	// The description for the model group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-modelpackagegroupdescription
	//
	ModelPackageGroupDescription *string `field:"optional" json:"modelPackageGroupDescription" yaml:"modelPackageGroupDescription"`
	// A resouce policy to control access to a model group.
	//
	// For information about resoure policies, see [Identity-based policies and resource-based policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html) in the *AWS Identity and Access Management User Guide.* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-modelpackagegrouppolicy
	//
	ModelPackageGroupPolicy interface{} `field:"optional" json:"modelPackageGroupPolicy" yaml:"modelPackageGroupPolicy"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

