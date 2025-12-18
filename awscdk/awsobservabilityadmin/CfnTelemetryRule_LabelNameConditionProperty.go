package awsobservabilityadmin


// Condition that matches based on WAF rule labels, with label names limited to 1024 characters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelNameConditionProperty := &LabelNameConditionProperty{
//   	LabelName: jsii.String("labelName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-labelnamecondition.html
//
type CfnTelemetryRule_LabelNameConditionProperty struct {
	// The label name to match, supporting alphanumeric characters, underscores, hyphens, and colons.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-labelnamecondition.html#cfn-observabilityadmin-telemetryrule-labelnamecondition-labelname
	//
	LabelName *string `field:"optional" json:"labelName" yaml:"labelName"`
}

