package previewawsdsqlmixins


// Defines the structure for multi-Region cluster configurations, containing the witness Region and peered cluster settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiRegionPropertiesProperty := &MultiRegionPropertiesProperty{
//   	Clusters: []*string{
//   		jsii.String("clusters"),
//   	},
//   	WitnessRegion: jsii.String("witnessRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dsql-cluster-multiregionproperties.html
//
type CfnClusterPropsMixin_MultiRegionPropertiesProperty struct {
	// The set of peered clusters that form the multi-Region cluster configuration.
	//
	// Each peered cluster represents a database instance in a different Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dsql-cluster-multiregionproperties.html#cfn-dsql-cluster-multiregionproperties-clusters
	//
	Clusters *[]*string `field:"optional" json:"clusters" yaml:"clusters"`
	// The Region that serves as the witness Region for a multi-Region cluster.
	//
	// The witness Region helps maintain cluster consistency and quorum.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dsql-cluster-multiregionproperties.html#cfn-dsql-cluster-multiregionproperties-witnessregion
	//
	WitnessRegion *string `field:"optional" json:"witnessRegion" yaml:"witnessRegion"`
}

