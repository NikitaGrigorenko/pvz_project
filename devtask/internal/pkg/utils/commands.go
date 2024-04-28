package utils

import "fmt"

const (
	COLOR_RESET = "\033[0m"
	COLOR_RED   = "\033[31m"
	COLOR_BLUE  = "\033[34m"
)

// PrintHelp printing the existing command with their flags
func PrintHelp() {
	fmt.Println("Commands which are allowed to the tool:")

	commands := map[string]map[string]interface{}{
		"take_order": {
			"description": "takes an order from courier. Command has 3 flags:",
			"parameters": []string{
				"--order_id : id of an order to be taken",
				"--client_id : id of a client owns an order",
				"--date_exp : date of expiration in pickup point",
				"--weight : weight of the order in kilograms",
				"--price : initial price of the order in rubles (without package)",
				"--package : type of the package (0-packet, 1-box, 2-tape)",
			},
			"example": "./main take_order --order_id=5 --client_id=4 --date_exp=2024-03-29 --weight=40 --cost=32 --package=2",
		},
		"return_order": {
			"description": "returns an order to courier. Command has a flag:",
			"parameters": []string{
				"--order_id : id of an order to be returned",
			},
			"example": "./main return_order --order_id=4",
		},
		"give_order": {
			"description": "gives slice of orders to the client. Command has a flag:",
			"parameters": []string{
				"--slice : slice of ids to be given",
			},
			"example": "./main give_order --slice=7,12,11",
		},
		"list_orders": {
			"description": "shows orders of the client. Command has 3 flags:",
			"parameters": []string{
				"--client_id : id of a client owns an orders",
				"--last_n : number of last orders to be shown",
				"--inpp : if true -> shows only orders which are in pickup point",
			},
			"example": "./main list_orders --client_id=3 --last_n=3 --inpp=false",
		},
		"client_refund": {
			"description": "refunds an order from the client to Pickup Point. Command has 2 flags:",
			"parameters": []string{
				"--order_id : id of an order to be refunded",
				"--client_id : id of a client owns an order",
			},
			"example": "./main client_refund --client_id=3 --order_id=2",
		},
		"list_refunds": {
			"description": "shows 5 refunds on each page. Command has a flag:",
			"parameters": []string{
				"--page_number : the number of page to be shown",
			},
			"example": "./main list_refunds --page_number=1",
		},
		"pvz": {
			"description": "run an interactive menu to work with pickup points.",
			"parameters":  []string{},
			"example":     "./main pvz",
		},
		"help": {
			"description": "shows a list of commands which can be used with their description.",
			"parameters":  []string{},
			"example":     "./main help",
		},
	}

	for command, details := range commands {
		fmt.Printf("* %s`%s`%s %s%s\n", COLOR_RED, command, COLOR_RESET, details["description"].(string), COLOR_RESET)
		parameters := details["parameters"].([]string)
		if len(parameters) > 0 {
			for _, param := range parameters {
				fmt.Printf("   %s%s%s\n", COLOR_BLUE, param, COLOR_RESET)
			}
		}
		fmt.Printf("  sample invocation %s%s%s\n\n\n", COLOR_RED, details["example"].(string), COLOR_RESET)
	}
}
