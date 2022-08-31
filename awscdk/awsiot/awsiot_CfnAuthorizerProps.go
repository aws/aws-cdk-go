package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAuthorizer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAuthorizerProps := &cfnAuthorizerProps{
//   	authorizerFunctionArn: jsii.String("authorizerFunctionArn"),
//
//   	// the properties below are optional
//   	authorizerName: jsii.String("authorizerName"),
//   	enableCachingForHttp: jsii.Boolean(false),
//   	signingDisabled: jsii.Boolean(false),
//   	status: jsii.String("status"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tokenKeyName: jsii.String("tokenKeyName"),
//   	tokenSigningPublicKeys: map[string]*string{
//   		"tokenSigningPublicKeysKey": jsii.String("tokenSigningPublicKeys"),
//   	},
//   }
//
type CfnAuthorizerProps struct {
	// The authorizer's Lambda function ARN.
	AuthorizerFunctionArn *string `field:"required" json:"authorizerFunctionArn" yaml:"authorizerFunctionArn"`
	// The authorizer name.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// `AWS::IoT::Authorizer.EnableCachingForHttp`.
	EnableCachingForHttp interface{} `field:"optional" json:"enableCachingForHttp" yaml:"enableCachingForHttp"`
	// Specifies whether AWS IoT validates the token signature in an authorization request.
	SigningDisabled interface{} `field:"optional" json:"signingDisabled" yaml:"signingDisabled"`
	// The status of the authorizer.
	//
	// Valid values: `ACTIVE` | `INACTIVE`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Metadata which can be used to manage the custom authorizer.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The key used to extract the token from the HTTP headers.
	TokenKeyName *string `field:"optional" json:"tokenKeyName" yaml:"tokenKeyName"`
	// The public keys used to validate the token signature returned by your custom authentication service.
	TokenSigningPublicKeys interface{} `field:"optional" json:"tokenSigningPublicKeys" yaml:"tokenSigningPublicKeys"`
}

