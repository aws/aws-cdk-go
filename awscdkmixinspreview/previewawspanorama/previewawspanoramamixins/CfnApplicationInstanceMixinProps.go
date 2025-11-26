package previewawspanoramamixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApplicationInstancePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationInstanceMixinProps := &CfnApplicationInstanceMixinProps{
//   	ApplicationInstanceIdToReplace: jsii.String("applicationInstanceIdToReplace"),
//   	DefaultRuntimeContextDevice: jsii.String("defaultRuntimeContextDevice"),
//   	Description: jsii.String("description"),
//   	ManifestOverridesPayload: &ManifestOverridesPayloadProperty{
//   		PayloadData: jsii.String("payloadData"),
//   	},
//   	ManifestPayload: &ManifestPayloadProperty{
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
type CfnApplicationInstanceMixinProps struct {
	// The ID of an application instance to replace with the new instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-applicationinstanceidtoreplace
	//
	ApplicationInstanceIdToReplace *string `field:"optional" json:"applicationInstanceIdToReplace" yaml:"applicationInstanceIdToReplace"`
	// The device's ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-defaultruntimecontextdevice
	//
	DefaultRuntimeContextDevice *string `field:"optional" json:"defaultRuntimeContextDevice" yaml:"defaultRuntimeContextDevice"`
	// A description for the application instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Setting overrides for the application manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-manifestoverridespayload
	//
	ManifestOverridesPayload interface{} `field:"optional" json:"manifestOverridesPayload" yaml:"manifestOverridesPayload"`
	// The application's manifest document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-manifestpayload
	//
	ManifestPayload interface{} `field:"optional" json:"manifestPayload" yaml:"manifestPayload"`
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

