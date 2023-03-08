package awspanorama

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnApplicationInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationInstanceProps := &cfnApplicationInstanceProps{
//   	defaultRuntimeContextDevice: jsii.String("defaultRuntimeContextDevice"),
//   	manifestPayload: &manifestPayloadProperty{
//   		payloadData: jsii.String("payloadData"),
//   	},
//
//   	// the properties below are optional
//   	applicationInstanceIdToReplace: jsii.String("applicationInstanceIdToReplace"),
//   	description: jsii.String("description"),
//   	deviceId: jsii.String("deviceId"),
//   	manifestOverridesPayload: &manifestOverridesPayloadProperty{
//   		payloadData: jsii.String("payloadData"),
//   	},
//   	name: jsii.String("name"),
//   	runtimeRoleArn: jsii.String("runtimeRoleArn"),
//   	statusFilter: jsii.String("statusFilter"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnApplicationInstanceProps struct {
	// The device's ID.
	DefaultRuntimeContextDevice *string `field:"required" json:"defaultRuntimeContextDevice" yaml:"defaultRuntimeContextDevice"`
	// The application's manifest document.
	ManifestPayload interface{} `field:"required" json:"manifestPayload" yaml:"manifestPayload"`
	// The ID of an application instance to replace with the new instance.
	ApplicationInstanceIdToReplace *string `field:"optional" json:"applicationInstanceIdToReplace" yaml:"applicationInstanceIdToReplace"`
	// A description for the application instance.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A device's ID.
	DeviceId *string `field:"optional" json:"deviceId" yaml:"deviceId"`
	// Setting overrides for the application manifest.
	ManifestOverridesPayload interface{} `field:"optional" json:"manifestOverridesPayload" yaml:"manifestOverridesPayload"`
	// A name for the application instance.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of a runtime role for the application instance.
	RuntimeRoleArn *string `field:"optional" json:"runtimeRoleArn" yaml:"runtimeRoleArn"`
	// Only include instances with a specific status.
	StatusFilter *string `field:"optional" json:"statusFilter" yaml:"statusFilter"`
	// Tags for the application instance.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

