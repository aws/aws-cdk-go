package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCustomization`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomizationProps := &CfnCustomizationProps{
//   	DefaultTheme: jsii.String("defaultTheme"),
//   	Namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-customization.html
//
type CfnCustomizationProps struct {
	// The ARN of the theme applied by default in the QuickSight console for this namespace.
	//
	// May be an AWS-managed starter theme such as arn:{Partition}:quicksight::aws:theme/MIDNIGHT or a theme owned by this account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-customization.html#cfn-quicksight-customization-defaulttheme
	//
	DefaultTheme *string `field:"required" json:"defaultTheme" yaml:"defaultTheme"`
	// The QuickSight namespace the customization applies to.
	//
	// One customization exists per (account, region, namespace), so this is create-only: changing it addresses a different resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-customization.html#cfn-quicksight-customization-namespace
	//
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Tags applied to the customization.
	//
	// QuickSight rejects any key prefixed aws: or quicksight:, so CloudFormation system tags are not propagated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-customization.html#cfn-quicksight-customization-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

