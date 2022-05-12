# AWS::CodeGuruProfiler Construct Library

Amazon CodeGuru Profiler collects runtime performance data from your live applications, and provides recommendations that can help you fine-tune your application performance.

## Installation

Import to your project:

```go
import codeguruprofiler "github.com/aws/aws-cdk-go/awscdk"
```

## Basic usage

Here's how to setup a profiling group and give your compute role permissions to publish to the profiling group to the profiling agent can publish profiling information:

```go
// The execution role of your application that publishes to the ProfilingGroup via CodeGuru Profiler Profiling Agent. (the following is merely an example)
publishAppRole := iam.NewRole(this, jsii.String("PublishAppRole"), &roleProps{
	assumedBy: iam.NewAccountRootPrincipal(),
})

profilingGroup := codeguruprofiler.NewProfilingGroup(this, jsii.String("MyProfilingGroup"))
profilingGroup.grantPublish(publishAppRole)
```

## Compute Platform configuration

Code Guru Profiler supports multiple compute environments.
They can be configured when creating a Profiling Group by using the `computePlatform` property:

```go
profilingGroup := codeguruprofiler.NewProfilingGroup(this, jsii.String("MyProfilingGroup"), &profilingGroupProps{
	computePlatform: codeguruprofiler.computePlatform_AWS_LAMBDA,
})
```
