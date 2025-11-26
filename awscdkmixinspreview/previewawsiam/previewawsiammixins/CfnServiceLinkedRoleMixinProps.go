package previewawsiammixins


// Properties for CfnServiceLinkedRolePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceLinkedRoleMixinProps := &CfnServiceLinkedRoleMixinProps{
//   	AwsServiceName: jsii.String("awsServiceName"),
//   	CustomSuffix: jsii.String("customSuffix"),
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-servicelinkedrole.html
//
type CfnServiceLinkedRoleMixinProps struct {
	// The service principal for the AWS service to which this role is attached.
	//
	// You use a string similar to a URL but without the http:// in front. For example: `elasticbeanstalk.amazonaws.com` .
	//
	// Service principals are unique and case-sensitive. To find the exact service principal for your service-linked role, see [AWS services that work with IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html) in the *IAM User Guide* . Look for the services that have *Yes* in the *Service-Linked Role* column. Choose the *Yes* link to view the service-linked role documentation for that service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-servicelinkedrole.html#cfn-iam-servicelinkedrole-awsservicename
	//
	AwsServiceName *string `field:"optional" json:"awsServiceName" yaml:"awsServiceName"`
	// A string that you provide, which is combined with the service-provided prefix to form the complete role name.
	//
	// If you make multiple requests for the same service, then you must supply a different `CustomSuffix` for each request. Otherwise the request fails with a duplicate role name error. For example, you could add `-1` or `-debug` to the suffix.
	//
	// Some services do not support the `CustomSuffix` parameter. If you provide an optional suffix and the operation fails, try the operation again without the suffix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-servicelinkedrole.html#cfn-iam-servicelinkedrole-customsuffix
	//
	CustomSuffix *string `field:"optional" json:"customSuffix" yaml:"customSuffix"`
	// The description of the role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-servicelinkedrole.html#cfn-iam-servicelinkedrole-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

