package main

import (
	"fmt"
)

func main() {
	//by:=[]byte{8 ,1 ,18 ,199 ,3 ,10 ,64 ,115 ,104 ,111 ,112 ,101 ,101 ,95 ,111 ,114 ,100 ,101 ,114 ,95 ,111 ,102 ,103 ,95 ,100 ,98 ,95 ,48 ,48 ,48 ,48 ,48 ,48 ,52 ,55 ,46 ,111 ,114 ,100 ,101 ,114 ,95 ,102 ,117 ,108 ,102 ,105 ,108 ,109 ,101 ,110 ,116 ,95 ,103 ,114 ,111 ,117 ,112 ,95 ,116 ,97 ,98 ,95 ,48 ,48 ,48 ,52 ,55 ,53 ,53 ,56 ,18 ,98 ,8 ,0 ,16 ,0 ,26 ,90 ,18 ,88 ,74 ,42 ,8 ,0 ,16 ,0 ,26 ,0 ,34 ,0 ,42 ,0 ,50 ,0 ,58 ,0 ,66 ,0 ,72 ,0 ,80 ,0 ,88 ,0 ,98 ,0 ,104 ,0 ,114 ,0 ,122 ,0 ,130 ,1 ,0 ,136 ,1 ,0 ,146 ,1 ,0 ,154 ,1 ,0 ,90 ,42 ,8 ,0 ,16 ,0 ,26 ,0 ,34 ,0 ,42 ,0 ,50 ,0 ,58 ,0 ,66 ,0 ,72 ,0 ,80 ,0 ,88 ,0 ,98 ,0 ,104 ,0 ,114 ,0 ,122 ,0 ,130 ,1 ,0 ,136 ,1 ,0 ,146 ,1 ,0 ,154 ,1 ,0 ,34 ,0 ,26 ,158 ,2 ,8 ,226 ,249 ,213 ,243 ,191 ,230 ,17 ,16 ,248 ,139 ,216 ,243 ,191 ,230 ,17 ,26 ,135 ,2 ,18 ,132 ,2 ,74 ,42 ,8 ,0 ,16 ,0 ,26 ,0 ,34 ,0 ,42 ,0 ,50 ,0 ,58 ,0 ,66 ,0 ,72 ,0 ,80 ,0 ,88 ,0 ,98 ,0 ,104 ,0 ,114 ,0 ,122 ,0 ,130 ,1 ,0 ,136 ,1 ,0 ,146 ,1 ,0 ,154 ,1 ,0 ,90 ,213 ,1 ,8 ,231 ,84 ,16 ,198 ,182 ,180 ,4 ,26 ,5 ,110 ,105 ,110 ,106 ,97 ,34 ,11 ,54 ,56 ,57 ,56 ,57 ,51 ,57 ,57 ,50 ,50 ,50 ,42 ,2 ,77 ,89 ,50 ,5 ,75 ,101 ,100 ,97 ,104 ,58 ,7 ,75 ,111 ,100 ,105 ,97 ,110 ,103 ,66 ,9 ,233 ,152 ,191 ,232 ,144 ,168 ,229 ,190 ,183 ,72 ,2 ,80 ,215 ,136 ,137 ,238 ,5 ,88 ,222 ,152 ,243 ,132 ,6 ,98 ,5 ,56 ,49 ,52 ,52 ,48 ,104 ,215 ,136 ,137 ,238 ,5 ,114 ,0 ,122 ,0 ,130 ,1 ,0 ,136 ,1 ,1 ,146 ,1 ,0 ,154 ,1 ,111 ,10 ,107 ,123 ,34 ,102 ,111 ,114 ,109 ,97 ,116 ,116 ,101 ,100 ,65 ,100 ,100 ,114 ,101 ,115 ,115 ,34 ,58 ,34 ,71 ,117 ,97 ,44 ,32 ,80 ,97 ,104 ,97 ,110 ,103 ,44 ,32 ,77 ,97 ,108 ,97 ,121 ,115 ,105 ,97 ,34 ,44 ,34 ,114 ,101 ,103 ,105 ,111 ,110 ,34 ,58 ,123 ,34 ,108 ,97 ,116 ,105 ,116 ,117 ,100 ,101 ,34 ,58 ,52 ,46 ,50 ,51 ,49 ,54 ,51 ,57 ,50 ,44 ,34 ,108 ,111 ,110 ,103 ,105 ,116 ,117 ,100 ,101 ,34 ,58 ,49 ,48 ,49 ,46 ,57 ,57 ,49 ,56 ,55 ,55 ,56 ,57 ,57 ,57 ,57 ,57 ,57 ,54 ,125 ,125 ,72 ,14 ,34 ,2 ,77 ,89 ,24 ,254 ,136 ,213 ,134 ,6 ,34 ,36 ,100 ,99 ,57 ,102 ,52 ,55 ,51 ,101 ,45 ,48 ,51 ,50 ,56 ,45 ,52 ,101 ,50 ,53 ,45 ,97 ,102 ,56 ,102 ,45 ,55 ,48 ,53 ,55 ,102 ,99 ,52 ,99 ,57 ,54 ,99 ,49}
	//fmt.Println(string(by))
	//str := `\u0008\u0001\u0012\ufffd\u0003\n@shopee_order_ofg_db_00000007.order_fulfilment_group_tab_00007176\u0012b\u0008\u0000\u0010\u0000\u001aZ\u0012XJ*\u0008\u0000\u0010\u0000\u001a\u0000\"\u0000*\u00002\u0000:\u0000B\u0000H\u0000P\u0000X\u0000b\u0000h\u0000r\u0000z\u0000\ufffd\u0001\u0000\ufffd\u0001\u0000\ufffd\u0001\u0000\ufffd\u0001\u0000Z*\u0008\u0000\u0010\u0000\u001a\u0000\"\u0000*\u00002\u0000:\u0000B\u0000H\u0000P\u0000X\u0000b\u0000h\u0000r\u0000z\u0000\ufffd\u0001\u0000\ufffd\u0001\u0000\ufffd\u0001\u0000\ufffd\u0001\u0000\"\u0000\u001a\ufffd\u0001\u0008\ufffd\ufffd\ufffd\ufffd\ufffd\ufffd\u0011\u0010\ufffd\ufffd\ufffd\ufffd\ufffd\ufffd\u0011\u001a\ufffd\u0001\u0012\ufffd\u0001J*\u0008\u0000\u0010\u0000\u001a\u0000\"\u0000*\u00002\u0000:\u0000B\u0000H\u0000P\u0000X\u0000b\u0000h\u0000r\u0000z\u0000\ufffd\u0001\u0000\ufffd\u0001\u0000\ufffd\u0001\u0000\ufffd\u0001\u0000Z\ufffd\u0001\u0008\ufffd\ufffdR\u0010\ufffd\ufffd\ufffd\ufffd\u0004\u001a\u0004test\"\n6582204512*\u0002SG2\u0000:\u0000B+Premier Electrical \u0026 Pool Maintenance,, #15H\u0002P\ufffd\ufffdІ\u0006X\ufffd\ufffdІ\u0006b\u0006533811h\ufffd\ufffdІ\u0006r\u0000z\u0000\ufffd\u0001\u0000\ufffd\u0001\u0000\ufffd\u0001\u0000\ufffd\u0001$\n\u0002{}*\u0018label_address_label_home2\u0000H\u0002X\u0002\"\u0002SG\u0018\ufffd\ufffdІ\u0006\"$9d032301-25ad-4c7a-bbae-1f0d09572c94`
	var mp=map[string]int{"A":1,"B":1}
	for key:=range mp{
		fmt.Println(key)
	}
}
