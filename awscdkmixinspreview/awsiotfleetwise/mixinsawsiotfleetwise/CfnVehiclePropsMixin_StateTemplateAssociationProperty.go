package mixinsawsiotfleetwise


// The state template associated with a vehicle.
//
// State templates contain state properties, which are signals that belong to a signal catalog that is synchronized between the AWS IoT FleetWise Edge and the AWS Cloud .
//
// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var onChange interface{}
//
//   stateTemplateAssociationProperty := &StateTemplateAssociationProperty{
//   	Identifier: jsii.String("identifier"),
//   	StateTemplateUpdateStrategy: &StateTemplateUpdateStrategyProperty{
//   		OnChange: onChange,
//   		Periodic: &PeriodicStateTemplateUpdateStrategyProperty{
//   			StateTemplateUpdateRate: &TimePeriodProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-vehicle-statetemplateassociation.html
//
type CfnVehiclePropsMixin_StateTemplateAssociationProperty struct {
	// The unique ID of the state template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-vehicle-statetemplateassociation.html#cfn-iotfleetwise-vehicle-statetemplateassociation-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-vehicle-statetemplateassociation.html#cfn-iotfleetwise-vehicle-statetemplateassociation-statetemplateupdatestrategy
	//
	StateTemplateUpdateStrategy interface{} `field:"optional" json:"stateTemplateUpdateStrategy" yaml:"stateTemplateUpdateStrategy"`
}

