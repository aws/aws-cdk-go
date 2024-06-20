package awscodepipeline


// The configuration that specifies the result, such as rollback, to occur upon stage failure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failureConditionsProperty := &FailureConditionsProperty{
//   	Result: jsii.String("result"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-failureconditions.html
//
type CfnPipeline_FailureConditionsProperty struct {
	// The specified result for when the failure conditions are met, such as rolling back the stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-failureconditions.html#cfn-codepipeline-pipeline-failureconditions-result
	//
	Result *string `field:"required" json:"result" yaml:"result"`
}

