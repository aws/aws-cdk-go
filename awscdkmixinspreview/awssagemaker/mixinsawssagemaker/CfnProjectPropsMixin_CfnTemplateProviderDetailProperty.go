package mixinsawssagemaker


// Details about a CloudFormation template provider configuration and associated provisioning information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTemplateProviderDetailProperty := &CfnTemplateProviderDetailProperty{
//   	Parameters: []interface{}{
//   		&CfnStackParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	TemplateName: jsii.String("templateName"),
//   	TemplateUrl: jsii.String("templateUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-cfntemplateproviderdetail.html
//
type CfnProjectPropsMixin_CfnTemplateProviderDetailProperty struct {
	// An array of CloudFormation stack parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-cfntemplateproviderdetail.html#cfn-sagemaker-project-cfntemplateproviderdetail-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The IAM role used by CloudFormation to create the stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-cfntemplateproviderdetail.html#cfn-sagemaker-project-cfntemplateproviderdetail-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The unique identifier of the template within the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-cfntemplateproviderdetail.html#cfn-sagemaker-project-cfntemplateproviderdetail-templatename
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// The Amazon S3 URL of the CloudFormation template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-cfntemplateproviderdetail.html#cfn-sagemaker-project-cfntemplateproviderdetail-templateurl
	//
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
}

