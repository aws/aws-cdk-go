package awsforecast

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDatasetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetGroupProps := &CfnDatasetGroupProps{
//   	DatasetGroupName: jsii.String("datasetGroupName"),
//   	Domain: jsii.String("domain"),
//
//   	// the properties below are optional
//   	DatasetArns: []*string{
//   		jsii.String("datasetArns"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDatasetGroupProps struct {
	// The name of the dataset group.
	DatasetGroupName *string `field:"required" json:"datasetGroupName" yaml:"datasetGroupName"`
	// The domain associated with the dataset group.
	//
	// When you add a dataset to a dataset group, this value and the value specified for the `Domain` parameter of the [CreateDataset](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html) operation must match.
	//
	// The `Domain` and `DatasetType` that you choose determine the fields that must be present in training data that you import to a dataset. For example, if you choose the `RETAIL` domain and `TARGET_TIME_SERIES` as the `DatasetType` , Amazon Forecast requires that `item_id` , `timestamp` , and `demand` fields are present in your data. For more information, see [Dataset groups](https://docs.aws.amazon.com/forecast/latest/dg/howitworks-datasets-groups.html) .
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.
	DatasetArns *[]*string `field:"optional" json:"datasetArns" yaml:"datasetArns"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

