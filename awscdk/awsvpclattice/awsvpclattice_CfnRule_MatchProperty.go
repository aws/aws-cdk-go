package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchProperty := &matchProperty{
//   	httpMatch: &httpMatchProperty{
//   		headerMatches: []interface{}{
//   			&headerMatchProperty{
//   				match: &headerMatchTypeProperty{
//   					contains: jsii.String("contains"),
//   					exact: jsii.String("exact"),
//   					prefix: jsii.String("prefix"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				caseSensitive: jsii.Boolean(false),
//   			},
//   		},
//   		method: jsii.String("method"),
//   		pathMatch: &pathMatchProperty{
//   			match: &pathMatchTypeProperty{
//   				exact: jsii.String("exact"),
//   				prefix: jsii.String("prefix"),
//   			},
//
//   			// the properties below are optional
//   			caseSensitive: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnRule_MatchProperty struct {
	// `CfnRule.MatchProperty.HttpMatch`.
	HttpMatch interface{} `field:"required" json:"httpMatch" yaml:"httpMatch"`
}

