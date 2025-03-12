package awscodepipelineactions


// Example:
//   jenkinsProvider := codepipeline_actions.NewJenkinsProvider(this, jsii.String("JenkinsProvider"), &JenkinsProviderProps{
//   	ProviderName: jsii.String("MyJenkinsProvider"),
//   	ServerUrl: jsii.String("http://my-jenkins.com:8080"),
//   	Version: jsii.String("2"),
//   })
//
type JenkinsProviderProps struct {
	// The name of the Jenkins provider that you set in the AWS CodePipeline plugin configuration of your Jenkins project.
	//
	// Example:
	//   "MyJenkinsProvider"
	//
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The base URL of your Jenkins server.
	//
	// Example:
	//   "http://myjenkins.com:8080"
	//
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// Whether to immediately register a Jenkins Provider for the build category.
	//
	// The Provider will always be registered if you create a `JenkinsAction`.
	// Default: false.
	//
	ForBuild *bool `field:"optional" json:"forBuild" yaml:"forBuild"`
	// Whether to immediately register a Jenkins Provider for the test category.
	//
	// The Provider will always be registered if you create a `JenkinsTestAction`.
	// Default: false.
	//
	ForTest *bool `field:"optional" json:"forTest" yaml:"forTest"`
	// The version of your provider.
	// Default: '1'.
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

