package awscustomerprofiles


// Training metrics by type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   metricsProperty := &MetricsProperty{
//   	Coverage: jsii.Number(123),
//   	Freshness: jsii.Number(123),
//   	Hit: jsii.Number(123),
//   	Popularity: jsii.Number(123),
//   	Recall: jsii.Number(123),
//   	Similarity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-metrics.html
//
type CfnRecommenderPropsMixin_MetricsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-metrics.html#cfn-customerprofiles-recommender-metrics-coverage
	//
	Coverage *float64 `field:"optional" json:"coverage" yaml:"coverage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-metrics.html#cfn-customerprofiles-recommender-metrics-freshness
	//
	Freshness *float64 `field:"optional" json:"freshness" yaml:"freshness"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-metrics.html#cfn-customerprofiles-recommender-metrics-hit
	//
	Hit *float64 `field:"optional" json:"hit" yaml:"hit"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-metrics.html#cfn-customerprofiles-recommender-metrics-popularity
	//
	Popularity *float64 `field:"optional" json:"popularity" yaml:"popularity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-metrics.html#cfn-customerprofiles-recommender-metrics-recall
	//
	Recall *float64 `field:"optional" json:"recall" yaml:"recall"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-metrics.html#cfn-customerprofiles-recommender-metrics-similarity
	//
	Similarity *float64 `field:"optional" json:"similarity" yaml:"similarity"`
}

