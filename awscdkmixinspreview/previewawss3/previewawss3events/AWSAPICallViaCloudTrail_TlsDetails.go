package previewawss3events


// Type definition for TlsDetails.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn Function
//
//
//   // Works with L2 Rule
//   // Works with L2 Rule
//   events.NewRule(scope, jsii.String("Rule"), &RuleProps{
//   	EventPattern: awscdkmixinspreview.AWSAPICallViaCloudTrail_EventPattern(&AWSAPICallViaCloudTrailProps{
//   		TlsDetails: &TlsDetails{
//   			TlsVersion: []*string{
//   				jsii.String("TLSv1.3"),
//   			},
//   		},
//   		EventMetadata: &AWSEventMetadataProps{
//   			Region: []*string{
//   				jsii.String("us-east-1"),
//   			},
//   		},
//   	}),
//   	Targets: []IRuleTarget{
//   		targets.NewLambdaFunction(fn),
//   	},
//   })
//
//   // Also works with L1 CfnRule
//   // Also works with L1 CfnRule
//   events.NewCfnRule(scope, jsii.String("CfnRule"), &CfnRuleProps{
//   	State: jsii.String("ENABLED"),
//   	EventPattern: awscdkmixinspreview.AWSAPICallViaCloudTrail_*EventPattern(&AWSAPICallViaCloudTrailProps{
//   		TlsDetails: &TlsDetails{
//   			TlsVersion: []*string{
//   				jsii.String("TLSv1.3"),
//   			},
//   		},
//   		EventMetadata: &AWSEventMetadataProps{
//   			Region: []*string{
//   				jsii.String("us-east-1"),
//   			},
//   		},
//   	}),
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			Arn: fn.functionArn,
//   			Id: jsii.String("L1"),
//   		},
//   	},
//   })
//
// Experimental.
type AWSAPICallViaCloudTrail_TlsDetails struct {
	// cipherSuite property.
	//
	// Specify an array of string values to match this event if the actual value of cipherSuite is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CipherSuite *[]*string `field:"optional" json:"cipherSuite" yaml:"cipherSuite"`
	// clientProvidedHostHeader property.
	//
	// Specify an array of string values to match this event if the actual value of clientProvidedHostHeader is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientProvidedHostHeader *[]*string `field:"optional" json:"clientProvidedHostHeader" yaml:"clientProvidedHostHeader"`
	// tlsVersion property.
	//
	// Specify an array of string values to match this event if the actual value of tlsVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TlsVersion *[]*string `field:"optional" json:"tlsVersion" yaml:"tlsVersion"`
}

