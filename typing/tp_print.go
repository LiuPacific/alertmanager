package typing

import (
	"github.com/prometheus/alertmanager/types"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var CurrentIP = "0.0.0.0"

var MonTraceLog *log.Logger

//var MonInAlertLog *log.Logger
//var MonOutAlertLog *log.Logger
var MonStoreLog *log.Logger

const (
	InterfaceName = "bond0.1000" // 网卡名
	//InterfaceName = "lo0" // TODO: 网卡名
	LogPath = "/data/log/alertmanager"
	//LogPath    = "/Users/taipingliu/data/logs" // TODO: 修改路径，带有服务名
	TimeFormat = "15:04:05.000"
	InMarker   = "am-in"  // TODO: 入标记 |11:55:55.234|10.129.101.85|am-in|adfasf|
	OutMarker  = "am-out" // TODO: 出标记 |11:55:55.234|10.129.101.85|am-out|adfasf|

	ExceptionID = "6305528762073529654_15240_critical"
)

func init() {
	initIP()
	initLog()
}

func initIP() {
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, iface := range ifaces {
		if iface.Name != InterfaceName {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if len(ip.To4()) == 4 {
				CurrentIP = ip.String()
			}
		}
	}
}

func initLog() {
	os.MkdirAll(LogPath, 0777)
	traceLogFile, err := os.OpenFile(LogPath+"/trace.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open log error：", err)
	}
	//alertInLogFile, err := os.OpenFile(LogPath+"/in.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalln("open in alert log error：", err)
	//}
	//alertOutLogFile, err := os.OpenFile(LogPath+"/out.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalln("open out alert log error：", err)
	//}
	storeLogFile, err := os.OpenFile(LogPath+"/store.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open store log error：", err)
	}

	//MonTraceLog = log.New(io.MultiWriter(os.Stderr, traceLogFile), "", log.Ldate|log.Ltime|log.Lshortfile)
	MonTraceLog = log.New(io.MultiWriter(traceLogFile), "", 0)
	//MonInAlertLog = log.New(io.MultiWriter(alertInLogFile), "", log.Ltime)
	//MonOutAlertLog = log.New(io.MultiWriter(alertOutLogFile), "", log.Ltime)
	MonStoreLog = log.New(io.MultiWriter(storeLogFile), "", log.Lmicroseconds)
}

func InAlert(alert *types.Alert) {
	tNow := time.Now()
	timeNow := tNow.Format(TimeFormat)
	MonTraceLog.Printf("|%s|%s|%s|%s|\n", timeNow, CurrentIP, InMarker, alert.Annotations["id"])
	//MonInAlertLog.Printf("annotations: %s\nlabels: %s\n", alert.Annotations.String(), alert.Labels.String())
}

func OutAlert(alerts []*types.Alert) {
	for _, alert := range alerts {
		tNow := time.Now()
		timeNow := tNow.Format(TimeFormat)
		MonTraceLog.Printf("|%s|%s|%s|%s|\n", timeNow, CurrentIP, OutMarker, alert.Annotations["id"])
		//MonOutAlertLog.Printf("annotations: %s\nlabels: %s\n", alert.Annotations.String(), alert.Labels.String())
	}
}

func FlushAlert(alert *types.Alert, position string) {
	if alert.Annotations["id"] == ExceptionID {
		tNow := time.Now()
		timeNow := tNow.Format(TimeFormat)
		MonTraceLog.Printf("|%s|%s|%s|%s|\n", timeNow, CurrentIP, position, alert.Annotations["id"])
	}
}
