package awstextract

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAdapterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAdapterMixinProps := &CfnAdapterMixinProps{
//   	AdapterName: jsii.String("adapterName"),
//   	AutoUpdate: jsii.String("autoUpdate"),
//   	Description: jsii.String("description"),
//   	FeatureTypes: []*string{
//   		jsii.String("featureTypes"),
//   	},
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
type CfnAdapterMixinProps struct {
	// The name to be assigned to the adapter being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-textract-adapter.html#cfn-textract-adapter-adaptername
	//
	AdapterName *string `field:"optional" json:"adapterName" yaml:"adapterName"`
	// Controls whether or not the adapter should automatically update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-textract-adapter.html#cfn-textract-adapter-autoupdate
	//
	AutoUpdate *string `field:"optional" json:"autoUpdate" yaml:"autoUpdate"`
	// The description to be assigned to the adapter being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-textract-adapter.html#cfn-textract-adapter-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The type of feature that the adapter is being trained on.
	//
	// Currently, supported feature types are: QUERIES.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-textract-adapter.html#cfn-textract-adapter-featuretypes
	//
	FeatureTypes *[]*string `field:"optional" json:"featureTypes" yaml:"featureTypes"`
	// A list of tags to be added to the adapter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-textract-adapter.html#cfn-textract-adapter-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

