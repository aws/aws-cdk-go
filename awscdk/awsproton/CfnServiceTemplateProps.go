package awsproton

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
	// A description of the service template.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The service template name as displayed in the developer interface.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The customer provided service template encryption key that's used to encrypt data.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The name of the service template.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// If `pipelineProvisioning` is `true` , a service pipeline is included in the service template.
	//
	// Otherwise, a service pipeline *isn't* included in the service template.
	PipelineProvisioning *string `field:"optional" json:"pipelineProvisioning" yaml:"pipelineProvisioning"`
	// An object that includes the template bundle S3 bucket path and name for the new version of a service template.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

