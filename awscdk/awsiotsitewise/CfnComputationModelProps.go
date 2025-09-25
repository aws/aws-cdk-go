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
//   var computationModelDataBindingValueProperty_ computationModelDataBindingValueProperty
//
//   cfnComputationModelProps := &CfnComputationModelProps{
//   	ComputationModelConfiguration: &ComputationModelConfigurationProperty{
//   		AnomalyDetection: &AnomalyDetectionComputationModelConfigurationProperty{
//   			InputProperties: jsii.String("inputProperties"),
//   			ResultProperty: jsii.String("resultProperty"),
//   		},
//   	},
//   	ComputationModelDataBinding: map[string]interface{}{
//   		"computationModelDataBindingKey": &computationModelDataBindingValueProperty{
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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html
//
type CfnComputationModelProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html#cfn-iotsitewise-computationmodel-computationmodelconfiguration
	//
	ComputationModelConfiguration interface{} `field:"required" json:"computationModelConfiguration" yaml:"computationModelConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html#cfn-iotsitewise-computationmodel-computationmodeldatabinding
	//
	ComputationModelDataBinding interface{} `field:"required" json:"computationModelDataBinding" yaml:"computationModelDataBinding"`
	// The name of the computation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html#cfn-iotsitewise-computationmodel-computationmodelname
	//
	ComputationModelName *string `field:"required" json:"computationModelName" yaml:"computationModelName"`
	// A description about the computation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html#cfn-iotsitewise-computationmodel-computationmodeldescription
	//
	ComputationModelDescription *string `field:"optional" json:"computationModelDescription" yaml:"computationModelDescription"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html#cfn-iotsitewise-computationmodel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

