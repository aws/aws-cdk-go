package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOpsItem`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOpsItemProps := &CfnOpsItemProps{
//   	Description: jsii.String("description"),
//   	Source: jsii.String("source"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	Category: jsii.String("category"),
//   	Priority: jsii.Number(123),
//   	Severity: jsii.String("severity"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html
//
type CfnOpsItemProps struct {
	// The description of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The origin of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// The title of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
	// The category of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-category
	//
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The priority of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The severity of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-severity
	//
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// Tags for the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

