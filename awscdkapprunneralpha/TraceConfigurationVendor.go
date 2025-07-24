package awscdkapprunneralpha


// The implementation provider chosen for tracing App Runner services.
//
// Example:
//   observabilityConfiguration := apprunner.NewObservabilityConfiguration(this, jsii.String("ObservabilityConfiguration"), &ObservabilityConfigurationProps{
//   	ObservabilityConfigurationName: jsii.String("MyObservabilityConfiguration"),
//   	TraceConfigurationVendor: apprunner.TraceConfigurationVendor_AWSXRAY,
//   })
//
//   apprunner.NewService(this, jsii.String("DemoService"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   	ObservabilityConfiguration: ObservabilityConfiguration,
//   })
//
// See: https://docs.aws.amazon.com/apprunner/latest/dg/monitor.html
//
// Experimental.
type TraceConfigurationVendor string

const (
	// Tracing (X-Ray).
	// Experimental.
	TraceConfigurationVendor_AWSXRAY TraceConfigurationVendor = "AWSXRAY"
)

