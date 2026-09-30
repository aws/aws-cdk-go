package awspersonalize

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFilterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnFilterMixinProps := &CfnFilterMixinProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	FilterExpression: jsii.String("filterExpression"),
//   	Name: jsii.String("name"),
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
type CfnFilterMixinProps struct {
	// The ARN of the dataset group that the filter belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-filter.html#cfn-personalize-filter-datasetgrouparn
	//
	DatasetGroupArn *string `field:"optional" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// The filter expression that defines which items are included or excluded from recommendations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-filter.html#cfn-personalize-filter-filterexpression
	//
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
	// The name of the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-filter.html#cfn-personalize-filter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags to associate with the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-filter.html#cfn-personalize-filter-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

