package awsdsql

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &CfnClusterProps{
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	KmsEncryptionKey: jsii.String("kmsEncryptionKey"),
//   	MultiRegionProperties: &MultiRegionPropertiesProperty{
//   		Clusters: []*string{
//   			jsii.String("clusters"),
//   		},
//   		WitnessRegion: jsii.String("witnessRegion"),
//   	},
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
type CfnClusterProps struct {
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
	// A map of key and value pairs this cluster is tagged with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dsql-cluster.html#cfn-dsql-cluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

