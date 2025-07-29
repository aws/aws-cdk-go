package awscustomerprofiles


// Specifies the circumstances under which the event should trigger the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventTriggerConditionProperty := &EventTriggerConditionProperty{
//   	EventTriggerDimensions: []interface{}{
//   		&EventTriggerDimensionProperty{
//   			ObjectAttributes: []interface{}{
//   				&ObjectAttributeProperty{
//   					ComparisonOperator: jsii.String("comparisonOperator"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//
//   					// the properties below are optional
//   					FieldName: jsii.String("fieldName"),
//   					Source: jsii.String("source"),
//   				},
//   			},
//   		},
//   	},
//   	LogicalOperator: jsii.String("logicalOperator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggercondition.html
//
type CfnEventTrigger_EventTriggerConditionProperty struct {
	// A list of dimensions to be evaluated for the event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggercondition.html#cfn-customerprofiles-eventtrigger-eventtriggercondition-eventtriggerdimensions
	//
	EventTriggerDimensions interface{} `field:"required" json:"eventTriggerDimensions" yaml:"eventTriggerDimensions"`
	// The operator used to combine multiple dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggercondition.html#cfn-customerprofiles-eventtrigger-eventtriggercondition-logicaloperator
	//
	LogicalOperator *string `field:"required" json:"logicalOperator" yaml:"logicalOperator"`
}

