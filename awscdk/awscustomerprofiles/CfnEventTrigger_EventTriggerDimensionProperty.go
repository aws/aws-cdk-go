package awscustomerprofiles


// A specific event dimension to be assessed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventTriggerDimensionProperty := &EventTriggerDimensionProperty{
//   	ObjectAttributes: []interface{}{
//   		&ObjectAttributeProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//
//   			// the properties below are optional
//   			FieldName: jsii.String("fieldName"),
//   			Source: jsii.String("source"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggerdimension.html
//
type CfnEventTrigger_EventTriggerDimensionProperty struct {
	// A list of object attributes to be evaluated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggerdimension.html#cfn-customerprofiles-eventtrigger-eventtriggerdimension-objectattributes
	//
	ObjectAttributes interface{} `field:"required" json:"objectAttributes" yaml:"objectAttributes"`
}

