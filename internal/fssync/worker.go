package fssync

import "runtime"

var nbOfCPU = runtime.NumCPU() / 2
var workers []*Worker = make([]*Worker, nbOfCPU)
var balance int = 0

type Worker struct {
	isRunning bool
	missions  chan *Mission
	executed  int
}

func (w *Worker) Run() {
	if !w.isRunning {
		w.isRunning = true
		for mission := range w.missions {
			mission.execute()
			w.executed++
		}
	}
}

func init() {
	for wn, w := range workers {
		if w == nil {
			workers[wn] = &Worker{
				isRunning: false,
				executed:  0,
				missions:  make(chan *Mission),
			}
		}
	}
}
func Run() {
	for wn, w := range workers {
		if w != nil {
			go workers[wn].Run()
		}
	}
}
func AddMission(mission *Mission) {
	//mission as a channel
	workers[balance%nbOfCPU].missions <- mission
	balance++
}
func TotalEnCours() int {
	var count int = 0
	for wn, w := range workers {
		if w != nil {
			count += workers[wn].executed
		}
	}
	return count
}
