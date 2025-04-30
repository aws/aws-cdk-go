package awsappsync


// Attributes for Event API imports.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventApiAttributes := &EventApiAttributes{
//   	ApiId: jsii.String("apiId"),
//   	HttpDns: jsii.String("httpDns"),
//   	RealtimeDns: jsii.String("realtimeDns"),
//
//   	// the properties below are optional
//   	ApiArn: jsii.String("apiArn"),
//   	ApiName: jsii.String("apiName"),
//   	AuthProviderTypes: []appSyncAuthorizationType{
//   		awscdk.Aws_appsync.*appSyncAuthorizationType_API_KEY,
//   	},
//   }
//
type EventApiAttributes struct {
	// an unique AWS AppSync Event API identifier i.e. 'lxz775lwdrgcndgz3nurvac7oa'.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// the domain name of the Api's HTTP endpoint.
	HttpDns *string `field:"required" json:"httpDns" yaml:"httpDns"`
	// the domain name of the Api's real-time endpoint.
	RealtimeDns *string `field:"required" json:"realtimeDns" yaml:"realtimeDns"`
	// the ARN of the Event API.
	// Default: - constructed arn.
	//
	ApiArn *string `field:"optional" json:"apiArn" yaml:"apiArn"`
	// the name of the Event API.
	// Default: - not needed to import API.
	//
	ApiName *string `field:"optional" json:"apiName" yaml:"apiName"`
	// The Authorization Types for this Event Api.
	// Default: - none, required to construct event rules from imported APIs.
	//
	AuthProviderTypes *[]AppSyncAuthorizationType `field:"optional" json:"authProviderTypes" yaml:"authProviderTypes"`
}

