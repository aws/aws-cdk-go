package awscognito


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   issuerConfigurationProperty := &IssuerConfigurationProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-issuerconfiguration.html
//
type CfnUserPool_IssuerConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-issuerconfiguration.html#cfn-cognito-userpool-issuerconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

