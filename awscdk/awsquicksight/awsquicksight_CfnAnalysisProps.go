package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAnalysis`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnalysisProps := &cfnAnalysisProps{
//   	analysisId: jsii.String("analysisId"),
//   	awsAccountId: jsii.String("awsAccountId"),
//   	sourceEntity: &analysisSourceEntityProperty{
//   		sourceTemplate: &analysisSourceTemplateProperty{
//   			arn: jsii.String("arn"),
//   			dataSetReferences: []interface{}{
//   				&dataSetReferenceProperty{
//   					dataSetArn: jsii.String("dataSetArn"),
//   					dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	errors: []interface{}{
//   		&analysisErrorProperty{
//   			message: jsii.String("message"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	parameters: &parametersProperty{
//   		dateTimeParameters: []interface{}{
//   			&dateTimeParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		decimalParameters: []interface{}{
//   			&decimalParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		integerParameters: []interface{}{
//   			&integerParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		stringParameters: []interface{}{
//   			&stringParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	themeArn: jsii.String("themeArn"),
//   }
//
type CfnAnalysisProps struct {
	// The ID for the analysis that you're creating.
	//
	// This ID displays in the URL of the analysis.
	AnalysisId *string `field:"required" json:"analysisId" yaml:"analysisId"`
	// The ID of the AWS account where you are creating an analysis.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// A source entity to use for the analysis that you're creating.
	//
	// This metadata structure contains details that describe a source template and one or more datasets.
	SourceEntity interface{} `field:"required" json:"sourceEntity" yaml:"sourceEntity"`
	// `AWS::QuickSight::Analysis.Errors`.
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// A descriptive name for the analysis that you're creating.
	//
	// This name displays for the analysis in the Amazon QuickSight console.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameter names and override values that you want to use.
	//
	// An analysis can have any parameter type, and some parameters might accept multiple values.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A structure that describes the principals and the resource-level permissions on an analysis.
	//
	// You can use the `Permissions` structure to grant permissions by providing a list of AWS Identity and Access Management (IAM) action information for each principal listed by Amazon Resource Name (ARN).
	//
	// To specify no permissions, omit `Permissions` .
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the analysis.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN for the theme to apply to the analysis that you're creating.
	//
	// To see the theme in the Amazon QuickSight console, make sure that you have access to it.
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
}

