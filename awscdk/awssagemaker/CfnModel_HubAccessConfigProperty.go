package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hubAccessConfigProperty := &HubAccessConfigProperty{
//   	HubContentArn: jsii.String("hubContentArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-hubaccessconfig.html
//
type CfnModel_HubAccessConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-hubaccessconfig.html#cfn-sagemaker-model-hubaccessconfig-hubcontentarn
	//
	HubContentArn *string `field:"required" json:"hubContentArn" yaml:"hubContentArn"`
}

