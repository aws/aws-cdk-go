package awsevents


// Contains the parameters needed for you to provide custom input to a target based on one or more pieces of data extracted from the event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputTransformerProperty := &inputTransformerProperty{
//   	inputTemplate: jsii.String("inputTemplate"),
//
//   	// the properties below are optional
//   	inputPathsMap: map[string]*string{
//   		"inputPathsMapKey": jsii.String("inputPathsMap"),
//   	},
//   }
//
type CfnRule_InputTransformerProperty struct {
	// Input template where you specify placeholders that will be filled with the values of the keys from `InputPathsMap` to customize the data sent to the target.
	//
	// Enclose each `InputPathsMaps` value in brackets: < *value* > The InputTemplate must be valid JSON.
	//
	// If `InputTemplate` is a JSON object (surrounded by curly braces), the following restrictions apply:
	//
	// - The placeholder cannot be used as an object key.
	//
	// The following example shows the syntax for using `InputPathsMap` and `InputTemplate` .
	//
	// `"InputTransformer":`
	//
	// `{`
	//
	// `"InputPathsMap": {"instance": "$.detail.instance","status": "$.detail.status"},`
	//
	// `"InputTemplate": "<instance> is in state <status>"`
	//
	// `}`
	//
	// To have the `InputTemplate` include quote marks within a JSON string, escape each quote marks with a slash, as in the following example:
	//
	// `"InputTransformer":`
	//
	// `{`
	//
	// `"InputPathsMap": {"instance": "$.detail.instance","status": "$.detail.status"},`
	//
	// `"InputTemplate": "<instance> is in state \"<status>\""`
	//
	// `}`
	//
	// The `InputTemplate` can also be valid JSON with varibles in quotes or out, as in the following example:
	//
	// `"InputTransformer":`
	//
	// `{`
	//
	// `"InputPathsMap": {"instance": "$.detail.instance","status": "$.detail.status"},`
	//
	// `"InputTemplate": '{"myInstance": <instance>,"myStatus": "<instance> is in state \"<status>\""}'`
	//
	// `}`.
	InputTemplate *string `field:"required" json:"inputTemplate" yaml:"inputTemplate"`
	// Map of JSON paths to be extracted from the event.
	//
	// You can then insert these in the template in `InputTemplate` to produce the output you want to be sent to the target.
	//
	// `InputPathsMap` is an array key-value pairs, where each value is a valid JSON path. You can have as many as 100 key-value pairs. You must use JSON dot notation, not bracket notation.
	//
	// The keys cannot start with " AWS ."
	InputPathsMap interface{} `field:"optional" json:"inputPathsMap" yaml:"inputPathsMap"`
}

