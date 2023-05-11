package awsproton

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnServiceTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceTemplateProps := &CfnServiceTemplateProps{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	Name: jsii.String("name"),
//   	PipelineProvisioning: jsii.String("pipelineProvisioning"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnServiceTemplateProps struct {
	// `AWS::Proton::ServiceTemplate.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Proton::ServiceTemplate.DisplayName`.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// `AWS::Proton::ServiceTemplate.EncryptionKey`.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// `AWS::Proton::ServiceTemplate.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Proton::ServiceTemplate.PipelineProvisioning`.
	PipelineProvisioning *string `field:"optional" json:"pipelineProvisioning" yaml:"pipelineProvisioning"`
	// `AWS::Proton::ServiceTemplate.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

