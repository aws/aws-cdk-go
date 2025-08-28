package awsses


// The union type representing the allowed types for the left hand side of a string condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressStringToEvaluateProperty := &IngressStringToEvaluateProperty{
//   	Analysis: &IngressAnalysisProperty{
//   		Analyzer: jsii.String("analyzer"),
//   		ResultField: jsii.String("resultField"),
//   	},
//   	Attribute: jsii.String("attribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressstringtoevaluate.html
//
type CfnMailManagerTrafficPolicy_IngressStringToEvaluateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressstringtoevaluate.html#cfn-ses-mailmanagertrafficpolicy-ingressstringtoevaluate-analysis
	//
	Analysis interface{} `field:"optional" json:"analysis" yaml:"analysis"`
	// The enum type representing the allowed attribute types for a string condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressstringtoevaluate.html#cfn-ses-mailmanagertrafficpolicy-ingressstringtoevaluate-attribute
	//
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

