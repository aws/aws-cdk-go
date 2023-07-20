package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAuthorizer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAuthorizerProps := &CfnAuthorizerProps{
//   	AuthorizerFunctionArn: jsii.String("authorizerFunctionArn"),
//
//   	// the properties below are optional
//   	AuthorizerName: jsii.String("authorizerName"),
//   	EnableCachingForHttp: jsii.Boolean(false),
//   	SigningDisabled: jsii.Boolean(false),
//   	Status: jsii.String("status"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TokenKeyName: jsii.String("tokenKeyName"),
//   	TokenSigningPublicKeys: map[string]*string{
//   		"tokenSigningPublicKeysKey": jsii.String("tokenSigningPublicKeys"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html
//
type CfnAuthorizerProps struct {
	// The authorizer's Lambda function ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-authorizerfunctionarn
	//
	AuthorizerFunctionArn *string `field:"required" json:"authorizerFunctionArn" yaml:"authorizerFunctionArn"`
	// The authorizer name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-authorizername
	//
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// When `true` , the result from the authorizer's Lambda function is cached for clients that use persistent HTTP connections.
	//
	// The results are cached for the time specified by the Lambda function in `refreshAfterInSeconds` . This value doesn't affect authorization of clients that use MQTT connections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-enablecachingforhttp
	//
	EnableCachingForHttp interface{} `field:"optional" json:"enableCachingForHttp" yaml:"enableCachingForHttp"`
	// Specifies whether AWS IoT validates the token signature in an authorization request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-signingdisabled
	//
	SigningDisabled interface{} `field:"optional" json:"signingDisabled" yaml:"signingDisabled"`
	// The status of the authorizer.
	//
	// Valid values: `ACTIVE` | `INACTIVE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Metadata which can be used to manage the custom authorizer.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The key used to extract the token from the HTTP headers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-tokenkeyname
	//
	TokenKeyName *string `field:"optional" json:"tokenKeyName" yaml:"tokenKeyName"`
	// The public keys used to validate the token signature returned by your custom authentication service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-authorizer.html#cfn-iot-authorizer-tokensigningpublickeys
	//
	TokenSigningPublicKeys interface{} `field:"optional" json:"tokenSigningPublicKeys" yaml:"tokenSigningPublicKeys"`
}

