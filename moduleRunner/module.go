package moduleRunner

// A module, new modules must extend this class
type versionModule interface {
	Init()
	Name() string
	Run(upstream string, params interface{}) (version string, err error)
}
