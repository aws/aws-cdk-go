package awswafv2


// Properties for defining a `CfnWebACLAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebACLAssociationProps := &CfnWebACLAssociationProps{
//   	ResourceArn: jsii.String("resourceArn"),
//   	WebAclArn: jsii.String("webAclArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html
//
type CfnWebACLAssociationProps struct {
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL.
	//
	// The ARN must be in one of the following formats:
	//
	// - For an Application Load Balancer: `arn: *partition* :elasticloadbalancing: *region* : *account-id* :loadbalancer/app/ *load-balancer-name* / *load-balancer-id*`
	// - For an Amazon API Gateway REST API: `arn: *partition* :apigateway: *region* ::/restapis/ *api-id* /stages/ *stage-name*`
	// - For an AWS AppSync GraphQL API: `arn: *partition* :appsync: *region* : *account-id* :apis/ *GraphQLApiId*`
	// - For an Amazon Cognito user pool: `arn: *partition* :cognito-idp: *region* : *account-id* :userpool/ *user-pool-id*`
	// - For an AWS App Runner service: `arn: *partition* :apprunner: *region* : *account-id* :service/ *apprunner-service-name* / *apprunner-service-id*`
	// - For an AWS Verified Access instance: `arn: *partition* :ec2: *region* : *account-id* :verified-access-instance/ *instance-id*`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html#cfn-wafv2-webaclassociation-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html#cfn-wafv2-webaclassociation-webaclarn
	//
	WebAclArn *string `field:"required" json:"webAclArn" yaml:"webAclArn"`
}

