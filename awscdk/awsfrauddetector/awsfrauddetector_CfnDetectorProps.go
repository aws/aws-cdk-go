package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDetector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDetectorProps := &cfnDetectorProps{
//   	detectorId: jsii.String("detectorId"),
//   	eventType: &eventTypeProperty{
//   		arn: jsii.String("arn"),
//   		createdTime: jsii.String("createdTime"),
//   		description: jsii.String("description"),
//   		entityTypes: []interface{}{
//   			&entityTypeProperty{
//   				arn: jsii.String("arn"),
//   				createdTime: jsii.String("createdTime"),
//   				description: jsii.String("description"),
//   				inline: jsii.Boolean(false),
//   				lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   				name: jsii.String("name"),
//   				tags: []cfnTag{
//   					&cfnTag{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		eventVariables: []interface{}{
//   			&eventVariableProperty{
//   				arn: jsii.String("arn"),
//   				createdTime: jsii.String("createdTime"),
//   				dataSource: jsii.String("dataSource"),
//   				dataType: jsii.String("dataType"),
//   				defaultValue: jsii.String("defaultValue"),
//   				description: jsii.String("description"),
//   				inline: jsii.Boolean(false),
//   				lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   				name: jsii.String("name"),
//   				tags: []*cfnTag{
//   					&cfnTag{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				variableType: jsii.String("variableType"),
//   			},
//   		},
//   		inline: jsii.Boolean(false),
//   		labels: []interface{}{
//   			&labelProperty{
//   				arn: jsii.String("arn"),
//   				createdTime: jsii.String("createdTime"),
//   				description: jsii.String("description"),
//   				inline: jsii.Boolean(false),
//   				lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   				name: jsii.String("name"),
//   				tags: []*cfnTag{
//   					&cfnTag{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   		name: jsii.String("name"),
//   		tags: []*cfnTag{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	rules: []interface{}{
//   		&ruleProperty{
//   			arn: jsii.String("arn"),
//   			createdTime: jsii.String("createdTime"),
//   			description: jsii.String("description"),
//   			detectorId: jsii.String("detectorId"),
//   			expression: jsii.String("expression"),
//   			language: jsii.String("language"),
//   			lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			outcomes: []interface{}{
//   				&outcomeProperty{
//   					arn: jsii.String("arn"),
//   					createdTime: jsii.String("createdTime"),
//   					description: jsii.String("description"),
//   					inline: jsii.Boolean(false),
//   					lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   					name: jsii.String("name"),
//   					tags: []*cfnTag{
//   						&cfnTag{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			ruleId: jsii.String("ruleId"),
//   			ruleVersion: jsii.String("ruleVersion"),
//   			tags: []*cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	associatedModels: []interface{}{
//   		&modelProperty{
//   			arn: jsii.String("arn"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	detectorVersionStatus: jsii.String("detectorVersionStatus"),
//   	ruleExecutionMode: jsii.String("ruleExecutionMode"),
//   	tags: []*cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDetectorProps struct {
	// The name of the detector.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The event type associated with this detector.
	EventType interface{} `field:"required" json:"eventType" yaml:"eventType"`
	// The rules to include in the detector version.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// The models to associate with this detector.
	//
	// You must provide the ARNs of all the models you want to associate.
	AssociatedModels interface{} `field:"optional" json:"associatedModels" yaml:"associatedModels"`
	// The detector description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The status of the detector version.
	//
	// If a value is not provided for this property, AWS CloudFormation assumes `DRAFT` status.
	//
	// Valid values: `ACTIVE | DRAFT`.
	DetectorVersionStatus *string `field:"optional" json:"detectorVersionStatus" yaml:"detectorVersionStatus"`
	// The rule execution mode for the rules included in the detector version.
	//
	// Valid values: `FIRST_MATCHED | ALL_MATCHED` Default value: `FIRST_MATCHED`
	//
	// You can define and edit the rule mode at the detector version level, when it is in draft status.
	//
	// If you specify `FIRST_MATCHED` , Amazon Fraud Detector evaluates rules sequentially, first to last, stopping at the first matched rule. Amazon Fraud dectector then provides the outcomes for that single rule.
	//
	// If you specifiy `ALL_MATCHED` , Amazon Fraud Detector evaluates all rules and returns the outcomes for all matched rules.
	RuleExecutionMode *string `field:"optional" json:"ruleExecutionMode" yaml:"ruleExecutionMode"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

