package mixinsawsappsync


// Use the `AwsIamConfig` property type to specify `AwsIamConfig` for a AWS AppSync authorizaton.
//
// `AwsIamConfig` is a property of the [AWS AppSync DataSource AuthorizationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-httpconfig-authorizationconfig.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   awsIamConfigProperty := &AwsIamConfigProperty{
//   	SigningRegion: jsii.String("signingRegion"),
//   	SigningServiceName: jsii.String("signingServiceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-awsiamconfig.html
//
type CfnDataSourcePropsMixin_AwsIamConfigProperty struct {
	// The signing Region for AWS Identity and Access Management authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-awsiamconfig.html#cfn-appsync-datasource-awsiamconfig-signingregion
	//
	SigningRegion *string `field:"optional" json:"signingRegion" yaml:"signingRegion"`
	// The signing service name for AWS Identity and Access Management authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-awsiamconfig.html#cfn-appsync-datasource-awsiamconfig-signingservicename
	//
	SigningServiceName *string `field:"optional" json:"signingServiceName" yaml:"signingServiceName"`
}

