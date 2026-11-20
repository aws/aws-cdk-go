package awscognito


// Token session settings that apply to every app client in the user pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tokenSessionProperty := &TokenSessionProperty{
//   	MaxPerUser: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolsessionconfigurationattachment-tokensession.html
//
type CfnUserPoolSessionConfigurationAttachmentPropsMixin_TokenSessionProperty struct {
	// The maximum number of concurrent token sessions per user across all app clients.
	//
	// When a new session would exceed this number, the least recently used session ends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolsessionconfigurationattachment-tokensession.html#cfn-cognito-userpoolsessionconfigurationattachment-tokensession-maxperuser
	//
	MaxPerUser *float64 `field:"optional" json:"maxPerUser" yaml:"maxPerUser"`
}

