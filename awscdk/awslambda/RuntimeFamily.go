package awslambda


type RuntimeFamily string

const (
	RuntimeFamily_NODEJS RuntimeFamily = "NODEJS"
	RuntimeFamily_JAVA RuntimeFamily = "JAVA"
	RuntimeFamily_PYTHON RuntimeFamily = "PYTHON"
	RuntimeFamily_DOTNET_CORE RuntimeFamily = "DOTNET_CORE"
	RuntimeFamily_GO RuntimeFamily = "GO"
	RuntimeFamily_RUBY RuntimeFamily = "RUBY"
	RuntimeFamily_OTHER RuntimeFamily = "OTHER"
)

