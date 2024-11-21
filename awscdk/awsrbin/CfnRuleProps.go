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
	// The resource type retained by the retention rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The retention period of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-retentionperiod
	//
	RetentionPeriod interface{} `field:"required" json:"retentionPeriod" yaml:"retentionPeriod"`
	// The description of the retention rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about the exclude resource tags used to identify resources that are excluded by the retention rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-excluderesourcetags
	//
	ExcludeResourceTags interface{} `field:"optional" json:"excludeResourceTags" yaml:"excludeResourceTags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-lockconfiguration
	//
	LockConfiguration interface{} `field:"optional" json:"lockConfiguration" yaml:"lockConfiguration"`
	// Information about the resource tags used to identify resources that are retained by the retention rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// The state of the retention rule.
	//
	// Only retention rules that are in the available state retain resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Information about the tags assigned to the retention rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rbin-rule.html#cfn-rbin-rule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

