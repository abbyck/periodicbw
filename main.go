package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// PingType of resp
type PingType struct {
	Jitter  float64 `json:"jitter"`
	Latency float64 `json:"latency"`
}

// DownloadType of resp
type DownloadType struct {
	Bandwidth int32 `json:"bandwidth"`
	Bytes     int32 `json:"bytes"`
	Elapsed   int32 `json:"elapsed"`
}

// UploadType of resp
type UploadType struct {
	Bandwidth int32 `json:"bandwidth"`
	Bytes     int32 `json:"bytes"`
	Elapsed   int32 `json:"elapsed"`
}

// InterfaceType of resp
type InterfaceType struct {
	InternalIP string `json:"internalIp"`
	Name       string `json:"name"`
	MacAddr    string `json:"macAddr"`
	IsVpn      bool   `json:"isVpn"`
	ExternalIP string `json:"externalIp"`
}

// ServerType of resp
type ServerType struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Country  string `json:"country"`
	Host     string `json:"host"`
	Port     int16  `json:"port"`
	IP       string `json:"ip"`
}

// ResultType of response
type ResultType struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// Resp contains response of speedtest cli JSON struct representation - Composite
type Resp struct {
	Type       string        `json:"type"`
	TimeStamp  time.Time     `json:"timestamp"`
	Ping       PingType      `json:"ping"`
	Download   DownloadType  `json:"download"`
	Upload     UploadType    `json:"upload"`
	PacketLoss float64       `json:"packetLoss"`
	Isp        string        `json:"isp"`
	Interface  InterfaceType `json:"interface"`
	Server     ServerType    `json:"server"`
	Result     ResultType    `json:"result"`
}

func execute() {

	/* here we perform the speedtest command with flags progress=no and format=json.
	// we can store the output of this in our out variable
	// and catch any errors in err
	*/
	out, err := exec.Command("speedtest", "--progress=no", "--format=json").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
	// out := string(out[:])

	var result Resp
	err = json.Unmarshal(out, &result)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("Date and Time: ", result.TimeStamp)
	// fmt.Println("Ping - Jitter: ", result.Ping.Jitter, "ms, Latency: ", result.Ping.Latency, "ms")
	// fmt.Println("ISP: ", result.Isp)
	// fmt.Println("External IP:", result.Interface.ExternalIP)
	// fmt.Println("Test server: ", result.Server.ID, result.Server.Name, result.Server.Location, result.Server.IP)
	// fmt.Println("Result URL: ", result.Result.URL)
	fmt.Println("Download bandwidth: ", float64(result.Download.Bandwidth)/float64(125000), "Mbps")
	fmt.Println("Upload bandwidth: ", float64(result.Upload.Bandwidth)/float64(125000), "Mbps")
	fmt.Println("Packet Loss: ", result.PacketLoss, "%")
	writeToFile(result)
}

func writeToFile(r Resp) {
	// create a csv file with the ISP's name in lowercase
	file, err := os.OpenFile(strings.ToLower(r.Isp)+".csv", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	//get file size and if it's empty, write headers
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Now()

	csvWriter := csv.NewWriter(file)
	if fi.Size() == 0 {
		header := []string{
			"Date",
			"Time",
			"Type",
			"Jitter",
			"Latency",
			"Download Bandwidth",
			"Upload Bandwidth",
			"PacketLoss",
			"ISP",
			"ExternalIP",
			"ServerID",
			"Server Name",
			"Server Location",
			"Country",
			"Host",
			"IP",
			"Result ID",
		}
		data := []string{
			r.TimeStamp.Format("06-Jan-02"),
			currentTime.Format("15:04 PM"),
			r.Type,
			FloatToString(r.Ping.Jitter),
			FloatToString(r.Ping.Latency),
			FloatToString(float64(r.Download.Bandwidth) / float64(125000)),
			FloatToString(float64(r.Upload.Bandwidth) / float64(125000)),
			FloatToString(r.PacketLoss),
			r.Isp,
			r.Interface.ExternalIP,
			strconv.Itoa(int(r.Server.ID)),
			r.Server.Name,
			r.Server.Location,
			r.Server.Country,
			r.Server.Host,
			r.Server.IP,
			r.Result.ID,
		}
		csvWriter.Write(header)
		csvWriter.Write(data)
		csvWriter.Flush()
		err := csvWriter.Error()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println("finished writing")
	} else {
		data := []string{
			r.TimeStamp.Format("02-Jan-2006"),
			currentTime.Format("03:04 PM"),
			r.Type,
			FloatToString(r.Ping.Jitter),
			FloatToString(r.Ping.Latency),
			FloatToString(float64(r.Download.Bandwidth) / float64(125000)),
			FloatToString(float64(r.Upload.Bandwidth) / float64(125000)),
			FloatToString(r.PacketLoss),
			r.Isp,
			r.Interface.ExternalIP,
			strconv.Itoa(int(r.Server.ID)),
			r.Server.Name,
			r.Server.Location,
			r.Server.Country,
			r.Server.Host,
			r.Server.IP,
			r.Result.ID,
		}
		csvWriter.Write(data)
		csvWriter.Flush()
		err := csvWriter.Error()
		if err != nil {
			fmt.Print(err)
		}
	}
	fmt.Println(r.Isp)
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}
