package awsfinspace

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentProps := &CfnEnvironmentProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	FederationMode: jsii.String("federationMode"),
//   	FederationParameters: &FederationParametersProperty{
//   		ApplicationCallBackUrl: jsii.String("applicationCallBackUrl"),
//   		AttributeMap: []interface{}{
//   			&AttributeMapItemsProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FederationProviderName: jsii.String("federationProviderName"),
//   		FederationUrn: jsii.String("federationUrn"),
//   		SamlMetadataDocument: jsii.String("samlMetadataDocument"),
//   		SamlMetadataUrl: jsii.String("samlMetadataUrl"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SuperuserParameters: &SuperuserParametersProperty{
//   		EmailAddress: jsii.String("emailAddress"),
//   		FirstName: jsii.String("firstName"),
//   		LastName: jsii.String("lastName"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEnvironmentProps struct {
	// The name of the FinSpace environment.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the FinSpace environment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The authentication mode for the environment.
	FederationMode *string `field:"optional" json:"federationMode" yaml:"federationMode"`
	// Configuration information when authentication mode is FEDERATED.
	FederationParameters interface{} `field:"optional" json:"federationParameters" yaml:"federationParameters"`
	// The KMS key id used to encrypt in the FinSpace environment.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Configuration information for the superuser.
	SuperuserParameters interface{} `field:"optional" json:"superuserParameters" yaml:"superuserParameters"`
	// `AWS::FinSpace::Environment.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

