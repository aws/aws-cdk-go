package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpMatchProperty := &httpMatchProperty{
//   	headerMatches: []interface{}{
//   		&headerMatchProperty{
//   			match: &headerMatchTypeProperty{
//   				contains: jsii.String("contains"),
//   				exact: jsii.String("exact"),
//   				prefix: jsii.String("prefix"),
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			caseSensitive: jsii.Boolean(false),
//   		},
//   	},
//   	method: jsii.String("method"),
//   	pathMatch: &pathMatchProperty{
//   		match: &pathMatchTypeProperty{
//   			exact: jsii.String("exact"),
//   			prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		caseSensitive: jsii.Boolean(false),
//   	},
//   }
//
type CfnRule_HttpMatchProperty struct {
	// `CfnRule.HttpMatchProperty.HeaderMatches`.
	HeaderMatches interface{} `field:"optional" json:"headerMatches" yaml:"headerMatches"`
	// `CfnRule.HttpMatchProperty.Method`.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// `CfnRule.HttpMatchProperty.PathMatch`.
	PathMatch interface{} `field:"optional" json:"pathMatch" yaml:"pathMatch"`
}

