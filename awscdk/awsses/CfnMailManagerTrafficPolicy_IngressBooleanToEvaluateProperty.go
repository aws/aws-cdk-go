package awsses


// The union type representing the allowed types of operands for a boolean condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressBooleanToEvaluateProperty := &IngressBooleanToEvaluateProperty{
//   	Analysis: &IngressAnalysisProperty{
//   		Analyzer: jsii.String("analyzer"),
//   		ResultField: jsii.String("resultField"),
//   	},
//   	IsInAddressList: &IngressIsInAddressListProperty{
//   		AddressLists: []*string{
//   			jsii.String("addressLists"),
//   		},
//   		Attribute: jsii.String("attribute"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressbooleantoevaluate.html
//
type CfnMailManagerTrafficPolicy_IngressBooleanToEvaluateProperty struct {
	// The structure type for a boolean condition stating the Add On ARN and its returned value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressbooleantoevaluate.html#cfn-ses-mailmanagertrafficpolicy-ingressbooleantoevaluate-analysis
	//
	Analysis interface{} `field:"optional" json:"analysis" yaml:"analysis"`
	// The structure type for a boolean condition that provides the address lists to evaluate incoming traffic on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressbooleantoevaluate.html#cfn-ses-mailmanagertrafficpolicy-ingressbooleantoevaluate-isinaddresslist
	//
	IsInAddressList interface{} `field:"optional" json:"isInAddressList" yaml:"isInAddressList"`
}

