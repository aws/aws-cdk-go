package awsapigateway


// `ThrottleSettings` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies the overall request rate (average requests per second) and burst capacity when users call your REST APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   throttleSettingsProperty := &ThrottleSettingsProperty{
//   	BurstLimit: jsii.Number(123),
//   	RateLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html
//
type CfnUsagePlan_ThrottleSettingsProperty struct {
	// The API target request burst rate limit.
	//
	// This allows more requests through for a period of time than the target rate limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html#cfn-apigateway-usageplan-throttlesettings-burstlimit
	//
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// The API target request rate limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html#cfn-apigateway-usageplan-throttlesettings-ratelimit
	//
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

