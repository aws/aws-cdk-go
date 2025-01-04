package awsrbin

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuleProps := &CfnRuleProps{
//   	ResourceType: jsii.String("resourceType"),
//   	RetentionPeriod: &RetentionPeriodProperty{
//   		RetentionPeriodUnit: jsii.String("retentionPeriodUnit"),
//   		RetentionPeriodValue: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ExcludeResourceTags: []interface{}{
//   		&ResourceTagProperty{
//   			ResourceTagKey: jsii.String("resourceTagKey"),
//   			ResourceTagValue: jsii.String("resourceTagValue"),
//   		},
//   	},
//   	LockConfiguration: &UnlockDelayProperty{
//   		UnlockDelayUnit: jsii.String("unlockDelayUnit"),
//   		UnlockDelayValue: jsii.Number(123),
//   	},
//   	ResourceTags: []interface{}{
//   		&ResourceTagProperty{
//   			ResourceTagKey: jsii.String("resourceTagKey"),
//   			ResourceTagValue: jsii.String("resourceTagValue"),
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html
//
type CfnRuleProps struct {
	// The resource type to be retained by the retention rule.
	//
	// Currently, only Amazon EBS snapshots and EBS-backed AMIs are supported. To retain snapshots, specify `EBS_SNAPSHOT` . To retain EBS-backed AMIs, specify `EC2_IMAGE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Information about the retention period for which the retention rule is to retain resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-retentionperiod
	//
	RetentionPeriod interface{} `field:"required" json:"retentionPeriod" yaml:"retentionPeriod"`
	// The retention rule description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// [Region-level retention rules only] Specifies the exclusion tags to use to identify resources that are to be excluded, or ignored, by a Region-level retention rule.
	//
	// Resources that have any of these tags are not retained by the retention rule upon deletion.
	//
	// You can't specify exclusion tags for tag-level retention rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-excluderesourcetags
	//
	ExcludeResourceTags interface{} `field:"optional" json:"excludeResourceTags" yaml:"excludeResourceTags"`
	// Information about the retention rule lock configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-lockconfiguration
	//
	LockConfiguration interface{} `field:"optional" json:"lockConfiguration" yaml:"lockConfiguration"`
	// [Tag-level retention rules only] Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule.
	//
	// For tag-level retention rules, only deleted resources, of the specified resource type, that have one or more of the specified tag key and value pairs are retained. If a resource is deleted, but it does not have any of the specified tag key and value pairs, it is immediately deleted without being retained by the retention rule.
	//
	// You can add the same tag key and value pair to a maximum or five retention rules.
	//
	// To create a Region-level retention rule, omit this parameter. A Region-level retention rule does not have any resource tags specified. It retains all deleted resources of the specified resource type in the Region in which the rule is created, even if the resources are not tagged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// The state of the retention rule.
	//
	// Only retention rules that are in the `available` state retain resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Information about the tags to assign to the retention rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

