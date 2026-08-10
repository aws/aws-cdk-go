package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnOpsItemPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnOpsItemMixinProps := &CfnOpsItemMixinProps{
//   	Category: jsii.String("category"),
//   	Description: jsii.String("description"),
//   	Priority: jsii.Number(123),
//   	Severity: jsii.String("severity"),
//   	Source: jsii.String("source"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Title: jsii.String("title"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html
//
type CfnOpsItemMixinProps struct {
	// The category of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-category
	//
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The description of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The priority of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The severity of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-severity
	//
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// The origin of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Tags for the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The title of the OpsItem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-opsitem.html#cfn-ssm-opsitem-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
}

