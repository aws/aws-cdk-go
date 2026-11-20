package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMetricDataQueryDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMetricDataQueryDefinitionProps := &CfnMetricDataQueryDefinitionProps{
//   	Expression: jsii.String("expression"),
//   	Path: jsii.String("path"),
//   	Type: jsii.String("type"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricdataquerydefinition.html
//
type CfnMetricDataQueryDefinitionProps struct {
	// The metric query.
	//
	// It is stored as-is and is not validated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricdataquerydefinition.html#cfn-cloudwatch-metricdataquerydefinition-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The name of the metric data query definition.
	//
	// Use '/' to organize definitions into folders. Must be unique in the account and Region. Changing it renames the definition in place.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricdataquerydefinition.html#cfn-cloudwatch-metricdataquerydefinition-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// The query language of Expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricdataquerydefinition.html#cfn-cloudwatch-metricdataquerydefinition-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// An array of key-value pairs to apply to the metric data query definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricdataquerydefinition.html#cfn-cloudwatch-metricdataquerydefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

