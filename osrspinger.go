package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"runtime"

	//"os"
	//"os/exec"
	//"net"
	//"github.com/fatih/color"

	"github.com/sparrc/go-ping"
)
//


func main() {

	//f2p := 	[]int{1, 8, 16, 26, 35, 81, 82, 83, 84, 85, 93, 94, 97, 98, 99, 113, 114, 117, 118, 119, 124, 125, 126,
	//	127, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 151, 152, 153, 154, 155, 156,
	//	157, 158, 159, 160, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 197, 198, 199,
	//	200, 201, 202, 203, 204, 205, 206}

	const serverFormat string = "oldschool%v.runescape.com"
	const numberOfWorlds int = 206 // 'rona 19
	const worldIdOffset int = 300

	fmt.Println("-- osrs pinger--")

	for worldId := 1; worldId < numberOfWorlds; worldId++ {

		curr := fmt.Sprintf(serverFormat, worldId)
		//ipaddr, err := net.ResolveIPAddr("ip", curr)

		//if err != nil {
		//	fmt.Println("Could not call world", curr, err)
		//}

		pinger, err := ping.NewPinger(curr)

		if err != nil {
			fmt.Println("fuq", err)
			os.Exit(1)
		}
		pinger.SetPrivileged(true)
		pinger.Count = 1
		pinger.Run()

		stats := pinger.Statistics()
		pinger.Statistics()
		worldnum := worldId + worldIdOffset
		if runtime.GOOS == "windows" {
			var rtoutput string
			if stats.MaxRtt.Milliseconds() > 60 {
				rtoutput = color.RedString("World %d", worldnum)
			} else {
				rtoutput = color.GreenString("World %d", worldnum)
			}
			//fmt.Fprintf(color.Output, "-> %s | %s | %s\n",
				//				curr, color.GreenString(ipaddr.IP.String()),
				//				color.CyanString(fmt.Sprintf("world %d", worldId + worldIdOffset)))

			//fmt.Fprintf(color.Output, "Windows support: %s", color.GreenString("PASS"))
			fmt.Fprintf(color.Output, "\n%5s %10s %5s %5s",
				rtoutput ,stats.Addr, stats.Rtts, stats.MaxRtt,
			)
		}
	}
}




		//if err != nil {
		//	fmt.Println("Could not call world", curr, err)
		//	os.Exit(1)
		//}
		//
		//resp := exec.Command("ping", "-n", "1",curr)
		//output, err := resp.Output()
		//if (err != nil) {
		//	fmt.Println("Could not call world", curr, err)
		//}

		//y := string(output)


		//fmt.Println("Stats", string(output))
//
//		if runtime.GOOS == "windows" {
//			fmt.Fprintf(color.Output, "-> %s | %s | %s\n",
//				curr, color.GreenString(ipaddr.IP.String()),
//				color.CyanString(fmt.Sprintf("world %d", worldId + worldIdOffset)))
//		} else {
//			//yellow := color.New(color.FgYellow).SprintFunc()
//			//red := color.New(color.FgRed).SprintFunc()
//			//fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))
//			//
//			//info := color.New(color.FgWhite, color.BgGreen).SprintFunc()
//			//fmt.Printf("This %s rocks!\n", info("package"))
//			//
//			//fmt.Fprintf(color.Output, "Windows support: %s", color.GreenString("PASS"))
//			//
//			//white := color.New(color.FgBlue).SprintFunc()
//		}
//	}
//
//	fmt.Println("-- done --")
//
//}
