package awspersonalize

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFilterProps := &CfnFilterProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	FilterExpression: jsii.String("filterExpression"),
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-filter.html
//
type CfnFilterProps struct {
	// The ARN of the dataset group that the filter belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-filter.html#cfn-personalize-filter-datasetgrouparn
	//
	DatasetGroupArn *string `field:"required" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// The filter expression that defines which items are included or excluded from recommendations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-filter.html#cfn-personalize-filter-filterexpression
	//
	FilterExpression *string `field:"required" json:"filterExpression" yaml:"filterExpression"`
	// The name of the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-filter.html#cfn-personalize-filter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tags to associate with the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-filter.html#cfn-personalize-filter-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

