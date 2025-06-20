package awsvpclattice


// Describes criteria that can be applied to incoming requests.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-httpmatch.html
//
type CfnRule_HttpMatchProperty struct {
	// The header matches.
	//
	// Matches incoming requests with rule based on request header value before applying rule action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-httpmatch.html#cfn-vpclattice-rule-httpmatch-headermatches
	//
	HeaderMatches interface{} `field:"optional" json:"headerMatches" yaml:"headerMatches"`
	// The HTTP method type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-httpmatch.html#cfn-vpclattice-rule-httpmatch-method
	//
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The path match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-httpmatch.html#cfn-vpclattice-rule-httpmatch-pathmatch
	//
	PathMatch interface{} `field:"optional" json:"pathMatch" yaml:"pathMatch"`
}

