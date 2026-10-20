package awsdevopsagent


// External identity provider configuration for the Operator App.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   idpAuthConfigurationProperty := &IdpAuthConfigurationProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	CreatedAt: jsii.String("createdAt"),
//   	IssuerUrl: jsii.String("issuerUrl"),
//   	OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//   	Provider: jsii.String("provider"),
//   	UpdatedAt: jsii.String("updatedAt"),
//   	UserIdClaim: jsii.String("userIdClaim"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html
//
type CfnAgentSpacePropsMixin_IdpAuthConfigurationProperty struct {
	// The OIDC client identifier for the external identity provider application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The OIDC client secret for the external identity provider application.
	//
	// Use a Secrets Manager dynamic reference to avoid storing the secret in the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The OIDC issuer URL for the external identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-issuerurl
	//
	IssuerUrl *string `field:"optional" json:"issuerUrl" yaml:"issuerUrl"`
	// The IAM role that Operator App users assume to access DevOps Agent APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-operatorapprolearn
	//
	OperatorAppRoleArn *string `field:"optional" json:"operatorAppRoleArn" yaml:"operatorAppRoleArn"`
	// The external identity provider type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-provider
	//
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-updatedat
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// The identity claim used to identify users.
	//
	// This value is selected by the service for the configured provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-useridclaim
	//
	UserIdClaim *string `field:"optional" json:"userIdClaim" yaml:"userIdClaim"`
}

