package awspersonalize


// The training data configuration to use when creating a domain recommender.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   trainingDataConfigProperty := &TrainingDataConfigProperty{
//   	ExcludedDatasetColumns: map[string][]*string{
//   		"excludedDatasetColumnsKey": []*string{
//   			jsii.String("excludedDatasetColumns"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-recommender-trainingdataconfig.html
//
type CfnRecommenderPropsMixin_TrainingDataConfigProperty struct {
	// Specifies the columns to exclude from training.
	//
	// Each key is a dataset type, and each value is a list of columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-recommender-trainingdataconfig.html#cfn-personalize-recommender-trainingdataconfig-excludeddatasetcolumns
	//
	ExcludedDatasetColumns interface{} `field:"optional" json:"excludedDatasetColumns" yaml:"excludedDatasetColumns"`
}

