package mixinsawssagemaker


// The configuration for a private hub model reference that points to a public SageMaker JumpStart model.
//
// For more information about private hubs, see [Private curated hubs for foundation model access control in JumpStart](https://docs.aws.amazon.com/sagemaker/latest/dg/jumpstart-curated-hubs.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hubAccessConfigProperty := &HubAccessConfigProperty{
//   	HubContentArn: jsii.String("hubContentArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-hubaccessconfig.html
//
type CfnModelPropsMixin_HubAccessConfigProperty struct {
	// The ARN of your private model hub content.
	//
	// This should be a `ModelReference` resource type that points to a SageMaker JumpStart public hub model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-hubaccessconfig.html#cfn-sagemaker-model-hubaccessconfig-hubcontentarn
	//
	HubContentArn *string `field:"optional" json:"hubContentArn" yaml:"hubContentArn"`
}

