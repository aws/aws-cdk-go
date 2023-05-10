package awscodepipelineactions


// Properties for importing an existing Jenkins provider.
//
// Example:
//   jenkinsProvider := codepipeline_actions.JenkinsProvider_FromJenkinsProviderAttributes(this, jsii.String("JenkinsProvider"), &JenkinsProviderAttributes{
//   	ProviderName: jsii.String("MyJenkinsProvider"),
//   	ServerUrl: jsii.String("http://my-jenkins.com:8080"),
//   	Version: jsii.String("2"),
//   })
//
// Experimental.
type JenkinsProviderAttributes struct {
	// The name of the Jenkins provider that you set in the AWS CodePipeline plugin configuration of your Jenkins project.
	//
	// Example:
	//   "MyJenkinsProvider"
	//
	// Experimental.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The base URL of your Jenkins server.
	//
	// Example:
	//   "http://myjenkins.com:8080"
	//
	// Experimental.
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// The version of your provider.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

