package main

import (
	"math"
	"fmt"
	//"log"
	"path"
	//"path/filepath"
	"runtime"
	//"time"
	cv "github.com/hybridgroup/go-opencv/opencv"
	"github.com/hybridgroup/gobot"
	//"github.com/hybridgroup/gobot/platforms/firmata"
	//"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/opencv"
)

func main() {
	gbot := gobot.NewGobot()
 
	window := opencv.NewWindowDriver("window")
        camera := opencv.NewCameraDriver("camera", 0)
	
	//e := edison.NewEdisonAdaptor("edison")
	//led1 := gpio.NewLedDriver(e, "led", "13")
	//firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/ttyACM0")
	//led1 := gpio.NewLedDriver(firmataAdaptor, "led", "13")
	//motor := gpio.NewMotorDriver(firmataAdaptor, "motor", "4")
	//led2 := gpio.NewLedDriver(firmataAdaptor, "led", "10")
	//led3 := gpio.NewLedDriver(firmataAdaptor, "led", "9")

	_, currentfile, _, _ := runtime.Caller(0)
	work := func() {
		//speed := byte(0)
		//fadeAmount := byte(15)
		gobot.On(camera.Event("frame"), func(data interface{}) {
			cascade := cv.LoadHaarClassifierCascade(path.Join(path.Dir(currentfile), "haarcascade_frontalface_alt.xml"))
			nestedcascade := cv.LoadHaarClassifierCascade(path.Join(path.Dir(currentfile), "smiled_04.xml"))
	
	
			i := data.(*cv.IplImage)	
        		faces := cascade.DetectObjects(i)
			teeth:=0	
			//flagcenter:= false		
			for _, value := range faces {
				/*cv.Rectangle(i,
				cv.Point{value.X() + value.Width(), value.Y()},
				cv.Point{value.X(), value.Y() + value.Height()},
				cv.ScalarAll(255.0), 1, 0, 0)*/
				
				mouths := nestedcascade.DetectObjects(i)
				for _, value1 := range mouths{
					if float64(value1.Y())>float64(value.Y())+float64(value.Height())*3/5 && float64(value1.Y())+float64(value1.Height())<float64(value.Y())+float64(value.Height()) && math.Abs((float64(value1.X())+float64(value1.Width())/2) - (float64(value.X())+float64(value.Width())/2)) < float64(value.Width())/10{
						
						cv.Rectangle(i,
						cv.Point{value1.X() + value1.Width(), value1.Y()},
						cv.Point{value1.X(), value1.Y() + value1.Height()},
						cv.ScalarAll(255.0), 1, 0, 0)
						teeth = teeth+1
						//if value1.X()>240 && value1.Y()>240 && value1.Height()>50 && value1.Width()>100{ flagcenter=true }else{ flagcenter=false }						
						fmt.Println(value1.X(),value1.Y(),value1.Height(),value1.Width())
					}  				
				}
				//if len(smiles) > 0 { fmt.Println("worked")}
			}
			window.ShowImage(i)	
			if teeth > 0 {
				//led1.On()
				//motor.Speed(speed)
				//speed = speed + fadeAmount
				//if speed == 0 || speed == 255 {
				//	fadeAmount = -fadeAmount
				//}
				fmt.Println("good")
				//led2.On()
				//led3.On()	
				/*f := faces[0]
				w := f.Width()
				fmt.Println(w)
				switch {
				case w < 250:
					fmt.Println("*")
					led2.On()
					led3.Off()
					//red3.Off()
					//red4.Off()
					//red5.Off()
				case w > 250 && w < 350:
					fmt.Println("**")
					led2.Off()
					led3.On()
					//red3.Off()
					//red4.Off()
					//red5.Off()
				case w > 350 && w < 400:
					fmt.Println("***")
					led2.On()
					led3.On()
					//red3.On()
					//red4.Off()
					//red5.Off()
				case w > 400 && w < 450:
					fmt.Println("****")
					led2.On()
					led3.On()
					//red3.On()
					//red4.On()
					//red5.Off()
				case w > 450:
					fmt.Println("*****")
					led2.Off()
					led3.Off()
					//red3.On()
					//red4.On()
					//red5.On()
				}*/
			} else {
				//led1.Off()
				//speed = byte(0)
				//fadeAmount := byte(15)
				fmt.Println("smile please")
				//led2.Off()
				//led3.Off()
				//red3.Off()
				//red4.Off()
				//red5.Off()
			}
			
		})

	}
	
	/*work1 := func() {
                gobot.Every(1*time.Second, func() {
                        
                })
        }*/
	

	robot := gobot.NewRobot("cameraBot",
		//[]gobot.Connection{firmataAdaptor},
                []gobot.Device{window,camera},
                work,
        )
	
	/*robot1 := gobot.NewRobot("ledBot",
                []gobot.Device{window,camera,led1},
                work,
        )*/

        gbot.AddRobot(robot)
	//gbot.AddRobot(robot1)

        gbot.Start()


}
