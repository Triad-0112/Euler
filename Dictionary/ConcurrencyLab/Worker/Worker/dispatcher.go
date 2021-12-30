package Dispatcher

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}
