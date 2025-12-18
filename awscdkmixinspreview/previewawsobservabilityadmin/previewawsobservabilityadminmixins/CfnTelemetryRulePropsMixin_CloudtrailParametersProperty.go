package previewawsobservabilityadminmixins


// Parameters specific to AWS CloudTrail telemetry configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudtrailParametersProperty := &CloudtrailParametersProperty{
//   	AdvancedEventSelectors: []interface{}{
//   		&AdvancedEventSelectorProperty{
//   			FieldSelectors: []interface{}{
//   				&AdvancedFieldSelectorProperty{
//   					EndsWith: []*string{
//   						jsii.String("endsWith"),
//   					},
//   					EqualTo: []*string{
//   						jsii.String("equalTo"),
//   					},
//   					Field: jsii.String("field"),
//   					NotEndsWith: []*string{
//   						jsii.String("notEndsWith"),
//   					},
//   					NotEquals: []*string{
//   						jsii.String("notEquals"),
//   					},
//   					NotStartsWith: []*string{
//   						jsii.String("notStartsWith"),
//   					},
//   					StartsWith: []*string{
//   						jsii.String("startsWith"),
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-cloudtrailparameters.html
//
type CfnTelemetryRulePropsMixin_CloudtrailParametersProperty struct {
	// The advanced event selectors to use for filtering AWS CloudTrail events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-cloudtrailparameters.html#cfn-observabilityadmin-telemetryrule-cloudtrailparameters-advancedeventselectors
	//
	AdvancedEventSelectors interface{} `field:"optional" json:"advancedEventSelectors" yaml:"advancedEventSelectors"`
}

