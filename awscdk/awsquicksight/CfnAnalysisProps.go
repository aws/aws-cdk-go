package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAnalysis`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnalysisProps := &CfnAnalysisProps{
//   	AnalysisId: jsii.String("analysisId"),
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	SourceEntity: &AnalysisSourceEntityProperty{
//   		SourceTemplate: &AnalysisSourceTemplateProperty{
//   			Arn: jsii.String("arn"),
//   			DataSetReferences: []interface{}{
//   				&DataSetReferenceProperty{
//   					DataSetArn: jsii.String("dataSetArn"),
//   					DataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Errors: []interface{}{
//   		&AnalysisErrorProperty{
//   			Message: jsii.String("message"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Parameters: &ParametersProperty{
//   		DateTimeParameters: []interface{}{
//   			&DateTimeParameterProperty{
//   				Name: jsii.String("name"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		DecimalParameters: []interface{}{
//   			&DecimalParameterProperty{
//   				Name: jsii.String("name"),
//   				Values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		IntegerParameters: []interface{}{
//   			&IntegerParameterProperty{
//   				Name: jsii.String("name"),
//   				Values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		StringParameters: []interface{}{
//   			&StringParameterProperty{
//   				Name: jsii.String("name"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThemeArn: jsii.String("themeArn"),
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
	//
	// Either a `SourceEntity` or a `Definition` must be provided in order for the request to be valid.
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

