package awsappsync


// Attributes for Event API imports.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventApiAttributes := &EventApiAttributes{
//   	ApiArn: jsii.String("apiArn"),
//   	ApiId: jsii.String("apiId"),
//   	ApiName: jsii.String("apiName"),
//   	HttpDns: jsii.String("httpDns"),
//   	RealtimeDns: jsii.String("realtimeDns"),
//
//   	// the properties below are optional
//   	AuthProviderTypes: []appSyncAuthorizationType{
//   		awscdk.Aws_appsync.*appSyncAuthorizationType_API_KEY,
//   	},
//   }
//
type EventApiAttributes struct {
	// the ARN of the Event API.
	ApiArn *string `field:"required" json:"apiArn" yaml:"apiArn"`
	// an unique AWS AppSync Event API identifier i.e. 'lxz775lwdrgcndgz3nurvac7oa'.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// the name of the Event API.
	ApiName *string `field:"required" json:"apiName" yaml:"apiName"`
	// the domain name of the Api's HTTP endpoint.
	HttpDns *string `field:"required" json:"httpDns" yaml:"httpDns"`
	// the domain name of the Api's real-time endpoint.
	RealtimeDns *string `field:"required" json:"realtimeDns" yaml:"realtimeDns"`
	// The Authorization Types for this Event Api.
	// Default: - none, required to construct event rules from imported APIs.
	//
	AuthProviderTypes *[]AppSyncAuthorizationType `field:"optional" json:"authProviderTypes" yaml:"authProviderTypes"`
}

