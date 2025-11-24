package mixinsawsrekognition

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCollectionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCollectionMixinProps := &CfnCollectionMixinProps{
//   	CollectionId: jsii.String("collectionId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rekognition-collection.html
//
type CfnCollectionMixinProps struct {
	// ID for the collection that you are creating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rekognition-collection.html#cfn-rekognition-collection-collectionid
	//
	CollectionId *string `field:"optional" json:"collectionId" yaml:"collectionId"`
	// A set of tags (key-value pairs) that you want to attach to the collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rekognition-collection.html#cfn-rekognition-collection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

