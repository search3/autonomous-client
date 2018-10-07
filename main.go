package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/space3/autonomous-client/jobmodels"
	"github.com/space3/autonomous-client/taskrunner"
	"github.com/space3/autonomous-client/resourcemanager"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

func main() {
	drone := tello.NewDriver("8888")
	keys := keyboard.NewDriver()

	///get job 
	curJob := jobmodels.JobSpec{....} //Not yet implemented

	resourcemanager.GetJobResources(curJob)


	work := func() {
		currentlyFlying := false
		//writeToFile := false
		const frameX = 400
		const frameY = 350
		const frameSize = frameX * frameY * 3

		mplayer := exec.Command("mplayer", "-fps", "25", "-")
		mplayerIn, _ := mplayer.StdinPipe()
		if err := mplayer.Start(); err != nil {
			fmt.Println(err)
			return
		}

		drone.On(tello.ConnectedEvent, func(data interface{}) {
			fmt.Println("Connected")
			drone.StartVideo()
			drone.SetVideoEncoderRate(tello.VideoBitRateAuto)
			drone.SetExposure(0)
			gobot.Every(100*time.Millisecond, func() {
				drone.StartVideo()
			})
		})

		drone.On(tello.VideoFrameEvent, func(data interface{}) {
			pkt := data.([]byte)
			if _, err := mplayerIn.Write(pkt); err != nil {
				fmt.Println(err)
			}
			if taskrunner.MatchFace(pkt) {
				sub := jobmodels.JobSubmission{curJob.jobHash, "./img002.jpg"}
				resourcemanager.UploadJobArtifiacts(sub)
			}
		})

		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)

			switch key.Key {
			case keyboard.A:
				fmt.Println("Left")
				drone.Left(20)
			case keyboard.W:
				fmt.Println("Forward")
				drone.Forward(20)
			case keyboard.S:
				fmt.Println("Backward")
				drone.Backward(20)
			case keyboard.D:
				fmt.Println("Right")
				drone.Right(20)
			case keyboard.E:
				fmt.Println("CW")
				drone.Clockwise(30)
			case keyboard.Q:
				fmt.Println("CCW")
				drone.CounterClockwise(30)
			case keyboard.F:
				fmt.Println("Up")
				drone.Up(20)
			case keyboard.V:
				fmt.Println("Down")
				drone.Up(20)
			case keyboard.Z:
				fmt.Println("Take Off")
				drone.TakeOff()
			case keyboard.X:
				fmt.Println("Land")
				drone.Land()
			default:
				fmt.Println("Nothing")
				if currentlyFlying {
					drone.Hover()
				}
			}
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone, keys},
		work,
	)

	robot.Start()
}
