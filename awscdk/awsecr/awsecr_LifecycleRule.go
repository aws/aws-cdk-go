package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// An ECR life cycle rule.
//
// Example:
//   var repository repository
//
//   repository.addLifecycleRule(&lifecycleRule{
//   	tagPrefixList: []*string{
//   		jsii.String("prod"),
//   	},
//   	maxImageCount: jsii.Number(9999),
//   })
//   repository.addLifecycleRule(&lifecycleRule{
//   	maxImageAge: awscdk.Duration.days(jsii.Number(30)),
//   })
//
type LifecycleRule struct {
	// Describes the purpose of the rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The maximum age of images to retain. The value must represent a number of days.
	//
	// Specify exactly one of maxImageCount and maxImageAge.
	MaxImageAge awscdk.Duration `field:"optional" json:"maxImageAge" yaml:"maxImageAge"`
	// The maximum number of images to retain.
	//
	// Specify exactly one of maxImageCount and maxImageAge.
	MaxImageCount *float64 `field:"optional" json:"maxImageCount" yaml:"maxImageCount"`
	// Controls the order in which rules are evaluated (low to high).
	//
	// All rules must have a unique priority, where lower numbers have
	// higher precedence. The first rule that matches is applied to an image.
	//
	// There can only be one rule with a tagStatus of Any, and it must have
	// the highest rulePriority.
	//
	// All rules without a specified priority will have incrementing priorities
	// automatically assigned to them, higher than any rules that DO have priorities.
	RulePriority *float64 `field:"optional" json:"rulePriority" yaml:"rulePriority"`
	// Select images that have ALL the given prefixes in their tag.
	//
	// Only if tagStatus == TagStatus.Tagged
	TagPrefixList *[]*string `field:"optional" json:"tagPrefixList" yaml:"tagPrefixList"`
	// Select images based on tags.
	//
	// Only one rule is allowed to select untagged images, and it must
	// have the highest rulePriority.
	TagStatus TagStatus `field:"optional" json:"tagStatus" yaml:"tagStatus"`
}

