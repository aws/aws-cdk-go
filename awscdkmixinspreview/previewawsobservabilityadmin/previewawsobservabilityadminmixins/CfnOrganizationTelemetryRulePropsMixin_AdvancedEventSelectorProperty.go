package previewawsobservabilityadminmixins


// Advanced event selectors let you create fine-grained selectors for management, data, and network activity events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   advancedEventSelectorProperty := &AdvancedEventSelectorProperty{
//   	FieldSelectors: []interface{}{
//   		&AdvancedFieldSelectorProperty{
//   			EndsWith: []*string{
//   				jsii.String("endsWith"),
//   			},
//   			EqualTo: []*string{
//   				jsii.String("equalTo"),
//   			},
//   			Field: jsii.String("field"),
//   			NotEndsWith: []*string{
//   				jsii.String("notEndsWith"),
//   			},
//   			NotEquals: []*string{
//   				jsii.String("notEquals"),
//   			},
//   			NotStartsWith: []*string{
//   				jsii.String("notStartsWith"),
//   			},
//   			StartsWith: []*string{
//   				jsii.String("startsWith"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-advancedeventselector.html
//
type CfnOrganizationTelemetryRulePropsMixin_AdvancedEventSelectorProperty struct {
	// Contains all selector statements in an advanced event selector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-advancedeventselector.html#cfn-observabilityadmin-organizationtelemetryrule-advancedeventselector-fieldselectors
	//
	FieldSelectors interface{} `field:"optional" json:"fieldSelectors" yaml:"fieldSelectors"`
	// An optional, descriptive name for an advanced event selector, such as "Log data events for only two S3 buckets".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-advancedeventselector.html#cfn-observabilityadmin-organizationtelemetryrule-advancedeventselector-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

