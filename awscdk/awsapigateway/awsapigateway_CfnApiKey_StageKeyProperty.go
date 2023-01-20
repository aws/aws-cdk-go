package awsapigateway


// `StageKey` is a property of the [AWS::ApiGateway::ApiKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html) resource that specifies the stage to associate with the API key. This association allows only clients with the key to make requests to methods in that stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageKeyProperty := &stageKeyProperty{
//   	restApiId: jsii.String("restApiId"),
//   	stageName: jsii.String("stageName"),
//   }
//
type CfnApiKey_StageKeyProperty struct {
	// The ID of a `RestApi` resource that includes the stage with which you want to associate the API key.
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
	// The name of the stage with which to associate the API key.
	//
	// The stage must be included in the `RestApi` resource that you specified in the `RestApiId` property.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

