package previewawsfrauddetectormixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDetectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDetectorMixinProps := &CfnDetectorMixinProps{
//   	AssociatedModels: []interface{}{
//   		&ModelProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DetectorId: jsii.String("detectorId"),
//   	DetectorVersionStatus: jsii.String("detectorVersionStatus"),
//   	EventType: &EventTypeProperty{
//   		Arn: jsii.String("arn"),
//   		CreatedTime: jsii.String("createdTime"),
//   		Description: jsii.String("description"),
//   		EntityTypes: []interface{}{
//   			&EntityTypeProperty{
//   				Arn: jsii.String("arn"),
//   				CreatedTime: jsii.String("createdTime"),
//   				Description: jsii.String("description"),
//   				Inline: jsii.Boolean(false),
//   				LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   				Name: jsii.String("name"),
//   				Tags: []CfnTag{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		EventVariables: []interface{}{
//   			&EventVariableProperty{
//   				Arn: jsii.String("arn"),
//   				CreatedTime: jsii.String("createdTime"),
//   				DataSource: jsii.String("dataSource"),
//   				DataType: jsii.String("dataType"),
//   				DefaultValue: jsii.String("defaultValue"),
//   				Description: jsii.String("description"),
//   				Inline: jsii.Boolean(false),
//   				LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   				Name: jsii.String("name"),
//   				Tags: []CfnTag{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				VariableType: jsii.String("variableType"),
//   			},
//   		},
//   		Inline: jsii.Boolean(false),
//   		Labels: []interface{}{
//   			&LabelProperty{
//   				Arn: jsii.String("arn"),
//   				CreatedTime: jsii.String("createdTime"),
//   				Description: jsii.String("description"),
//   				Inline: jsii.Boolean(false),
//   				LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   				Name: jsii.String("name"),
//   				Tags: []CfnTag{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   		Name: jsii.String("name"),
//   		Tags: []CfnTag{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	RuleExecutionMode: jsii.String("ruleExecutionMode"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			Arn: jsii.String("arn"),
//   			CreatedTime: jsii.String("createdTime"),
//   			Description: jsii.String("description"),
//   			DetectorId: jsii.String("detectorId"),
//   			Expression: jsii.String("expression"),
//   			Language: jsii.String("language"),
//   			LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			Outcomes: []interface{}{
//   				&OutcomeProperty{
//   					Arn: jsii.String("arn"),
//   					CreatedTime: jsii.String("createdTime"),
//   					Description: jsii.String("description"),
//   					Inline: jsii.Boolean(false),
//   					LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   					Name: jsii.String("name"),
//   					Tags: []CfnTag{
//   						&CfnTag{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			RuleId: jsii.String("ruleId"),
//   			RuleVersion: jsii.String("ruleVersion"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html
//
type CfnDetectorMixinProps struct {
	// The models to associate with this detector.
	//
	// You must provide the ARNs of all the models you want to associate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-associatedmodels
	//
	AssociatedModels interface{} `field:"optional" json:"associatedModels" yaml:"associatedModels"`
	// The detector description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorid
	//
	DetectorId *string `field:"optional" json:"detectorId" yaml:"detectorId"`
	// The status of the detector version.
	//
	// If a value is not provided for this property, AWS CloudFormation assumes `DRAFT` status.
	//
	// Valid values: `ACTIVE | DRAFT`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorversionstatus
	//
	DetectorVersionStatus *string `field:"optional" json:"detectorVersionStatus" yaml:"detectorVersionStatus"`
	// The event type associated with this detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-eventtype
	//
	EventType interface{} `field:"optional" json:"eventType" yaml:"eventType"`
	// The rule execution mode for the rules included in the detector version.
	//
	// Valid values: `FIRST_MATCHED | ALL_MATCHED` Default value: `FIRST_MATCHED`
	//
	// You can define and edit the rule mode at the detector version level, when it is in draft status.
	//
	// If you specify `FIRST_MATCHED` , Amazon Fraud Detector evaluates rules sequentially, first to last, stopping at the first matched rule. Amazon Fraud dectector then provides the outcomes for that single rule.
	//
	// If you specifiy `ALL_MATCHED` , Amazon Fraud Detector evaluates all rules and returns the outcomes for all matched rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-ruleexecutionmode
	//
	RuleExecutionMode *string `field:"optional" json:"ruleExecutionMode" yaml:"ruleExecutionMode"`
	// The rules to include in the detector version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

