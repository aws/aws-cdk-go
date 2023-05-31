package awsproton

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
	// A description of the environment template.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the environment template as displayed in the developer interface.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The customer provided encryption key for the environment template.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The name of the environment template.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// When included, indicates that the environment template is for customer provisioned and managed infrastructure.
	Provisioning *string `field:"optional" json:"provisioning" yaml:"provisioning"`
	// An optional list of metadata items that you can associate with the AWS Proton environment template.
	//
	// A tag is a key-value pair.
	//
	// For more information, see [AWS Proton resources and tagging](https://docs.aws.amazon.com/proton/latest/userguide/resources.html) in the *AWS Proton User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

