package providers

type Provider interface {
	Run()
	Exit()
	Bootstrap()
}
