package awsmgn


// An AWS Systems Manager Document to execute as a post launch action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmDocumentProperty := &SsmDocumentProperty{
//   	ActionName: jsii.String("actionName"),
//   	SsmDocumentName: jsii.String("ssmDocumentName"),
//
//   	// the properties below are optional
//   	ExternalParameters: map[string]interface{}{
//   		"externalParametersKey": &SsmExternalParameterProperty{
//   			"dynamicPath": jsii.String("dynamicPath"),
//   		},
//   	},
//   	MustSucceedForCutover: jsii.Boolean(false),
//   	Parameters: map[string]interface{}{
//   		"parametersKey": []interface{}{
//   			&SsmParameterStoreParameterProperty{
//   				"parameterName": jsii.String("parameterName"),
//   				"parameterType": jsii.String("parameterType"),
//   			},
//   		},
//   	},
//   	TimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmdocument.html
//
type CfnLaunchConfigurationTemplate_SsmDocumentProperty struct {
	// User-friendly name for the AWS Systems Manager Document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmdocument.html#cfn-mgn-launchconfigurationtemplate-ssmdocument-actionname
	//
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// AWS Systems Manager Document name or full ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmdocument.html#cfn-mgn-launchconfigurationtemplate-ssmdocument-ssmdocumentname
	//
	SsmDocumentName *string `field:"required" json:"ssmDocumentName" yaml:"ssmDocumentName"`
	// AWS Systems Manager Document external parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmdocument.html#cfn-mgn-launchconfigurationtemplate-ssmdocument-externalparameters
	//
	ExternalParameters interface{} `field:"optional" json:"externalParameters" yaml:"externalParameters"`
	// Whether Cutover is blocked when the document has failed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmdocument.html#cfn-mgn-launchconfigurationtemplate-ssmdocument-mustsucceedforcutover
	//
	MustSucceedForCutover interface{} `field:"optional" json:"mustSucceedForCutover" yaml:"mustSucceedForCutover"`
	// AWS Systems Manager Document parameters, each resolved from AWS Systems Manager Parameter Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmdocument.html#cfn-mgn-launchconfigurationtemplate-ssmdocument-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// AWS Systems Manager Document timeout, in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmdocument.html#cfn-mgn-launchconfigurationtemplate-ssmdocument-timeoutseconds
	//
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

