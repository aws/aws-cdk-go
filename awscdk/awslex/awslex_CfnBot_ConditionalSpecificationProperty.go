package awslex


// Provides a list of conditional branches.
//
// Branches are evaluated in the order that they are entered in the list. The first branch with a condition that evaluates to true is executed. The last branch in the list is the default branch. The default branch should not have any condition expression. The default branch is executed if no other branch has a matching condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   conditionalSpecificationProperty := &conditionalSpecificationProperty{
//   	conditionalBranches: []interface{}{
//   		&conditionalBranchProperty{
//   			condition: &conditionProperty{
//   				expressionString: jsii.String("expressionString"),
//   			},
//   			name: jsii.String("name"),
//   			nextStep: &dialogStateProperty{
//   				dialogAction: &dialogActionProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					slotToElicit: jsii.String("slotToElicit"),
//   					suppressNextMessage: jsii.Boolean(false),
//   				},
//   				intent: &intentOverrideProperty{
//   					name: jsii.String("name"),
//   					slots: []interface{}{
//   						&slotValueOverrideMapProperty{
//   							slotName: jsii.String("slotName"),
//   							slotValueOverride: &slotValueOverrideProperty{
//   								shape: jsii.String("shape"),
//   								value: &slotValueProperty{
//   									interpretedValue: jsii.String("interpretedValue"),
//   								},
//   								values: []interface{}{
//   									slotValueOverrideProperty_,
//   								},
//   							},
//   						},
//   					},
//   				},
//   				sessionAttributes: []interface{}{
//   					&sessionAttributeProperty{
//   						key: jsii.String("key"),
//
//   						// the properties below are optional
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			response: &responseSpecificationProperty{
//   				messageGroupsList: []interface{}{
//   					&messageGroupProperty{
//   						message: &messageProperty{
//   							customPayload: &customPayloadProperty{
//   								value: jsii.String("value"),
//   							},
//   							imageResponseCard: &imageResponseCardProperty{
//   								title: jsii.String("title"),
//
//   								// the properties below are optional
//   								buttons: []interface{}{
//   									&buttonProperty{
//   										text: jsii.String("text"),
//   										value: jsii.String("value"),
//   									},
//   								},
//   								imageUrl: jsii.String("imageUrl"),
//   								subtitle: jsii.String("subtitle"),
//   							},
//   							plainTextMessage: &plainTextMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   							ssmlMessage: &sSMLMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						variations: []interface{}{
//   							&messageProperty{
//   								customPayload: &customPayloadProperty{
//   									value: jsii.String("value"),
//   								},
//   								imageResponseCard: &imageResponseCardProperty{
//   									title: jsii.String("title"),
//
//   									// the properties below are optional
//   									buttons: []interface{}{
//   										&buttonProperty{
//   											text: jsii.String("text"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   									imageUrl: jsii.String("imageUrl"),
//   									subtitle: jsii.String("subtitle"),
//   								},
//   								plainTextMessage: &plainTextMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   								ssmlMessage: &sSMLMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				allowInterrupt: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	defaultBranch: &defaultConditionalBranchProperty{
//   		nextStep: &dialogStateProperty{
//   			dialogAction: &dialogActionProperty{
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				slotToElicit: jsii.String("slotToElicit"),
//   				suppressNextMessage: jsii.Boolean(false),
//   			},
//   			intent: &intentOverrideProperty{
//   				name: jsii.String("name"),
//   				slots: []interface{}{
//   					&slotValueOverrideMapProperty{
//   						slotName: jsii.String("slotName"),
//   						slotValueOverride: &slotValueOverrideProperty{
//   							shape: jsii.String("shape"),
//   							value: &slotValueProperty{
//   								interpretedValue: jsii.String("interpretedValue"),
//   							},
//   							values: []interface{}{
//   								slotValueOverrideProperty_,
//   							},
//   						},
//   					},
//   				},
//   			},
//   			sessionAttributes: []interface{}{
//   				&sessionAttributeProperty{
//   					key: jsii.String("key"),
//
//   					// the properties below are optional
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		response: &responseSpecificationProperty{
//   			messageGroupsList: []interface{}{
//   				&messageGroupProperty{
//   					message: &messageProperty{
//   						customPayload: &customPayloadProperty{
//   							value: jsii.String("value"),
//   						},
//   						imageResponseCard: &imageResponseCardProperty{
//   							title: jsii.String("title"),
//
//   							// the properties below are optional
//   							buttons: []interface{}{
//   								&buttonProperty{
//   									text: jsii.String("text"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   							imageUrl: jsii.String("imageUrl"),
//   							subtitle: jsii.String("subtitle"),
//   						},
//   						plainTextMessage: &plainTextMessageProperty{
//   							value: jsii.String("value"),
//   						},
//   						ssmlMessage: &sSMLMessageProperty{
//   							value: jsii.String("value"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					variations: []interface{}{
//   						&messageProperty{
//   							customPayload: &customPayloadProperty{
//   								value: jsii.String("value"),
//   							},
//   							imageResponseCard: &imageResponseCardProperty{
//   								title: jsii.String("title"),
//
//   								// the properties below are optional
//   								buttons: []interface{}{
//   									&buttonProperty{
//   										text: jsii.String("text"),
//   										value: jsii.String("value"),
//   									},
//   								},
//   								imageUrl: jsii.String("imageUrl"),
//   								subtitle: jsii.String("subtitle"),
//   							},
//   							plainTextMessage: &plainTextMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   							ssmlMessage: &sSMLMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			allowInterrupt: jsii.Boolean(false),
//   		},
//   	},
//   	isActive: jsii.Boolean(false),
//   }
//
type CfnBot_ConditionalSpecificationProperty struct {
	// A list of conditional branches.
	//
	// A conditional branch is made up of a condition, a response and a next step. The response and next step are executed when the condition is true.
	ConditionalBranches interface{} `field:"required" json:"conditionalBranches" yaml:"conditionalBranches"`
	// The conditional branch that should be followed when the conditions for other branches are not satisfied.
	//
	// A conditional branch is made up of a condition, a response and a next step.
	DefaultBranch interface{} `field:"required" json:"defaultBranch" yaml:"defaultBranch"`
	// Determines whether a conditional branch is active.
	//
	// When `IsActive` is false, the conditions are not evaluated.
	IsActive interface{} `field:"required" json:"isActive" yaml:"isActive"`
}

