package awslambda


// Controls how your Lambda function handles multi-tenant execution environments.
//
// Enable this setting to configure specific tenant isolation strategies for your function invocations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tenancyConfigProperty := &TenancyConfigProperty{
//   	TenantIsolationMode: jsii.String("tenantIsolationMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-tenancyconfig.html
//
type CfnFunction_TenancyConfigProperty struct {
	// Determines how your Lambda function isolates execution environments between tenants.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-tenancyconfig.html#cfn-lambda-function-tenancyconfig-tenantisolationmode
	//
	TenantIsolationMode *string `field:"required" json:"tenantIsolationMode" yaml:"tenantIsolationMode"`
}

