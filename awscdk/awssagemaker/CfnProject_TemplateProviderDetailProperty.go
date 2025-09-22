package awssagemaker


// Details about a template provider configuration and associated provisioning information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateProviderDetailProperty := &TemplateProviderDetailProperty{
//   	CfnTemplateProviderDetail: &CfnTemplateProviderDetailProperty{
//   		TemplateName: jsii.String("templateName"),
//   		TemplateUrl: jsii.String("templateUrl"),
//
//   		// the properties below are optional
//   		Parameters: []interface{}{
//   			&CfnStackParameterProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-templateproviderdetail.html
//
type CfnProject_TemplateProviderDetailProperty struct {
	// Details about a CloudFormation template provider configuration and associated provisioning information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-templateproviderdetail.html#cfn-sagemaker-project-templateproviderdetail-cfntemplateproviderdetail
	//
	CfnTemplateProviderDetail interface{} `field:"required" json:"cfnTemplateProviderDetail" yaml:"cfnTemplateProviderDetail"`
}

