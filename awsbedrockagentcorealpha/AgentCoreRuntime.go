package awsbedrockagentcorealpha


// Bedrock AgentCore runtime environment for code execution Allowed values: PYTHON_3_10 | PYTHON_3_11 | PYTHON_3_12 | PYTHON_3_13 | PYTHON_3_14 | NODE_22.
// Experimental.
type AgentCoreRuntime string

const (
	// Python 3.10 runtime.
	// Experimental.
	AgentCoreRuntime_PYTHON_3_10 AgentCoreRuntime = "PYTHON_3_10"
	// Python 3.11 runtime.
	// Experimental.
	AgentCoreRuntime_PYTHON_3_11 AgentCoreRuntime = "PYTHON_3_11"
	// Python 3.12 runtime.
	// Experimental.
	AgentCoreRuntime_PYTHON_3_12 AgentCoreRuntime = "PYTHON_3_12"
	// Python 3.13 runtime.
	// Experimental.
	AgentCoreRuntime_PYTHON_3_13 AgentCoreRuntime = "PYTHON_3_13"
	// Python 3.14 runtime.
	// Experimental.
	AgentCoreRuntime_PYTHON_3_14 AgentCoreRuntime = "PYTHON_3_14"
	// Node.js 22 runtime.
	// Experimental.
	AgentCoreRuntime_NODE_22 AgentCoreRuntime = "NODE_22"
)

