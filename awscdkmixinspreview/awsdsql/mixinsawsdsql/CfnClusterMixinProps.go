package mixinsawsdsql

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnClusterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterMixinProps := &CfnClusterMixinProps{
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	KmsEncryptionKey: jsii.String("kmsEncryptionKey"),
//   	MultiRegionProperties: &MultiRegionPropertiesProperty{
//   		Clusters: []*string{
//   			jsii.String("clusters"),
//   		},
//   		WitnessRegion: jsii.String("witnessRegion"),
//   	},
//   	PolicyDocument: jsii.String("policyDocument"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dsql-cluster.html
//
type CfnClusterMixinProps struct {
	// Whether deletion protection is enabled on this cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dsql-cluster.html#cfn-dsql-cluster-deletionprotectionenabled
	//
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// The KMS key that encrypts data on the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dsql-cluster.html#cfn-dsql-cluster-kmsencryptionkey
	//
	KmsEncryptionKey *string `field:"optional" json:"kmsEncryptionKey" yaml:"kmsEncryptionKey"`
	// Defines the structure for multi-Region cluster configurations, containing the witness Region and peered cluster settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dsql-cluster.html#cfn-dsql-cluster-multiregionproperties
	//
	MultiRegionProperties interface{} `field:"optional" json:"multiRegionProperties" yaml:"multiRegionProperties"`
	// A resource-based policy document in JSON format.
	//
	// Length constraints: Minimum length of 1. Maximum length of 20480 characters (approximately 20KB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dsql-cluster.html#cfn-dsql-cluster-policydocument
	//
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// A map of key and value pairs this cluster is tagged with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dsql-cluster.html#cfn-dsql-cluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

