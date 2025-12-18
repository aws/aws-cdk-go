package awsobservabilityadmin


// Parameters specific to AWS CloudTrail telemetry configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//
//   			// the properties below are optional
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-cloudtrailparameters.html
//
type CfnOrganizationTelemetryRule_CloudtrailParametersProperty struct {
	// The advanced event selectors to use for filtering AWS CloudTrail events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-cloudtrailparameters.html#cfn-observabilityadmin-organizationtelemetryrule-cloudtrailparameters-advancedeventselectors
	//
	AdvancedEventSelectors interface{} `field:"required" json:"advancedEventSelectors" yaml:"advancedEventSelectors"`
}

