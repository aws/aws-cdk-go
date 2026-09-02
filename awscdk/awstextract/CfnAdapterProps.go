package awstextract

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAdapter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAdapterProps := &CfnAdapterProps{
//   	AdapterName: jsii.String("adapterName"),
//   	FeatureTypes: []*string{
//   		jsii.String("featureTypes"),
//   	},
//
//   	// the properties below are optional
//   	AutoUpdate: jsii.String("autoUpdate"),
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-textract-adapter.html
//
type CfnAdapterProps struct {
	// The name to be assigned to the adapter being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-textract-adapter.html#cfn-textract-adapter-adaptername
	//
	AdapterName *string `field:"required" json:"adapterName" yaml:"adapterName"`
	// The type of feature that the adapter is being trained on.
	//
	// Currently, supported feature types are: QUERIES.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-textract-adapter.html#cfn-textract-adapter-featuretypes
	//
	FeatureTypes *[]*string `field:"required" json:"featureTypes" yaml:"featureTypes"`
	// Controls whether or not the adapter should automatically update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-textract-adapter.html#cfn-textract-adapter-autoupdate
	//
	AutoUpdate *string `field:"optional" json:"autoUpdate" yaml:"autoUpdate"`
	// The description to be assigned to the adapter being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-textract-adapter.html#cfn-textract-adapter-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of tags to be added to the adapter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-textract-adapter.html#cfn-textract-adapter-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

