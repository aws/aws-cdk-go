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
type SynthesisMessage struct {
	Entry *cloudassemblyschema.MetadataEntry `field:"required" json:"entry" yaml:"entry"`
	Id *string `field:"required" json:"id" yaml:"id"`
	Level SynthesisMessageLevel `field:"required" json:"level" yaml:"level"`
}

