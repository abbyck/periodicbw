out := Resp{
	Type:      "result",
	TimeStamp: time.Now(),
	Ping: PingType{
		Jitter:  2.2149999999999999,
		Latency: 39.049999999999997,
	},
	Download: DownloadType{
		Bandwidth: 8785817,
		Bytes:     101747840,
		Elapsed:   13905,
	},
	Upload: UploadType{
		Bandwidth: 7917794,
		Bytes:     111661778,
		Elapsed:   15011,
	},
	PacketLoss: 0,
	Isp:        "BSNL",
	Interface: InterfaceType{
		InternalIP: "192.168.1.110",
		Name:       "wlp3s0",
		MacAddr:    "78:2B:46:D9:B2:28",
		IsVpn:      false,
		ExternalIP: "117.213.85.44",
	},
	Server: ServerType{
		ID:       16512,
		Name:     "DeeNet Services",
		Location: "Mangalore",
		Country:  "India",
		Host:     "mangalore.ganeshcable.net",
		Port:     8080,
		IP:       "103.89.233.30",
	},
	Result: ResultType{
		ID:  "a5a8160a-db26-483c-9e1b-67726d44e287",
		URL: "https://www.speedtest.net/result/c/a5a8160a-db26-483c-9e1b-67726d44e287",
	},
}