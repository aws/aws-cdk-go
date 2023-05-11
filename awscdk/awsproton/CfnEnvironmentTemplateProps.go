package awsproton

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEnvironmentTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentTemplateProps := &CfnEnvironmentTemplateProps{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	Name: jsii.String("name"),
//   	Provisioning: jsii.String("provisioning"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEnvironmentTemplateProps struct {
	// `AWS::Proton::EnvironmentTemplate.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Proton::EnvironmentTemplate.DisplayName`.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// `AWS::Proton::EnvironmentTemplate.EncryptionKey`.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// `AWS::Proton::EnvironmentTemplate.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Proton::EnvironmentTemplate.Provisioning`.
	Provisioning *string `field:"optional" json:"provisioning" yaml:"provisioning"`
	// `AWS::Proton::EnvironmentTemplate.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

