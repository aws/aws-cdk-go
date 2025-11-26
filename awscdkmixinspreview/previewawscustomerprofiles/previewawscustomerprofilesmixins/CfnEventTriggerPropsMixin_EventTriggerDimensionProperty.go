package previewawscustomerprofilesmixins


// A specific event dimension to be assessed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventTriggerDimensionProperty := &EventTriggerDimensionProperty{
//   	ObjectAttributes: []interface{}{
//   		&ObjectAttributeProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			FieldName: jsii.String("fieldName"),
//   			Source: jsii.String("source"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggerdimension.html
//
type CfnEventTriggerPropsMixin_EventTriggerDimensionProperty struct {
	// A list of object attributes to be evaluated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggerdimension.html#cfn-customerprofiles-eventtrigger-eventtriggerdimension-objectattributes
	//
	ObjectAttributes interface{} `field:"optional" json:"objectAttributes" yaml:"objectAttributes"`
}

