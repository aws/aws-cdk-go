package previewawsgreengrassmixins


// A list of the [resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.html) in the group that the function can access, with the corresponding read-only or read-write permissions. The maximum is 10 resources.
//
// > This property applies only to Lambda functions that run in a Greengrass container.
//
// In an CloudFormation template, `ResourceAccessPolicy` is a property of the [`Environment`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-environment.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceAccessPolicyProperty := &ResourceAccessPolicyProperty{
//   	Permission: jsii.String("permission"),
//   	ResourceId: jsii.String("resourceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-resourceaccesspolicy.html
//
type CfnFunctionDefinitionPropsMixin_ResourceAccessPolicyProperty struct {
	// The read-only or read-write access that the Lambda function has to the resource.
	//
	// Valid values are `ro` or `rw` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-resourceaccesspolicy.html#cfn-greengrass-functiondefinition-resourceaccesspolicy-permission
	//
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
	// The ID of the resource.
	//
	// This ID is assigned to the resource when you create the resource definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-resourceaccesspolicy.html#cfn-greengrass-functiondefinition-resourceaccesspolicy-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

