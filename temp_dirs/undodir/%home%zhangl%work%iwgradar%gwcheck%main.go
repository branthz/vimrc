Vim�UnDo� ������ 7��� ���I_���&啍�Q��"                                     Z���    _�                            ����                                                                                                                                                                                                                                                                                                                                                             Z��     �                	"iwgradar/agent/g"5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z��     �               	g.ParseConfig(*cfg)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z��     �               	if g.Config().Debug {5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z��     �               	if .Config().Debug {5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z��     �               		g.InitLog("debug")5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z��     �               		.InitLog("debug")5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z��     �               		g.InitLog("warn")5�_�      	                     ����                                                                                                                                                                                                                                                                                                                                                             Z��     �               		.InitLog("warn")5�_�      
           	          ����                                                                                                                                                                                                                                                                                                                                                             Z��     �               	g.ConfigShow()5�_�   	              
          ����                                                                                                                                                                                                                                                                                                                                                             Z��     �               	.ConfigShow()5�_�   
                         ����                                                                                                                                                                                                                                                                                                                                                             Z��    �                  package main       import (   	"flag"   	"fmt"   	"iwgradar/agent/http"   	"os"   )       func main() {   :	cfg := flag.String("c", "cfg.json", "configuration file")   1	version := flag.Bool("v", false, "show version")       	flag.Parse()       	if *version {   		fmt.Println(g.VERSION)   		os.Exit(0)   	}   	ParseConfig(*cfg)       	if Config().Debug {   		InitLog("debug")   		} else {   		InitLog("warn")   	}   	ConfigShow()   	go http.Start()   
	select {}   }5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Z�{    �                  package main       import (   	"flag"   	"fmt"   	"iwgradar/agent/g"   	"iwgradar/agent/http"   	"os"   )       func main() {   :	cfg := flag.String("c", "cfg.json", "configuration file")   1	version := flag.Bool("v", false, "show version")       	flag.Parse()       	if *version {   		fmt.Println(g.VERSION)   		os.Exit(0)   	}   	ParseConfig(*cfg)       	if Config().Debug {   		InitLog("debug")   		} else {   		InitLog("warn")   	}   	ConfigShow()   	go http.Start()   
	select {}   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�S     �                
	select {}5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�U     �               	go http.Start()5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�Z     �               	http.Start()5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Z�\    �                  package main       import (   	"flag"   	"fmt"   	"iwgradar/agent/g"   	"iwgradar/agent/http"   	"os"   )       func main() {   :	cfg := flag.String("c", "cfg.json", "configuration file")   1	version := flag.Bool("v", false, "show version")       	flag.Parse()       	if *version {   		fmt.Println(g.VERSION)   		os.Exit(0)   	}   	ParseConfig(*cfg)       	if Config().Debug {   		InitLog("debug")   		} else {   		InitLog("warn")   	}   	ConfigShow()   	Start()   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z���     �               	�             5�_�                       	    ����                                                                                                                                                                                                                                                                                                                                                             Z���     �               
	InitLog()5�_�                            ����                                                                                                                                                                                                                                                                                                                                                V       Z���     �                	if Config().Debug {   		InitLog("debug")   		} else {   		InitLog("warn")   	}5�_�                           ����                                                                                                                                                                                                                                                                                                                                                V       Z���    �                  package main       import (   	"flag"   	"fmt"   	"iwgradar/agent/g"   	"os"   )       func main() {   :	cfg := flag.String("c", "cfg.json", "configuration file")   1	version := flag.Bool("v", false, "show version")       	flag.Parse()       	if *version {   		fmt.Println(g.VERSION)   		os.Exit(0)   	}   	ParseConfig(*cfg)       	InitLog(Config().Loglevel)   	ConfigShow()   	Start()   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                V       Z���    �                  package main       import (   	"flag"   	"fmt"   	"iwgradar/agent/g"   	"os"   )       func main() {   :	cfg := flag.String("c", "cfg.json", "configuration file")   1	version := flag.Bool("v", false, "show version")       	flag.Parse()       	if *version {   		fmt.Println(g.VERSION)   		os.Exit(0)   	}   	ParseConfig(*cfg)       	InitLog(Config().Loglevel)   	ConfigShow()   	Start()   }5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Z�c�    �                  package main       import (   	"flag"   	"fmt"   	"iwgradar/agent/g"   	"os"   )       func main() {   :	cfg := flag.String("c", "cfg.json", "configuration file")   1	version := flag.Bool("v", false, "show version")       	flag.Parse()       	if *version {   		fmt.Println(g.VERSION)   		os.Exit(0)   	}   	ParseConfig(*cfg)       	InitLog(Config().Loglevel)   	ConfigShow()   	Start()   }5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Z���    �                  package main       import (   	"flag"   	"fmt"   	"iwgradar/agent/g"   	"os"   )       func main() {   :	cfg := flag.String("c", "cfg.json", "configuration file")   1	version := flag.Bool("v", false, "show version")       	flag.Parse()       	if *version {   		fmt.Println(g.VERSION)   		os.Exit(0)   	}   	ParseConfig(*cfg)       	InitLog(Config().Loglevel)   	ConfigShow()   	Start()   }5��