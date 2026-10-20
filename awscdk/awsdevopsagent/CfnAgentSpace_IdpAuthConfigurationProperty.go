package awsdevopsagent


// External identity provider configuration for the Operator App.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idpAuthConfigurationProperty := &IdpAuthConfigurationProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	IssuerUrl: jsii.String("issuerUrl"),
//   	OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//   	Provider: jsii.String("provider"),
//
//   	// the properties below are optional
//   	CreatedAt: jsii.String("createdAt"),
//   	UpdatedAt: jsii.String("updatedAt"),
//   	UserIdClaim: jsii.String("userIdClaim"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html
//
type CfnAgentSpace_IdpAuthConfigurationProperty struct {
	// The OIDC client identifier for the external identity provider application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OIDC client secret for the external identity provider application.
	//
	// Use a Secrets Manager dynamic reference to avoid storing the secret in the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-clientsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer URL for the external identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-issuerurl
	//
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// The IAM role that Operator App users assume to access DevOps Agent APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-operatorapprolearn
	//
	OperatorAppRoleArn *string `field:"required" json:"operatorAppRoleArn" yaml:"operatorAppRoleArn"`
	// The external identity provider type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-provider
	//
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idpauthconfiguration.html#cfn-devopsagent-agentspace-idpauthconfiguration-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
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

