package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnComputationModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var computationModelDataBindingValueProperty_ ComputationModelDataBindingValueProperty
//
//   cfnComputationModelProps := &CfnComputationModelProps{
//   	ComputationModelConfiguration: &ComputationModelConfigurationProperty{
//   		AnomalyDetection: &AnomalyDetectionComputationModelConfigurationProperty{
//   			InputProperties: jsii.String("inputProperties"),
//   			ResultProperty: jsii.String("resultProperty"),
//   		},
//   	},
//   	ComputationModelDataBinding: map[string]interface{}{
//   		"computationModelDataBindingKey": &ComputationModelDataBindingValueProperty{
//   			"assetModelProperty": &AssetModelPropertyBindingValueProperty{
//   				"assetModelId": jsii.String("assetModelId"),
//   				"propertyId": jsii.String("propertyId"),
//   			},
//   			"assetProperty": &AssetPropertyBindingValueProperty{
//   				"assetId": jsii.String("assetId"),
//   				"propertyId": jsii.String("propertyId"),
//   			},
//   			"list": []interface{}{
//   				computationModelDataBindingValueProperty_,
//   			},
//   		},
//   	},
//   	ComputationModelName: jsii.String("computationModelName"),
//
//   	// the properties below are optional
//   	ComputationModelDescription: jsii.String("computationModelDescription"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html
//
type CfnComputationModelProps struct {
	// The configuration for the computation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html#cfn-iotsitewise-computationmodel-computationmodelconfiguration
	//
	ComputationModelConfiguration interface{} `field:"required" json:"computationModelConfiguration" yaml:"computationModelConfiguration"`
	// The data binding for the computation model.
	//
	// Key is a variable name defined in configuration. Value is a `ComputationModelDataBindingValue` referenced by the variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html#cfn-iotsitewise-computationmodel-computationmodeldatabinding
	//
	ComputationModelDataBinding interface{} `field:"required" json:"computationModelDataBinding" yaml:"computationModelDataBinding"`
	// The name of the computation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html#cfn-iotsitewise-computationmodel-computationmodelname
	//
	ComputationModelName *string `field:"required" json:"computationModelName" yaml:"computationModelName"`
	// The description of the computation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html#cfn-iotsitewise-computationmodel-computationmodeldescription
	//
	ComputationModelDescription *string `field:"optional" json:"computationModelDescription" yaml:"computationModelDescription"`
	// A list of key-value pairs that contain metadata for the asset.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html#cfn-iotsitewise-computationmodel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

