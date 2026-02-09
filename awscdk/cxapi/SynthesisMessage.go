package cxapi

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   synthesisMessage := &SynthesisMessage{
//   	Entry: &MetadataEntry{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Data: jsii.String("data"),
//   		Trace: []*string{
//   			jsii.String("trace"),
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	Level: awscdk.Cx_api.SynthesisMessageLevel_INFO,
//   }
//
// Deprecated: The official definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type SynthesisMessage struct {
	// Deprecated: The official definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Entry *cloudassemblyschema.MetadataEntry `field:"required" json:"entry" yaml:"entry"`
	// Deprecated: The official definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Deprecated: The official definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Level SynthesisMessageLevel `field:"required" json:"level" yaml:"level"`
}

