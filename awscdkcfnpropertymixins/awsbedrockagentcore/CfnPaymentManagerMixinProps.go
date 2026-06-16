package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPaymentManagerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPaymentManagerMixinProps := &CfnPaymentManagerMixinProps{
//   	AuthorizerConfiguration: &AuthorizerConfigurationProperty{
//   		CustomJwtAuthorizer: &CustomJWTAuthorizerConfigurationProperty{
//   			AllowedAudience: []*string{
//   				jsii.String("allowedAudience"),
//   			},
//   			AllowedClients: []*string{
//   				jsii.String("allowedClients"),
//   			},
//   			AllowedScopes: []*string{
//   				jsii.String("allowedScopes"),
//   			},
//   			CustomClaims: []interface{}{
//   				&CustomClaimValidationTypeProperty{
//   					AuthorizingClaimMatchValue: &AuthorizingClaimMatchValueTypeProperty{
//   						ClaimMatchOperator: jsii.String("claimMatchOperator"),
//   						ClaimMatchValue: &ClaimMatchValueTypeProperty{
//   							MatchValueString: jsii.String("matchValueString"),
//   							MatchValueStringList: []*string{
//   								jsii.String("matchValueStringList"),
//   							},
//   						},
//   					},
//   					InboundTokenClaimName: jsii.String("inboundTokenClaimName"),
//   					InboundTokenClaimValueType: jsii.String("inboundTokenClaimValueType"),
//   				},
//   			},
//   			DiscoveryUrl: jsii.String("discoveryUrl"),
//   		},
//   	},
//   	AuthorizerType: jsii.String("authorizerType"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentmanager.html
//
type CfnPaymentManagerMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentmanager.html#cfn-bedrockagentcore-paymentmanager-authorizerconfiguration
	//
	AuthorizerConfiguration interface{} `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentmanager.html#cfn-bedrockagentcore-paymentmanager-authorizertype
	//
	AuthorizerType *string `field:"optional" json:"authorizerType" yaml:"authorizerType"`
	// A description of the payment manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentmanager.html#cfn-bedrockagentcore-paymentmanager-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the payment manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentmanager.html#cfn-bedrockagentcore-paymentmanager-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of the IAM role for the payment manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentmanager.html#cfn-bedrockagentcore-paymentmanager-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags to assign to the payment manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentmanager.html#cfn-bedrockagentcore-paymentmanager-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

