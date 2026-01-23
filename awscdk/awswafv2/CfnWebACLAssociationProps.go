package awswafv2


// Properties for defining a `CfnWebACLAssociation`.
//
// Example:
//   var api EventApi
//   var webAcl CfnWebACL
//
//
//   // Associate waf with Event API
//   // Associate waf with Event API
//   wafv2.NewCfnWebACLAssociation(this, jsii.String("WafAssociation"), &CfnWebACLAssociationProps{
//   	ResourceArn: api.ApiArn,
//   	WebAclArn: webAcl.attrArn,
//   })
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
	// - For an AWS Verified Access instance: `arn: *partition* :ec2: *region* : *account-id* :verified-access-instance/ *instance-id*`
	// - For an AWS Amplify instance: `arn: *partition* :amplify: *region* : *account-id* :apps/ *app-id*`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html#cfn-wafv2-webaclassociation-resourcearn
	//
	ResourceArn interface{} `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html#cfn-wafv2-webaclassociation-webaclarn
	//
	WebAclArn *string `field:"required" json:"webAclArn" yaml:"webAclArn"`
}

