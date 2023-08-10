package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aWSManagedRulesACFPRuleSetProperty := &AWSManagedRulesACFPRuleSetProperty{
//   	CreationPath: jsii.String("creationPath"),
//   	RegistrationPagePath: jsii.String("registrationPagePath"),
//   	RequestInspection: &RequestInspectionACFPProperty{
//   		PayloadType: jsii.String("payloadType"),
//
//   		// the properties below are optional
//   		AddressFields: []interface{}{
//   			&FieldIdentifierProperty{
//   				Identifier: jsii.String("identifier"),
//   			},
//   		},
//   		EmailField: &FieldIdentifierProperty{
//   			Identifier: jsii.String("identifier"),
//   		},
//   		PasswordField: &FieldIdentifierProperty{
//   			Identifier: jsii.String("identifier"),
//   		},
//   		PhoneNumberFields: []interface{}{
//   			&FieldIdentifierProperty{
//   				Identifier: jsii.String("identifier"),
//   			},
//   		},
//   		UsernameField: &FieldIdentifierProperty{
//   			Identifier: jsii.String("identifier"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	EnableRegexInPath: jsii.Boolean(false),
//   	ResponseInspection: &ResponseInspectionProperty{
//   		BodyContains: &ResponseInspectionBodyContainsProperty{
//   			FailureStrings: []*string{
//   				jsii.String("failureStrings"),
//   			},
//   			SuccessStrings: []*string{
//   				jsii.String("successStrings"),
//   			},
//   		},
//   		Header: &ResponseInspectionHeaderProperty{
//   			FailureValues: []*string{
//   				jsii.String("failureValues"),
//   			},
//   			Name: jsii.String("name"),
//   			SuccessValues: []*string{
//   				jsii.String("successValues"),
//   			},
//   		},
//   		Json: &ResponseInspectionJsonProperty{
//   			FailureValues: []*string{
//   				jsii.String("failureValues"),
//   			},
//   			Identifier: jsii.String("identifier"),
//   			SuccessValues: []*string{
//   				jsii.String("successValues"),
//   			},
//   		},
//   		StatusCode: &ResponseInspectionStatusCodeProperty{
//   			FailureCodes: []interface{}{
//   				jsii.Number(123),
//   			},
//   			SuccessCodes: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html
//
type CfnWebACL_AWSManagedRulesACFPRuleSetProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html#cfn-wafv2-webacl-awsmanagedrulesacfpruleset-creationpath
	//
	CreationPath *string `field:"required" json:"creationPath" yaml:"creationPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html#cfn-wafv2-webacl-awsmanagedrulesacfpruleset-registrationpagepath
	//
	RegistrationPagePath *string `field:"required" json:"registrationPagePath" yaml:"registrationPagePath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html#cfn-wafv2-webacl-awsmanagedrulesacfpruleset-requestinspection
	//
	RequestInspection interface{} `field:"required" json:"requestInspection" yaml:"requestInspection"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html#cfn-wafv2-webacl-awsmanagedrulesacfpruleset-enableregexinpath
	//
	EnableRegexInPath interface{} `field:"optional" json:"enableRegexInPath" yaml:"enableRegexInPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html#cfn-wafv2-webacl-awsmanagedrulesacfpruleset-responseinspection
	//
	ResponseInspection interface{} `field:"optional" json:"responseInspection" yaml:"responseInspection"`
}

