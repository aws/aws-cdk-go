package awsrtbfabric


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   moduleConfigurationProperty := &ModuleConfigurationProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DependsOn: []*string{
//   		jsii.String("dependsOn"),
//   	},
//   	ModuleParameters: &ModuleParametersProperty{
//   		NoBid: &NoBidModuleParametersProperty{
//   			PassThroughPercentage: jsii.Number(123),
//   			Reason: jsii.String("reason"),
//   			ReasonCode: jsii.Number(123),
//   		},
//   		OpenRtbAttribute: &OpenRtbAttributeModuleParametersProperty{
//   			Action: &ActionProperty{
//   				HeaderTag: &HeaderTagActionProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   				NoBid: &NoBidActionProperty{
//   					NoBidReasonCode: jsii.Number(123),
//   				},
//   			},
//   			FilterConfiguration: []interface{}{
//   				&FilterProperty{
//   					Criteria: []interface{}{
//   						&FilterCriterionProperty{
//   							Path: jsii.String("path"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			FilterType: jsii.String("filterType"),
//   			HoldbackPercentage: jsii.Number(123),
//   		},
//   	},
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-moduleconfiguration.html
//
type CfnLink_ModuleConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-moduleconfiguration.html#cfn-rtbfabric-link-moduleconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-moduleconfiguration.html#cfn-rtbfabric-link-moduleconfiguration-dependson
	//
	DependsOn *[]*string `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-moduleconfiguration.html#cfn-rtbfabric-link-moduleconfiguration-moduleparameters
	//
	ModuleParameters interface{} `field:"optional" json:"moduleParameters" yaml:"moduleParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-moduleconfiguration.html#cfn-rtbfabric-link-moduleconfiguration-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

