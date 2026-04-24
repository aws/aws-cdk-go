package awsses


// The string to evaluate in a string condition expression.
//
// > This data type is a UNION, so only one of the following members can be specified when used or returned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ruleStringToEvaluateProperty := &RuleStringToEvaluateProperty{
//   	Analysis: &AnalysisProperty{
//   		Analyzer: jsii.String("analyzer"),
//   		ResultField: jsii.String("resultField"),
//   	},
//   	Attribute: jsii.String("attribute"),
//   	ClientCertificateAttribute: jsii.String("clientCertificateAttribute"),
//   	MimeHeaderAttribute: jsii.String("mimeHeaderAttribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulestringtoevaluate.html
//
type CfnMailManagerRuleSetPropsMixin_RuleStringToEvaluateProperty struct {
	// The Add On ARN and its returned value to evaluate in a string condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulestringtoevaluate.html#cfn-ses-mailmanagerruleset-rulestringtoevaluate-analysis
	//
	Analysis interface{} `field:"optional" json:"analysis" yaml:"analysis"`
	// The email attribute to evaluate in a string condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulestringtoevaluate.html#cfn-ses-mailmanagerruleset-rulestringtoevaluate-attribute
	//
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulestringtoevaluate.html#cfn-ses-mailmanagerruleset-rulestringtoevaluate-clientcertificateattribute
	//
	ClientCertificateAttribute *string `field:"optional" json:"clientCertificateAttribute" yaml:"clientCertificateAttribute"`
	// The email MIME X-Header attribute to evaluate in a string condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulestringtoevaluate.html#cfn-ses-mailmanagerruleset-rulestringtoevaluate-mimeheaderattribute
	//
	MimeHeaderAttribute *string `field:"optional" json:"mimeHeaderAttribute" yaml:"mimeHeaderAttribute"`
}

