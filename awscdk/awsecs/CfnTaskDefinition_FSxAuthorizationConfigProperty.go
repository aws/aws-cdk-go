package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fSxAuthorizationConfigProperty := &FSxAuthorizationConfigProperty{
//   	CredentialsParameter: jsii.String("credentialsParameter"),
//   	Domain: jsii.String("domain"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-fsxauthorizationconfig.html
//
type CfnTaskDefinition_FSxAuthorizationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-fsxauthorizationconfig.html#cfn-ecs-taskdefinition-fsxauthorizationconfig-credentialsparameter
	//
	CredentialsParameter *string `field:"required" json:"credentialsParameter" yaml:"credentialsParameter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-fsxauthorizationconfig.html#cfn-ecs-taskdefinition-fsxauthorizationconfig-domain
	//
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

