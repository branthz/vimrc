Vim�UnDo� ���ۍU�������������PY����J�                    F       F   F   F    Z�2�    _�                             ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �                 �              5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �                  �               5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �                 import()5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �                5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �               	""5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �                  �               5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �                 func TestPing(){}5�_�      	                     ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �         	      func TestPing(){5�_�      
           	           ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �      	   	       5�_�   	              
           ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �      	   	      	5�_�   
                        ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �      	   	      	NewPinger()	5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �      	   	      	NewPinger('')	5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �      	   	      	NewPinger('223.5.5.5')	5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �      
   
      	�      
   	    5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �         
      	if err!=nil{}5�_�                    
       ����                                                                                                                                                                                                                                                                                                                                                             Z�	�     �   	            	5�_�                    
   
    ����                                                                                                                                                                                                                                                                                                                                                             Z�
     �   	            		t.Fatal()5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�
    �             5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�     �               	�             5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�,     �               		p.Bind()5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�1     �      	         	p,err:=NewPinger('223.5.5.5')	5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�4     �      	         	p,err:=NewPinger("223.5.5.5')	5�_�                       	    ����                                                                                                                                                                                                                                                                                                                                                             Z�8     �               	p.Bind("")5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�?     �               	�             5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�`     �               	�             5�_�                       #    ����                                                                                                                                                                                                                                                                                                                                                             Z�y     �               $	if p.Statistics().PacketsRecv <2 {}5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z�y     �               	5�_�                       
    ����                                                                                                                                                                                                                                                                                                                                                             Z��     �               		t.Fatal()5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Z��     �               		t.Fatal("")5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Z��    �                  package main       import(   
	"testing"   )       func TestPing(t *testing.T){   	p,err:=NewPinger("223.5.5.5")	   	if err!=nil{   		t.Fatal(err)   	}   	p.Bind("192.168.34.41")   	p.Count =5   	p.Timeout=1e10    	p.Run()   #	if p.Statistics().PacketsRecv <2 {   		t.Fatal("ping test failed")   	}       }5�_�                             ����                                                                                                                                                                                                                                                                                                                                                             Z�0     �                 5�_�      !                       ����                                                                                                                                                                                                                                                                                                                                                             Z�0    �                  package main       import (   
	"testing"   )       func TestPing(t *testing.T) {   !	p, err := NewPinger("223.5.5.5")   	if err != nil {   		t.Fatal(err)   	}   	p.Bind("192.168.34.41")   	p.Count = 5   	p.Timeout = 1e10   	p.Run()   $	if p.Statistics().PacketsRecv < 2 {   		t.Fatal("ping test failed")   	}   }5�_�       "           !          ����                                                                                                                                                                                                                                                                                                                                                             Z�b�    �                  package main       import (   
	"testing"   )       func TestPing(t *testing.T) {   !	p, err := NewPinger("223.5.5.5")   	if err != nil {   		t.Fatal(err)   	}   	p.Bind("192.168.34.41")   	p.Count = 5   	p.Timeout = 1e10   	p.Run()   $	if p.Statistics().PacketsRecv < 2 {   		t.Fatal("ping test failed")   	}   }5�_�   !   #           "           ����                                                                                                                                                                                                                                                                                                                                                             Z�c1     �      	          5�_�   "   $           #           ����                                                                                                                                                                                                                                                                                                                                                             Z�c3     �                5�_�   #   %           $          ����                                                                                                                                                                                                                                                                                                                                                             Z�c7     �      
         func init(){}5�_�   $   &           %           ����                                                                                                                                                                                                                                                                                                                                                             Z�c9     �      	          5�_�   %   '           &      	    ����                                                                                                                                                                                                                                                                                                                                                             Z�cL     �      	         
	InitLog()5�_�   &   (           '      
    ����                                                                                                                                                                                                                                                                                                                                                             Z�cO     �      	         	InitLog("")5�_�   '   )           (           ����                                                                                                                                                                                                                                                                                                                                                             Z�cW    �                  package main       import (   
	"testing"   )       func init(){   	InitLog("debug")   }       func TestPing(t *testing.T) {   !	p, err := NewPinger("223.5.5.5")   	if err != nil {   		t.Fatal(err)   	}   	p.Bind("192.168.34.41")   	p.Count = 5   	p.Timeout = 1e10   	p.Run()   $	if p.Statistics().PacketsRecv < 2 {   		t.Fatal("ping test failed")   	}   }5�_�   (   *           )          ����                                                                                                                                                                                                                                                                                                                                                             Z�d      �               		t.Fatal("ping test failed")5�_�   )   +           *      "    ����                                                                                                                                                                                                                                                                                                                                                             Z�d'     �               #		t.Fatal("ping test failed:%d,%d")5�_�   *   ,           +           ����                                                                                                                                                                                                                                                                                                                                                             Z�dC    �                  package main       import (   
	"testing"   )       func init() {   	InitLog("debug")   }       func TestPing(t *testing.T) {   !	p, err := NewPinger("223.5.5.5")   	if err != nil {   		t.Fatal(err)   	}   	p.Bind("192.168.34.41")   	p.Count = 5   	p.Timeout = 1e10   	p.Run()   $	if p.Statistics().PacketsRecv < 2 {   Y		t.Fatal("ping test failed:%d,%d",p.Statistics().PacketsSent,p.Statistics().PacketsRecv)   	}   }5�_�   +   -           ,          ����                                                                                                                                                                                                                                                                                                                                                             Z�dX     �               [		t.Fatal("ping test failed:%d,%d", p.Statistics().PacketsSent, p.Statistics().PacketsRecv)5�_�   ,   .           -          ����                                                                                                                                                                                                                                                                                                                                                             Z�dX     �               Z		t.Fatal("ping test failed:d,%d", p.Statistics().PacketsSent, p.Statistics().PacketsRecv)5�_�   -   /           .          ����                                                                                                                                                                                                                                                                                                                                                             Z�dY     �               Y		t.Fatal("ping test failed:,%d", p.Statistics().PacketsSent, p.Statistics().PacketsRecv)5�_�   .   0           /          ����                                                                                                                                                                                                                                                                                                                                                             Z�dY     �               X		t.Fatal("ping test failed:%d", p.Statistics().PacketsSent, p.Statistics().PacketsRecv)5�_�   /   1           0          ����                                                                                                                                                                                                                                                                                                                                                             Z�dY     �               W		t.Fatal("ping test failed:d", p.Statistics().PacketsSent, p.Statistics().PacketsRecv)5�_�   0   2           1          ����                                                                                                                                                                                                                                                                                                                                                             Z�dZ    �                  package main       import (   
	"testing"   )       func init() {   	InitLog("debug")   }       func TestPing(t *testing.T) {   !	p, err := NewPinger("223.5.5.5")   	if err != nil {   		t.Fatal(err)   	}   	p.Bind("192.168.34.41")   	p.Count = 5   	p.Timeout = 1e10   	p.Run()   $	if p.Statistics().PacketsRecv < 2 {   V		t.Fatal("ping test failed:", p.Statistics().PacketsSent, p.Statistics().PacketsRecv)   	}   }5�_�   1   3           2          ����                                                                                                                                                                                                                                                                                                                                                             Z�di     �               !	p, err := NewPinger("223.5.5.5")5�_�   2   4           3          ����                                                                                                                                                                                                                                                                                                                                                             Z�du    �                  package main       import (   
	"testing"   )       func init() {   	InitLog("debug")   }       func TestPing(t *testing.T) {   !	p, err := NewPinger("222.3.3.3")   	if err != nil {   		t.Fatal(err)   	}   	p.Bind("192.168.34.41")   	p.Count = 5   	p.Timeout = 1e10   	p.Run()   $	if p.Statistics().PacketsRecv < 2 {   V		t.Fatal("ping test failed:", p.Statistics().PacketsSent, p.Statistics().PacketsRecv)   	}   }5�_�   3   5           4          ����                                                                                                                                                                                                                                                                                                                                                v   !    Z�d�     �               !	p, err := NewPinger("222.3.3.3")5�_�   4   6           5          ����                                                                                                                                                                                                                                                                                                                                                v   !    Z�d�   	 �                  package main       import (   
	"testing"   )       func init() {   	InitLog("debug")   }       func TestPing(t *testing.T) {   '	p, err := NewPinger("114.114.114.114")   	if err != nil {   		t.Fatal(err)   	}   	p.Bind("192.168.34.41")   	p.Count = 5   	p.Timeout = 1e10   	p.Run()   $	if p.Statistics().PacketsRecv < 2 {   V		t.Fatal("ping test failed:", p.Statistics().PacketsSent, p.Statistics().PacketsRecv)   	}   }5�_�   5   7           6          ����                                                                                                                                                                                                                                                                                                                                                             Z�     �               	�             5�_�   6   8           7          ����                                                                                                                                                                                                                                                                                                                                                             Z�     �               	log.Debug()5�_�   7   9           8          ����                                                                                                                                                                                                                                                                                                                                                             Z�      �               	log.Debug("")5�_�   8   :           9           ����                                                                                                                                                                                                                                                                                                                                                             Z�#   
 �                  package main       import (   
	"testing"   )       func init() {   	InitLog("debug")   }       func TestPing(t *testing.T) {   '	p, err := NewPinger("114.114.114.114")   	if err != nil {   		t.Fatal(err)   	}   	p.Bind("192.168.34.41")   #	log.Debug("---------------------")   	p.Count = 5   	p.Timeout = 1e10   	p.Run()   $	if p.Statistics().PacketsRecv < 2 {   V		t.Fatal("ping test failed:", p.Statistics().PacketsSent, p.Statistics().PacketsRecv)   	}   }5�_�   9   ;           :          ����                                                                                                                                                                                                                                                                                                                                                             Z�2B     �               '	p, err := NewPinger("114.114.114.114")5�_�   :   <           ;          ����                                                                                                                                                                                                                                                                                                                                                             Z�2F     �               *	p, err := NewPinger("","114.114.114.114")5�_�   ;   =           <          ����                                                                                                                                                                                                                                                                                                                                                             Z�2N     �               7	p, err := NewPinger("192.168.34.41","114.114.114.114")5�_�   <   >           =          ����                                                                                                                                                                                                                                                                                                                                                             Z�2S     �                	p.Bind("192.168.34.41")5�_�   =   ?           >          ����                                                                                                                                                                                                                                                                                                                                                             Z�2T     �                #	log.Debug("---------------------")5�_�   >   @           ?           ����                                                                                                                                                                                                                                                                                                                                                V       Z�2Y     �                	if err != nil {   		t.Fatal(err)   	}5�_�   ?   A           @          ����                                                                                                                                                                                                                                                                                                                                                V       Z�2[     �             �             5�_�   @   B           A          ����                                                                                                                                                                                                                                                                                                                                                V       Z�2]     �               	p.Run()5�_�   A   C           B          ����                                                                                                                                                                                                                                                                                                                                                v        Z�2f     �               $	if p.Statistics().PacketsRecv < 2 {5�_�   B   D           C          ����                                                                                                                                                                                                                                                                                                                                                v        Z�2g     �               	if p. < 2 {5�_�   C   E           D          ����                                                                                                                                                                                                                                                                                                                                         T       v   e    Z�2�     �               V		t.Fatal("ping test failed:", p.Statistics().PacketsSent, p.Statistics().PacketsRecv)5�_�   D   F           E          ����                                                                                                                                                                                                                                                                                                                                         T       v   e    Z�2�     �                		t.Fatal("ping test failed:", )5�_�   E               F           ����                                                                                                                                                                                                                                                                                                                                         T       v   e    Z�2�    �                  package main       import (   
	"testing"       %	"github.com/branthz/utarrow/lib/log"   )       func init() {   	InitLog("debug")   }       func TestPing(t *testing.T) {   1	p:= NewPinger("192.168.34.41","114.114.114.114")   	p.Count = 5   	p.Timeout = 1e10   	err:=p.Run()   	if err != nil {   		t.Fatal(err)   	}   	if p.Static.Recv < 2 {   		t.Fatal("ping test failed:")   	}   }5��