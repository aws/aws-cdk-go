package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpMatchProperty := &HttpMatchProperty{
//   	HeaderMatches: []interface{}{
//   		&HeaderMatchProperty{
//   			Match: &HeaderMatchTypeProperty{
//   				Contains: jsii.String("contains"),
//   				Exact: jsii.String("exact"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			CaseSensitive: jsii.Boolean(false),
//   		},
//   	},
//   	Method: jsii.String("method"),
//   	PathMatch: &PathMatchProperty{
//   		Match: &PathMatchTypeProperty{
//   			Exact: jsii.String("exact"),
//   			Prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		CaseSensitive: jsii.Boolean(false),
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

