package awspanorama

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplicationInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationInstanceProps := &CfnApplicationInstanceProps{
//   	DefaultRuntimeContextDevice: jsii.String("defaultRuntimeContextDevice"),
//   	ManifestPayload: &ManifestPayloadProperty{
//   		PayloadData: jsii.String("payloadData"),
//   	},
//
//   	// the properties below are optional
//   	ApplicationInstanceIdToReplace: jsii.String("applicationInstanceIdToReplace"),
//   	Description: jsii.String("description"),
//   	ManifestOverridesPayload: &ManifestOverridesPayloadProperty{
//   		PayloadData: jsii.String("payloadData"),
//   	},
//   	Name: jsii.String("name"),
//   	RuntimeRoleArn: jsii.String("runtimeRoleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html
//
type CfnApplicationInstanceProps struct {
	// The device's ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-defaultruntimecontextdevice
	//
	DefaultRuntimeContextDevice *string `field:"required" json:"defaultRuntimeContextDevice" yaml:"defaultRuntimeContextDevice"`
	// The application's manifest document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-manifestpayload
	//
	ManifestPayload interface{} `field:"required" json:"manifestPayload" yaml:"manifestPayload"`
	// The ID of an application instance to replace with the new instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-applicationinstanceidtoreplace
	//
	ApplicationInstanceIdToReplace *string `field:"optional" json:"applicationInstanceIdToReplace" yaml:"applicationInstanceIdToReplace"`
	// A description for the application instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Setting overrides for the application manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-manifestoverridespayload
	//
	ManifestOverridesPayload interface{} `field:"optional" json:"manifestOverridesPayload" yaml:"manifestOverridesPayload"`
	// A name for the application instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of a runtime role for the application instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-runtimerolearn
	//
	RuntimeRoleArn *string `field:"optional" json:"runtimeRoleArn" yaml:"runtimeRoleArn"`
	// Tags for the application instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

