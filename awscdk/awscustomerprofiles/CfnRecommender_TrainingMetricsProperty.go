package awscustomerprofiles


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainingMetricsProperty := &TrainingMetricsProperty{
//   	Metrics: &MetricsProperty{
//   		Coverage: jsii.Number(123),
//   		Freshness: jsii.Number(123),
//   		Hit: jsii.Number(123),
//   		Popularity: jsii.Number(123),
//   		Recall: jsii.Number(123),
//   		Similarity: jsii.Number(123),
//   	},
//   	Time: jsii.String("time"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-trainingmetrics.html
//
type CfnRecommender_TrainingMetricsProperty struct {
	// Training metrics by type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-trainingmetrics.html#cfn-customerprofiles-recommender-trainingmetrics-metrics
	//
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// Timestamp of the training metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommender-trainingmetrics.html#cfn-customerprofiles-recommender-trainingmetrics-time
	//
	Time *string `field:"optional" json:"time" yaml:"time"`
}

