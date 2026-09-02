package awsdataexchange

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSetProps := &CfnDataSetProps{
//   	AssetType: jsii.String("assetType"),
//   	Description: jsii.String("description"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-dataset.html
//
type CfnDataSetProps struct {
	// The type of asset that is added to a data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-dataset.html#cfn-dataexchange-dataset-assettype
	//
	AssetType *string `field:"required" json:"assetType" yaml:"assetType"`
	// A description for the data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-dataset.html#cfn-dataexchange-dataset-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-dataset.html#cfn-dataexchange-dataset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tags for the data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-dataset.html#cfn-dataexchange-dataset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

