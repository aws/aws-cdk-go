package awsapigateway


// Options to the UsagePlan.addApiKey() method.
//
// Example:
//   var usageplan usagePlan
//   var apiKey apiKey
//
//
//   usageplan.addApiKey(apiKey, &addApiKeyOptions{
//   	overrideLogicalId: jsii.String("..."),
//   })
//
type AddApiKeyOptions struct {
	// Override the CloudFormation logical id of the AWS::ApiGateway::UsagePlanKey resource.
	OverrideLogicalId *string `field:"optional" json:"overrideLogicalId" yaml:"overrideLogicalId"`
}

