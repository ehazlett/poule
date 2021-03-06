package configuration

const (
	// JenkinsBaseUrl is the base URL for the Jenkins CI server.
	JenkinsBaseUrl = "https://leeroy.dockerproject.org/build/retry"

	// FailingCILabel is the label that indicates that a pull request is
	// failing for a legitimate reason and should be ignored.
	FailingCILabel = "status/failing-ci"

	// PouleToken is injected as an HTML comment in the body of all messages
	// posted by the tool itself.
	PouleToken = "AUTOMATED:POULE"
)
