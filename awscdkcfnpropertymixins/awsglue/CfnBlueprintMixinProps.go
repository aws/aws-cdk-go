package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnBlueprintPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnBlueprintMixinProps := &CfnBlueprintMixinProps{
//   	BlueprintLocation: jsii.String("blueprintLocation"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-blueprint.html
//
type CfnBlueprintMixinProps struct {
	// Specifies a path in Amazon S3 where the blueprint is published.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-blueprint.html#cfn-glue-blueprint-blueprintlocation
	//
	BlueprintLocation *string `field:"optional" json:"blueprintLocation" yaml:"blueprintLocation"`
	// A description of the blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-blueprint.html#cfn-glue-blueprint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-blueprint.html#cfn-glue-blueprint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags to be applied to this blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-blueprint.html#cfn-glue-blueprint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

