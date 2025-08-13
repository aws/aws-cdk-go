package awssagemaker


// Configuration for processing job outputs in Amazon SageMaker Feature Store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   featureStoreOutputProperty := &FeatureStoreOutputProperty{
//   	FeatureGroupName: jsii.String("featureGroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-featurestoreoutput.html
//
type CfnProcessingJob_FeatureStoreOutputProperty struct {
	// The name of the Amazon SageMaker FeatureGroup to use as the destination for processing job output.
	//
	// Note that your processing script is responsible for putting records into your Feature Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-featurestoreoutput.html#cfn-sagemaker-processingjob-featurestoreoutput-featuregroupname
	//
	FeatureGroupName *string `field:"required" json:"featureGroupName" yaml:"featureGroupName"`
}

